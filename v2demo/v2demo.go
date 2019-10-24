package v2demo

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ditcraft/indexer/database"
	"gopkg.in/mgo.v2/bson"

	"github.com/ditcraft/indexer/smartcontracts_v2/KNWToken"
	"github.com/ditcraft/indexer/smartcontracts_v2/KNWVoting"
	"github.com/ditcraft/indexer/smartcontracts_v2/ditDemoCoordinator"
	"github.com/ditcraft/indexer/smartcontracts_v2/ditToken"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/golang/glog"
)

var mutex = &sync.Mutex{}

type logKNWTokenEvent struct {
	Who   common.Address
	Id    *big.Int
	Value *big.Int
}

type logKNWTokenNewLabelEvent struct {
	Id    *big.Int
	Label string
}

type logERC20TransferEvent struct {
	From  common.Address
	To    common.Address
	Value *big.Int
}

type logDitCoordinatorInit struct {
	Repository [32]byte
	Who        common.Address
}

type logDitCoordinatorPropose struct {
	Repository  [32]byte
	Proposal    *big.Int
	Who         common.Address
	KnowledgeID *big.Int
	NumberOfKNW *big.Int
}

type logDitCoordinatorCommit struct {
	Repository    [32]byte
	Proposal      *big.Int
	Who           common.Address
	KnowledgeID   *big.Int
	Stake         *big.Int
	NumberOfKNW   *big.Int
	NumberOfVotes *big.Int
}

type logDitCoordinatorOpen struct {
	Repository    [32]byte
	Proposal      *big.Int
	Who           common.Address
	KnowledgeID   *big.Int
	Accept        bool
	NumberOfVotes *big.Int
}

type logDitCoordinatorFinalizeVote struct {
	Repository  [32]byte
	Proposal    *big.Int
	Who         common.Address
	KnowledgeID *big.Int
	VotedRight  bool
	NumberOfKNW *big.Int
}

type logDitCoordinatorFinalizeProposal struct {
	Repository  [32]byte
	Proposal    *big.Int
	KnowledgeID *big.Int
	Accepted    bool
}

var logTransferSig = []byte("Transfer(address,address,uint256)")
var logTransferSigHash = crypto.Keccak256Hash(logTransferSig)
var logDitCoordinatorInitSig = []byte("InitializeRepository(bytes32,address)")
var logDitCoordinatorInitSigHash = crypto.Keccak256Hash(logDitCoordinatorInitSig)
var logDitCoordinatorProposeSig = []byte("ProposeCommit(bytes32,uint256,address,uint256,uint256)")
var logDitCoordinatorProposeSigHash = crypto.Keccak256Hash(logDitCoordinatorProposeSig)
var logDitCoordinatorCommitSig = []byte("CommitVote(bytes32,uint256,address,uint256,uint256,uint256,uint256)")
var logDitCoordinatorCommitSigHash = crypto.Keccak256Hash(logDitCoordinatorCommitSig)
var logDitCoordinatorOpenSig = []byte("OpenVote(bytes32,uint256,address,uint256,bool,uint256)")
var logDitCoordinatorOpenSigHash = crypto.Keccak256Hash(logDitCoordinatorOpenSig)
var logDitCoordinatorFinalizeVoteSig = []byte("FinalizeVote(bytes32,uint256,address,uint256,bool,uint256)")
var logDitCoordinatorFinalizeVoteSigHash = crypto.Keccak256Hash(logDitCoordinatorFinalizeVoteSig)
var logDitCoordinatorFinalizeProposalSig = []byte("FinalizeProposal(bytes32,uint256,uint256,bool)")
var logDitCoordinatorFinalizeProposalSigHash = crypto.Keccak256Hash(logDitCoordinatorFinalizeProposalSig)
var logKNWTokenMint = []byte("Mint(address,uint256,uint256)")
var logKNWTokenMintSig = crypto.Keccak256Hash(logKNWTokenMint)
var logKNWTokenBurn = []byte("Burn(address,uint256,uint256)")
var logKNWTokenBurnSig = crypto.Keccak256Hash(logKNWTokenBurn)
var logKNWNewLabel = []byte("NewLabel(uint256,string)")
var logKNWNewLabelSig = crypto.Keccak256Hash(logKNWNewLabel)
var logERC20Transfer = []byte("Transfer(address,address,uint256)")
var logERC20TransferSig = crypto.Keccak256Hash(logERC20Transfer)

var coordinatorAddress common.Address
var knwTokenAddress common.Address
var ditTokenAddress common.Address

var ditDemoCoordinatorABI abi.ABI
var knwTokenABI abi.ABI
var ditTokenABI abi.ABI

// Init func
func Init() {
	coordinatorAddress = common.HexToAddress(os.Getenv("CONTRACT_DIT_COORDINATOR_DEMO_V2"))
	knwTokenAddress = common.HexToAddress(os.Getenv("CONTRACT_KNW_TOKEN_DEMO_V2"))
	ditTokenAddress = common.HexToAddress(os.Getenv("CONTRACT_DIT_TOKEN_DEMO"))

	var err error
	ditDemoCoordinatorABI, err = abi.JSON(strings.NewReader(os.Getenv("DIT_DEMO_COORDINATOR_ABI_V2")))
	if err != nil {
		glog.Fatal(err)
	}

	knwTokenABI, err = abi.JSON(strings.NewReader(os.Getenv("KNW_TOKEN_ABI_V2")))
	if err != nil {
		glog.Fatal(err)
	}

	ditTokenABI, err = abi.JSON(strings.NewReader(os.Getenv("DIT_TOKEN_ABI")))
	if err != nil {
		glog.Fatal(err)
	}
}

// GetCurrentBlock func
func GetCurrentBlock() (int64, error) {
	connection, err := getConnection()
	if err != nil {
		return 0, err
	}

	header, err := connection.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return 0, err
	}

	return header.Number.Int64(), nil
}

