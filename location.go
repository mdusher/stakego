package stakego

import (
	"encoding/json"
)

// LocationDataDateFormat - format for dates in the location data
var LocationDataDateFormat = "2006-01-02"
var LocationDataTimeFormat = "2006-01-02 15:04"

// NewLocationFromJSON - creates a Location from a json byte slice
func NewLocationFromJSON(jsonStr []byte) *LocationData {
	var l LocationData
	err := json.Unmarshal(jsonStr, &l)
	if err != nil {
		return nil
	}
	return &l
}

type LocationData struct {
	TradingLimits struct {
		USTRADING struct {
			ExtendedHoursMarketOrders int `json:"extendedHoursMarketOrders"`
		} `json:"US_TRADING"`
	} `json:"tradingLimits"`
	Calendar struct {
		AUTRADING struct {
			TradingHolidays []struct {
				Date      string `json:"date"`
				Name      string `json:"name"`
				ShortName string `json:"shortName"`
			} `json:"tradingHolidays"`
			EarlyClose []struct {
				Date         string `json:"date"`
				Name         string `json:"name"`
				ShortName    string `json:"shortName"`
				TradingOpen  string `json:"trading_open"`
				TradingClose string `json:"trading_close"`
				PreCspaClose string `json:"pre_cspa_close"`
			} `json:"earlyClose"`
		} `json:"AU_TRADING"`
		USTRADING struct {
			TradingHolidays []struct {
				Date      string `json:"date"`
				Name      string `json:"name"`
				ShortName string `json:"short_name"`
			} `json:"trading_holidays"`
			EarlyClose []struct {
				Date            string `json:"date"`
				Name            string `json:"name"`
				ShortName       string `json:"short_name"`
				TradingOpen     string `json:"trading_open"`
				TradingClose    string `json:"trading_close"`
				AfterHoursClose string `json:"after_hours_close"`
			} `json:"earlyClose"`
		} `json:"US_TRADING"`
	} `json:"calendar"`
}
