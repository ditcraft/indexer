// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package KNWVoting

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// KNWTokenContractABI is the input ABI used to generate the binding from.
const KNWTokenContractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"balanceOfID\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"lockTokens\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"freeBalanceOfID\",\"outputs\":[{\"name\":\"freeBalance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"unlockTokens\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_id\",\"type\":\"uint256\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// KNWTokenContractFuncSigs maps the 4-byte function signature to its string representation.
var KNWTokenContractFuncSigs = map[string]string{
	"3938400b": "balanceOfID(address,uint256)",
	"f5298aca": "burn(address,uint256,uint256)",
	"b5c2cdba": "freeBalanceOfID(address,uint256)",
	"a25983e5": "lockTokens(address,uint256,uint256)",
	"156e29f6": "mint(address,uint256,uint256)",
	"cd3877df": "unlockTokens(address,uint256,uint256)",
}

// KNWTokenContract is an auto generated Go binding around an Ethereum contract.
type KNWTokenContract struct {
	KNWTokenContractCaller     // Read-only binding to the contract
	KNWTokenContractTransactor // Write-only binding to the contract
	KNWTokenContractFilterer   // Log filterer for contract events
}

// KNWTokenContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type KNWTokenContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KNWTokenContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KNWTokenContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KNWTokenContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KNWTokenContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KNWTokenContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KNWTokenContractSession struct {
	Contract     *KNWTokenContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KNWTokenContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KNWTokenContractCallerSession struct {
	Contract *KNWTokenContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// KNWTokenContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KNWTokenContractTransactorSession struct {
	Contract     *KNWTokenContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// KNWTokenContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type KNWTokenContractRaw struct {
	Contract *KNWTokenContract // Generic contract binding to access the raw methods on
}

// KNWTokenContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KNWTokenContractCallerRaw struct {
	Contract *KNWTokenContractCaller // Generic read-only contract binding to access the raw methods on
}

// KNWTokenContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KNWTokenContractTransactorRaw struct {
	Contract *KNWTokenContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKNWTokenContract creates a new instance of KNWTokenContract, bound to a specific deployed contract.
func NewKNWTokenContract(address common.Address, backend bind.ContractBackend) (*KNWTokenContract, error) {
	contract, err := bindKNWTokenContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KNWTokenContract{KNWTokenContractCaller: KNWTokenContractCaller{contract: contract}, KNWTokenContractTransactor: KNWTokenContractTransactor{contract: contract}, KNWTokenContractFilterer: KNWTokenContractFilterer{contract: contract}}, nil
}

// NewKNWTokenContractCaller creates a new read-only instance of KNWTokenContract, bound to a specific deployed contract.
func NewKNWTokenContractCaller(address common.Address, caller bind.ContractCaller) (*KNWTokenContractCaller, error) {
	contract, err := bindKNWTokenContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KNWTokenContractCaller{contract: contract}, nil
}

// NewKNWTokenContractTransactor creates a new write-only instance of KNWTokenContract, bound to a specific deployed contract.
func NewKNWTokenContractTransactor(address common.Address, transactor bind.ContractTransactor) (*KNWTokenContractTransactor, error) {
	contract, err := bindKNWTokenContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KNWTokenContractTransactor{contract: contract}, nil
}

// NewKNWTokenContractFilterer creates a new log filterer instance of KNWTokenContract, bound to a specific deployed contract.
func NewKNWTokenContractFilterer(address common.Address, filterer bind.ContractFilterer) (*KNWTokenContractFilterer, error) {
	contract, err := bindKNWTokenContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KNWTokenContractFilterer{contract: contract}, nil
}

// bindKNWTokenContract binds a generic wrapper to an already deployed contract.
func bindKNWTokenContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KNWTokenContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KNWTokenContract *KNWTokenContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KNWTokenContract.Contract.KNWTokenContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KNWTokenContract *KNWTokenContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KNWTokenContract.Contract.KNWTokenContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KNWTokenContract *KNWTokenContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KNWTokenContract.Contract.KNWTokenContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KNWTokenContract *KNWTokenContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KNWTokenContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KNWTokenContract *KNWTokenContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KNWTokenContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KNWTokenContract *KNWTokenContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KNWTokenContract.Contract.contract.Transact(opts, method, params...)
}

// BalanceOfID is a free data retrieval call binding the contract method 0x3938400b.
//
// Solidity: function balanceOfID(address _address, uint256 _id) constant returns(uint256 balance)
func (_KNWTokenContract *KNWTokenContractCaller) BalanceOfID(opts *bind.CallOpts, _address common.Address, _id *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KNWTokenContract.contract.Call(opts, out, "balanceOfID", _address, _id)
	return *ret0, err
}

// BalanceOfID is a free data retrieval call binding the contract method 0x3938400b.
//
// Solidity: function balanceOfID(address _address, uint256 _id) constant returns(uint256 balance)
func (_KNWTokenContract *KNWTokenContractSession) BalanceOfID(_address common.Address, _id *big.Int) (*big.Int, error) {
	return _KNWTokenContract.Contract.BalanceOfID(&_KNWTokenContract.CallOpts, _address, _id)
}

// BalanceOfID is a free data retrieval call binding the contract method 0x3938400b.
//
// Solidity: function balanceOfID(address _address, uint256 _id) constant returns(uint256 balance)
func (_KNWTokenContract *KNWTokenContractCallerSession) BalanceOfID(_address common.Address, _id *big.Int) (*big.Int, error) {
	return _KNWTokenContract.Contract.BalanceOfID(&_KNWTokenContract.CallOpts, _address, _id)
}

// FreeBalanceOfID is a free data retrieval call binding the contract method 0xb5c2cdba.
//
// Solidity: function freeBalanceOfID(address _address, uint256 _id) constant returns(uint256 freeBalance)
func (_KNWTokenContract *KNWTokenContractCaller) FreeBalanceOfID(opts *bind.CallOpts, _address common.Address, _id *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KNWTokenContract.contract.Call(opts, out, "freeBalanceOfID", _address, _id)
	return *ret0, err
}

// FreeBalanceOfID is a free data retrieval call binding the contract method 0xb5c2cdba.
//
// Solidity: function freeBalanceOfID(address _address, uint256 _id) constant returns(uint256 freeBalance)
func (_KNWTokenContract *KNWTokenContractSession) FreeBalanceOfID(_address common.Address, _id *big.Int) (*big.Int, error) {
	return _KNWTokenContract.Contract.FreeBalanceOfID(&_KNWTokenContract.CallOpts, _address, _id)
}

// FreeBalanceOfID is a free data retrieval call binding the contract method 0xb5c2cdba.
//
// Solidity: function freeBalanceOfID(address _address, uint256 _id) constant returns(uint256 freeBalance)
func (_KNWTokenContract *KNWTokenContractCallerSession) FreeBalanceOfID(_address common.Address, _id *big.Int) (*big.Int, error) {
	return _KNWTokenContract.Contract.FreeBalanceOfID(&_KNWTokenContract.CallOpts, _address, _id)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address _address, uint256 _id, uint256 _amount) returns(bool success)
