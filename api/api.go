package api

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ditcraft/indexer/database"
	"github.com/ditcraft/indexer/ethereum"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var srv *http.Server

type requestProposal struct {
	APIKey          string `json:"api_key" bson:"api_key"`
	Mode            string `json:"mode" bson:"mode"`
	KNWVoteID       int    `json:"knw_vote_id" bson:"knw_vote_id"`
	UserAddress     string `json:"user_address" bson:"user_address"`
	UserTwitterID   string `json:"user_twitter_id" bson:"user_twitter_id"`
	UserIsProposer  bool   `json:"user_is_proposer" bson:"user_is_proposer"`
	UserIsValidator bool   `json:"user_is_validator" bson:"user_is_validator"`
	SinceDate       string `json:"since_date" bson:"since_date"`
	Amount          int    `json:"amount" bson:"amount"`
	OnlyFinalized   bool   `json:"only_finalized" bson:"only_finalized"`
	OnlyActive      bool   `json:"only_active" bson:"only_active"`
	RepositoryName  string `json:"repository_name" bson:"repository_name"`
	RepositoryHash  string `json:"repository_hash" bson:"repository_hash"`
}

type requestRepository struct {
	APIKey          string `json:"api_key" bson:"api_key"`
	Mode            string `json:"mode" bson:"mode"`
	Name            string `json:"name" bson:"name"`
	Hash            string `json:"hash" bson:"hash"`
	Provider        string `json:"provider" bson:"provider"`
	KnowledgeLabel  string `json:"knowledge_label" bson:"knowledge_label"`
	UserAddress     string `json:"user_address" bson:"user_address"`
	UserTwitterID   string `json:"user_twitter_id" bson:"user_twitter_id"`
	UserIsProposer  bool   `json:"user_is_proposer" bson:"user_is_proposer"`
	UserIsValidator bool   `json:"user_is_validator" bson:"user_is_validator"`
	OnlyActive      bool   `json:"only_active" bson:"only_active"`
	Amount          int    `json:"amount" bson:"amount"`
}

type requestUser struct {
	APIKey    string `json:"api_key" bson:"api_key"`
	Mode      string `json:"mode" bson:"mode"`
	Address   string `json:"address" bson:"address"`
	TwitterID string `json:"twitter_id" bson:"twitter_id"`
}

type requestKYC struct {
	APIKey  string `json:"api_key" bson:"api_key"`
	Address string `json:"address" bson:"address"`
}

type sendLog struct {
	APIKey  string `json:"api_key" bson:"api_key"`
	Address string `json:"address" bson:"address"`
	Version string `json:"version" bson:"version"`
	LogData string `json:"log_data" bson:"log_data"`
}

type responseStatus struct {
	Status string `json:"status" bson:"status"`
}

// Start the API listener
func Start(address string, timeoutInSeconds int) {
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)

	router := mux.NewRouter()
	router.HandleFunc("/api/proposal", handleProposal)
	router.HandleFunc("/api/user", handleUser)
	router.HandleFunc("/api/repository", handleRepository)
	router.HandleFunc("/api/kyc", handleKYC)
	router.HandleFunc("/api/uploadlog", handleLogUpload)
	// router.HandleFunc("/indexer/metric", handleDeviceList)

	srv = &http.Server{
		Addr:              address,
		WriteTimeout:      time.Second * time.Duration(timeoutInSeconds),
		ReadTimeout:       time.Second * time.Duration(timeoutInSeconds),
		ReadHeaderTimeout: time.Second * time.Duration(timeoutInSeconds),
		IdleTimeout:       time.Second * time.Duration(timeoutInSeconds*4),
		Handler:           router,
		ErrorLog:          logger,
	}
	var listenErr error
	go func() {
		listenErr = srv.ListenAndServe()
		if listenErr != nil {
			if strings.Contains(listenErr.Error(), "Server closed") {
				glog.Info("Successfully closed")
			} else {
				glog.Error(listenErr)
			}
		}
	}()
}

// Stop the API listener
func Stop() {
	err := srv.Shutdown(nil)
	if err != nil {
		glog.Error(err)
	}
}