// GetOldEvents func
func GetOldEvents(_fromBlock int64, _toBlock int64) {
	connection, err := getConnection()
	if err != nil {
		glog.Error(err)
	}

	offset := int64(100000)
	offsetChanged := false
	for _fromBlock < _toBlock {
		toBlock := _toBlock
		if _toBlock-_fromBlock > offset {
			toBlock = _fromBlock + offset
		}
		glog.Info("[OldEvents-Demo] Going from " + strconv.Itoa(int(_fromBlock)) + " to " + strconv.Itoa(int(toBlock)))
		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(_fromBlock),
			ToBlock:   big.NewInt(toBlock),
			Addresses: []common.Address{
				coordinatorAddress,
				knwTokenAddress,
				ditTokenAddress,
			},
		}

		events, err := connection.FilterLogs(context.Background(), query)
		if err != nil {
			if strings.Contains(err.Error(), "unexpected EOF") {
				glog.Info("[OldEvents-Demo] Got 'unexpected EOF', reducing the offset to " + strconv.Itoa(int(offset)))
				connection, err = getConnection()
				if err != nil {
					glog.Error(err)
				}
				offsetChanged = true
				offset = offset * 75 / 100
			} else {
				glog.Fatal(err)
			}
		} else {
			if !offsetChanged && offset < int64(100000) {
				offset = offset * 100 / 80
			}
			offsetChanged = false
			_fromBlock = toBlock + 1
		}

		for _, event := range events {
			switch event.Topics[0].Hex() {
			case logKNWNewLabelSig.Hex():
				handleNewLabel(&event, connection)
			case logERC20TransferSig.Hex():
				handleERC20Transfer(&event, connection)
			case logKNWTokenMintSig.Hex():
				handleKNWTokenMint(&event, connection)
			case logKNWTokenBurnSig.Hex():
				handleKNWTokenBurn(&event, connection)
			case logDitCoordinatorInitSigHash.Hex():
				handleDitDemoCoordinatorInit(&event, connection)
			case logDitCoordinatorProposeSigHash.Hex():
				handleDitDemoCoordinatorPropose(&event, connection)
			case logDitCoordinatorCommitSigHash.Hex():
				handleDitDemoCoordinatorCommit(&event, connection)
			case logDitCoordinatorOpenSigHash.Hex():
				handleDitDemoCoordinatorOpen(&event, connection)
			case logDitCoordinatorFinalizeVoteSigHash.Hex():
				handleDitDemoCoordinatorFinalizeVote(&event, connection)
			case logDitCoordinatorFinalizeProposalSigHash.Hex():
				handleDitDemoCoordinatorFinalizeProposal(&event, connection)
			}
		}
		time.Sleep(100 * time.Millisecond)
	}
	glog.Info("[OldEvents-Demo] Finished indexing old blocks for demo mode (last block was " + strconv.FormatInt(_toBlock, 10) + ")")
}

// WatchEvents func
func WatchEvents() {
	glog.Info("Starting watcher for demo mode...")
	connection, err := getConnection()
	if err != nil {
		glog.Error(err)
		time.Sleep(1 * time.Second)
		go WatchEvents()
		return
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{
			coordinatorAddress,
			knwTokenAddress,
			ditTokenAddress,
		},
	}

	eventChan := make(chan types.Log)
	sub, err := connection.SubscribeFilterLogs(context.Background(), query, eventChan)
	if err != nil {
		glog.Error(err)
		time.Sleep(500 * time.Millisecond)
		go WatchEvents()
		return
	}

	glog.Info("Successfully started watcher for demo mode.")

	for {
		select {
		case err := <-sub.Err():
			glog.Error(err)
			time.Sleep(500 * time.Millisecond)
			go WatchEvents()
			return
		case event := <-eventChan:
			switch event.Topics[0].Hex() {
			case logKNWNewLabelSig.Hex():
				handleNewLabel(&event, connection)
			case logERC20TransferSig.Hex():
				handleERC20Transfer(&event, connection)
			case logKNWTokenMintSig.Hex():
				handleKNWTokenMint(&event, connection)
			case logKNWTokenBurnSig.Hex():
				handleKNWTokenBurn(&event, connection)
			case logDitCoordinatorInitSigHash.Hex():
				handleDitDemoCoordinatorInit(&event, connection)
			case logDitCoordinatorProposeSigHash.Hex():
				handleDitDemoCoordinatorPropose(&event, connection)
			case logDitCoordinatorCommitSigHash.Hex():
				handleDitDemoCoordinatorCommit(&event, connection)
			case logDitCoordinatorOpenSigHash.Hex():
				handleDitDemoCoordinatorOpen(&event, connection)
			case logDitCoordinatorFinalizeVoteSigHash.Hex():
				handleDitDemoCoordinatorFinalizeVote(&event, connection)
			case logDitCoordinatorFinalizeProposalSigHash.Hex():
				handleDitDemoCoordinatorFinalizeProposal(&event, connection)
			}
		}
	}
}

func handleNewLabel(event *types.Log, connection *ethclient.Client) {
	glog.Infof("[%d] New KNW Label Event", event.BlockNumber)
	var newLabelEvent logKNWTokenNewLabelEvent
	err := knwTokenABI.Unpack(&newLabelEvent, "NewLabel", event.Data)
	if err != nil {
		glog.Error(err)
		return
	}
	newLabelEvent.Id = event.Topics[1].Big()

	block, err := connection.BlockByNumber(context.Background(), big.NewInt(int64(event.BlockNumber)))
	if err != nil {
		glog.Error(err)
		return
	}

	var newLabel database.KnowledgeLabel
	newLabel.ID = int(newLabelEvent.Id.Int64())
	newLabel.Label = newLabelEvent.Label
	newLabel.TotalSupply = 0.0
	newLabel.LastActivityDate = time.Unix(int64(block.Time()), 0)

	err = database.Insert("knowledge_labels_demo", newLabel)
	if err != nil {
		glog.Error(err)
	}

}

func handleERC20Transfer(event *types.Log, connection *ethclient.Client) {
	glog.Infof("[%d] Token Transfer Event", event.BlockNumber)
	var transferEvent logERC20TransferEvent
	err := ditTokenABI.Unpack(&transferEvent, "Transfer", event.Data)
	if err != nil {
		glog.Error(err)
		return
	}
	transferEvent.From = common.HexToAddress(event.Topics[1].Hex())
	transferEvent.To = common.HexToAddress(event.Topics[2].Hex())

	err = processDitTokenRefresh(transferEvent.From, connection)
	if err != nil {
		glog.Error(err)
	}
	err = processDitTokenRefresh(transferEvent.To, connection)
	if err != nil {
		glog.Error(err)
	}
}

func processDitTokenRefresh(_address common.Address, connection *ethclient.Client) error {
	if strings.Contains(os.Getenv("AUTOVALIDATORS"), _address.Hex()) {
		return nil
	}

	var user database.User
	userBytes, err := database.Get("users", []string{"dit_address"}, []string{"=="}, []interface{}{_address.Hex()})
	if err != nil {
		return err
	}

	if len(userBytes) > 1 {
		return fmt.Errorf("Expected one user to be found but found %d", len(userBytes))
	} else if len(userBytes) < 1 {
		return nil
	}

	bsonErr := bson.Unmarshal(userBytes[0], &user)
	if bsonErr != nil {
		return bsonErr
	}

	ditTokenInstance, err := getDitTokenInstance(connection)
	if err != nil {
		return err
	}

	balance, err := ditTokenInstance.BalanceOf(nil, _address)
	if err != nil {
		return err
	}

	currentDaiBalance, err := connection.BalanceAt(context.Background(), _address, nil)
	if err != nil {
		return err
	}

	floatIndividualDitBalance, _ := (new(big.Float).Quo((new(big.Float).SetInt(balance)), big.NewFloat(1000000000000000000))).Float64()
	floatIndividualDaiBalance, _ := (new(big.Float).Quo((new(big.Float).SetInt(currentDaiBalance)), big.NewFloat(1000000000000000000))).Float64()

	user.XDITBalance = floatIndividualDitBalance
	user.XDAIBalance = floatIndividualDaiBalance

	err = database.Replace("users", "dit_address", _address.Hex(), user)
	if err != nil {
		return err
	}

	return nil
}