func (_KNWTokenContract *KNWTokenContractTransactor) Burn(opts *bind.TransactOpts, _address common.Address, _id *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _KNWTokenContract.contract.Transact(opts, "burn", _address, _id, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address _address, uint256 _id, uint256 _amount) returns(bool success)
func (_KNWTokenContract *KNWTokenContractSession) Burn(_address common.Address, _id *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _KNWTokenContract.Contract.Burn(&_KNWTokenContract.TransactOpts, _address, _id, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address _address, uint256 _id, uint256 _amount) returns(bool success)
func (_KNWTokenContract *KNWTokenContractTransactorSession) Burn(_address common.Address, _id *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _KNWTokenContract.Contract.Burn(&_KNWTokenContract.TransactOpts, _address, _id, _amount)
}

// LockTokens is a paid mutator transaction binding the contract method 0xa25983e5.
//
// Solidity: function lockTokens(address _address, uint256 _id, uint256 _amount) returns(bool success)
func (_KNWTokenContract *KNWTokenContractTransactor) LockTokens(opts *bind.TransactOpts, _address common.Address, _id *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _KNWTokenContract.contract.Transact(opts, "lockTokens", _address, _id, _amount)
}

// LockTokens is a paid mutator transaction binding the contract method 0xa25983e5.
//
// Solidity: function lockTokens(address _address, uint256 _id, uint256 _amount) returns(bool success)
func (_KNWTokenContract *KNWTokenContractSession) LockTokens(_address common.Address, _id *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _KNWTokenContract.Contract.LockTokens(&_KNWTokenContract.TransactOpts, _address, _id, _amount)
}

// LockTokens is a paid mutator transaction binding the contract method 0xa25983e5.
//
// Solidity: function lockTokens(address _address, uint256 _id, uint256 _amount) returns(bool success)
func (_KNWTokenContract *KNWTokenContractTransactorSession) LockTokens(_address common.Address, _id *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _KNWTokenContract.Contract.LockTokens(&_KNWTokenContract.TransactOpts, _address, _id, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x156e29f6.
//
// Solidity: function mint(address _address, uint256 _id, uint256 _amount) returns(bool success)
func (_KNWTokenContract *KNWTokenContractTransactor) Mint(opts *bind.TransactOpts, _address common.Address, _id *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _KNWTokenContract.contract.Transact(opts, "mint", _address, _id, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x156e29f6.
//
// Solidity: function mint(address _address, uint256 _id, uint256 _amount) returns(bool success)
func (_KNWTokenContract *KNWTokenContractSession) Mint(_address common.Address, _id *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _KNWTokenContract.Contract.Mint(&_KNWTokenContract.TransactOpts, _address, _id, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x156e29f6.
//
// Solidity: function mint(address _address, uint256 _id, uint256 _amount) returns(bool success)
func (_KNWTokenContract *KNWTokenContractTransactorSession) Mint(_address common.Address, _id *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _KNWTokenContract.Contract.Mint(&_KNWTokenContract.TransactOpts, _address, _id, _amount)
}

// UnlockTokens is a paid mutator transaction binding the contract method 0xcd3877df.
//
// Solidity: function unlockTokens(address _address, uint256 _id, uint256 _amount) returns(bool success)
func (_KNWTokenContract *KNWTokenContractTransactor) UnlockTokens(opts *bind.TransactOpts, _address common.Address, _id *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _KNWTokenContract.contract.Transact(opts, "unlockTokens", _address, _id, _amount)
}

// UnlockTokens is a paid mutator transaction binding the contract method 0xcd3877df.
//
// Solidity: function unlockTokens(address _address, uint256 _id, uint256 _amount) returns(bool success)
func (_KNWTokenContract *KNWTokenContractSession) UnlockTokens(_address common.Address, _id *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _KNWTokenContract.Contract.UnlockTokens(&_KNWTokenContract.TransactOpts, _address, _id, _amount)
}

// UnlockTokens is a paid mutator transaction binding the contract method 0xcd3877df.
//
// Solidity: function unlockTokens(address _address, uint256 _id, uint256 _amount) returns(bool success)
func (_KNWTokenContract *KNWTokenContractTransactorSession) UnlockTokens(_address common.Address, _id *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _KNWTokenContract.Contract.UnlockTokens(&_KNWTokenContract.TransactOpts, _address, _id, _amount)
}

// KNWVotingABI is the input ABI used to generate the binding from.
const KNWVotingABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"nextKNWVoting\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakesOfVote\",\"outputs\":[{\"name\":\"proposersStake\",\"type\":\"uint256\"},{\"name\":\"proposersReward\",\"type\":\"uint256\"},{\"name\":\"returnPool\",\"type\":\"uint256\"},{\"name\":\"rewardPool\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newManager\",\"type\":\"address\"}],\"name\":\"replaceManager\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_voteID\",\"type\":\"uint256\"},{\"name\":\"_voteOption\",\"type\":\"uint256\"},{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"finalizeVote\",\"outputs\":[{\"name\":\"reward\",\"type\":\"uint256\"},{\"name\":\"winningSide\",\"type\":\"bool\"},{\"name\":\"numberOfKNW\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"ditCoordinatorContracts\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"manager\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_voteID\",\"type\":\"uint256\"}],\"name\":\"isPassed\",\"outputs\":[{\"name\":\"passed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentVoteID\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"votes\",\"outputs\":[{\"name\":\"repository\",\"type\":\"bytes32\"},{\"name\":\"knowledgeID\",\"type\":\"uint256\"},{\"name\":\"commitEndDate\",\"type\":\"uint256\"},{\"name\":\"openEndDate\",\"type\":\"uint256\"},{\"name\":\"neededMajority\",\"type\":\"uint256\"},{\"name\":\"winningPercentage\",\"type\":\"uint256\"},{\"name\":\"votesFor\",\"type\":\"uint256\"},{\"name\":\"votesAgainst\",\"type\":\"uint256\"},{\"name\":\"votesUnrevealed\",\"type\":\"uint256\"},{\"name\":\"participantsFor\",\"type\":\"uint256\"},{\"name\":\"participantsAgainst\",\"type\":\"uint256\"},{\"name\":\"participantsUnrevealed\",\"type\":\"uint256\"},{\"name\":\"isResolved\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_voteID\",\"type\":\"uint256\"}],\"name\":\"didOpen\",\"outputs\":[{\"name\":\"revealed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_voteID\",\"type\":\"uint256\"}],\"name\":\"voteExists\",\"outputs\":[{\"name\":\"exists\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_voteID\",\"type\":\"uint256\"}],\"name\":\"didCommit\",\"outputs\":[{\"name\":\"committed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newRepository\",\"type\":\"bytes32\"},{\"name\":\"_majority\",\"type\":\"uint256\"}],\"name\":\"addNewRepository\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"startingVoteID\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_voteID\",\"type\":\"uint256\"}],\"name\":\"endVote\",\"outputs\":[{\"name\":\"votePassed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_voteID\",\"type\":\"uint256\"}],\"name\":\"getUsedKNW\",\"outputs\":[{\"name\":\"usedKNW\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"KNWTokenAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newDitCoordinatorAddress\",\"type\":\"address\"}],\"name\":\"addDitCoordinator\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_voteID\",\"type\":\"uint256\"}],\"name\":\"commitPeriodActive\",\"outputs\":[{\"name\":\"active\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_voteID\",\"type\":\"uint256\"}],\"name\":\"voteEnded\",\"outputs\":[{\"name\":\"ended\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_voteID\",\"type\":\"uint256\"}],\"name\":\"getGrossStake\",\"outputs\":[{\"name\":\"grossStake\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BURNING_METHOD\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_voteID\",\"type\":\"uint256\"}],\"name\":\"isResolved\",\"outputs\":[{\"name\":\"resolved\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_voteID\",\"type\":\"uint256\"}],\"name\":\"getAmountOfVotes\",\"outputs\":[{\"name\":\"numberOfVotes\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_voteID\",\"type\":\"uint256\"},{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_voteOption\",\"type\":\"uint256\"},{\"name\":\"_salt\",\"type\":\"uint256\"}],\"name\":\"openVote\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_voteID\",\"type\":\"uint256\"},{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_secretHash\",\"type\":\"bytes32\"},{\"name\":\"_numberOfKNW\",\"type\":\"uint256\"}],\"name\":\"commitVote\",\"outputs\":[{\"name\":\"numberOfVotes\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_terminationDate\",\"type\":\"uint256\"}],\"name\":\"isExpired\",\"outputs\":[{\"name\":\"expired\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_voteID\",\"type\":\"uint256\"}],\"name\":\"getNetStake\",\"outputs\":[{\"name\":\"netStake\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_repository\",\"type\":\"bytes32\"},{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_knowledgeID\",\"type\":\"uint256\"},{\"name\":\"_voteDuration\",\"type\":\"uint256\"},{\"name\":\"_proposersStake\",\"type\":\"uint256\"},{\"name\":\"_numberOfKNW\",\"type\":\"uint256\"}],\"name\":\"startVote\",\"outputs\":[{\"name\":\"voteID\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_voteID\",\"type\":\"uint256\"}],\"name\":\"openPeriodActive\",\"outputs\":[{\"name\":\"active\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"upgradeContract\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MINTING_METHOD\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastKNWVoting\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_KNWTokenAddress\",\"type\":\"address\"},{\"name\":\"_lastKNWVoting\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// KNWVotingFuncSigs maps the 4-byte function signature to its string representation.
var KNWVotingFuncSigs = map[string]string{
	"c814af1f": "BURNING_METHOD()",
	"985dbfc5": "KNWTokenAddress()",
	"effb21e1": "MINTING_METHOD()",
	"a326690a": "addDitCoordinator(address)",
	"828c1653": "addNewRepository(bytes32,uint256)",
	"a4439dc5": "commitPeriodActive(uint256)",
	"d4e0ac95": "commitVote(uint256,address,bytes32,uint256)",
	"5a981129": "currentVoteID()",
	"7f97e836": "didCommit(address,uint256)",
	"70d18152": "didOpen(address,uint256)",
	"3fdf3dab": "ditCoordinatorContracts(address)",
	"865df0ad": "endVote(uint256)",
	"36bf4c91": "finalizeVote(uint256,uint256,address)",
	"cd6ef46b": "getAmountOfVotes(address,uint256)",
	"c3eb9708": "getGrossStake(uint256)",
	"dcfde092": "getNetStake(uint256)",
	"8e1b660f": "getUsedKNW(address,uint256)",
	"d9548e53": "isExpired(uint256)",
	"49403183": "isPassed(uint256)",
	"cc4e1954": "isResolved(uint256)",
	"ff77f4b8": "lastKNWVoting()",
	"481c6a75": "manager()",
	"0bf682e6": "nextKNWVoting()",
	"ea521a52": "openPeriodActive(uint256)",
	"cdd6ceb9": "openVote(uint256,address,uint256,uint256)",
	"23447982": "replaceManager(address)",
	"1f84950f": "stakesOfVote(uint256)",
	"e5023ff2": "startVote(bytes32,address,uint256,uint256,uint256,uint256)",
	"82ec45a3": "startingVoteID()",
	"eb2c0223": "upgradeContract(address)",
	"b43a401d": "voteEnded(uint256)",
	"7e4173a8": "voteExists(uint256)",
	"5df81330": "votes(uint256)",
}

// KNWVotingBin is the compiled bytecode used for deploying new contracts.
var KNWVotingBin = "0x60806040523480156200001157600080fd5b5060405162002ddf38038062002ddf833981810160405260408110156200003757600080fd5b5080516020909101516001600160a01b038216620000b657604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f4b4e57546f6b656e20616464726573732063616e277420626520656d70747900604482015290519081900360640190fd5b600580546001600160a01b038085166001600160a01b03199283161792839055600680549092169281169290921790558116156200019d57600380546001600160a01b0319166001600160a01b038381169190911791829055604080517f5a981129000000000000000000000000000000000000000000000000000000008152905192909116918291635a981129916004808301926020929190829003018186803b1580156200016557600080fd5b505afa1580156200017a573d6000803e3d6000fd5b505050506040513d60208110156200019157600080fd5b505160075550620001a3565b60006007555b5050600754600855600280546001600160a01b03191633179055612c1280620001cd6000396000f3fe608060405234801561001057600080fd5b50600436106101fb5760003560e01c8063985dbfc51161011a578063cdd6ceb9116100ad578063e5023ff21161007c578063e5023ff214610643578063ea521a5214610687578063eb2c0223146106a4578063effb21e114610548578063ff77f4b8146106ca576101fb565b8063cdd6ceb914610599578063d4e0ac95146105d1578063d9548e5314610609578063dcfde09214610626576101fb565b8063c3eb9708116100e9578063c3eb97081461052b578063c814af1f14610548578063cc4e195414610550578063cd6ef46b1461056d576101fb565b8063985dbfc5146104c3578063a326690a146104cb578063a4439dc5146104f1578063b43a401d1461050e576101fb565b80635df8133011610192578063828c165311610161578063828c16531461044f57806382ec45a314610472578063865df0ad1461047a5780638e1b660f14610497576101fb565b80635df813301461035657806370d18152146103da5780637e4173a8146104065780637f97e83614610423576101fb565b80633fdf3dab116101ce5780633fdf3dab146102f1578063481c6a7514610317578063494031831461031f5780635a9811291461033c576101fb565b80630bf682e6146102005780631f84950f14610224578063234479821461026757806336bf4c91146102a1575b600080fd5b6102086106d2565b604080516001600160a01b039092168252519081900360200190f35b6102416004803603602081101561023a57600080fd5b50356106e1565b604080519485526020850193909352838301919091526060830152519081900360800190f35b61028d6004803603602081101561027d57600080fd5b50356001600160a01b0316610708565b604080519115158252519081900360200190f35b6102d3600480360360608110156102b757600080fd5b50803590602081013590604001356001600160a01b0316610759565b60408051938452911515602084015282820152519081900360600190f35b61028d6004803603602081101561030757600080fd5b50356001600160a01b0316610b83565b610208610b98565b61028d6004803603602081101561033557600080fd5b5035610ba7565b610344610cc2565b60408051918252519081900360200190f35b6103736004803603602081101561036c57600080fd5b5035610cc8565b604080519d8e5260208e019c909c528c8c019a909a5260608c019890985260808b019690965260a08a019490945260c089019290925260e0880152610100870152610120860152610140850152610160840152151561018083015251908190036101a00190f35b61028d600480360360408110156103f057600080fd5b506001600160a01b038135169060200135610d33565b61028d6004803603602081101561041c57600080fd5b5035610db9565b61028d6004803603604081101561043957600080fd5b506001600160a01b038135169060200135610dce565b61028d6004803603604081101561046557600080fd5b5080359060200135610e4d565b610344610eb5565b61028d6004803603602081101561049057600080fd5b5035610ebb565b610344600480360360408110156104ad57600080fd5b506001600160a01b03813516906020013561133d565b61020861136b565b61028d600480360360208110156104e157600080fd5b50356001600160a01b031661137a565b61028d6004803603602081101561050757600080fd5b503561144a565b61028d6004803603602081101561052457600080fd5b50356114bc565b6103446004803603602081101561054157600080fd5b5035611527565b610344611539565b61028d6004803603602081101561056657600080fd5b503561153e565b6103446004803603604081101561058357600080fd5b506001600160a01b038135169060200135611556565b61028d600480360360808110156105af57600080fd5b508035906001600160a01b036020820135169060408101359060600135611584565b610344600480360360808110156105e757600080fd5b508035906001600160a01b03602082013516906040810135906060013561191d565b61028d6004803603602081101561061f57600080fd5b5035611ec0565b6103446004803603602081101561063c57600080fd5b5035611ec5565b610344600480360360c081101561065957600080fd5b508035906001600160a01b036020820135169060408101359060608101359060808101359060a00135611f3e565b61028d6004803603602081101561069d57600080fd5b50356123e3565b61028d600480360360208110156106ba57600080fd5b50356001600160a01b031661245f565b6102086124b1565b6004546001600160a01b031681565b600a6020526000908152604090208054600182015460028301546003909301549192909184565b6002546000906001600160a01b0316331461072257600080fd5b6001600160a01b03821661073557600080fd5b50600280546001600160a01b0319166001600160a01b03831617905560015b919050565b336000818152602081905260408120549091829182919060ff166107ae5760405162461bcd60e51b815260040180806020018281038252602b815260200180612a9a602b913960400191505060405180910390fd5b6000878152600960205260409020600c81015460ff16610815576040805162461bcd60e51b815260206004820152601760248201527f506f6c6c2068617320746f206265207265736f6c766564000000000000000000604482015290519081900360640190fd5b6001600160a01b0386166000908152600d82016020526040902060010154156108e7576006546001828101546001600160a01b038981166000818152600d87016020908152604080832090960154865163cd3877df60e01b81526004810194909452602484019590955260448301949094529351919094169363cd3877df9360648083019493928390030190829087803b1580156108b257600080fd5b505af11580156108c6573d6000803e3d6000fd5b505050506040513d60208110156108dc57600080fd5b50516108e757600080fd5b60006108f289610ba7565b9050600081610902576000610905565b60015b6001600160a01b0389166000908152600d8501602052604090205460ff9182168b14925062010000900416156109e957811561097757600183015460058401546001600160a01b038a166000908152600d86016020526040902060020154610970928b9290916124c0565b94506109ce565b60008a8152600a60205260409020600101546109ce576109cb88846001015485600d0160008c6001600160a01b03166001600160a01b031681526020019081526020016000206001015486600501546125a5565b94505b506000898152600a6020526040902060010154955080610b75565b6109f3888b610d33565b15610abb5781158015610a0d575082600701548360060154145b15610a2557610a1e8a8260016126f2565b9650610ab6565b610a318a8260006126f2565b96508015610a7557600183015460058401546001600160a01b038a166000908152600d86016020526040902060020154610a6e928b9290916124c0565b9450610ab6565b610ab388846001015485600d0160008c6001600160a01b03166001600160a01b031681526020019081526020016000206001015486600501546125a5565b94505b610b75565b610ac5888b610d33565b158015610ad75750610ad7888b610dce565b15610b2857610ae88a6000806126f2565b9650610ab388846001015485600d0160008c6001600160a01b03166001600160a01b031681526020019081526020016000206001015486600501546125a5565b6040805162461bcd60e51b815260206004820152601d60248201527f4e6f742061207061727469636970616e74206f662074686520766f7465000000604482015290519081900360640190fd5b945050505093509350939050565b60006020819052908152604090205460ff1681565b6002546001600160a01b031681565b6000610bb2826114bc565b610bfc576040805162461bcd60e51b8152602060048201526016602482015275141bdb1b081a185cc81d1bc81a185d9948195b99195960521b604482015290519081900360640190fd5b610c04612a05565b505060009081526009602081815260409283902083516101a0810185528154815260018201549281019290925260028101549382019390935260038301546060820152600483015460808201819052600584015460a0830152600684015460c08301819052600785015460e08401819052600886015461010085015293850154610120840152600a850154610140840152600b850154610160840152600c9094015460ff161515610180909201919091529082010260649091021190565b60075481565b600960208190526000918252604090912080546001820154600283015460038401546004850154600586015460068701546007880154600889015499890154600a8a0154600b8b0154600c909b0154999b989a97999698959794969395929492939192909160ff168d565b6000610d3e82610db9565b610d83576040805162461bcd60e51b8152602060048201526011602482015270141bdb1b081a185cc81d1bc8195e1a5cdd607a1b604482015290519081900360640190fd5b5060008181526009602090815260408083206001600160a01b0386168452600d01909152902054610100900460ff165b92915050565b60008115801590610db3575050600754101590565b6000610dd982610db9565b610e1e576040805162461bcd60e51b8152602060048201526011602482015270141bdb1b081a185cc81d1bc8195e1a5cdd607a1b604482015290519081900360640190fd5b5060009081526009602090815260408083206001600160a01b03949094168352600d9093019052205460ff1690565b3360008181526020819052604081205490919060ff16610e9e5760405162461bcd60e51b815260040180806020018281038252602b815260200180612a9a602b913960400191505060405180910390fd5b505060009182526001602081905260409092205590565b60085481565b3360008181526020819052604081205490919060ff16610f0c5760405162461bcd60e51b815260040180806020018281038252602b815260200180612a9a602b913960400191505060405180910390fd5b610f15836114bc565b610f5f576040805162461bcd60e51b8152602060048201526016602482015275141bdb1b081a185cc81d1bc81a185d9948195b99195960521b604482015290519081900360640190fd5b60008381526009602052604081206007810154600690910154610f879163ffffffff6128a316565b60008581526009602081905260408220600b81015491810154600a909101549394509192610fcb92610fbf919063ffffffff6128a316565b9063ffffffff6128a316565b9050806110105750506000838152600a6020908152604080832080546001918201556009909252822060058101839055600c01805460ff191690911790559150611337565b6000858152600a602052604090205461105f9061105290611037908463ffffffff6128ee16565b6000888152600a60205260409020549063ffffffff61294f16565b829063ffffffff61299616565b6000868152600a602052604081206002019190915561107d86610ba7565b94508415611107576000868152600960205260409020600601546110ba9084906110ae90606463ffffffff61299616565b9063ffffffff6128ee16565b60008781526009602052604090206005810191909155600b810154600a909101546110ea9163ffffffff6128a316565b6000878152600a60205260409020805460019091015590506111c9565b600086815260096020526040902060078101546006909101541461119b5760008681526009602052604090206007015461114e9084906110ae90606463ffffffff61299616565b60008781526009602081905260409091206005810192909255600b82015491015461117e9163ffffffff6128a316565b6000878152600a60205260409020805460039091015590506111c9565b50600085815260096020908152604080832060326005820155600b0154600a90925290912080546001909101555b6000868152600a602052604090206002015415611267576000868152600a602052604090206002015461125490611236906112299061120e908663ffffffff6128ee16565b60008a8152600a60205260409020549063ffffffff61294f16565b839063ffffffff61299616565b6000888152600a60205260409020600301549063ffffffff6128a316565b6000878152600a60205260409020600301555b6000868152600960209081526040808320600c01805460ff19166001908117909155600a9092529091200154156113335760006112d06112b26001610fbf868663ffffffff61294f16565b6000898152600a60205260409020600301549063ffffffff6128ee16565b6000888152600a60205260409020600101549091506112f5908263ffffffff6128a316565b6000888152600a6020526040902060018101919091556003015461131f908263ffffffff61294f16565b6000888152600a6020526040902060030155505b5050505b50919050565b60009081526009602090815260408083206001600160a01b03949094168352600d9093019052206001015490565b6005546001600160a01b031681565b6002546000906001600160a01b031633146113dc576040805162461bcd60e51b815260206004820152601e60248201527f4f6e6c7920746865206d616e616765722063616e2063616c6c20746869730000604482015290519081900360640190fd5b6001600160a01b0382166114215760405162461bcd60e51b815260040180806020018281038252603a815260200180612b44603a913960400191505060405180910390fd5b506001600160a01b03166000908152602081905260409020805460ff1916600190811790915590565b600061145582610db9565b61149a576040805162461bcd60e51b8152602060048201526011602482015270141bdb1b081a185cc81d1bc8195e1a5cdd607a1b604482015290519081900360640190fd5b6000828152600960205260409020600201546114b590611ec0565b1592915050565b60006114c782610db9565b61150c576040805162461bcd60e51b8152602060048201526011602482015270141bdb1b081a185cc81d1bc8195e1a5cdd607a1b604482015290519081900360640190fd5b600082815260096020526040902060030154610db390611ec0565b6000908152600a602052604090205490565b600081565b6000908152600960205260409020600c015460ff1690565b60009081526009602090815260408083206001600160a01b03949094168352600d9093019052206003015490565b3360008181526020819052604081205490919060ff166115d55760405162461bcd60e51b815260040180806020018281038252602b815260200180612a9a602b913960400191505060405180910390fd5b6115de866123e3565b61162f576040805162461bcd60e51b815260206004820152601e60248201527f52657665616c20706572696f642068617320746f206265206163746976650000604482015290519081900360640190fd5b60008681526009602090815260408083206001600160a01b0389168452600d0190915290205460ff166116935760405162461bcd60e51b8152600401808060200182810382526027815260200180612a736027913960400191505060405180910390fd5b60008681526009602090815260408083206001600160a01b0389168452600d01909152902054610100900460ff16156116fd5760405162461bcd60e51b8152600401808060200182810382526022815260200180612ac56022913960400191505060405180910390fd5b60008681526009602090815260408083206001600160a01b0389168452600d018252918290206004015482518083018890528084018790528351808203850181526060909101909352825192909101919091201461178c5760405162461bcd60e51b8152600401808060200182810382526036815260200180612ba86036913960400191505060405180910390fd5b60008681526009602081815260408084206001600160a01b038a168552600d8101835290842060030154938a9052919052600801546117d1908263ffffffff61294f16565b60008881526009602052604090206008810191909155600b01546117fc90600163ffffffff61294f16565b6000888152600960205260409020600b0155600185141561187d57600087815260096020526040902060060154611839908263ffffffff6128a316565b60008881526009602081905260409091206006810192909255015461186590600163ffffffff6128a316565b600088815260096020819052604090912001556118dd565b60008781526009602052604090206007015461189f908263ffffffff6128a316565b60008881526009602052604090206007810191909155600a01546118ca90600163ffffffff6128a316565b6000888152600960205260409020600a01555b5060008681526009602090815260408083206001600160a01b0389168452600d019091529020805461ff0019166101001790556001915050949350505050565b3360008181526020819052604081205490919060ff1661196e5760405162461bcd60e51b815260040180806020018281038252602b815260200180612a9a602b913960400191505060405180910390fd5b856119b7576040805162461bcd60e51b8152602060048201526014602482015273766f746549442063616e2774206265207a65726f60601b604482015290519081900360640190fd5b6119c08661144a565b611a11576040805162461bcd60e51b815260206004820152601e60248201527f436f6d6d697420706572696f642068617320746f206265206163746976650000604482015290519081900360640190fd5b611a1b8587610dce565b15611a6d576040805162461bcd60e51b815260206004820152601f60248201527f43616e277420636f6d6d6974206d6f7265207468616e206f6e6520766f746500604482015290519081900360640190fd5b83611abf576040805162461bcd60e51b815260206004820152601b60248201527f43616e277420766f746520776974682061207a65726f20686173680000000000604482015290519081900360640190fd5b6000868152600a6020908152604080832054600654600754855260098452828520600101548351635ae166dd60e11b81526001600160a01b038c8116600483015260248201929092529351929750169263b5c2cdba9260448082019391829003018186803b158015611b3057600080fd5b505afa158015611b44573d6000803e3d6000fd5b505050506040513d6020811015611b5a57600080fd5b5051600654600089815260096020908152604080832060010154815163a25983e560e01b81526001600160a01b038d811660048301526024820192909252604481018b905291519596509093169363a25983e593606480820194918390030190829087803b158015611bcb57600080fd5b505af1158015611bdf573d6000803e3d6000fd5b505050506040513d6020811015611bf557600080fd5b5051611c0057600080fd5b60008781526009602090815260408083206001600160a01b038a168452600d01909152902060010184905580841415611c655760008781526009602090815260408083206001600160a01b038a168452600d0190915290206064600290910155611caa565b8015611caa57611c80816110ae86606463ffffffff61299616565b60008881526009602090815260408083206001600160a01b038b168452600d019091529020600201555b60075460009081526009602090815260408083206001600160a01b038a168452600d0190915290206002015460011115611d2b576040805162461bcd60e51b815260206004820181905260248201527f43616e277420766f746520776974686f7574207573696e6720616e79204b4e57604482015290519081900360640190fd5b6000611d428564e8d4a5100063ffffffff6128ee16565b73__$fb7607f2ea6af1e84a77318294713ee952$__63677342ce90916040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b158015611d9257600080fd5b505af4158015611da6573d6000803e3d6000fd5b505050506040513d6020811015611dbc57600080fd5b50519050611dd78566038d7ea4c6800063ffffffff6128ee16565b8110611df757611df48566038d7ea4c6800063ffffffff6128ee16565b90505b611e1d611e106103e86110ae848863ffffffff61299616565b859063ffffffff6128a316565b60008981526009602081815260408084206001600160a01b038d168552600d8101835290842060038101869055600481018c9055805460ff19166001179055928c90525260080154909450611e78908563ffffffff6128a316565b60008981526009602052604090206008810191909155600b0154611ea390600163ffffffff6128a316565b6000898152600960205260409020600b0155505050949350505050565b421190565b60008181526009602081905260408220600b81015491810154600a909101548392611efb929091610fbf9163ffffffff6128a316565b90508015611f2a576000838152600a6020526040902054611f22908263ffffffff6128ee16565b915050610754565b50506000908152600a602052604090205490565b3360008181526020819052604081205490919060ff16611f8f5760405162461bcd60e51b815260040180806020018281038252602b815260200180612a9a602b913960400191505060405180910390fd5b600754611fa390600163ffffffff6128a316565b600755604080516101a08101825289815260208101889052908101611fce428863ffffffff6128a316565b8152602001611fe787610fbf428263ffffffff6128a316565b8152602001600160008b81526020019081526020016000206000015481526020016000815260200160008152602001600081526020016000815260200160008152602001600081526020016000815260200160001515815250600960006007548152602001908152602001600020600082015181600001556020820151816001015560408201518160020155606082015181600301556080820151816004015560a0820151816005015560c0820151816006015560e082015181600701556101008201518160080155610120820151816009015561014082015181600a015561016082015181600b015561018082015181600c0160006101000a81548160ff021916908315150217905550905050604051806080016040528085815260200160008152602001600081526020016000815250600a60006007548152602001908152602001600020600082015181600001556020820151816001015560408201518160020155606082015181600301559050506000600660009054906101000a90046001600160a01b03166001600160a01b031663b5c2cdba89896040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b031681526020018281526020019250505060206040518083038186803b1580156121ce57600080fd5b505afa1580156121e2573d6000803e3d6000fd5b505050506040513d60208110156121f857600080fd5b50516006546040805163a25983e560e01b81526001600160a01b038c81166004830152602482018c905260448201899052915193945091169163a25983e5916064808201926020929091908290030181600087803b15801561225957600080fd5b505af115801561226d573d6000803e3d6000fd5b505050506040513d602081101561228357600080fd5b505161228e57600080fd5b6007805460009081526009602081815260408084206001600160a01b038e16808652600d91820184528286206001018b90559554855292825280842094845293909101905220805462ff0000191662010000179055808414156123205760075460009081526009602090815260408083206001600160a01b038c168452600d0190915290206064600290910155612368565b80156123685761233b816110ae86606463ffffffff61299616565b60075460009081526009602090815260408083206001600160a01b038d168452600d019091529020600201555b60075460009081526009602090815260408083206001600160a01b038c168452600d01909152902060020154600111156123d35760405162461bcd60e51b8152600401808060200182810382526028815260200180612ae76028913960400191505060405180910390fd5b5050600754979650505050505050565b60006123ee82610db9565b612433576040805162461bcd60e51b8152602060048201526011602482015270141bdb1b081a185cc81d1bc8195e1a5cdd607a1b604482015290519081900360640190fd5b60008281526009602052604090206003015461244e90611ec0565b158015610db357506114b58261144a565b6002546000906001600160a01b0316331461247957600080fd5b6001600160a01b03821661248c57600080fd5b50600480546001600160a01b0383166001600160a01b03199091161790556001919050565b6003546001600160a01b031681565b6000806124eb66470de4df8200006124df86603263ffffffff61294f16565b9063ffffffff61299616565b9050801561259c5761250860646110ae838663ffffffff61299616565b60065460408051630ab714fb60e11b81526001600160a01b038a81166004830152602482018a905260448201859052915193945091169163156e29f6916064808201926020929091908290030181600087803b15801561256757600080fd5b505af115801561257b573d6000803e3d6000fd5b505050506040513d602081101561259157600080fd5b505161259c57600080fd5b95945050505050565b600080831561268f57600061265866038d7ea4c680006125d08764e8d4a5100063ffffffff6128ee16565b73__$fb7607f2ea6af1e84a77318294713ee952$__63677342ce90916040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b15801561262057600080fd5b505af4158015612634573d6000803e3d6000fd5b505050506040513d602081101561264a57600080fd5b50519063ffffffff61299616565b90508481101561267957612672858263ffffffff61294f16565b915061268d565b61268a85600263ffffffff6128ee16565b91505b505b801561259c5760065460408051637a94c56560e11b81526001600160a01b03898116600483015260248201899052604482018590529151919092169163f5298aca9160648083019260209291908290030181600087803b15801561256757600080fd5b60008381526009602081905260408220600b81015491810154600a909101548392612728929091610fbf9163ffffffff6128a316565b90506000836127e4576000868152600a6020526040812060020154612753908463ffffffff6128ee16565b905080915085156127de57600061276988610ba7565b1561278757506000878152600960208190526040909120015461279b565b506000878152600960205260409020600a01545b6000888152600a60205260409020600301546127da906127c1908363ffffffff6128ee16565b610fbf6127cd8b611ec5565b869063ffffffff6128a316565b9250505b5061259c565b506000858152600a60209081526040808320546009909252909120600b01541580159061282157506000868152600960205260409020600b015482115b1561259c576000868152600960205260409020600b01546128999061287e9061285190859063ffffffff61294f16565b6000898152600a60209081526040808320546009909252909120600b01546110ae9163ffffffff61299616565b6000888152600a60205260409020549063ffffffff6128a316565b9695505050505050565b6000828201838110156128e75760405162461bcd60e51b815260040180806020018281038252602a815260200180612b7e602a913960400191505060405180910390fd5b9392505050565b600080821161293b576040805162461bcd60e51b815260206004820152601460248201527343616e277420646976696465206279207a65726f60601b604482015290519081900360640190fd5b600082848161294657fe5b04949350505050565b6000828211156129905760405162461bcd60e51b8152600401808060200182810382526035815260200180612b0f6035913960400191505060405180910390fd5b50900390565b6000826129a557506000610db3565b828202828482816129b257fe5b04146128e7576040805162461bcd60e51b815260206004820152601f60248201527f466c6177656420696e70757420666f72206d756c7469706c69636174696f6e00604482015290519081900360640190fd5b604051806101a00160405280600080191681526020016000815260200160008152602001600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081526020016000815260200160008152602001600015158152509056fe5061727469636970616e742068617320746f2068617665206120766f746520636f6d6d697465644f6e6c79206120646974436f6f7264696e61746f7220697320616c6c6f7720746f2063616c6c207468697343616e27742072657665616c206120766f7465206d6f7265207468616e206f6e636543616e2774207374617274206120766f746520776974686f7574207573696e6720616e79204b4e5743616e27742073756274726163742061206e756d6265722066726f6d206120736d616c6c6572206f6e6520776974682075696e7473646974436f6f7264696e61746f7220616464726573732063616e206f6e6c792062652061646465642069662069742773206e6f7420656d707479526573756c742068617320746f20626520626967676572207468616e20626f74682073756d6d616e647343686f69636520616e642053616c74206861766520746f206265207468652073616d6520617320696e2074686520766f746568617368a265627a7a7230582017fa6514c5c71639ec3610506f925a663bbe2108a4a12420a3a5990b65f066d964736f6c634300050a0032"

// DeployKNWVoting deploys a new Ethereum contract, binding an instance of KNWVoting to it.
func DeployKNWVoting(auth *bind.TransactOpts, backend bind.ContractBackend, _KNWTokenAddress common.Address, _lastKNWVoting common.Address) (common.Address, *types.Transaction, *KNWVoting, error) {
	parsed, err := abi.JSON(strings.NewReader(KNWVotingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	safeMathAddr, _, _, _ := DeploySafeMath(auth, backend)
	KNWVotingBin = strings.Replace(KNWVotingBin, "__$fb7607f2ea6af1e84a77318294713ee952$__", safeMathAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KNWVotingBin), backend, _KNWTokenAddress, _lastKNWVoting)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KNWVoting{KNWVotingCaller: KNWVotingCaller{contract: contract}, KNWVotingTransactor: KNWVotingTransactor{contract: contract}, KNWVotingFilterer: KNWVotingFilterer{contract: contract}}, nil
}

// KNWVoting is an auto generated Go binding around an Ethereum contract.
type KNWVoting struct {
	KNWVotingCaller     // Read-only binding to the contract
	KNWVotingTransactor // Write-only binding to the contract
	KNWVotingFilterer   // Log filterer for contract events
}

// KNWVotingCaller is an auto generated read-only Go binding around an Ethereum contract.
type KNWVotingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KNWVotingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KNWVotingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KNWVotingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KNWVotingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KNWVotingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KNWVotingSession struct {
	Contract     *KNWVoting        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KNWVotingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KNWVotingCallerSession struct {
	Contract *KNWVotingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// KNWVotingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KNWVotingTransactorSession struct {
	Contract     *KNWVotingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// KNWVotingRaw is an auto generated low-level Go binding around an Ethereum contract.
type KNWVotingRaw struct {
	Contract *KNWVoting // Generic contract binding to access the raw methods on
}

// KNWVotingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KNWVotingCallerRaw struct {
	Contract *KNWVotingCaller // Generic read-only contract binding to access the raw methods on
}

// KNWVotingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KNWVotingTransactorRaw struct {
	Contract *KNWVotingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKNWVoting creates a new instance of KNWVoting, bound to a specific deployed contract.
func NewKNWVoting(address common.Address, backend bind.ContractBackend) (*KNWVoting, error) {
	contract, err := bindKNWVoting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KNWVoting{KNWVotingCaller: KNWVotingCaller{contract: contract}, KNWVotingTransactor: KNWVotingTransactor{contract: contract}, KNWVotingFilterer: KNWVotingFilterer{contract: contract}}, nil
}

// NewKNWVotingCaller creates a new read-only instance of KNWVoting, bound to a specific deployed contract.
func NewKNWVotingCaller(address common.Address, caller bind.ContractCaller) (*KNWVotingCaller, error) {
	contract, err := bindKNWVoting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KNWVotingCaller{contract: contract}, nil
}

// NewKNWVotingTransactor creates a new write-only instance of KNWVoting, bound to a specific deployed contract.
func NewKNWVotingTransactor(address common.Address, transactor bind.ContractTransactor) (*KNWVotingTransactor, error) {
	contract, err := bindKNWVoting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KNWVotingTransactor{contract: contract}, nil
}

// NewKNWVotingFilterer creates a new log filterer instance of KNWVoting, bound to a specific deployed contract.
func NewKNWVotingFilterer(address common.Address, filterer bind.ContractFilterer) (*KNWVotingFilterer, error) {
	contract, err := bindKNWVoting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KNWVotingFilterer{contract: contract}, nil
}

// bindKNWVoting binds a generic wrapper to an already deployed contract.
func bindKNWVoting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KNWVotingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KNWVoting *KNWVotingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KNWVoting.Contract.KNWVotingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KNWVoting *KNWVotingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KNWVoting.Contract.KNWVotingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KNWVoting *KNWVotingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KNWVoting.Contract.KNWVotingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KNWVoting *KNWVotingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KNWVoting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KNWVoting *KNWVotingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KNWVoting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KNWVoting *KNWVotingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KNWVoting.Contract.contract.Transact(opts, method, params...)
}

// BURNINGMETHOD is a free data retrieval call binding the contract method 0xc814af1f.
//
// Solidity: function BURNING_METHOD() constant returns(uint256)
func (_KNWVoting *KNWVotingCaller) BURNINGMETHOD(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "BURNING_METHOD")
	return *ret0, err
}

// BURNINGMETHOD is a free data retrieval call binding the contract method 0xc814af1f.
//
// Solidity: function BURNING_METHOD() constant returns(uint256)
func (_KNWVoting *KNWVotingSession) BURNINGMETHOD() (*big.Int, error) {
	return _KNWVoting.Contract.BURNINGMETHOD(&_KNWVoting.CallOpts)
}

// BURNINGMETHOD is a free data retrieval call binding the contract method 0xc814af1f.
//
// Solidity: function BURNING_METHOD() constant returns(uint256)
func (_KNWVoting *KNWVotingCallerSession) BURNINGMETHOD() (*big.Int, error) {
	return _KNWVoting.Contract.BURNINGMETHOD(&_KNWVoting.CallOpts)
}

// KNWTokenAddress is a free data retrieval call binding the contract method 0x985dbfc5.
//
// Solidity: function KNWTokenAddress() constant returns(address)
func (_KNWVoting *KNWVotingCaller) KNWTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "KNWTokenAddress")
	return *ret0, err
}

// KNWTokenAddress is a free data retrieval call binding the contract method 0x985dbfc5.
//
// Solidity: function KNWTokenAddress() constant returns(address)
func (_KNWVoting *KNWVotingSession) KNWTokenAddress() (common.Address, error) {
	return _KNWVoting.Contract.KNWTokenAddress(&_KNWVoting.CallOpts)
}

// KNWTokenAddress is a free data retrieval call binding the contract method 0x985dbfc5.
//
// Solidity: function KNWTokenAddress() constant returns(address)
func (_KNWVoting *KNWVotingCallerSession) KNWTokenAddress() (common.Address, error) {
	return _KNWVoting.Contract.KNWTokenAddress(&_KNWVoting.CallOpts)
}

// MINTINGMETHOD is a free data retrieval call binding the contract method 0xeffb21e1.
//
// Solidity: function MINTING_METHOD() constant returns(uint256)
func (_KNWVoting *KNWVotingCaller) MINTINGMETHOD(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "MINTING_METHOD")
	return *ret0, err
}

// MINTINGMETHOD is a free data retrieval call binding the contract method 0xeffb21e1.
//
// Solidity: function MINTING_METHOD() constant returns(uint256)
func (_KNWVoting *KNWVotingSession) MINTINGMETHOD() (*big.Int, error) {
	return _KNWVoting.Contract.MINTINGMETHOD(&_KNWVoting.CallOpts)
}

// MINTINGMETHOD is a free data retrieval call binding the contract method 0xeffb21e1.
//
// Solidity: function MINTING_METHOD() constant returns(uint256)
func (_KNWVoting *KNWVotingCallerSession) MINTINGMETHOD() (*big.Int, error) {
	return _KNWVoting.Contract.MINTINGMETHOD(&_KNWVoting.CallOpts)
}

// CommitPeriodActive is a free data retrieval call binding the contract method 0xa4439dc5.
//
// Solidity: function commitPeriodActive(uint256 _voteID) constant returns(bool active)
func (_KNWVoting *KNWVotingCaller) CommitPeriodActive(opts *bind.CallOpts, _voteID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "commitPeriodActive", _voteID)
	return *ret0, err
}

// CommitPeriodActive is a free data retrieval call binding the contract method 0xa4439dc5.
//
// Solidity: function commitPeriodActive(uint256 _voteID) constant returns(bool active)
func (_KNWVoting *KNWVotingSession) CommitPeriodActive(_voteID *big.Int) (bool, error) {
	return _KNWVoting.Contract.CommitPeriodActive(&_KNWVoting.CallOpts, _voteID)
}

// CommitPeriodActive is a free data retrieval call binding the contract method 0xa4439dc5.
//
// Solidity: function commitPeriodActive(uint256 _voteID) constant returns(bool active)
func (_KNWVoting *KNWVotingCallerSession) CommitPeriodActive(_voteID *big.Int) (bool, error) {
	return _KNWVoting.Contract.CommitPeriodActive(&_KNWVoting.CallOpts, _voteID)
}

// CurrentVoteID is a free data retrieval call binding the contract method 0x5a981129.
//
// Solidity: function currentVoteID() constant returns(uint256)
func (_KNWVoting *KNWVotingCaller) CurrentVoteID(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "currentVoteID")
	return *ret0, err
}

// CurrentVoteID is a free data retrieval call binding the contract method 0x5a981129.
//
// Solidity: function currentVoteID() constant returns(uint256)
func (_KNWVoting *KNWVotingSession) CurrentVoteID() (*big.Int, error) {
	return _KNWVoting.Contract.CurrentVoteID(&_KNWVoting.CallOpts)
}

// CurrentVoteID is a free data retrieval call binding the contract method 0x5a981129.
//
// Solidity: function currentVoteID() constant returns(uint256)
func (_KNWVoting *KNWVotingCallerSession) CurrentVoteID() (*big.Int, error) {
	return _KNWVoting.Contract.CurrentVoteID(&_KNWVoting.CallOpts)
}

// DidCommit is a free data retrieval call binding the contract method 0x7f97e836.
//
// Solidity: function didCommit(address _address, uint256 _voteID) constant returns(bool committed)
func (_KNWVoting *KNWVotingCaller) DidCommit(opts *bind.CallOpts, _address common.Address, _voteID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "didCommit", _address, _voteID)
	return *ret0, err
}

// DidCommit is a free data retrieval call binding the contract method 0x7f97e836.
//
// Solidity: function didCommit(address _address, uint256 _voteID) constant returns(bool committed)
func (_KNWVoting *KNWVotingSession) DidCommit(_address common.Address, _voteID *big.Int) (bool, error) {
	return _KNWVoting.Contract.DidCommit(&_KNWVoting.CallOpts, _address, _voteID)
}

// DidCommit is a free data retrieval call binding the contract method 0x7f97e836.
//
// Solidity: function didCommit(address _address, uint256 _voteID) constant returns(bool committed)
func (_KNWVoting *KNWVotingCallerSession) DidCommit(_address common.Address, _voteID *big.Int) (bool, error) {
	return _KNWVoting.Contract.DidCommit(&_KNWVoting.CallOpts, _address, _voteID)
}

// DidOpen is a free data retrieval call binding the contract method 0x70d18152.
//
// Solidity: function didOpen(address _address, uint256 _voteID) constant returns(bool revealed)
func (_KNWVoting *KNWVotingCaller) DidOpen(opts *bind.CallOpts, _address common.Address, _voteID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "didOpen", _address, _voteID)
	return *ret0, err
}

// DidOpen is a free data retrieval call binding the contract method 0x70d18152.
//
// Solidity: function didOpen(address _address, uint256 _voteID) constant returns(bool revealed)
func (_KNWVoting *KNWVotingSession) DidOpen(_address common.Address, _voteID *big.Int) (bool, error) {
	return _KNWVoting.Contract.DidOpen(&_KNWVoting.CallOpts, _address, _voteID)
}

// DidOpen is a free data retrieval call binding the contract method 0x70d18152.
//
// Solidity: function didOpen(address _address, uint256 _voteID) constant returns(bool revealed)
func (_KNWVoting *KNWVotingCallerSession) DidOpen(_address common.Address, _voteID *big.Int) (bool, error) {
	return _KNWVoting.Contract.DidOpen(&_KNWVoting.CallOpts, _address, _voteID)
}

// DitCoordinatorContracts is a free data retrieval call binding the contract method 0x3fdf3dab.
//
// Solidity: function ditCoordinatorContracts(address ) constant returns(bool)
func (_KNWVoting *KNWVotingCaller) DitCoordinatorContracts(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "ditCoordinatorContracts", arg0)
	return *ret0, err
}