func handleProposal(w http.ResponseWriter, r *http.Request) {
	if r.ContentLength > 0 {
		request := requestProposal{UserIsProposer: true, UserIsValidator: true, Amount: 10, OnlyFinalized: false, OnlyActive: false}
		jsonErr := json.NewDecoder(r.Body).Decode(&request)
		if jsonErr != nil {
			glog.Error(jsonErr)
			returnErrorStatus(w, r, "invalid request")
			return
		}

		err := validateAPIKey(request.APIKey, false)
		if err != nil {
			glog.Error(errors.New("Unauthorized access"))
			returnErrorStatus(w, r, "unauthorized")
			return
		}

		collection := "proposals_"
		if request.Mode == "live" {
			collection += "live"
		} else if request.Mode == "demo" {
			collection += "demo"
		} else {
			returnErrorStatus(w, r, "invalid mode")
			return
		}

		var bsonReturn []bson.M
		mgoErr := database.MgoRequest(collection, func(c *mgo.Collection) error {
			requestBson := bson.M{}
			if request.KNWVoteID > 0 {
				requestBson["knw_vote_id"] = bson.M{"$eq": request.KNWVoteID}
			}
			if len(request.UserAddress) > 0 {
				if request.UserIsProposer && request.UserIsValidator {
					requestBson["participants"] = bson.M{"$elemMatch": bson.M{"address": request.UserAddress}}
				} else if !request.UserIsProposer && request.UserIsValidator {
					requestBson["participants"] = bson.M{"$elemMatch": bson.M{"address": request.UserAddress}}
					requestBson["proposer"] = bson.M{"$ne": request.UserAddress}
				} else if request.UserIsProposer && !request.UserIsValidator {
					requestBson["proposer"] = bson.M{"$eq": request.UserAddress}
				}
			} else if len(request.UserTwitterID) > 0 {
				address, err := database.GetAddressFromID(request.UserTwitterID)
				if err != nil {
					glog.Error(err)
				}
				if request.UserIsProposer && request.UserIsValidator {
					requestBson["participants"] = bson.M{"$elemMatch": bson.M{"address": address}}
				} else if !request.UserIsProposer && request.UserIsValidator {
					requestBson["participants"] = bson.M{"$elemMatch": bson.M{"address": address}}
					requestBson["proposer"] = bson.M{"$ne": address}
				} else if request.UserIsProposer && !request.UserIsValidator {
					requestBson["proposer"] = bson.M{"$eq": address}
				}
			}

			if len(request.RepositoryHash) > 0 {
				requestBson["repository"] = bson.M{"$eq": request.RepositoryHash}
			} else if len(request.RepositoryName) > 0 {
				hash, err := database.GetRepoHashFromName(false, request.RepositoryName)
				if err != nil {
					glog.Error(err)
				}
				requestBson["repository"] = bson.M{"$eq": hash}
			}
			if request.OnlyFinalized {
				requestBson["finalized"] = bson.M{"$eq": true}
			}
			if len(request.SinceDate) > 0 {
				intTime, err := strconv.Atoi(request.SinceDate)
				if err == nil {
					sinceDate := time.Unix(int64(intTime), 0)
					requestBson["creation_date"] = bson.M{"$gte": sinceDate}
				}
			}
			if request.OnlyActive {
				requestBson["reveal_phase_end"] = bson.M{"$gt": time.Now()}
			}
			return c.Find(requestBson).Limit(request.Amount).Sort("-creation_date").All(&bsonReturn)
		})
		if mgoErr != nil {
			glog.Error(mgoErr)
		}
		var bsonBytes [][]byte
		for _, element := range bsonReturn {
			bytes, bsonErr := bson.Marshal(element)
			if bsonErr != nil {
				glog.Error(bsonErr)
			}
			bsonBytes = append(bsonBytes, bytes)
		}

		var responseProposals []database.Proposal
		for _, element := range bsonBytes {
			var foundProposal database.Proposal
			bsonErr := bson.Unmarshal(element, &foundProposal)
			if bsonErr != nil {
				glog.Error(bsonErr)
				returnErrorStatus(w, r, "error during processing")
				return
			}
			responseProposals = append(responseProposals, foundProposal)
		}

		jsonString, jsonErr := json.Marshal(responseProposals)
		if jsonErr != nil {
			glog.Error(jsonErr)
			returnErrorStatus(w, r, "error during processing")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, string(jsonString))
	}
}

