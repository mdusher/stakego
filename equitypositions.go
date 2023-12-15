package stakego

import (
  "encoding/json"
)

// NewEquityPositionsFromJSON - create an EquityPositions from a json byte slice
func NewEquityPositionsFromJSON(jsonStr []byte) *EquityPositions {
  var e EquityPositions
  err := json.Unmarshal(jsonStr, &e)
  if err != nil {
    return nil
  }
  return &e
}

// EquityPostitions - stores response from the EquityPositions request
type EquityPositions struct {
  PageNum         int  `json:"pageNum"`
  HasNext         bool `json:"hasNext"`
  EquityPositions []EquityPositionItem `json:"equityPositions"`
}

// EquityPositionItem - Stores information about an equity item
type EquityPositionItem struct {
  InstrumentID           string      `json:"instrumentId"`
  Symbol                 string      `json:"symbol"`
  Name                   string      `json:"name"`
  OpenQty                int         `json:"openQty,string"`
  AvailableForTradingQty int         `json:"availableForTradingQty,string"`
  AveragePrice           float64     `json:"averagePrice,string"`
  MarketValue            float64     `json:"marketValue,string"`
  MktPrice               float64     `json:"mktPrice,string"`
  PriorClose             float64     `json:"priorClose,string"`
  UnrealizedDayPL        float64     `json:"unrealizedDayPL,string"`
  UnrealizedDayPLPercent float64     `json:"unrealizedDayPLPercent,string"`
  UnrealizedPL           float64     `json:"unrealizedPL,string"`
  UnrealizedPLPercent    float64     `json:"unrealizedPLPercent,string"`
  RecentAnnouncement     bool        `json:"recentAnnouncement"`
  Sensitive              bool        `json:"sensitive"`
  CapitalRaise           interface{} `json:"capitalRaise"`
}

func (e *EquityPositions) GetTotal() float64 {
  total := 0.0
  for _, ep := range e.EquityPositions {
    total += ep.MarketValue
  }
  return total
}
