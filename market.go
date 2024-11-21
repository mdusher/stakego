package stakego

import (
	"encoding/json"
	"fmt"
	"time"
)

type MarketTime struct {
	Hour   int
	Minute int
}

const MarketStatusOpen = "OPEN"
const MarketStatusClosed = "CLOSED"

// Note: this won't be accurate for trading holidays,
// but it's a start.
var MarketDefaultOpenASX = MarketTime{Hour: 10, Minute: 0}
var MarketDefaultCloseASX = MarketTime{Hour: 16, Minute: 0}

// Couple constants for comparison
var MarketAU = "AU"

// NewMarket - creates an empty Market
func NewMarket() *Market {
	m := Market{}
	m.Market = MarketAU
	return &m
}

// NewMarketFromJSON - creates a Market from a json byte slice
func NewMarketFromJSON(jsonStr []byte) *Market {
	m := NewMarket()
	err := json.Unmarshal(jsonStr, &m)
	if err != nil {
		return nil
	}
	return m
}

// NewMarketFromJSON - creates a Market from a json byte slice
func NewMarketWithLocationData(l *LocationData) *Market {
	m := NewMarket()
	m.LocationData = l
	return m
}

// Market - stores response when getting the market status
type Market struct {
	LastTradingDate string `json:"lastTradingDate"`
	Status          struct {
		Current string `json:"current"`
	} `json:"status"`
	MarketLimits  [][]float64 `json:"marketLimits"`
	PassiveLimits [][]float64 `json:"passiveLimits"`
	Market        string
	LocationData  *LocationData
}

// GetStatus - returns the current market status as a string using static values
func (m *Market) GetStatus() string {
	if m.LocationData != nil {
		if m.IsNormalHours() == false {
			return MarketStatusClosed
		}
		if m.IsTradingHoliday() == false && m.HasClosedEarly() == false {
			return MarketStatusOpen
		}
	} else {
		if m.IsNormalHours() {
			return MarketStatusOpen
		}
	}
	return MarketStatusClosed
}

// GetAUTime - returns the current time in Sydney
func GetAUTime() (*time.Time, error) {
	loc, err := time.LoadLocation("Australia/Sydney")
	if err != nil {
		return nil, err
	}
	now := time.Now().In(loc)
	return &now, err
}

// IsNormalHours - checks to see if it is currently within "normal" hours for the market
func (m *Market) IsNormalHours() bool {
	if m.Market == MarketAU {
		now, err := GetAUTime()
		if err != nil {
			return false
		}

		// Weekends closed
		if now.Weekday() == 0 || now.Weekday() == 6 {
			return false
		}

		openTime := time.Date(now.Year(), now.Month(), now.Day(), MarketDefaultOpenASX.Hour, MarketDefaultOpenASX.Minute, 0, 0, now.Location())
		closeTime := time.Date(now.Year(), now.Month(), now.Day(), MarketDefaultCloseASX.Hour, MarketDefaultCloseASX.Minute, 0, 0, now.Location())
		if now.After(openTime) && now.Before(closeTime) {
			return true
		}
	}

	return false
}

func (m *Market) IsTradingHoliday() bool {
	if m.LocationData != nil {
		if m.Market == MarketAU {
			now, err := GetAUTime()
			if err != nil {
				return false
			}
			for _, h := range m.LocationData.Calendar.AUTRADING.TradingHolidays {
				th, err := time.Parse(LocationDataDateFormat, h.Date)
				if err == nil && DatesEqual(*now, th) {
					return true
				}
			}
		}
	}

	return false
}

// HasClosedEarly - check to see if there is an early close today
func (m *Market) HasClosedEarly() bool {
	if m.LocationData != nil {
		if m.Market == MarketAU {
			now, err := GetAUTime()
			if err != nil {
				return false
			}
			// Check early closes
			for _, e := range m.LocationData.Calendar.AUTRADING.EarlyClose {
				eOpen, err := time.Parse(LocationDataTimeFormat, fmt.Sprintf("%s %s", e.Date, e.TradingOpen))
				if err != nil {
					return false
				}
				eClose, err := time.Parse(LocationDataTimeFormat, fmt.Sprintf("%s %s", e.Date, e.TradingClose))
				if err != nil {
					return false
				}
				// Closed if the date matches, and it's outside of hours
				if DatesEqual(*now, eOpen) && now.After(eClose) && now.Before(eOpen) {
					return true
				}
			}
		}
	}
	return false
}