func handleKNWTokenMint(event *types.Log, connection *ethclient.Client) {
	glog.Infof("[%d] Token Mint Event", event.BlockNumber)
	var tokenEvent logKNWTokenEvent
	err := knwTokenABI.Unpack(&tokenEvent, "Mint", event.Data)
	if err != nil {
		glog.Error(err)
		return
	}
	tokenEvent.Who = common.HexToAddress(event.Topics[1].Hex())

	KNWTokenInstance, err := getKNWTokenInstance(connection)
	if err != nil {
		glog.Error(err)
		return
	}

	totalSupply, err := KNWTokenInstance.TotalIDSupply(nil, tokenEvent.Id)
	if err != nil {
		glog.Error(err)
		return
	}

	floatTotalSupply, _ := (new(big.Float).Quo((new(big.Float).SetInt(totalSupply)), big.NewFloat(1000000000000000000))).Float64()

	block, err := connection.BlockByNumber(context.Background(), big.NewInt(int64(event.BlockNumber)))
	if err != nil {
		glog.Error(err)
		return
	}

	err = database.Update("knowledge_labels_demo", "id", interface{}(int(tokenEvent.Id.Int64())), []string{"total_supply", "last_activity_date"}, []interface{}{floatTotalSupply, time.Unix(int64(block.Time()), 0)})
	if err != nil {
		glog.Error(err)
		return
	}

	if strings.Contains(os.Getenv("AUTOVALIDATORS"), common.HexToAddress(event.Topics[1].Hex()).Hex()) {
		return
	}

	userBytes, err := database.Get("users", []string{"dit_address"}, []string{"=="}, []interface{}{tokenEvent.Who.Hex()})
	if err != nil {
		glog.Error(err)
		return
	}

	var user database.User
	if len(userBytes) == 0 {
		user.DitAddress = tokenEvent.Who.Hex()
	} else if len(userBytes) == 1 {
		bsonErr := bson.Unmarshal(userBytes[0], &user)
		if bsonErr != nil {
			glog.Error(bsonErr)
			return
		}
	}

	knwLabel, err := database.GetKnowledgeLabel(false, tokenEvent.Id)
	if err != nil {
		glog.Error(err)
		return
	}

	found := false
	for i, label := range user.KNWTokensDemo {
		if label.Label == knwLabel {
			found = true
			amount, ok := new(big.Int).SetString(label.RawBalance, 10)
			if !ok {
				glog.Error("Conversion not okay")
				return
			}
			newValue := new(big.Int).Add(amount, tokenEvent.Value)
			user.KNWTokensDemo[i].RawBalance = newValue.String()
			floatBalance, _ := (new(big.Float).Quo((new(big.Float).SetInt(newValue)), big.NewFloat(1000000000000000000))).Float64()
			user.KNWTokensDemo[i].Balance = floatBalance
		}
	}
	if !found {
		floatBalance, _ := (new(big.Float).Quo((new(big.Float).SetInt(tokenEvent.Value)), big.NewFloat(1000000000000000000))).Float64()
		newLabel := database.KNWLabels{Label: knwLabel, RawBalance: tokenEvent.Value.String(), Balance: floatBalance}
		user.KNWTokensDemo = append(user.KNWTokensDemo, newLabel)
	}

	err = database.Replace("users", "dit_address", tokenEvent.Who.Hex(), user)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			err = database.Insert("users", user)
			if err != nil {
				glog.Error(err)
				return
			}
		}
		if err != nil {
			glog.Error(err)
			return
		}
	}
}

func handleKNWTokenBurn(event *types.Log, connection *ethclient.Client) {
	glog.Infof("[%d] Token Burn Event", event.BlockNumber)
	var tokenEvent logKNWTokenEvent
	err := knwTokenABI.Unpack(&tokenEvent, "Burn", event.Data)
	if err != nil {
		glog.Error(err)
		return
	}
	tokenEvent.Who = common.HexToAddress(event.Topics[1].Hex())

	KNWTokenInstance, err := getKNWTokenInstance(connection)
	if err != nil {
		glog.Error(err)
		return
	}

	totalSupply, err := KNWTokenInstance.TotalIDSupply(nil, tokenEvent.Id)
	if err != nil {
		glog.Error(err)
		return
	}

	floatTotalSupply, _ := (new(big.Float).Quo((new(big.Float).SetInt(totalSupply)), big.NewFloat(1000000000000000000))).Float64()

	block, err := connection.BlockByNumber(context.Background(), big.NewInt(int64(event.BlockNumber)))
	if err != nil {
		glog.Error(err)
		return
	}

	err = database.Update("knowledge_labels_demo", "id", interface{}(int(tokenEvent.Id.Int64())), []string{"total_supply", "last_activity_date"}, []interface{}{floatTotalSupply, time.Unix(int64(block.Time()), 0)})
	if err != nil {
		glog.Error(err)
		return
	}

	if strings.Contains(os.Getenv("AUTOVALIDATORS"), common.HexToAddress(event.Topics[1].Hex()).Hex()) {
		return
	}

	userBytes, err := database.Get("users", []string{"dit_address"}, []string{"=="}, []interface{}{tokenEvent.Who.Hex()})
	if err != nil {
		glog.Error(err)
		return
	}

	var user database.User
	if len(userBytes) == 0 {
		user.DitAddress = tokenEvent.Who.Hex()
	} else if len(userBytes) == 1 {
		bsonErr := bson.Unmarshal(userBytes[0], &user)
		if bsonErr != nil {
			glog.Error(bsonErr)
			return
		}
	}

	knwLabel, err := database.GetKnowledgeLabel(false, tokenEvent.Id)
	if err != nil {
		glog.Error(err)
		return
	}

	found := false
	for i, label := range user.KNWTokensDemo {
		if label.Label == knwLabel {
			found = true
			amount, ok := new(big.Int).SetString(label.RawBalance, 10)
			if !ok {
				glog.Error("Conversion not okay")
				return
			}
			newValue := new(big.Int).Sub(amount, tokenEvent.Value)
			user.KNWTokensDemo[i].RawBalance = newValue.String()
			floatBalance, _ := (new(big.Float).Quo((new(big.Float).SetInt(newValue)), big.NewFloat(1000000000000000000))).Float64()
			user.KNWTokensDemo[i].Balance = floatBalance
		}
	}
	if !found {
		floatBalance, _ := (new(big.Float).Quo((new(big.Float).SetInt(tokenEvent.Value)), big.NewFloat(1000000000000000000))).Float64()
		newLabel := database.KNWLabels{Label: knwLabel, RawBalance: tokenEvent.Value.String(), Balance: floatBalance}
		user.KNWTokensDemo = append(user.KNWTokensDemo, newLabel)
	}

	err = database.Replace("users", "dit_address", tokenEvent.Who.Hex(), user)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			err = database.Insert("users", user)
			if err != nil {
				glog.Error(err)
				return
			}
		}
		if err != nil {
			glog.Error(err)
			return
		}
	}
}