// DitCoordinatorContracts is a free data retrieval call binding the contract method 0x3fdf3dab.
//
// Solidity: function ditCoordinatorContracts(address ) constant returns(bool)
func (_KNWVoting *KNWVotingSession) DitCoordinatorContracts(arg0 common.Address) (bool, error) {
	return _KNWVoting.Contract.DitCoordinatorContracts(&_KNWVoting.CallOpts, arg0)
}

// DitCoordinatorContracts is a free data retrieval call binding the contract method 0x3fdf3dab.
//
// Solidity: function ditCoordinatorContracts(address ) constant returns(bool)
func (_KNWVoting *KNWVotingCallerSession) DitCoordinatorContracts(arg0 common.Address) (bool, error) {
	return _KNWVoting.Contract.DitCoordinatorContracts(&_KNWVoting.CallOpts, arg0)
}

// GetAmountOfVotes is a free data retrieval call binding the contract method 0xcd6ef46b.
//
// Solidity: function getAmountOfVotes(address _address, uint256 _voteID) constant returns(uint256 numberOfVotes)
func (_KNWVoting *KNWVotingCaller) GetAmountOfVotes(opts *bind.CallOpts, _address common.Address, _voteID *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "getAmountOfVotes", _address, _voteID)
	return *ret0, err
}

