package stakego

import (
	"encoding/json"
)

// NewInstrumentFromJSON - creates an Instrument from a JSON string
func NewInstrumentFromJSON(jsonStr []byte) *Instrument {
	var i Instrument
	err := json.Unmarshal(jsonStr, &i)
	if err != nil {
		return nil
	}
	return &i
}

// NewInstrumentResponseFromJSON - creates an InstrumentResponse from a JSON string
func NewInstrumentResponseFromJSON(jsonStr []byte) *InstrumentResponse {
	var i InstrumentResponse
	err := json.Unmarshal(jsonStr, &i)
	if err != nil {
		return nil
	}
	return &i
}

type Instrument struct {
	InstrumentID       string `json:"instrumentId"`
	Symbol             string `json:"symbol"`
	Name               string `json:"name"`
	Type               string `json:"type"`
	RecentAnnouncement bool   `json:"recentAnnouncement"`
	Sensitive          bool   `json:"sensitive"`
	GfdOnly            bool   `json:"gfdOnly"`
	MarketCap          string `json:"marketCap"`
	Score              int    `json:"score"`
}

type InstrumentResponse struct {
	Instruments    []Instrument  `json:"instruments"`
	InstrumentTags []interface{} `json:"instrumentTags"`
}
