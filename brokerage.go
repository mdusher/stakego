package stakego

import (
    "encoding/json"
)

// NewBrokerageFromJSON - create a Brokerage item from a json byte slice
func NewBrokerageFromJSON(jsonStr []byte) *Brokerage {
  var b Brokerage
  err := json.Unmarshal(jsonStr, &b)
  if err != nil {
    return nil
  }
  return &b
}

// Brokerage - store result from a brokerage request
type Brokerage struct {
  BrokerageFee          float64 `json:"brokerageFee"`
  BrokerageDiscount     float64 `json:"brokerageDiscount"`
  FixedFee              float64 `json:"fixedFee"`
  VariableFeePercentage float64 `json:"variableFeePercentage"`
  VariableLimit         int     `json:"variableLimit"`
}