// GetAmountOfVotes is a free data retrieval call binding the contract method 0xcd6ef46b.
//
// Solidity: function getAmountOfVotes(address _address, uint256 _voteID) constant returns(uint256 numberOfVotes)
func (_KNWVoting *KNWVotingSession) GetAmountOfVotes(_address common.Address, _voteID *big.Int) (*big.Int, error) {
	return _KNWVoting.Contract.GetAmountOfVotes(&_KNWVoting.CallOpts, _address, _voteID)
}

// GetAmountOfVotes is a free data retrieval call binding the contract method 0xcd6ef46b.
//
// Solidity: function getAmountOfVotes(address _address, uint256 _voteID) constant returns(uint256 numberOfVotes)
func (_KNWVoting *KNWVotingCallerSession) GetAmountOfVotes(_address common.Address, _voteID *big.Int) (*big.Int, error) {
	return _KNWVoting.Contract.GetAmountOfVotes(&_KNWVoting.CallOpts, _address, _voteID)
}

// GetGrossStake is a free data retrieval call binding the contract method 0xc3eb9708.
//
// Solidity: function getGrossStake(uint256 _voteID) constant returns(uint256 grossStake)
func (_KNWVoting *KNWVotingCaller) GetGrossStake(opts *bind.CallOpts, _voteID *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "getGrossStake", _voteID)
	return *ret0, err
}