func handleDitDemoCoordinatorInit(event *types.Log, connection *ethclient.Client) {
	glog.Infof("[%d] Repo Init Event", event.BlockNumber)
	var initEvent logDitCoordinatorInit
	initEvent.Repository = event.Topics[1]
	initEvent.Who = common.HexToAddress(event.Topics[2].Hex())

	connection, err := getConnection()
	if err != nil {
		glog.Error(err)
		return
	}

	ditDemoCoordinatorInstance, err := getDitDemoCoordinatorInstance(connection)
	if err != nil {
		glog.Error(err)
		return
	}

	KNWTokenInstance, err := getKNWTokenInstance(connection)
	if err != nil {
		glog.Error(err)
		return
	}

	repository, err := ditDemoCoordinatorInstance.Repositories(nil, initEvent.Repository)
	if err != nil {
		glog.Error(err)
		return
	}

	var newRepository database.Repository
	newRepository.Hash = event.Topics[1].Hex()
	nameParts := strings.SplitN(repository.Name, "/", 2)
	newRepository.Name = nameParts[1]
	provider := strings.ToLower(nameParts[0])
	if strings.Contains(provider, "github") {
		newRepository.Provider = "GitHub"
	} else if strings.Contains(provider, "bitbucket") {
		newRepository.Provider = "Bitbucket"
	} else if strings.Contains(provider, "gitlab") {
		newRepository.Provider = "GitLab"
	} else {
		newRepository.Provider = "Unknown"
	}
	newRepository.Majority = int(repository.VotingMajority.Int64())
	newRepository.URL = "https://" + repository.Name

	block, err := connection.BlockByNumber(context.Background(), big.NewInt(int64(event.BlockNumber)))
	if err != nil {
		glog.Error(err)
		return
	}

	newRepository.CreationDate = time.Unix(int64(block.Time()), 0)
	newRepository.LastActivityDate = newRepository.CreationDate

	ids, err := ditDemoCoordinatorInstance.GetKnowledgeIDs(nil, initEvent.Repository)
	if err != nil {
		glog.Error(err)
		return
	}
	for i := 0; i < len(ids); i++ {
		label, err := KNWTokenInstance.LabelOfID(nil, ids[i])
		if err != nil {
			glog.Error(err)
			return
		}
		if len(label) > 0 {
			newRepository.KnowledgeLabels = append(newRepository.KnowledgeLabels, label)
		}
	}

	err = database.Insert("repositories_demo", newRepository)
	if err != nil {
		glog.Error(err)
		return
	}

	userBytes, err := database.Get("users", []string{"dit_address"}, []string{"=="}, []interface{}{initEvent.Who.Hex()})
	if err != nil {
		glog.Error(err)
	}

	var user database.User
	if len(userBytes) == 0 {
		user.DitAddress = initEvent.Who.Hex()
	} else if len(userBytes) == 1 {
		bsonErr := bson.Unmarshal(userBytes[0], &user)
		if bsonErr != nil {
			glog.Error(bsonErr)
			return
		}
	}

	currentDaiBalance, err := connection.BalanceAt(context.Background(), initEvent.Who, nil)
	if err != nil {
		glog.Error(err)
	} else {
		floatIndividualDaiBalance, _ := (new(big.Float).Quo((new(big.Float).SetInt(currentDaiBalance)), big.NewFloat(1000000000000000000))).Float64()
		user.XDAIBalance = floatIndividualDaiBalance
	}

	var newUserRepository database.RepositoryDetailsUser
	newUserRepository.Hash = newRepository.Hash
	newUserRepository.Name = newRepository.Name
	newUserRepository.Provider = newRepository.Provider
	newUserRepository.URL = newRepository.URL
	newUserRepository.Notifications = true
	newUserRepository.LastActivityDate = time.Unix(int64(block.Time()), 0)

	user.RepositoriesDemo = append(user.RepositoriesDemo, newUserRepository)

	err = database.Replace("users", "dit_address", initEvent.Who.Hex(), user)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			err = database.Insert("users", user)
			if err != nil {
				glog.Error(err)
				return
			}
		}
		if err != nil {
			glog.Error(err)
			return
		}
	}
}

