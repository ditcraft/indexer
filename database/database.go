package database

import (
	"errors"
	"fmt"
	"hash/fnv"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/golang/glog"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var mgoSession *mgo.Session
var databaseAddress string
var databaseName = "ditIndexV2"

// Proposal struct
type Proposal struct {
	ProposalID          int                    `json:"id" bson:"id"`
	KNWVoteID           int                    `json:"knw_vote_id" bson:"knw_vote_id"`
	KNWLabel            string                 `json:"knw_label" bson:"knw_label"`
	Repository          string                 `json:"repository" bson:"repository"`
	Proposer            string                 `json:"proposer" bson:"proposer"`
	Topic               string                 `json:"topic" bson:"topic"`
	Identifier          string                 `json:"identifier" bson:"identifier"`
	CreationDate        time.Time              `json:"creation_date" bson:"creation_date"`
	CommitPhaseEnd      time.Time              `json:"commit_phase_end" bson:"commit_phase_end"`
	RevealPhaseEnd      time.Time              `json:"reveal_phase_end" bson:"reveal_phase_end"`
	Finalized           bool                   `json:"finalized" bson:"finalized"`
	Accepted            bool                   `json:"accepted" bson:"accepted"`
	Votes               ProposalVotes          `json:"votes" bson:"votes"`
	Participants        []ProposalParticipants `json:"participants" bson:"participants"`
	StakePerParticipant float64                `json:"stake_per_participant" bson:"stake_per_participant"`
	TotalStake          float64                `json:"total_stake" bson:"total_stake"`
	TotalMintedKNW      float64                `json:"total_minted_knw" bson:"total_minted_knw"`
}

// ProposalVotes struct
type ProposalVotes struct {
	ParticipantsFor        int     `json:"participants_for" bson:"participants_for"`
	ParticipantsAgainst    int     `json:"participants_against" bson:"participants_against"`
	ParticipantsUnrevealed int     `json:"participants_unrevealed" bson:"participants_unrevealed"`
	VotesFor               float64 `json:"votes_for" bson:"votes_for"`
	VotesAgainst           float64 `json:"votes_against" bson:"votes_against"`
	VotesUnrevealed        float64 `json:"votes_unrevealed" bson:"votes_unrevealed"`
}

// ProposalParticipants struct
type ProposalParticipants struct {
	Address       string  `json:"address" bson:"address"`
	UsedKNW       float64 `json:"used_knw" bson:"used_knw"`
	Votes         float64 `json:"votes" bson:"votes"`
	Opened        bool    `json:"opened" bson:"opened"`
	Finalized     bool    `json:"finalized" bson:"finalized"`
	VotedRight    bool    `json:"voted_right" bson:"voted_right"`
	KNWDifference float64 `json:"knw_difference" bson:"knw_difference"`
}

// User struct
type User struct {
	DitAddress          string                  `json:"dit_address" bson:"dit_address"`
	AuthorizedAddresses AuthorizedAddress       `json:"authorized_adresses" bson:"authorized_adresses"`
	TwitterID           string                  `json:"twitter_id" bson:"twitter_id"`
	GitHubID            string                  `json:"github_id" bson:"github_id"`
	GitHubToken         string                  `json:"github_token" bson:"github_token"`
	MainAccount         string                  `json:"main_account" bson:"main_account"`
	XDAIBalance         float64                 `json:"xdai_balance" bson:"xdai_balance"`
	XDITBalance         float64                 `json:"xdit_balance" bson:"xdit_balance"`
	KNWTokensLive       []KNWLabels             `json:"knw_tokens_live" bson:"knw_tokens_live"`
	ProposalsLive       []ProposalDetailsUser   `json:"proposals_live" bson:"proposals_live"`
	RepositoriesLive    []RepositoryDetailsUser `json:"repositories_live" bson:"repositories_live"`
	KNWTokensDemo       []KNWLabels             `json:"knw_tokens_demo" bson:"knw_tokens_demo"`
	ProposalsDemo       []ProposalDetailsUser   `json:"proposals_demo" bson:"proposals_demo"`
	RepositoriesDemo    []RepositoryDetailsUser `json:"repositories_demo" bson:"repositories_demo"`
}

// AuthorizedAddress struct
type AuthorizedAddress struct {
	DitCLI      string `json:"dit_cli" bson:"dit_cli"`
	DitExplorer string `json:"dit_explorer" bson:"dit_explorer"`
	Alice       string `json:"alice" bson:"alice"`
}

// KNWLabels struct
type KNWLabels struct {
	Label      string  `json:"label" bson:"label"`
	Balance    float64 `json:"balance" bson:"balance"`
	RawBalance string  `json:"raw_balance" bson:"raw_balance"`
}

// ProposalDetailsUser struct
type ProposalDetailsUser struct {
	KNWVoteID     int       `json:"knw_vote_id" bson:"knw_vote_id"`
	VoteDate      time.Time `json:"vote_date" bson:"vote_date"`
	UsedKNW       float64   `json:"used_knw" bson:"used_knw"`
	Stake         float64   `json:"stake" bson:"stake"`
	IsProposer    bool      `json:"is_proposer" bson:"is_proposer"`
	Opened        bool      `json:"opened" bson:"opened"`
	Finalized     bool      `json:"finalized" bson:"finalized"`
	VotedRight    bool      `json:"voted_right" bson:"voted_right"`
	KNWDifference float64   `json:"knw_difference" bson:"knw_difference"`
}

// RepositoryDetailsUser struct
type RepositoryDetailsUser struct {
	Hash                string    `json:"hash" bson:"hash"`
	Name                string    `json:"name" bson:"name"`
	Provider            string    `json:"provider" bson:"provider"`
	URL                 string    `json:"url" bson:"url"`
	Notifications       bool      `json:"notifications" bson:"notifications"`
	LastActivityDate    time.Time `json:"last_activity_date" bson:"last_activity_date"`
	EarnedKNW           float64   `json:"earned_knw" bson:"earned_knw"`
	AmountOfProposals   int       `json:"amount_of_proposals" bson:"amount_of_proposals"`
	AmountOfValidations int       `json:"amount_of_validations" bson:"amount_of_validations"`
}

// Repository struct
type Repository struct {
	Hash             string                       `json:"hash" bson:"hash"`
	Name             string                       `json:"name" bson:"name"`
	Provider         string                       `json:"provider" bson:"provider"`
	URL              string                       `json:"url" bson:"url"`
	CreationDate     time.Time                    `json:"creation_date" bson:"creation_date"`
	LastActivityDate time.Time                    `json:"last_activity_date" bson:"last_activity_date"`
	Majority         int                          `json:"majority" bson:"majority"`
	KnowledgeLabels  []string                     `json:"knw_labels" bson:"knw_labels"`
	Proposals        []ProposalDetailsRepository  `json:"proposals" bson:"proposals"`
	Contributors     []ProposalDetailsContributor `json:"contributors" bson:"contributors"`
}

// ProposalDetailsRepository struct
type ProposalDetailsRepository struct {
	KNWVoteID int       `json:"knw_vote_id" bson:"knw_vote_id"`
	VoteDate  time.Time `json:"vote_date" bson:"vote_date"`
}

// ProposalDetailsContributor struct
type ProposalDetailsContributor struct {
	Address             string    `json:"address" bson:"address"`
	LastActivityDate    time.Time `json:"last_activity_date" bson:"last_activity_date"`
	EarnedKNW           float64   `json:"earned_knw" bson:"earned_knw"`
	AmountOfProposals   int       `json:"amount_of_proposals" bson:"amount_of_proposals"`
	AmountOfValidations int       `json:"amount_of_validations" bson:"amount_of_validations"`
}

// KnowledgeLabel Struct
type KnowledgeLabel struct {
	Label            string    `json:"label" bson:"label"`
	ID               int       `json:"id" bson:"id"`
	TotalSupply      float64   `json:"total_supply" bson:"total_supply"`
	LastActivityDate time.Time `json:"last_activity_date" bson:"last_activity_date"`
}

// NotificationProposalStarted struct
type NotificationProposalStarted struct {
	TwitterID         string    `json:"twitter_id" bson:"twitter_id"`
	LiveMode          bool      `json:"live_mode" bson:"live_mode"`
	RepositoryHash    string    `json:"repository_hash" bson:"repository_hash"`
	RepositoryName    string    `json:"repository_name" bson:"repository_name"`
	KNWVoteID         int       `json:"knw_vote_id" bson:"knw_vote_id"`
	KnowledgeLabel    string    `json:"knowledge_label" bson:"knowledge_label"`
	ProposerTwitterID string    `json:"proposer_twitter_handle" bson:"proposer_twitter_handle"`
	Description       string    `json:"description" bson:"description"`
	Identifier        string    `json:"identifier" bson:"identifier"`
	CommitUntil       time.Time `json:"commit_until" bson:"commit_until"`
	RevealUntil       time.Time `json:"reveal_until" bson:"reveal_until"`
}

// LogEntry struct
type LogEntry struct {
	ID      string `json:"id" bson:"id"`
	Address string `json:"address" bson:"address"`
	Version string `json:"version" bson:"version"`
	LogData string `json:"log_data" bson:"log_data"`
}

// Get returns a bson-byte array with the result of the database request
func Get(collection string, key []string, operator []string, value []interface{}) ([][]byte, error) {
	if len(key) != len(value) && (len(operator) == 0 || len(operator) == 1 || len(operator) == len(value)) {
		return nil, errors.New("Key- and Value-Arrays have to be the same length")
	}
	goOperators, err := convertToGoOperator(operator)
	if err != nil {
		return nil, err
	}
	var bsonReturn []bson.M
	mgoErr := MgoRequest(collection, func(c *mgo.Collection) error {
		requestBson := bson.M{}
		for index, element := range key {
			var currentOperator string
			if len(goOperators) == 0 {
				currentOperator = "$eq"
			} else if len(goOperators) == 1 {
				currentOperator = goOperators[0]
			} else {
				currentOperator = goOperators[index]
			}
			requestBson[element] = bson.M{currentOperator: value[index]}
		}
		return c.Find(requestBson).All(&bsonReturn)
	})
	if mgoErr != nil {
		return nil, mgoErr
	}
	var bsonBytes [][]byte
	for _, element := range bsonReturn {
		bytes, bsonErr := bson.Marshal(element)
		if bsonErr != nil {
			return nil, bsonErr
		}
		bsonBytes = append(bsonBytes, bytes)
	}
	return bsonBytes, nil
}

// Delete deletes an entry from a collection, returns an error if the request failed
func Delete(collection string, key string, operator string, value interface{}) error {
	goOperators, err := convertToGoOperator([]string{operator})
	if err != nil {
		return err
	}
	mgoErr := MgoRequest(collection, func(c *mgo.Collection) error {
		_, mgoErr := c.RemoveAll(bson.M{key: bson.M{goOperators[0]: value}})
		return mgoErr
	})
	if mgoErr != nil {
		return mgoErr
	}
	return nil
}

// Insert inserts a new entry into the database, returns an error if the request failed
func Insert(collection string, structInterface interface{}) error {
	mgoErr := MgoRequest(collection, func(c *mgo.Collection) error {
		return c.Insert(structInterface)
	})
	if mgoErr != nil {
		return mgoErr
	}
	return nil
}

// Replace replaces an existing entry in the database, returns an error if the request failed
func Replace(collection string, whereKey string, whereValue interface{}, structInterface interface{}) error {
	if len(whereKey) == 0 || whereValue == nil {
		return errors.New("Arguments can not be empty")
	}
	mgoErr := MgoRequest(collection, func(c *mgo.Collection) error {
		return c.Update(bson.M{whereKey: whereValue}, structInterface)
	})
	if mgoErr != nil {
		return mgoErr
	}
	return nil
}

// Update updates an existing entry in the database, returns an error if the request failed
func Update(collection string, whereKey string, whereValue interface{}, changeKey []string, changeValue []interface{}) error {
	if len(changeKey) == 0 || len(changeValue) == 0 || len(whereKey) == 0 || whereValue == nil {
		return errors.New("Arguments can not be empty")
	}
	if len(changeKey) != len(changeValue) {
		return errors.New("Key- and Value-Arrays have to be the same length")
	}
	mgoErr := MgoRequest(collection, func(c *mgo.Collection) error {
		changeBson := bson.M{}
		for index, element := range changeKey {
			changeBson[element] = changeValue[index]
		}
		return c.Update(bson.M{whereKey: whereValue}, bson.M{"$set": changeBson})
	})
	if mgoErr != nil {
		return mgoErr
	}
	return nil
}

// StoreLog stores a new log entry into the database
func StoreLog(_address string, _logData string, _version string) (string, error) {
	var newLogEntry LogEntry
	newLogEntry.Address = _address
	newLogEntry.Version = _version
	newLogEntry.LogData = _logData

	dt := time.Now()
	h := fnv.New32a()
	h.Write([]byte(newLogEntry.Address + dt.Format("01-02-2006 15:04:05")))
	id := fmt.Sprint(h.Sum32())
	newLogEntry.ID = id

	err := Insert("logs", newLogEntry)
	if err != nil {
		return "", err
	}
	return id, err
}

// GetRepoHashFromName func
func GetRepoHashFromName(_live bool, _name string) (string, error) {
	collection := "repositories_"
	if _live {
		collection += "live"
	} else {
		collection += "demo"
	}
	var bsonReturn bson.M
	mgoErr := MgoRequest(collection, func(c *mgo.Collection) error {
		requestBson := bson.M{}
		requestBson["name"] = bson.M{"$eq": _name}
		return c.Find(requestBson).One(&bsonReturn)
	})
	if mgoErr != nil {
		return "", mgoErr
	}
	return bsonReturn["hash"].(string), nil
}

// GetRepoNameFromHash func
func GetRepoNameFromHash(_live bool, _hash string) (string, error) {
	collection := "repositories_"
	if _live {
		collection += "live"
	} else {
		collection += "demo"
	}
	var bsonReturn bson.M
	mgoErr := MgoRequest(collection, func(c *mgo.Collection) error {
		requestBson := bson.M{}
		requestBson["hash"] = bson.M{"$eq": _hash}
		return c.Find(requestBson).One(&bsonReturn)
	})
	if mgoErr != nil {
		return "", mgoErr
	}
	return bsonReturn["name"].(string), nil
}

// WasUserProposer func
func WasUserProposer(_live bool, _hash string, _address string) (bool, error) {
	collection := "proposals_"
	if _live {
		collection += "live"
	} else {
		collection += "demo"
	}
	var bsonReturn bson.M
	mgoErr := MgoRequest(collection, func(c *mgo.Collection) error {
		requestBson := bson.M{}
		requestBson["proposer"] = bson.M{"$eq": _address}
		requestBson["repository"] = bson.M{"$eq": _hash}
		return c.Find(requestBson).One(&bsonReturn)
	})
	if mgoErr != nil {
		if strings.Contains(mgoErr.Error(), "not found") {
			return false, nil
		}
		return false, mgoErr
	}
	if bsonReturn["id"].(int) > 0 {
		return true, nil
	}
	return false, nil
}

// WasUserValidator func
func WasUserValidator(_live bool, _hash string, _address string) (bool, error) {
	collection := "proposals_"
	if _live {
		collection += "live"
	} else {
		collection += "demo"
	}
	var bsonReturn bson.M
	mgoErr := MgoRequest(collection, func(c *mgo.Collection) error {
		requestBson := bson.M{}
		requestBson["participants"] = bson.M{"$elemMatch": bson.M{"address": _address}}
		requestBson["proposer"] = bson.M{"$ne": _address}
		requestBson["repository"] = bson.M{"$eq": _hash}
		return c.Find(requestBson).One(&bsonReturn)
	})
	if mgoErr != nil {
		if strings.Contains(mgoErr.Error(), "not found") {
			return false, nil
		}
		return false, mgoErr
	}
	if bsonReturn["id"].(int) > 0 {
		return true, nil
	}
	return false, nil
}

// GetAddressFromID func
func GetAddressFromID(_id string) (string, error) {
	var bsonReturn bson.M
	mgoErr := MgoRequest("users", func(c *mgo.Collection) error {
		requestBson := bson.M{}
		requestBson["twitter_id"] = bson.M{"$eq": _id}
		return c.Find(requestBson).One(&bsonReturn)
	})
	if mgoErr != nil {
		return "", mgoErr
	}
	return bsonReturn["dit_address"].(string), nil
}

// GetIDFromAddress func
func GetIDFromAddress(_address string) (string, error) {
	var bsonReturn bson.M
	mgoErr := MgoRequest("users", func(c *mgo.Collection) error {
		requestBson := bson.M{}
		requestBson["dit_address"] = bson.M{"$eq": _address}
		return c.Find(requestBson).One(&bsonReturn)
	})
	if mgoErr != nil {
		return "", mgoErr
	}
	return bsonReturn["twitter_id"].(string), nil
}

// AddressExists func
func AddressExists(_address string) error {
	var bsonReturn bson.M
	mgoErr := MgoRequest("users", func(c *mgo.Collection) error {
		requestBson := bson.M{}
		requestBson["dit_address"] = bson.M{"$eq": _address}
		return c.Find(requestBson).One(&bsonReturn)
	})
	if mgoErr != nil {
		return mgoErr
	}
	if len(bsonReturn["twitter_id"].(string)) == 0 && len(bsonReturn["github_id"].(string)) == 0 {
		return errors.New("Not found")
	}
	return nil
}

// APIKeyExists func
func APIKeyExists(_key string, _kycRights bool) error {
	var bsonReturn bson.M
	mgoErr := MgoRequest("api_keys", func(c *mgo.Collection) error {
		requestBson := bson.M{}
		requestBson["api_key"] = bson.M{"$eq": _key}
		if _kycRights {
			requestBson["has_kyc_right"] = bson.M{"$eq": _kycRights}
		}
		return c.Find(requestBson).One(&bsonReturn)
	})
	if mgoErr != nil {
		return mgoErr
	}
	if len(bsonReturn["user"].(string)) == 0 {
		return errors.New("Not found")
	}
	return nil
}

// ClearCollection clears a collection, returns an error if the request failed
func ClearCollection(collections ...string) error {
	for _, collection := range collections {
		mgoErr := MgoRequest(collection, func(c *mgo.Collection) error {
			c.RemoveAll(nil)
			return nil
		})
		if mgoErr != nil {
			return mgoErr
		}
	}
	return nil
}

// GetKnowledgeLabel func
func GetKnowledgeLabel(_live bool, _id *big.Int) (string, error) {
	collection := "knowledge_labels_"
	if _live {
		collection += "live"
	} else {
		collection += "demo"
	}
	var bsonReturn bson.M
	mgoErr := MgoRequest(collection, func(c *mgo.Collection) error {
		requestBson := bson.M{}
		requestBson["id"] = bson.M{"$eq": int(_id.Int64())}
		return c.Find(requestBson).One(&bsonReturn)
	})
	if mgoErr != nil {
		return "", mgoErr
	}
	return bsonReturn["label"].(string), nil
}

// GetKnowledgeID func
func GetKnowledgeID(_live bool, _label string) (int, error) {
	collection := "knowledge_labels_"
	if _live {
		collection += "live"
	} else {
		collection += "demo"
	}
	var bsonReturn bson.M
	mgoErr := MgoRequest(collection, func(c *mgo.Collection) error {
		requestBson := bson.M{}
		requestBson["label"] = bson.M{"$eq": _label}
		return c.Find(requestBson).One(&bsonReturn)
	})
	if mgoErr != nil {
		return 0, mgoErr
	}
	return bsonReturn["id"].(int), nil
}

// Notify func
func Notify(_repoHash string, _notification NotificationProposalStarted) error {
	collection := "users_"
	if _notification.LiveMode {
		collection += "live"
	} else {
		collection += "demo"
	}
	var bsonReturn []bson.M
	mgoErr := MgoRequest(collection, func(c *mgo.Collection) error {
		requestBson := bson.M{}
		requestBson["repositories"] = bson.M{"$elemMatch": bson.M{"hash": _repoHash, "notifications": true}}
		return c.Find(requestBson).All(&bsonReturn)
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

	var users []User
	for _, element := range bsonBytes {
		var foundUser User
		bsonErr := bson.Unmarshal(element, &foundUser)
		if bsonErr != nil {
			glog.Error(bsonErr)
			return bsonErr
		}
		users = append(users, foundUser)
	}

	for _, user := range users {
		if user.TwitterID != _notification.ProposerTwitterID {
			newNotification := _notification
			newNotification.TwitterID = user.TwitterID
			mgoErr := MgoRequest("notifications", func(c *mgo.Collection) error {
				return c.Insert(newNotification)
			})
			if mgoErr != nil {
				return mgoErr
			}
		}
	}

	return nil
}

// InitDatabase will return an error if the database does not respond to a ping
func InitDatabase() error {
	session, mgoErr := GetSession()
	if mgoErr != nil {
		return mgoErr
	}
	defer session.Close()
	pingError := session.Ping()
	if pingError != nil {
		return errors.New("Database is not responding")
	}
	log.Info("Successfully connected to database")
	return nil
}

func convertToGoOperator(operatorArray []string) ([]string, error) {
	var goOperatorArray []string
	for _, element := range operatorArray {
		switch element {
		case "=", "==":
			goOperatorArray = append(goOperatorArray, "$eq")
		case "<":
			goOperatorArray = append(goOperatorArray, "$lt")
		case ">":
			goOperatorArray = append(goOperatorArray, "$gt")
		case "<=":
			goOperatorArray = append(goOperatorArray, "$lte")
		case ">=":
			goOperatorArray = append(goOperatorArray, "$gte")
		case "!=":
			goOperatorArray = append(goOperatorArray, "$ne")
		default:
			return []string{}, errors.New("Undefined operator in query")
		}
	}
	return goOperatorArray, nil
}

// GetSession do
func GetSession() (*mgo.Session, error) {
	if mgoSession == nil {
		var mgoErr error
		mgoSession, mgoErr = mgo.DialWithInfo(
			&mgo.DialInfo{Addrs: []string{os.Getenv("MONGO_DB_ADDRESS")},
				Timeout:  10 * time.Second,
				Database: "admin",
				Username: os.Getenv("MONGO_DB_USER"),
				Password: os.Getenv("MONGO_DB_PASSWORD")})
		if mgoErr != nil {
			return nil, mgoErr
		}
		mgoSession.SetMode(mgo.Monotonic, true)
	}
	return mgoSession.Clone(), nil
}

// MgoRequest do
func MgoRequest(collection string, s func(*mgo.Collection) error, database ...string) error {
	session, mgoErr := GetSession()
	if mgoErr != nil {
		return mgoErr
	}
	defer session.Close()
	var c *mgo.Collection
	if len(database) == 1 {
		c = session.DB(database[0]).C(collection)
	} else {
		c = session.DB(databaseName).C(collection)
	}
	return s(c)
}