// GetGrossStake is a free data retrieval call binding the contract method 0xc3eb9708.
//
// Solidity: function getGrossStake(uint256 _voteID) constant returns(uint256 grossStake)
func (_KNWVoting *KNWVotingSession) GetGrossStake(_voteID *big.Int) (*big.Int, error) {
	return _KNWVoting.Contract.GetGrossStake(&_KNWVoting.CallOpts, _voteID)
}

// GetGrossStake is a free data retrieval call binding the contract method 0xc3eb9708.
//
// Solidity: function getGrossStake(uint256 _voteID) constant returns(uint256 grossStake)
func (_KNWVoting *KNWVotingCallerSession) GetGrossStake(_voteID *big.Int) (*big.Int, error) {
	return _KNWVoting.Contract.GetGrossStake(&_KNWVoting.CallOpts, _voteID)
}

// GetNetStake is a free data retrieval call binding the contract method 0xdcfde092.
//
// Solidity: function getNetStake(uint256 _voteID) constant returns(uint256 netStake)
func (_KNWVoting *KNWVotingCaller) GetNetStake(opts *bind.CallOpts, _voteID *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "getNetStake", _voteID)
	return *ret0, err
}

// GetNetStake is a free data retrieval call binding the contract method 0xdcfde092.
//
// Solidity: function getNetStake(uint256 _voteID) constant returns(uint256 netStake)
func (_KNWVoting *KNWVotingSession) GetNetStake(_voteID *big.Int) (*big.Int, error) {
	return _KNWVoting.Contract.GetNetStake(&_KNWVoting.CallOpts, _voteID)
}