func handleRepository(w http.ResponseWriter, r *http.Request) {
	if r.ContentLength > 0 {
		request := requestRepository{UserIsProposer: true, UserIsValidator: true, Amount: 10, OnlyActive: false}
		jsonErr := json.NewDecoder(r.Body).Decode(&request)
		if jsonErr != nil {
			glog.Error(jsonErr)
			returnErrorStatus(w, r, "invalid request")
			return
		}

		err := validateAPIKey(request.APIKey, false)
		if err != nil {
			glog.Error(errors.New("Unauthorized access"))
			returnErrorStatus(w, r, "unauthorized")
			return
		}

		collection := "repositories_"
		if request.Mode == "live" {
			collection += "live"
		} else if request.Mode == "demo" {
			collection += "demo"
		} else {
			returnErrorStatus(w, r, "invalid mode")
			return
		}

		var bsonReturn []bson.M
		mgoErr := database.MgoRequest(collection, func(c *mgo.Collection) error {
			requestBson := bson.M{}
			if len(request.Hash) > 0 {
				requestBson["hash"] = bson.M{"$eq": request.Hash}
			} else if len(request.Name) > 0 {
				hash, err := database.GetRepoHashFromName(false, request.Name)
				if err != nil {
					glog.Error(err)
				}
				requestBson["hash"] = bson.M{"$eq": hash}
			}
			if len(request.Provider) > 0 {
				requestBson["provider"] = bson.M{"$eq": request.Provider}
			}
			if len(request.KnowledgeLabel) > 0 {
				requestBson["knw_labels"] = request.KnowledgeLabel
			}
			if len(request.UserTwitterID) > 0 {
				address, err := database.GetAddressFromID(request.UserTwitterID)
				if err != nil {
					glog.Error(err)
				} else {
					request.UserAddress = address
				}
			}
			if len(request.UserAddress) > 0 {
				if request.UserIsProposer && request.UserIsValidator {
					requestBson["contributors"] = bson.M{"$elemMatch": bson.M{"address": request.UserAddress}}
				} else if !request.UserIsProposer && request.UserIsValidator {
					requestBson["contributors"] = bson.M{"$elemMatch": bson.M{"address": request.UserAddress, "amount_of_proposals": 0}}
				} else if request.UserIsProposer && !request.UserIsValidator {
					requestBson["contributors"] = bson.M{"$elemMatch": bson.M{"address": request.UserAddress, "amount_of_validations": 0}}
				}
			}
			if request.OnlyActive {
				requestBson["proposals.0"] = bson.M{"$exists": true}
			}
			return c.Find(requestBson).Limit(request.Amount).Sort("-creation_date").All(&bsonReturn)
		})
		if mgoErr != nil {
			glog.Error(mgoErr)
		}
		var bsonBytes [][]byte
		for _, element := range bsonReturn {
			bytes, bsonErr := bson.Marshal(element)
			if bsonErr != nil {
				glog.Error(bsonErr)
			}
			bsonBytes = append(bsonBytes, bytes)
		}

		var responseRepositories []database.Repository
		for _, element := range bsonBytes {
			var foundRepository database.Repository
			bsonErr := bson.Unmarshal(element, &foundRepository)
			if bsonErr != nil {
				glog.Error(bsonErr)
				returnErrorStatus(w, r, "error during processing")
				return
			}
			responseRepositories = append(responseRepositories, foundRepository)
		}

		jsonString, jsonErr := json.Marshal(responseRepositories)
		if jsonErr != nil {
			glog.Error(jsonErr)
			returnErrorStatus(w, r, "error during processing")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, string(jsonString))
	}
}

func handleUser(w http.ResponseWriter, r *http.Request) {
	if r.ContentLength > 0 {
		request := requestUser{}
		jsonErr := json.NewDecoder(r.Body).Decode(&request)
		if jsonErr != nil {
			glog.Error(jsonErr)
			returnErrorStatus(w, r, "invalid request")
			return
		}

		err := validateAPIKey(request.APIKey, false)
		if err != nil {
			glog.Error(errors.New("Unauthorized access"))
			returnErrorStatus(w, r, "unauthorized")
			return
		}

		collection := "users"
		subcollection := ""
		if request.Mode == "live" {
			subcollection += "live"
		} else if request.Mode == "demo" {
			subcollection += "demo"
		} else {
			returnErrorStatus(w, r, "invalid mode")
			return
		}

		var bsonReturn bson.M
		mgoErr := database.MgoRequest(collection, func(c *mgo.Collection) error {
			requestBson := bson.M{}
			if len(request.Address) > 0 {
				requestBson["dit_address"] = bson.M{"$eq": request.Address}
			} else if len(request.TwitterID) > 0 {
				requestBson["twitter_id"] = bson.M{"$eq": request.TwitterID}
			} else {
				return errors.New("invalid request")
			}

			return c.Find(requestBson).One(&bsonReturn)
		})

		if mgoErr != nil {
			if strings.Contains(mgoErr.Error(), "not found") {
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprintf(w, string([]byte{}))
				return
			}
			glog.Error(mgoErr)
			returnErrorStatus(w, r, "error during processing")
			return
		}

		bytes, bsonErr := bson.Marshal(bsonReturn)
		if bsonErr != nil {
			glog.Error(bsonErr)
			returnErrorStatus(w, r, "error during processing")
			return
		}

		var responseUser database.User
		bsonErr = bson.Unmarshal(bytes, &responseUser)
		if bsonErr != nil {
			glog.Error(bsonErr)
			returnErrorStatus(w, r, "error during processing")
			return
		}

		jsonString, jsonErr := json.Marshal(responseUser)
		if jsonErr != nil {
			glog.Error(jsonErr)
			returnErrorStatus(w, r, "error during processing")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, string(jsonString))
	}
}