func handleDitDemoCoordinatorPropose(event *types.Log, connection *ethclient.Client) {
	glog.Infof("[%d] Propose Event", event.BlockNumber)
	var proposeEvent logDitCoordinatorPropose
	err := ditDemoCoordinatorABI.Unpack(&proposeEvent, "ProposeCommit", event.Data)
	if err != nil {
		glog.Error(err)
	}
	proposeEvent.Repository = event.Topics[1]
	proposeEvent.Proposal = event.Topics[2].Big()
	proposeEvent.Who = common.HexToAddress(event.Topics[3].Hex())

	ditDemoCoordinatorInstance, err := getDitDemoCoordinatorInstance(connection)
	if err != nil {
		glog.Error(err)
		return
	}

	knwVotingInstance, err := getKNWVotingInstance(connection)
	if err != nil {
		glog.Error(err)
		return
	}

	block, err := connection.BlockByNumber(context.Background(), big.NewInt(int64(event.BlockNumber)))
	if err != nil {
		glog.Error(err)
		return
	}

	repositoryBytes, err := database.Get("repositories_demo", []string{"hash"}, []string{"=="}, []interface{}{event.Topics[1].Hex()})
	if err != nil {
		glog.Error(err)
		return
	}

	if len(repositoryBytes) != 1 {
		glog.Errorf("Expected one repository to be found but found %d\n", len(repositoryBytes))
		return
	}

	var repository database.Repository
	bsonErr := bson.Unmarshal(repositoryBytes[0], &repository)
	if bsonErr != nil {
		glog.Error(bsonErr)
		return
	}

	proposal, err := ditDemoCoordinatorInstance.ProposalsOfRepository(nil, event.Topics[1], proposeEvent.Proposal)
	if err != nil {
		glog.Error(err)
		return
	}

	vote, err := knwVotingInstance.Votes(nil, proposal.KNWVoteID)
	if err != nil {
		glog.Error(err)
		return
	}

	var newVote database.ProposalDetailsRepository
	newVote.KNWVoteID = int(proposal.KNWVoteID.Int64())
	newVote.VoteDate = time.Unix(int64(block.Time()), 0)
	repository.Proposals = append(repository.Proposals, newVote)

	found := false
	for i, contributor := range repository.Contributors {
		if contributor.Address == proposeEvent.Who.Hex() {
			repository.Contributors[i].AmountOfProposals++
			repository.Contributors[i].LastActivityDate = time.Unix(int64(block.Time()), 0)
			found = true
			break
		}
	}
	if !found {
		if !strings.Contains(os.Getenv("AUTOVALIDATORS"), proposeEvent.Who.Hex()) {
			newContributor := database.ProposalDetailsContributor{
				Address:             proposeEvent.Who.Hex(),
				LastActivityDate:    time.Unix(int64(block.Time()), 0),
				EarnedKNW:           0,
				AmountOfProposals:   1,
				AmountOfValidations: 0,
			}
			repository.Contributors = append(repository.Contributors, newContributor)
		}
	}
	err = database.Replace("repositories_demo", "hash", event.Topics[1].Hex(), repository)
	if err != nil {
		glog.Error(err)
		return
	}

	knwLabel, err := database.GetKnowledgeLabel(false, proposeEvent.KnowledgeID)
	if err != nil {
		glog.Error(err)
		return
	}

	var newProposal database.Proposal
	newProposal.ProposalID = int(proposeEvent.Proposal.Int64())
	newProposal.KNWVoteID = int(proposal.KNWVoteID.Int64())
	newProposal.KNWLabel = knwLabel
	newProposal.Repository = event.Topics[1].Hex()
	newProposal.Proposer = proposeEvent.Who.Hex()
	newProposal.Topic = proposal.Description
	newProposal.Identifier = proposal.Identifier
	newProposal.CreationDate = time.Unix(int64(block.Time()), 0)
	newProposal.CommitPhaseEnd = time.Unix(vote.CommitEndDate.Int64(), 0)
	newProposal.RevealPhaseEnd = time.Unix(vote.OpenEndDate.Int64(), 0)
	newProposal.Finalized = false
	newProposal.Accepted = false
	newProposal.Votes = database.ProposalVotes{}
	newProposal.Participants = []database.ProposalParticipants{}
	floatIndividualStake, _ := (new(big.Float).Quo((new(big.Float).SetInt(proposal.IndividualStake)), big.NewFloat(1000000000000000000))).Float64()
	newProposal.StakePerParticipant = floatIndividualStake
	newProposal.TotalStake = floatIndividualStake
	newProposal.TotalMintedKNW = 0

	var newParticipant database.ProposalParticipants
	newParticipant.Address = proposeEvent.Who.Hex()
	newParticipant.Opened = true
	newParticipant.Finalized = false
	newParticipant.KNWDifference = 0.0
	floatUsedKNW, _ := (new(big.Float).Quo((new(big.Float).SetInt(proposeEvent.NumberOfKNW)), big.NewFloat(1000000000000000000))).Float64()
	newParticipant.UsedKNW = floatUsedKNW
	newParticipant.VotedRight = false
	newParticipant.Votes = 0

	newProposal.Participants = append(newProposal.Participants, newParticipant)

	err = database.Insert("proposals_demo", &newProposal)
	if err != nil {
		glog.Error(err)
	}

	userBytes, err := database.Get("users", []string{"dit_address"}, []string{"=="}, []interface{}{proposeEvent.Who.Hex()})
	if err != nil {
		glog.Error(err)
	}

	var user database.User
	if len(userBytes) == 0 {
		user.DitAddress = proposeEvent.Who.Hex()
	} else if len(userBytes) == 1 {
		bsonErr := bson.Unmarshal(userBytes[0], &user)
		if bsonErr != nil {
			glog.Error(bsonErr)
			return
		}
	}

	currentDaiBalance, err := connection.BalanceAt(context.Background(), proposeEvent.Who, nil)
	if err != nil {
		glog.Error(err)
	} else {
		floatIndividualDaiBalance, _ := (new(big.Float).Quo((new(big.Float).SetInt(currentDaiBalance)), big.NewFloat(1000000000000000000))).Float64()
		user.XDAIBalance = floatIndividualDaiBalance
	}

	var newUserProposal database.ProposalDetailsUser
	newUserProposal.Finalized = false
	newUserProposal.IsProposer = true
	newUserProposal.KNWDifference = 0.0
	newUserProposal.KNWVoteID = newProposal.KNWVoteID
	newUserProposal.Opened = true
	newUserProposal.UsedKNW = floatUsedKNW
	newUserProposal.Stake = floatIndividualStake
	newUserProposal.VoteDate = newProposal.CreationDate
	newUserProposal.VotedRight = false

	user.ProposalsDemo = append(user.ProposalsDemo, newUserProposal)

	index := -1
	for i, singleRepository := range user.RepositoriesDemo {
		if singleRepository.Hash == repository.Hash {
			index = i
			break
		}
	}
	if index == -1 {
		var newRepository database.RepositoryDetailsUser
		newRepository.Hash = repository.Hash
		newRepository.Name = repository.Name
		newRepository.URL = repository.URL
		newRepository.Provider = repository.Provider
		user.RepositoriesDemo = append(user.RepositoriesDemo, newRepository)
		index = len(user.RepositoriesDemo) - 1
	}

	user.RepositoriesDemo[index].LastActivityDate = time.Unix(int64(block.Time()), 0)
	user.RepositoriesDemo[index].AmountOfProposals++

	err = database.Replace("users", "dit_address", proposeEvent.Who.Hex(), user)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			err = database.Insert("users", user)
			if err != nil {
				glog.Error(err)
				return
			}
		}
		if err != nil {
			glog.Error(err)
			return
		}
	}

	var newNotification database.NotificationProposalStarted
	newNotification.LiveMode = false
	newNotification.KnowledgeLabel = newProposal.KNWLabel
	newNotification.KNWVoteID = newProposal.KNWVoteID
	newNotification.Description = newProposal.Topic
	newNotification.Identifier = newProposal.Identifier
	newNotification.ProposerTwitterID = user.TwitterID
	newNotification.RepositoryHash = repository.Hash
	newNotification.RepositoryName = repository.Name
	newNotification.CommitUntil = newProposal.CommitPhaseEnd
	newNotification.RevealUntil = newProposal.RevealPhaseEnd

	err = database.Notify(repository.Hash, newNotification)
	if err != nil {
		glog.Error(err)
		return
	}
}

