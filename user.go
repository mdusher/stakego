package stakego

import (
	"encoding/json"
)

// NewUserFromJSON - creates a User from a json byte slice
func NewUserFromJSON(jsonStr []byte) *User {
	var u User
    err := json.Unmarshal(jsonStr, &u)
    if err != nil {
        return nil
    }
	return &u
}

// User - user profile information
type User struct {
	CanTradeOnUnsettledFunds   bool          `json:"canTradeOnUnsettledFunds"`
	CpfValue                   interface{}   `json:"cpfValue"`
	EmailVerified              bool          `json:"emailVerified"`
	HasFunded                  bool          `json:"hasFunded"`
	HasTraded                  bool          `json:"hasTraded"`
	UserID                     string        `json:"userId"`
	Username                   interface{}   `json:"username"`
	EmailAddress               string        `json:"emailAddress"`
	DwAccountID                interface{}   `json:"dw_AccountId"`
	DwAccountNumber            interface{}   `json:"dw_AccountNumber"`
	MacAccountNumber           interface{}   `json:"macAccountNumber"`
	Status                     interface{}   `json:"status"`
	MacStatus                  string        `json:"macStatus"`
	DwStatus                   interface{}   `json:"dwStatus"`
	TruliooStatus              string        `json:"truliooStatus"`
	TruliooStatusWithWatchlist interface{}   `json:"truliooStatusWithWatchlist"`
	FirstName                  string        `json:"firstName"`
	MiddleName                 interface{}   `json:"middleName"`
	LastName                   string        `json:"lastName"`
	PhoneNumber                string        `json:"phoneNumber"`
	SignUpPhase                int           `json:"signUpPhase"`
	AckSignedWhen              string        `json:"ackSignedWhen"`
	CreatedDate                int64         `json:"createdDate"`
	StakeApprovedDate          int64         `json:"stakeApprovedDate"`
	AccountType                string        `json:"accountType"`
	MasterAccountID            interface{}   `json:"masterAccountId"`
	ReferralCode               string        `json:"referralCode"`
	ReferredByCode             interface{}   `json:"referredByCode"`
	RegionIdentifier           string        `json:"regionIdentifier"`
	AssetSummary               interface{}   `json:"assetSummary"`
	FundingStatistics          interface{}   `json:"fundingStatistics"`
	TradingStatistics          interface{}   `json:"tradingStatistics"`
	W8File                     []interface{} `json:"w8File"`
	RewardJourneyTimestamp     interface{}   `json:"rewardJourneyTimestamp"`
	RewardJourneyStatus        interface{}   `json:"rewardJourneyStatus"`
	UserProfile                struct {
		ResidentialAddress interface{} `json:"residentialAddress"`
		PostalAddress      interface{} `json:"postalAddress"`
	} `json:"userProfile"`
	LedgerBalance            float64     `json:"ledgerBalance"`
	InvestorAccreditations   interface{} `json:"investorAccreditations"`
	FxSpeed                  interface{} `json:"fxSpeed"`
	DateOfBirth              interface{} `json:"dateOfBirth"`
	UpToDateDetails2021      string      `json:"upToDateDetails2021"`
	StakeKycStatus           string      `json:"stakeKycStatus"`
	AwxMigrationDocsRequired interface{} `json:"awxMigrationDocsRequired"`
	DocumentsStatus          string      `json:"documentsStatus"`
	AccountStatus            string      `json:"accountStatus"`
	Mfaenabled               bool        `json:"mfaenabled"`
}