// GetNetStake is a free data retrieval call binding the contract method 0xdcfde092.
//
// Solidity: function getNetStake(uint256 _voteID) constant returns(uint256 netStake)
func (_KNWVoting *KNWVotingCallerSession) GetNetStake(_voteID *big.Int) (*big.Int, error) {
	return _KNWVoting.Contract.GetNetStake(&_KNWVoting.CallOpts, _voteID)
}

// GetUsedKNW is a free data retrieval call binding the contract method 0x8e1b660f.
//
// Solidity: function getUsedKNW(address _address, uint256 _voteID) constant returns(uint256 usedKNW)
func (_KNWVoting *KNWVotingCaller) GetUsedKNW(opts *bind.CallOpts, _address common.Address, _voteID *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "getUsedKNW", _address, _voteID)
	return *ret0, err
}

// GetUsedKNW is a free data retrieval call binding the contract method 0x8e1b660f.
//
// Solidity: function getUsedKNW(address _address, uint256 _voteID) constant returns(uint256 usedKNW)
func (_KNWVoting *KNWVotingSession) GetUsedKNW(_address common.Address, _voteID *big.Int) (*big.Int, error) {
	return _KNWVoting.Contract.GetUsedKNW(&_KNWVoting.CallOpts, _address, _voteID)
}

// GetUsedKNW is a free data retrieval call binding the contract method 0x8e1b660f.
//
// Solidity: function getUsedKNW(address _address, uint256 _voteID) constant returns(uint256 usedKNW)
func (_KNWVoting *KNWVotingCallerSession) GetUsedKNW(_address common.Address, _voteID *big.Int) (*big.Int, error) {
	return _KNWVoting.Contract.GetUsedKNW(&_KNWVoting.CallOpts, _address, _voteID)
}

// IsExpired is a free data retrieval call binding the contract method 0xd9548e53.
//
// Solidity: function isExpired(uint256 _terminationDate) constant returns(bool expired)
func (_KNWVoting *KNWVotingCaller) IsExpired(opts *bind.CallOpts, _terminationDate *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "isExpired", _terminationDate)
	return *ret0, err
}

// IsExpired is a free data retrieval call binding the contract method 0xd9548e53.
//
// Solidity: function isExpired(uint256 _terminationDate) constant returns(bool expired)
func (_KNWVoting *KNWVotingSession) IsExpired(_terminationDate *big.Int) (bool, error) {
	return _KNWVoting.Contract.IsExpired(&_KNWVoting.CallOpts, _terminationDate)
}

// IsExpired is a free data retrieval call binding the contract method 0xd9548e53.
//
// Solidity: function isExpired(uint256 _terminationDate) constant returns(bool expired)
func (_KNWVoting *KNWVotingCallerSession) IsExpired(_terminationDate *big.Int) (bool, error) {
	return _KNWVoting.Contract.IsExpired(&_KNWVoting.CallOpts, _terminationDate)
}

// IsPassed is a free data retrieval call binding the contract method 0x49403183.
//
// Solidity: function isPassed(uint256 _voteID) constant returns(bool passed)
func (_KNWVoting *KNWVotingCaller) IsPassed(opts *bind.CallOpts, _voteID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "isPassed", _voteID)
	return *ret0, err
}

// IsPassed is a free data retrieval call binding the contract method 0x49403183.
//
// Solidity: function isPassed(uint256 _voteID) constant returns(bool passed)
func (_KNWVoting *KNWVotingSession) IsPassed(_voteID *big.Int) (bool, error) {
	return _KNWVoting.Contract.IsPassed(&_KNWVoting.CallOpts, _voteID)
}

// IsPassed is a free data retrieval call binding the contract method 0x49403183.
//
// Solidity: function isPassed(uint256 _voteID) constant returns(bool passed)
func (_KNWVoting *KNWVotingCallerSession) IsPassed(_voteID *big.Int) (bool, error) {
	return _KNWVoting.Contract.IsPassed(&_KNWVoting.CallOpts, _voteID)
}

// IsResolved is a free data retrieval call binding the contract method 0xcc4e1954.
//
// Solidity: function isResolved(uint256 _voteID) constant returns(bool resolved)
func (_KNWVoting *KNWVotingCaller) IsResolved(opts *bind.CallOpts, _voteID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "isResolved", _voteID)
	return *ret0, err
}

// IsResolved is a free data retrieval call binding the contract method 0xcc4e1954.
//
// Solidity: function isResolved(uint256 _voteID) constant returns(bool resolved)
func (_KNWVoting *KNWVotingSession) IsResolved(_voteID *big.Int) (bool, error) {
	return _KNWVoting.Contract.IsResolved(&_KNWVoting.CallOpts, _voteID)
}

// IsResolved is a free data retrieval call binding the contract method 0xcc4e1954.
//
// Solidity: function isResolved(uint256 _voteID) constant returns(bool resolved)
func (_KNWVoting *KNWVotingCallerSession) IsResolved(_voteID *big.Int) (bool, error) {
	return _KNWVoting.Contract.IsResolved(&_KNWVoting.CallOpts, _voteID)
}

// LastKNWVoting is a free data retrieval call binding the contract method 0xff77f4b8.
//
// Solidity: function lastKNWVoting() constant returns(address)
func (_KNWVoting *KNWVotingCaller) LastKNWVoting(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "lastKNWVoting")
	return *ret0, err
}

// LastKNWVoting is a free data retrieval call binding the contract method 0xff77f4b8.
//
// Solidity: function lastKNWVoting() constant returns(address)
func (_KNWVoting *KNWVotingSession) LastKNWVoting() (common.Address, error) {
	return _KNWVoting.Contract.LastKNWVoting(&_KNWVoting.CallOpts)
}

// LastKNWVoting is a free data retrieval call binding the contract method 0xff77f4b8.
//
// Solidity: function lastKNWVoting() constant returns(address)
func (_KNWVoting *KNWVotingCallerSession) LastKNWVoting() (common.Address, error) {
	return _KNWVoting.Contract.LastKNWVoting(&_KNWVoting.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() constant returns(address)
func (_KNWVoting *KNWVotingCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "manager")
	return *ret0, err
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() constant returns(address)
func (_KNWVoting *KNWVotingSession) Manager() (common.Address, error) {
	return _KNWVoting.Contract.Manager(&_KNWVoting.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() constant returns(address)
func (_KNWVoting *KNWVotingCallerSession) Manager() (common.Address, error) {
	return _KNWVoting.Contract.Manager(&_KNWVoting.CallOpts)
}

// NextKNWVoting is a free data retrieval call binding the contract method 0x0bf682e6.
//
// Solidity: function nextKNWVoting() constant returns(address)
func (_KNWVoting *KNWVotingCaller) NextKNWVoting(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "nextKNWVoting")
	return *ret0, err
}

// NextKNWVoting is a free data retrieval call binding the contract method 0x0bf682e6.
//
// Solidity: function nextKNWVoting() constant returns(address)
func (_KNWVoting *KNWVotingSession) NextKNWVoting() (common.Address, error) {
	return _KNWVoting.Contract.NextKNWVoting(&_KNWVoting.CallOpts)
}

// NextKNWVoting is a free data retrieval call binding the contract method 0x0bf682e6.
//
// Solidity: function nextKNWVoting() constant returns(address)
func (_KNWVoting *KNWVotingCallerSession) NextKNWVoting() (common.Address, error) {
	return _KNWVoting.Contract.NextKNWVoting(&_KNWVoting.CallOpts)
}

// OpenPeriodActive is a free data retrieval call binding the contract method 0xea521a52.
//
// Solidity: function openPeriodActive(uint256 _voteID) constant returns(bool active)
func (_KNWVoting *KNWVotingCaller) OpenPeriodActive(opts *bind.CallOpts, _voteID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "openPeriodActive", _voteID)
	return *ret0, err
}

// OpenPeriodActive is a free data retrieval call binding the contract method 0xea521a52.
//
// Solidity: function openPeriodActive(uint256 _voteID) constant returns(bool active)
func (_KNWVoting *KNWVotingSession) OpenPeriodActive(_voteID *big.Int) (bool, error) {
	return _KNWVoting.Contract.OpenPeriodActive(&_KNWVoting.CallOpts, _voteID)
}

// OpenPeriodActive is a free data retrieval call binding the contract method 0xea521a52.
//
// Solidity: function openPeriodActive(uint256 _voteID) constant returns(bool active)
func (_KNWVoting *KNWVotingCallerSession) OpenPeriodActive(_voteID *big.Int) (bool, error) {
	return _KNWVoting.Contract.OpenPeriodActive(&_KNWVoting.CallOpts, _voteID)
}

// StakesOfVote is a free data retrieval call binding the contract method 0x1f84950f.
//
// Solidity: function stakesOfVote(uint256 ) constant returns(uint256 proposersStake, uint256 proposersReward, uint256 returnPool, uint256 rewardPool)
func (_KNWVoting *KNWVotingCaller) StakesOfVote(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ProposersStake  *big.Int
	ProposersReward *big.Int
	ReturnPool      *big.Int
	RewardPool      *big.Int
}, error) {
	ret := new(struct {
		ProposersStake  *big.Int
		ProposersReward *big.Int
		ReturnPool      *big.Int
		RewardPool      *big.Int
	})
	out := ret
	err := _KNWVoting.contract.Call(opts, out, "stakesOfVote", arg0)
	return *ret, err
}

// StakesOfVote is a free data retrieval call binding the contract method 0x1f84950f.
//
// Solidity: function stakesOfVote(uint256 ) constant returns(uint256 proposersStake, uint256 proposersReward, uint256 returnPool, uint256 rewardPool)
func (_KNWVoting *KNWVotingSession) StakesOfVote(arg0 *big.Int) (struct {
	ProposersStake  *big.Int
	ProposersReward *big.Int
	ReturnPool      *big.Int
	RewardPool      *big.Int
}, error) {
	return _KNWVoting.Contract.StakesOfVote(&_KNWVoting.CallOpts, arg0)
}

// StakesOfVote is a free data retrieval call binding the contract method 0x1f84950f.
//
// Solidity: function stakesOfVote(uint256 ) constant returns(uint256 proposersStake, uint256 proposersReward, uint256 returnPool, uint256 rewardPool)
func (_KNWVoting *KNWVotingCallerSession) StakesOfVote(arg0 *big.Int) (struct {
	ProposersStake  *big.Int
	ProposersReward *big.Int
	ReturnPool      *big.Int
	RewardPool      *big.Int
}, error) {
	return _KNWVoting.Contract.StakesOfVote(&_KNWVoting.CallOpts, arg0)
}

// StartingVoteID is a free data retrieval call binding the contract method 0x82ec45a3.
//
// Solidity: function startingVoteID() constant returns(uint256)
func (_KNWVoting *KNWVotingCaller) StartingVoteID(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "startingVoteID")
	return *ret0, err
}