func handleDitDemoCoordinatorCommit(event *types.Log, connection *ethclient.Client) {
	glog.Infof("[%d] Commit Vote Event", event.BlockNumber)
	var commitEvent logDitCoordinatorCommit
	err := ditDemoCoordinatorABI.Unpack(&commitEvent, "CommitVote", event.Data)
	if err != nil {
		glog.Error(err)
	}
	commitEvent.Repository = event.Topics[1]
	commitEvent.Proposal = event.Topics[2].Big()
	commitEvent.Who = common.HexToAddress(event.Topics[3].Hex())

	block, err := connection.BlockByNumber(context.Background(), big.NewInt(int64(event.BlockNumber)))
	if err != nil {
		glog.Error(err)
		return
	}

	repositoryBytes, err := database.Get("repositories_demo", []string{"hash"}, []string{"=="}, []interface{}{event.Topics[1].Hex()})
	if err != nil {
		glog.Error(err)
		return
	}

	if len(repositoryBytes) != 1 {
		glog.Errorf("Expected one repository to be found but found %d\n", len(repositoryBytes))
		return
	}

	var repository database.Repository
	bsonErr := bson.Unmarshal(repositoryBytes[0], &repository)
	if bsonErr != nil {
		glog.Error(bsonErr)
		return
	}

	repository.LastActivityDate = time.Unix(int64(block.Time()), 0)

	found := false
	for i, contributor := range repository.Contributors {
		if contributor.Address == commitEvent.Who.Hex() {
			repository.Contributors[i].AmountOfValidations++
			repository.Contributors[i].LastActivityDate = time.Unix(int64(block.Time()), 0)
			found = true
			break
		}
	}
	if !found {
		if !strings.Contains(os.Getenv("AUTOVALIDATORS"), commitEvent.Who.Hex()) {
			newContributor := database.ProposalDetailsContributor{
				Address:             commitEvent.Who.Hex(),
				LastActivityDate:    time.Unix(int64(block.Time()), 0),
				EarnedKNW:           0,
				AmountOfProposals:   0,
				AmountOfValidations: 1,
			}
			repository.Contributors = append(repository.Contributors, newContributor)
		}
	}
	err = database.Replace("repositories_demo", "hash", event.Topics[1].Hex(), repository)
	if err != nil {
		glog.Error(err)
		return
	}

	proposalBytes, err := database.Get("proposals_demo", []string{"repository", "id"}, []string{"=="}, []interface{}{event.Topics[1].Hex(), int(commitEvent.Proposal.Int64())})
	if err != nil {
		glog.Error(err)
		return
	}

	if len(proposalBytes) != 1 {
		glog.Errorf("Expected one proposal to be found but found %d\n", len(proposalBytes))
		return
	}

	var proposal database.Proposal
	bsonErr = bson.Unmarshal(proposalBytes[0], &proposal)
	if bsonErr != nil {
		glog.Error(bsonErr)
		return
	}

	proposal.Votes.ParticipantsUnrevealed++
	floatVotes, _ := (new(big.Float).Quo((new(big.Float).SetInt(commitEvent.NumberOfVotes)), big.NewFloat(1000000000000000000))).Float64()
	proposal.Votes.VotesUnrevealed += floatVotes
	floatStake, _ := (new(big.Float).Quo((new(big.Float).SetInt(commitEvent.Stake)), big.NewFloat(1000000000000000000))).Float64()
	proposal.TotalStake += floatStake

	var newParticipant database.ProposalParticipants
	newParticipant.Address = commitEvent.Who.Hex()
	newParticipant.Opened = false
	newParticipant.Finalized = false
	newParticipant.KNWDifference = 0.0
	floatUsedKNW, _ := (new(big.Float).Quo((new(big.Float).SetInt(commitEvent.NumberOfKNW)), big.NewFloat(1000000000000000000))).Float64()
	newParticipant.UsedKNW = floatUsedKNW
	newParticipant.VotedRight = false
	newParticipant.Votes = floatVotes

	proposal.Participants = append(proposal.Participants, newParticipant)

	err = database.Replace("proposals_demo", "knw_vote_id", proposal.KNWVoteID, proposal)
	if err != nil {
		glog.Error(err)
		return
	}

	userBytes, err := database.Get("users", []string{"dit_address"}, []string{"=="}, []interface{}{commitEvent.Who.Hex()})
	if err != nil {
		glog.Error(err)
	}

	if strings.Contains(os.Getenv("AUTOVALIDATORS"), commitEvent.Who.Hex()) {
		return
	}

	var user database.User
	if len(userBytes) == 0 {
		user.DitAddress = commitEvent.Who.Hex()
	} else if len(userBytes) == 1 {
		bsonErr := bson.Unmarshal(userBytes[0], &user)
		if bsonErr != nil {
			glog.Error(bsonErr)
			return
		}
	}

	currentDaiBalance, err := connection.BalanceAt(context.Background(), commitEvent.Who, nil)
	if err != nil {
		glog.Error(err)
	} else {
		floatIndividualDaiBalance, _ := (new(big.Float).Quo((new(big.Float).SetInt(currentDaiBalance)), big.NewFloat(1000000000000000000))).Float64()
		user.XDAIBalance = floatIndividualDaiBalance
	}

	var newUserProposal database.ProposalDetailsUser
	newUserProposal.Finalized = false
	newUserProposal.IsProposer = false
	newUserProposal.KNWDifference = 0.0
	newUserProposal.KNWVoteID = proposal.KNWVoteID
	newUserProposal.Opened = false
	newUserProposal.UsedKNW = floatUsedKNW
	newUserProposal.Stake = floatStake
	newUserProposal.VoteDate = time.Unix(int64(block.Time()), 0)
	newUserProposal.VotedRight = false

	user.ProposalsDemo = append(user.ProposalsDemo, newUserProposal)

	index := -1
	for i, singleRepository := range user.RepositoriesDemo {
		if singleRepository.Hash == repository.Hash {
			index = i
			break
		}
	}
	if index == -1 {
		var newRepository database.RepositoryDetailsUser
		newRepository.Hash = repository.Hash
		newRepository.Name = repository.Name
		newRepository.URL = repository.URL
		newRepository.Provider = repository.Provider
		user.RepositoriesDemo = append(user.RepositoriesDemo, newRepository)
		index = len(user.RepositoriesDemo) - 1
	}

	user.RepositoriesDemo[index].LastActivityDate = time.Unix(int64(block.Time()), 0)
	user.RepositoriesDemo[index].AmountOfValidations++

	err = database.Replace("users", "dit_address", commitEvent.Who.Hex(), user)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			err = database.Insert("users", user)
			if err != nil {
				glog.Error(err)
			}
		}
		if err != nil {
			glog.Error(err)
		}
	}
}

