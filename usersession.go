package stakego

import (
  "encoding/json"
)

// NewUserSessionFromJSON - creates a UserSession from a json byte slice
func NewUserSessionFromJSON(jsonStr []byte) *UserSession {
  var u UserSession
  err := json.Unmarshal(jsonStr, &u)
  if err != nil {
    return nil
  }
  return &u
}

// UserSession - stores the response from the createSession API
type UserSession struct {
  UserID                   string      `json:"userID"`
  FirstName                string      `json:"firstName"`
  LastName                 string      `json:"lastName"`
  Username                 string      `json:"username"`
  Email                    string      `json:"email"`
  LoginState               interface{} `json:"loginState"`
  CommissionRate           interface{} `json:"commissionRate"`
  WlpID                    interface{} `json:"wlpID"`
  ReferralCode             string      `json:"referralCode"`
  Guest                    interface{} `json:"guest"`
  SessionKey               string      `json:"sessionKey"`
  AppTypeID                interface{} `json:"appTypeID"`
  DefaultProductDetailPage interface{} `json:"defaultProductDetailPage"`
  Status                   string      `json:"status"`
  TruliooStatus            string      `json:"truliooStatus"`
  DwStatus                 interface{} `json:"dwStatus"`
  MacStatus                string      `json:"macStatus"`
  AccountType              string      `json:"accountType"`
  RegionIdentifier         string      `json:"regionIdentifier"`
}