// StartingVoteID is a free data retrieval call binding the contract method 0x82ec45a3.
//
// Solidity: function startingVoteID() constant returns(uint256)
func (_KNWVoting *KNWVotingSession) StartingVoteID() (*big.Int, error) {
	return _KNWVoting.Contract.StartingVoteID(&_KNWVoting.CallOpts)
}

// StartingVoteID is a free data retrieval call binding the contract method 0x82ec45a3.
//
// Solidity: function startingVoteID() constant returns(uint256)
func (_KNWVoting *KNWVotingCallerSession) StartingVoteID() (*big.Int, error) {
	return _KNWVoting.Contract.StartingVoteID(&_KNWVoting.CallOpts)
}

// VoteEnded is a free data retrieval call binding the contract method 0xb43a401d.
//
// Solidity: function voteEnded(uint256 _voteID) constant returns(bool ended)
func (_KNWVoting *KNWVotingCaller) VoteEnded(opts *bind.CallOpts, _voteID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "voteEnded", _voteID)
	return *ret0, err
}

// VoteEnded is a free data retrieval call binding the contract method 0xb43a401d.
//
// Solidity: function voteEnded(uint256 _voteID) constant returns(bool ended)
func (_KNWVoting *KNWVotingSession) VoteEnded(_voteID *big.Int) (bool, error) {
	return _KNWVoting.Contract.VoteEnded(&_KNWVoting.CallOpts, _voteID)
}

// VoteEnded is a free data retrieval call binding the contract method 0xb43a401d.
//
// Solidity: function voteEnded(uint256 _voteID) constant returns(bool ended)
func (_KNWVoting *KNWVotingCallerSession) VoteEnded(_voteID *big.Int) (bool, error) {
	return _KNWVoting.Contract.VoteEnded(&_KNWVoting.CallOpts, _voteID)
}

// VoteExists is a free data retrieval call binding the contract method 0x7e4173a8.
//
// Solidity: function voteExists(uint256 _voteID) constant returns(bool exists)
func (_KNWVoting *KNWVotingCaller) VoteExists(opts *bind.CallOpts, _voteID *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KNWVoting.contract.Call(opts, out, "voteExists", _voteID)
	return *ret0, err
}

// VoteExists is a free data retrieval call binding the contract method 0x7e4173a8.
//
// Solidity: function voteExists(uint256 _voteID) constant returns(bool exists)
func (_KNWVoting *KNWVotingSession) VoteExists(_voteID *big.Int) (bool, error) {
	return _KNWVoting.Contract.VoteExists(&_KNWVoting.CallOpts, _voteID)
}

// VoteExists is a free data retrieval call binding the contract method 0x7e4173a8.
//
// Solidity: function voteExists(uint256 _voteID) constant returns(bool exists)
func (_KNWVoting *KNWVotingCallerSession) VoteExists(_voteID *big.Int) (bool, error) {
	return _KNWVoting.Contract.VoteExists(&_KNWVoting.CallOpts, _voteID)
}

// Votes is a free data retrieval call binding the contract method 0x5df81330.
//
// Solidity: function votes(uint256 ) constant returns(bytes32 repository, uint256 knowledgeID, uint256 commitEndDate, uint256 openEndDate, uint256 neededMajority, uint256 winningPercentage, uint256 votesFor, uint256 votesAgainst, uint256 votesUnrevealed, uint256 participantsFor, uint256 participantsAgainst, uint256 participantsUnrevealed, bool isResolved)
func (_KNWVoting *KNWVotingCaller) Votes(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Repository             [32]byte
	KnowledgeID            *big.Int
	CommitEndDate          *big.Int
	OpenEndDate            *big.Int
	NeededMajority         *big.Int
	WinningPercentage      *big.Int
	VotesFor               *big.Int
	VotesAgainst           *big.Int
	VotesUnrevealed        *big.Int
	ParticipantsFor        *big.Int
	ParticipantsAgainst    *big.Int
	ParticipantsUnrevealed *big.Int
	IsResolved             bool
}, error) {
	ret := new(struct {
		Repository             [32]byte
		KnowledgeID            *big.Int
		CommitEndDate          *big.Int
		OpenEndDate            *big.Int
		NeededMajority         *big.Int
		WinningPercentage      *big.Int
		VotesFor               *big.Int
		VotesAgainst           *big.Int
		VotesUnrevealed        *big.Int
		ParticipantsFor        *big.Int
		ParticipantsAgainst    *big.Int
		ParticipantsUnrevealed *big.Int
		IsResolved             bool
	})
	out := ret
	err := _KNWVoting.contract.Call(opts, out, "votes", arg0)
	return *ret, err
}

// Votes is a free data retrieval call binding the contract method 0x5df81330.
//
// Solidity: function votes(uint256 ) constant returns(bytes32 repository, uint256 knowledgeID, uint256 commitEndDate, uint256 openEndDate, uint256 neededMajority, uint256 winningPercentage, uint256 votesFor, uint256 votesAgainst, uint256 votesUnrevealed, uint256 participantsFor, uint256 participantsAgainst, uint256 participantsUnrevealed, bool isResolved)
func (_KNWVoting *KNWVotingSession) Votes(arg0 *big.Int) (struct {
	Repository             [32]byte
	KnowledgeID            *big.Int
	CommitEndDate          *big.Int
	OpenEndDate            *big.Int
	NeededMajority         *big.Int
	WinningPercentage      *big.Int
	VotesFor               *big.Int
	VotesAgainst           *big.Int
	VotesUnrevealed        *big.Int
	ParticipantsFor        *big.Int
	ParticipantsAgainst    *big.Int
	ParticipantsUnrevealed *big.Int
	IsResolved             bool
}, error) {
	return _KNWVoting.Contract.Votes(&_KNWVoting.CallOpts, arg0)
}

// Votes is a free data retrieval call binding the contract method 0x5df81330.
//
// Solidity: function votes(uint256 ) constant returns(bytes32 repository, uint256 knowledgeID, uint256 commitEndDate, uint256 openEndDate, uint256 neededMajority, uint256 winningPercentage, uint256 votesFor, uint256 votesAgainst, uint256 votesUnrevealed, uint256 participantsFor, uint256 participantsAgainst, uint256 participantsUnrevealed, bool isResolved)
func (_KNWVoting *KNWVotingCallerSession) Votes(arg0 *big.Int) (struct {
	Repository             [32]byte
	KnowledgeID            *big.Int
	CommitEndDate          *big.Int
	OpenEndDate            *big.Int
	NeededMajority         *big.Int
	WinningPercentage      *big.Int
	VotesFor               *big.Int
	VotesAgainst           *big.Int
	VotesUnrevealed        *big.Int
	ParticipantsFor        *big.Int
	ParticipantsAgainst    *big.Int
	ParticipantsUnrevealed *big.Int
	IsResolved             bool
}, error) {
	return _KNWVoting.Contract.Votes(&_KNWVoting.CallOpts, arg0)
}

// AddDitCoordinator is a paid mutator transaction binding the contract method 0xa326690a.
//
// Solidity: function addDitCoordinator(address _newDitCoordinatorAddress) returns(bool success)
func (_KNWVoting *KNWVotingTransactor) AddDitCoordinator(opts *bind.TransactOpts, _newDitCoordinatorAddress common.Address) (*types.Transaction, error) {
	return _KNWVoting.contract.Transact(opts, "addDitCoordinator", _newDitCoordinatorAddress)
}

// AddDitCoordinator is a paid mutator transaction binding the contract method 0xa326690a.
//
// Solidity: function addDitCoordinator(address _newDitCoordinatorAddress) returns(bool success)
func (_KNWVoting *KNWVotingSession) AddDitCoordinator(_newDitCoordinatorAddress common.Address) (*types.Transaction, error) {
	return _KNWVoting.Contract.AddDitCoordinator(&_KNWVoting.TransactOpts, _newDitCoordinatorAddress)
}

// AddDitCoordinator is a paid mutator transaction binding the contract method 0xa326690a.
//
// Solidity: function addDitCoordinator(address _newDitCoordinatorAddress) returns(bool success)
func (_KNWVoting *KNWVotingTransactorSession) AddDitCoordinator(_newDitCoordinatorAddress common.Address) (*types.Transaction, error) {
	return _KNWVoting.Contract.AddDitCoordinator(&_KNWVoting.TransactOpts, _newDitCoordinatorAddress)
}

// AddNewRepository is a paid mutator transaction binding the contract method 0x828c1653.
//
// Solidity: function addNewRepository(bytes32 _newRepository, uint256 _majority) returns(bool success)
func (_KNWVoting *KNWVotingTransactor) AddNewRepository(opts *bind.TransactOpts, _newRepository [32]byte, _majority *big.Int) (*types.Transaction, error) {
	return _KNWVoting.contract.Transact(opts, "addNewRepository", _newRepository, _majority)
}

// AddNewRepository is a paid mutator transaction binding the contract method 0x828c1653.
//
// Solidity: function addNewRepository(bytes32 _newRepository, uint256 _majority) returns(bool success)
func (_KNWVoting *KNWVotingSession) AddNewRepository(_newRepository [32]byte, _majority *big.Int) (*types.Transaction, error) {
	return _KNWVoting.Contract.AddNewRepository(&_KNWVoting.TransactOpts, _newRepository, _majority)
}

// AddNewRepository is a paid mutator transaction binding the contract method 0x828c1653.
//
// Solidity: function addNewRepository(bytes32 _newRepository, uint256 _majority) returns(bool success)
func (_KNWVoting *KNWVotingTransactorSession) AddNewRepository(_newRepository [32]byte, _majority *big.Int) (*types.Transaction, error) {
	return _KNWVoting.Contract.AddNewRepository(&_KNWVoting.TransactOpts, _newRepository, _majority)
}

// CommitVote is a paid mutator transaction binding the contract method 0xd4e0ac95.
//
// Solidity: function commitVote(uint256 _voteID, address _address, bytes32 _secretHash, uint256 _numberOfKNW) returns(uint256 numberOfVotes)
func (_KNWVoting *KNWVotingTransactor) CommitVote(opts *bind.TransactOpts, _voteID *big.Int, _address common.Address, _secretHash [32]byte, _numberOfKNW *big.Int) (*types.Transaction, error) {
	return _KNWVoting.contract.Transact(opts, "commitVote", _voteID, _address, _secretHash, _numberOfKNW)
}

// CommitVote is a paid mutator transaction binding the contract method 0xd4e0ac95.
//
// Solidity: function commitVote(uint256 _voteID, address _address, bytes32 _secretHash, uint256 _numberOfKNW) returns(uint256 numberOfVotes)
func (_KNWVoting *KNWVotingSession) CommitVote(_voteID *big.Int, _address common.Address, _secretHash [32]byte, _numberOfKNW *big.Int) (*types.Transaction, error) {
	return _KNWVoting.Contract.CommitVote(&_KNWVoting.TransactOpts, _voteID, _address, _secretHash, _numberOfKNW)
}