func handleDitDemoCoordinatorOpen(event *types.Log, connection *ethclient.Client) {
	glog.Infof("[%d] Open Vote Event", event.BlockNumber)
	var openEvent logDitCoordinatorOpen
	err := ditDemoCoordinatorABI.Unpack(&openEvent, "OpenVote", event.Data)
	if err != nil {
		glog.Error(err)
	}
	openEvent.Repository = event.Topics[1]
	openEvent.Proposal = event.Topics[2].Big()
	openEvent.Who = common.HexToAddress(event.Topics[3].Hex())

	block, err := connection.BlockByNumber(context.Background(), big.NewInt(int64(event.BlockNumber)))
	if err != nil {
		glog.Error(err)
		return
	}

	repositoryBytes, err := database.Get("repositories_demo", []string{"hash"}, []string{"=="}, []interface{}{event.Topics[1].Hex()})
	if err != nil {
		glog.Error(err)
		return
	}

	if len(repositoryBytes) != 1 {
		glog.Errorf("Expected one repository to be found but found %d\n", len(repositoryBytes))
		return
	}

	var repository database.Repository
	bsonErr := bson.Unmarshal(repositoryBytes[0], &repository)
	if bsonErr != nil {
		glog.Error(bsonErr)
		return
	}

	repository.LastActivityDate = time.Unix(int64(block.Time()), 0)

	for i, contributor := range repository.Contributors {
		if contributor.Address == openEvent.Who.Hex() {
			repository.Contributors[i].LastActivityDate = time.Unix(int64(block.Time()), 0)
			break
		}
	}

	err = database.Replace("repositories_demo", "hash", event.Topics[1].Hex(), repository)
	if err != nil {
		glog.Error(err)
		return
	}

	proposalBytes, err := database.Get("proposals_demo", []string{"repository", "id"}, []string{"=="}, []interface{}{event.Topics[1].Hex(), int(openEvent.Proposal.Int64())})
	if err != nil {
		glog.Error(err)
		return
	}

	if len(proposalBytes) != 1 {
		glog.Errorf("Expected one proposal to be found but found %d\n", len(proposalBytes))
		return
	}

	var proposal database.Proposal
	bsonErr = bson.Unmarshal(proposalBytes[0], &proposal)
	if bsonErr != nil {
		glog.Error(bsonErr)
		return
	}

	proposal.Votes.ParticipantsUnrevealed--

	for i, participant := range proposal.Participants {
		if participant.Address == openEvent.Who.Hex() {
			proposal.Votes.VotesUnrevealed -= participant.Votes
			if openEvent.Accept {
				proposal.Votes.ParticipantsFor++
				proposal.Votes.VotesFor += participant.Votes
			} else {
				proposal.Votes.ParticipantsAgainst++
				proposal.Votes.VotesAgainst += participant.Votes
			}
			proposal.Participants[i].Opened = true
			break
		}
	}
	if proposal.Votes.ParticipantsUnrevealed == 0 && proposal.Votes.VotesUnrevealed < 0.00001 {
		proposal.Votes.VotesUnrevealed = 0
	}

	err = database.Replace("proposals_demo", "knw_vote_id", proposal.KNWVoteID, proposal)
	if err != nil {
		glog.Error(err)
		return
	}

	if strings.Contains(os.Getenv("AUTOVALIDATORS"), openEvent.Who.Hex()) {
		return
	}

	userBytes, err := database.Get("users", []string{"dit_address"}, []string{"=="}, []interface{}{openEvent.Who.Hex()})
	if err != nil {
		glog.Error(err)
	}

	if len(userBytes) != 1 {
		glog.Error(fmt.Errorf("Expected one user to be found but found %d", len(userBytes)))
		return
	}

	var user database.User
	bsonErr = bson.Unmarshal(userBytes[0], &user)
	if bsonErr != nil {
		glog.Error(bsonErr)
		return
	}

	currentDaiBalance, err := connection.BalanceAt(context.Background(), openEvent.Who, nil)
	if err != nil {
		glog.Error(err)
	} else {
		floatIndividualDaiBalance, _ := (new(big.Float).Quo((new(big.Float).SetInt(currentDaiBalance)), big.NewFloat(1000000000000000000))).Float64()
		user.XDAIBalance = floatIndividualDaiBalance
	}

	for i, singleProposal := range user.ProposalsDemo {
		if singleProposal.KNWVoteID == proposal.KNWVoteID {
			user.ProposalsDemo[i].Opened = true
			break
		}
	}

	for i, singleRepository := range user.RepositoriesDemo {
		if singleRepository.Hash == repository.Hash {
			user.RepositoriesDemo[i].LastActivityDate = time.Unix(int64(block.Time()), 0)
			break
		}
	}

	err = database.Replace("users", "dit_address", openEvent.Who.Hex(), user)
	if err != nil {
		glog.Error(err)
	}
}

func handleDitDemoCoordinatorFinalizeVote(event *types.Log, connection *ethclient.Client) {
	glog.Infof("[%d] Finalize Vote Event", event.BlockNumber)
	var finalizeVoteEvent logDitCoordinatorFinalizeVote
	err := ditDemoCoordinatorABI.Unpack(&finalizeVoteEvent, "FinalizeVote", event.Data)
	if err != nil {
		glog.Error(err)
		return
	}
	finalizeVoteEvent.Repository = event.Topics[1]
	finalizeVoteEvent.Proposal = event.Topics[2].Big()
	finalizeVoteEvent.Who = common.HexToAddress(event.Topics[3].Hex())
	floatKNWDifference, _ := (new(big.Float).Quo((new(big.Float).SetInt(finalizeVoteEvent.NumberOfKNW)), big.NewFloat(1000000000000000000))).Float64()

	block, err := connection.BlockByNumber(context.Background(), big.NewInt(int64(event.BlockNumber)))
	if err != nil {
		glog.Error(err)
		return
	}

	repositoryBytes, err := database.Get("repositories_demo", []string{"hash"}, []string{"=="}, []interface{}{event.Topics[1].Hex()})
	if err != nil {
		glog.Error(err)
		return
	}

	if len(repositoryBytes) != 1 {
		glog.Errorf("Expected one repository to be found but found %d\n", len(repositoryBytes))
		return
	}

	var repository database.Repository
	bsonErr := bson.Unmarshal(repositoryBytes[0], &repository)
	if bsonErr != nil {
		glog.Error(bsonErr)
		return
	}

	proposalBytes, err := database.Get("proposals_demo", []string{"repository", "id"}, []string{"=="}, []interface{}{event.Topics[1].Hex(), int(finalizeVoteEvent.Proposal.Int64())})
	if err != nil {
		glog.Error(err)
		return
	}

	if len(proposalBytes) != 1 {
		glog.Errorf("Expected one proposal to be found but found %d\n", len(proposalBytes))
		return
	}

	var proposal database.Proposal
	bsonErr = bson.Unmarshal(proposalBytes[0], &proposal)
	if bsonErr != nil {
		glog.Error(bsonErr)
		return
	}

	votedRight := false
	for i, participant := range proposal.Participants {
		if participant.Address == finalizeVoteEvent.Who.Hex() {
			if proposal.Proposer == finalizeVoteEvent.Who.Hex() {
				votedRight = proposal.Accepted
				proposal.Participants[i].VotedRight = votedRight
			} else {
				votedRight = finalizeVoteEvent.VotedRight
				proposal.Participants[i].VotedRight = votedRight
			}
			proposal.Participants[i].Finalized = true

			proposal.Participants[i].KNWDifference = floatKNWDifference
			if votedRight {
				proposal.TotalMintedKNW += floatKNWDifference
			}
			break
		}
	}

	repository.LastActivityDate = time.Unix(int64(block.Time()), 0)

	var currentEarnedKNW float64
	for i, contributor := range repository.Contributors {
		if contributor.Address == finalizeVoteEvent.Who.Hex() {
			repository.Contributors[i].LastActivityDate = time.Unix(int64(block.Time()), 0)
			if votedRight {
				repository.Contributors[i].EarnedKNW += floatKNWDifference
			} else {
				repository.Contributors[i].EarnedKNW -= floatKNWDifference
			}
			currentEarnedKNW = repository.Contributors[i].EarnedKNW
			break
		}
	}

	err = database.Replace("repositories_demo", "hash", event.Topics[1].Hex(), repository)
	if err != nil {
		glog.Error(err)
		return
	}

	err = database.Replace("proposals_demo", "knw_vote_id", proposal.KNWVoteID, proposal)
	if err != nil {
		glog.Error(err)
		return
	}

	if strings.Contains(os.Getenv("AUTOVALIDATORS"), finalizeVoteEvent.Who.Hex()) {
		return
	}

	userBytes, err := database.Get("users", []string{"dit_address"}, []string{"=="}, []interface{}{finalizeVoteEvent.Who.Hex()})
	if err != nil {
		glog.Error(err)
	}

	if len(userBytes) != 1 {
		glog.Error(fmt.Errorf("Expected one user to be found but found %d", len(userBytes)))
		return
	}

	var user database.User
	bsonErr = bson.Unmarshal(userBytes[0], &user)
	if bsonErr != nil {
		glog.Error(bsonErr)
		return
	}

	currentDaiBalance, err := connection.BalanceAt(context.Background(), finalizeVoteEvent.Who, nil)
	if err != nil {
		glog.Error(err)
	} else {
		floatIndividualDaiBalance, _ := (new(big.Float).Quo((new(big.Float).SetInt(currentDaiBalance)), big.NewFloat(1000000000000000000))).Float64()
		user.XDAIBalance = floatIndividualDaiBalance
	}

	for i, singleProposal := range user.ProposalsDemo {
		if singleProposal.KNWVoteID == proposal.KNWVoteID {
			user.ProposalsDemo[i].Finalized = true
			user.ProposalsDemo[i].VotedRight = votedRight
			user.ProposalsDemo[i].KNWDifference = floatKNWDifference
			break
		}
	}

	for i, singleRepository := range user.RepositoriesDemo {
		if singleRepository.Hash == repository.Hash {
			user.RepositoriesDemo[i].LastActivityDate = time.Unix(int64(block.Time()), 0)
			user.RepositoriesDemo[i].EarnedKNW = currentEarnedKNW
			break
		}
	}

	err = database.Replace("users", "dit_address", finalizeVoteEvent.Who.Hex(), user)
	if err != nil {
		glog.Error(err)
	}
}

