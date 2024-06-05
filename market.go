package stakego

import (
  "encoding/json"
  "time"
)

type MarketTime struct {
  Hour int
  Minute int
}

const MarketStatusOpen = "OPEN"
const MarketStatusClosed = "CLOSED"

// Note: this won't be accurate for trading holidays,
// but it's a start.
var MarketDefaultOpenASX = MarketTime{Hour: 7, Minute: 0}
var MarketDefaultCloseASX = MarketTime{Hour: 14, Minute: 0}

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
  loc, err := time.LoadLocation("Australia/Sydney")
  if err != nil {
    // Return closed, if we can't figure out if it's open
    return MarketStatusClosed
  }

  now := time.Now().In(loc)
  openTime := time.Date(now.Year(), now.Month(), now.Day(), MarketDefaultOpenASX.Hour, MarketDefaultOpenASX.Minute, 0, 0, now.Location())
  closeTime := time.Date(now.Year(), now.Month(), now.Day(), MarketDefaultCloseASX.Hour, MarketDefaultCloseASX.Minute, 0, 0, now.Location())

  if now.After(openTime) && now.Before(closeTime) {
    return MarketStatusOpen
  }
  return MarketStatusClosed
}