// CommitVote is a paid mutator transaction binding the contract method 0xd4e0ac95.
//
// Solidity: function commitVote(uint256 _voteID, address _address, bytes32 _secretHash, uint256 _numberOfKNW) returns(uint256 numberOfVotes)
func (_KNWVoting *KNWVotingTransactorSession) CommitVote(_voteID *big.Int, _address common.Address, _secretHash [32]byte, _numberOfKNW *big.Int) (*types.Transaction, error) {
	return _KNWVoting.Contract.CommitVote(&_KNWVoting.TransactOpts, _voteID, _address, _secretHash, _numberOfKNW)
}

// EndVote is a paid mutator transaction binding the contract method 0x865df0ad.
//
// Solidity: function endVote(uint256 _voteID) returns(bool votePassed)
func (_KNWVoting *KNWVotingTransactor) EndVote(opts *bind.TransactOpts, _voteID *big.Int) (*types.Transaction, error) {
	return _KNWVoting.contract.Transact(opts, "endVote", _voteID)
}

// EndVote is a paid mutator transaction binding the contract method 0x865df0ad.
//
// Solidity: function endVote(uint256 _voteID) returns(bool votePassed)
func (_KNWVoting *KNWVotingSession) EndVote(_voteID *big.Int) (*types.Transaction, error) {
	return _KNWVoting.Contract.EndVote(&_KNWVoting.TransactOpts, _voteID)
}

// EndVote is a paid mutator transaction binding the contract method 0x865df0ad.
//
// Solidity: function endVote(uint256 _voteID) returns(bool votePassed)
func (_KNWVoting *KNWVotingTransactorSession) EndVote(_voteID *big.Int) (*types.Transaction, error) {
	return _KNWVoting.Contract.EndVote(&_KNWVoting.TransactOpts, _voteID)
}

// FinalizeVote is a paid mutator transaction binding the contract method 0x36bf4c91.
//
// Solidity: function finalizeVote(uint256 _voteID, uint256 _voteOption, address _address) returns(uint256 reward, bool winningSide, uint256 numberOfKNW)
func (_KNWVoting *KNWVotingTransactor) FinalizeVote(opts *bind.TransactOpts, _voteID *big.Int, _voteOption *big.Int, _address common.Address) (*types.Transaction, error) {
	return _KNWVoting.contract.Transact(opts, "finalizeVote", _voteID, _voteOption, _address)
}

// FinalizeVote is a paid mutator transaction binding the contract method 0x36bf4c91.
//
// Solidity: function finalizeVote(uint256 _voteID, uint256 _voteOption, address _address) returns(uint256 reward, bool winningSide, uint256 numberOfKNW)
func (_KNWVoting *KNWVotingSession) FinalizeVote(_voteID *big.Int, _voteOption *big.Int, _address common.Address) (*types.Transaction, error) {
	return _KNWVoting.Contract.FinalizeVote(&_KNWVoting.TransactOpts, _voteID, _voteOption, _address)
}

// FinalizeVote is a paid mutator transaction binding the contract method 0x36bf4c91.
//
// Solidity: function finalizeVote(uint256 _voteID, uint256 _voteOption, address _address) returns(uint256 reward, bool winningSide, uint256 numberOfKNW)
func (_KNWVoting *KNWVotingTransactorSession) FinalizeVote(_voteID *big.Int, _voteOption *big.Int, _address common.Address) (*types.Transaction, error) {
	return _KNWVoting.Contract.FinalizeVote(&_KNWVoting.TransactOpts, _voteID, _voteOption, _address)
}

// OpenVote is a paid mutator transaction binding the contract method 0xcdd6ceb9.
//
// Solidity: function openVote(uint256 _voteID, address _address, uint256 _voteOption, uint256 _salt) returns(bool success)
func (_KNWVoting *KNWVotingTransactor) OpenVote(opts *bind.TransactOpts, _voteID *big.Int, _address common.Address, _voteOption *big.Int, _salt *big.Int) (*types.Transaction, error) {
	return _KNWVoting.contract.Transact(opts, "openVote", _voteID, _address, _voteOption, _salt)
}

// OpenVote is a paid mutator transaction binding the contract method 0xcdd6ceb9.
//
// Solidity: function openVote(uint256 _voteID, address _address, uint256 _voteOption, uint256 _salt) returns(bool success)
func (_KNWVoting *KNWVotingSession) OpenVote(_voteID *big.Int, _address common.Address, _voteOption *big.Int, _salt *big.Int) (*types.Transaction, error) {
	return _KNWVoting.Contract.OpenVote(&_KNWVoting.TransactOpts, _voteID, _address, _voteOption, _salt)
}

// OpenVote is a paid mutator transaction binding the contract method 0xcdd6ceb9.
//
// Solidity: function openVote(uint256 _voteID, address _address, uint256 _voteOption, uint256 _salt) returns(bool success)
func (_KNWVoting *KNWVotingTransactorSession) OpenVote(_voteID *big.Int, _address common.Address, _voteOption *big.Int, _salt *big.Int) (*types.Transaction, error) {
	return _KNWVoting.Contract.OpenVote(&_KNWVoting.TransactOpts, _voteID, _address, _voteOption, _salt)
}

// ReplaceManager is a paid mutator transaction binding the contract method 0x23447982.
//
// Solidity: function replaceManager(address _newManager) returns(bool success)
func (_KNWVoting *KNWVotingTransactor) ReplaceManager(opts *bind.TransactOpts, _newManager common.Address) (*types.Transaction, error) {
	return _KNWVoting.contract.Transact(opts, "replaceManager", _newManager)
}

// ReplaceManager is a paid mutator transaction binding the contract method 0x23447982.
//
// Solidity: function replaceManager(address _newManager) returns(bool success)
func (_KNWVoting *KNWVotingSession) ReplaceManager(_newManager common.Address) (*types.Transaction, error) {
	return _KNWVoting.Contract.ReplaceManager(&_KNWVoting.TransactOpts, _newManager)
}

// ReplaceManager is a paid mutator transaction binding the contract method 0x23447982.
//
// Solidity: function replaceManager(address _newManager) returns(bool success)
func (_KNWVoting *KNWVotingTransactorSession) ReplaceManager(_newManager common.Address) (*types.Transaction, error) {
	return _KNWVoting.Contract.ReplaceManager(&_KNWVoting.TransactOpts, _newManager)
}

// StartVote is a paid mutator transaction binding the contract method 0xe5023ff2.
//
// Solidity: function startVote(bytes32 _repository, address _address, uint256 _knowledgeID, uint256 _voteDuration, uint256 _proposersStake, uint256 _numberOfKNW) returns(uint256 voteID)
func (_KNWVoting *KNWVotingTransactor) StartVote(opts *bind.TransactOpts, _repository [32]byte, _address common.Address, _knowledgeID *big.Int, _voteDuration *big.Int, _proposersStake *big.Int, _numberOfKNW *big.Int) (*types.Transaction, error) {
	return _KNWVoting.contract.Transact(opts, "startVote", _repository, _address, _knowledgeID, _voteDuration, _proposersStake, _numberOfKNW)
}

// StartVote is a paid mutator transaction binding the contract method 0xe5023ff2.
//
// Solidity: function startVote(bytes32 _repository, address _address, uint256 _knowledgeID, uint256 _voteDuration, uint256 _proposersStake, uint256 _numberOfKNW) returns(uint256 voteID)
func (_KNWVoting *KNWVotingSession) StartVote(_repository [32]byte, _address common.Address, _knowledgeID *big.Int, _voteDuration *big.Int, _proposersStake *big.Int, _numberOfKNW *big.Int) (*types.Transaction, error) {
	return _KNWVoting.Contract.StartVote(&_KNWVoting.TransactOpts, _repository, _address, _knowledgeID, _voteDuration, _proposersStake, _numberOfKNW)
}

// StartVote is a paid mutator transaction binding the contract method 0xe5023ff2.
//
// Solidity: function startVote(bytes32 _repository, address _address, uint256 _knowledgeID, uint256 _voteDuration, uint256 _proposersStake, uint256 _numberOfKNW) returns(uint256 voteID)
func (_KNWVoting *KNWVotingTransactorSession) StartVote(_repository [32]byte, _address common.Address, _knowledgeID *big.Int, _voteDuration *big.Int, _proposersStake *big.Int, _numberOfKNW *big.Int) (*types.Transaction, error) {
	return _KNWVoting.Contract.StartVote(&_KNWVoting.TransactOpts, _repository, _address, _knowledgeID, _voteDuration, _proposersStake, _numberOfKNW)
}

// UpgradeContract is a paid mutator transaction binding the contract method 0xeb2c0223.
//
// Solidity: function upgradeContract(address _address) returns(bool success)
func (_KNWVoting *KNWVotingTransactor) UpgradeContract(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _KNWVoting.contract.Transact(opts, "upgradeContract", _address)
}

// UpgradeContract is a paid mutator transaction binding the contract method 0xeb2c0223.
//
// Solidity: function upgradeContract(address _address) returns(bool success)
func (_KNWVoting *KNWVotingSession) UpgradeContract(_address common.Address) (*types.Transaction, error) {
	return _KNWVoting.Contract.UpgradeContract(&_KNWVoting.TransactOpts, _address)
}

// UpgradeContract is a paid mutator transaction binding the contract method 0xeb2c0223.
//
// Solidity: function upgradeContract(address _address) returns(bool success)
func (_KNWVoting *KNWVotingTransactorSession) UpgradeContract(_address common.Address) (*types.Transaction, error) {
	return _KNWVoting.Contract.UpgradeContract(&_KNWVoting.TransactOpts, _address)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"a\",\"type\":\"uint256\"}],\"name\":\"sqrt\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// SafeMathFuncSigs maps the 4-byte function signature to its string representation.
var SafeMathFuncSigs = map[string]string{
	"677342ce": "sqrt(uint256)",
}

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x610129610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361060335760003560e01c8063677342ce146038575b600080fd5b605260048036036020811015604c57600080fd5b50356064565b60408051918252519081900360200190f35b60008160715750600060ef565b81826001011160bf576040805162461bcd60e51b8152602060048201526015602482015274119b185dd959081a5b9c1d5d08199bdc881cdc5c9d605a1b604482015290519081900360640190fd5b60026001830104825b8082101560eb57508060028180868160dc57fe5b04018160e457fe5b04915060c8565b5090505b91905056fea265627a7a72305820953020e84b9f63c03b694ac9fe34b1131c3ea11c415e206ee9a52f32f26c3bd364736f6c634300050a0032"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// Sqrt is a free data retrieval call binding the contract method 0x677342ce.
//
// Solidity: function sqrt(uint256 a) constant returns(uint256)
func (_SafeMath *SafeMathCaller) Sqrt(opts *bind.CallOpts, a *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SafeMath.contract.Call(opts, out, "sqrt", a)
	return *ret0, err
}

// Sqrt is a free data retrieval call binding the contract method 0x677342ce.
//
// Solidity: function sqrt(uint256 a) constant returns(uint256)
func (_SafeMath *SafeMathSession) Sqrt(a *big.Int) (*big.Int, error) {
	return _SafeMath.Contract.Sqrt(&_SafeMath.CallOpts, a)
}

// Sqrt is a free data retrieval call binding the contract method 0x677342ce.
//
// Solidity: function sqrt(uint256 a) constant returns(uint256)
func (_SafeMath *SafeMathCallerSession) Sqrt(a *big.Int) (*big.Int, error) {
	return _SafeMath.Contract.Sqrt(&_SafeMath.CallOpts, a)
}