func handleDitDemoCoordinatorFinalizeProposal(event *types.Log, connection *ethclient.Client) {
	glog.Infof("[%d] Finalize Proposal Event", event.BlockNumber)
	var finalizeProposalEvent logDitCoordinatorFinalizeProposal
	err := ditDemoCoordinatorABI.Unpack(&finalizeProposalEvent, "FinalizeProposal", event.Data)
	if err != nil {
		glog.Error(err)
		return
	}
	finalizeProposalEvent.Repository = event.Topics[1]
	finalizeProposalEvent.Proposal = event.Topics[2].Big()

	block, err := connection.BlockByNumber(context.Background(), big.NewInt(int64(event.BlockNumber)))
	if err != nil {
		glog.Error(err)
	}

	repositoryBytes, err := database.Get("repositories_demo", []string{"hash"}, []string{"=="}, []interface{}{event.Topics[1].Hex()})
	if err != nil {
		glog.Error(err)
		return
	}

	if len(repositoryBytes) != 1 {
		glog.Errorf("Expected one repository to be found but found %d\n", len(repositoryBytes))
		return
	}

	var repository database.Repository
	bsonErr := bson.Unmarshal(repositoryBytes[0], &repository)
	if bsonErr != nil {
		glog.Error(bsonErr)
		return
	}

	proposalBytes, err := database.Get("proposals_demo", []string{"repository", "id"}, []string{"=="}, []interface{}{event.Topics[1].Hex(), int(finalizeProposalEvent.Proposal.Int64())})
	if err != nil {
		glog.Error(err)
		return
	}

	if len(proposalBytes) != 1 {
		glog.Errorf("Expected one proposal to be found but found %d\n", len(proposalBytes))
		return
	}

	var proposal database.Proposal
	bsonErr = bson.Unmarshal(proposalBytes[0], &proposal)
	if bsonErr != nil {
		glog.Error(bsonErr)
		return
	}

	repository.LastActivityDate = time.Unix(int64(block.Time()), 0)

	for i, proposal := range repository.Proposals {
		if proposal.KNWVoteID == proposal.KNWVoteID {
			repository.Proposals[i].VoteDate = time.Unix(int64(block.Time()), 0)
			break
		}
	}

	err = database.Replace("repositories_demo", "hash", event.Topics[1].Hex(), repository)
	if err != nil {
		glog.Error(err)
		return
	}

	proposal.Finalized = true
	proposal.Accepted = finalizeProposalEvent.Accepted

	err = database.Replace("proposals_demo", "knw_vote_id", proposal.KNWVoteID, proposal)
	if err != nil {
		glog.Error(err)
		return
	}
}

func getDitDemoCoordinatorInstance(_connection *ethclient.Client) (*ditDemoCoordinator.DitDemoCoordinator, error) {
	// Convertig the hex-string-formatted address into an address object
	ditDemoCoordinatorAddressObject := common.HexToAddress(os.Getenv("CONTRACT_DIT_COORDINATOR_DEMO_V2"))

	// Create a new instance of the ditDemoCoordinator to access it
	ditDemoCoordinatorInstance, err := ditDemoCoordinator.NewDitDemoCoordinator(ditDemoCoordinatorAddressObject, _connection)
	if err != nil {
		return nil, errors.New("Failed to find ditDemoCoordinator at provided address")
	}

	return ditDemoCoordinatorInstance, nil
}

func getDitTokenInstance(_connection *ethclient.Client) (*ditToken.MintableERC20, error) {
	// Convertig the hex-string-formatted address into an address object
	ditTokenObject := common.HexToAddress(os.Getenv("CONTRACT_DIT_TOKEN_DEMO"))

	// Create a new instance of the ditDemoCoordinator to access it
	ditTokenInstance, err := ditToken.NewMintableERC20(ditTokenObject, _connection)
	if err != nil {
		return nil, errors.New("Failed to find ditTokenContract at provided address")
	}

	return ditTokenInstance, nil
}

func getKNWVotingInstance(_connection *ethclient.Client) (*KNWVoting.KNWVoting, error) {
	knwVotingAddressObject := common.HexToAddress(os.Getenv("CONTRACT_KNW_VOTING_DEMO_V2"))

	// Create a new instance of the KNWVoting contract to access it
	KNWVotingInstance, err := KNWVoting.NewKNWVoting(knwVotingAddressObject, _connection)
	if err != nil {
		return nil, errors.New("Failed to find KNWVoting at provided address")
	}

	return KNWVotingInstance, nil
}

func getKNWTokenInstance(_connection *ethclient.Client) (*KNWToken.KNWToken, error) {
	KNWTokenAddress := common.HexToAddress(os.Getenv("CONTRACT_KNW_TOKEN_DEMO_V2"))

	// Create a new instance of the KNWToken contract to access it
	KNWTokenInstance, err := KNWToken.NewKNWToken(KNWTokenAddress, _connection)
	if err != nil {
		return nil, errors.New("Failed to find KNWToken at provided address")
	}

	return KNWTokenInstance, nil
}

// getConnection will return a connection to the ethereum blockchain
func getConnection() (*ethclient.Client, error) {
	connection, err := ethclient.Dial(os.Getenv("ETHEREUM_RPC"))
	if err != nil {
		if strings.Contains(err.Error(), "bad status") {
			glog.Warning("Received a bad status from node, waiting for five seconds")
			time.Sleep(5 * time.Second)
			connection, err := getConnection()
			return connection, err
		}
		return nil, err
	}
	return connection, nil
}