func handleKYC(w http.ResponseWriter, r *http.Request) {
	if r.ContentLength > 0 {
		var request requestKYC
		jsonErr := json.NewDecoder(r.Body).Decode(&request)
		if jsonErr != nil {
			glog.Error(jsonErr)
			returnErrorStatus(w, r, "invalid request")
			return
		}

		err := validateAPIKey(request.APIKey, true)
		if err != nil {
			glog.Error(errors.New("Unauthorized access"))
			returnErrorStatus(w, r, "unauthorized")
			return
		}

		if len(request.Address) != 42 || !strings.HasPrefix(request.Address, "0x") {
			glog.Error(jsonErr)
			returnErrorStatus(w, r, "error during kyc process")
			return
		}

		ethereum.Mutex.Lock()
		defer ethereum.Mutex.Unlock()

		err = ethereum.KYCPassed(request.Address, false)
		if err != nil {
			glog.Error(jsonErr)
			returnErrorStatus(w, r, "error during kyc process")
			return
		}

		err = ethereum.KYCPassed(request.Address, false)
		if err != nil {
			glog.Error(jsonErr)
			returnErrorStatus(w, r, "error during kyc process")
			return
		}

		err = ethereum.SendDitTokens(request.Address)
		if err != nil {
			glog.Error(jsonErr)
			returnErrorStatus(w, r, "error during kyc process")
			return
		}

		err = ethereum.SendXDaiCent(request.Address)
		if err != nil {
			glog.Error(jsonErr)
			returnErrorStatus(w, r, "error during kyc process")
			return
		}

		// time.Sleep(5 * time.Second)

		var responseCode responseStatus
		responseCode.Status = "ok"

		jsonString, jsonErr := json.Marshal(responseCode)
		if jsonErr != nil {
			glog.Error(jsonErr)
			returnErrorStatus(w, r, "error during processing")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, string(jsonString))
	}
}

func handleLogUpload(w http.ResponseWriter, r *http.Request) {
	if r.ContentLength > 0 {
		var sendLog sendLog
		jsonErr := json.NewDecoder(r.Body).Decode(&sendLog)
		if jsonErr != nil {
			glog.Error(jsonErr)
			returnErrorStatus(w, r, "invalid request")
			return
		}

		err := validateAPIKey(sendLog.APIKey, true)
		if err != nil {
			glog.Error(errors.New("Unauthorized access"))
			returnErrorStatus(w, r, "unauthorized")
			return
		}

		id, err := database.StoreLog(sendLog.Address, sendLog.LogData, sendLog.Version)
		if err != nil {
			glog.Error(errors.New("Unauthorized access"))
			returnErrorStatus(w, r, "unauthorized")
			return
		}

		var responseCode responseStatus
		responseCode.Status = id

		jsonString, jsonErr := json.Marshal(responseCode)
		if jsonErr != nil {
			glog.Error(jsonErr)
			returnErrorStatus(w, r, "error during processing")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, string(jsonString))
	}
}

func validateAPIKey(_apiKey string, _needsKycRights bool) error {
	if len(_apiKey) == 130 {
		hash := crypto.Keccak256Hash([]byte("api"))
		bytesSignature, err := hex.DecodeString(_apiKey)
		if err != nil {
			return err
		}
		pubKey, err := crypto.SigToPub(hash.Bytes(), bytesSignature)
		if err != nil {
			return err
		}

		address := crypto.PubkeyToAddress(*pubKey)

		err = database.AddressExists(address.Hex())
		return err
	}

	err := database.APIKeyExists(_apiKey, _needsKycRights)
	return err
}

func returnErrorStatus(w http.ResponseWriter, r *http.Request, statusCode string) {
	var newResponseStatus responseStatus
	newResponseStatus.Status = statusCode
	jsonString, jsonErr := json.Marshal(newResponseStatus)
	if jsonErr != nil {
		glog.Error(jsonErr)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(jsonString))
}
