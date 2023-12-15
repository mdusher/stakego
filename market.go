package stakego

import (
  "encoding/json"
)

const MarketStatusOpen = "OPEN"
const MarketStatusClosed = "CLOSED"

// NewMarketFromJSON - creates a Market from a json byte slice
func NewMarketFromJSON(jsonStr []byte) *Market {
  var m Market
  err := json.Unmarshal(jsonStr, &m)
  if err != nil {
    return nil
  }
  return &m
}

// Market - stores response when getting the market status
type Market struct {
  LastTradingDate string `json:"lastTradingDate"`
  Status          struct {
    Current string `json:"current"`
  } `json:"status"`
  MarketLimits  [][]float64 `json:"marketLimits"`
  PassiveLimits [][]float64 `json:"passiveLimits"`
}

// GetStatus - returns the current market status as a string
func (m *Market) GetStatus() string {
  if m.Status.Current == "open" {
    return MarketStatusOpen
  }
  return MarketStatusClosed
}