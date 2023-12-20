package stakego

import (
	"encoding/json"
	"time"
)

const OrderBUY = "BUY"
const OrderSELL = "SELL"
const OrderTypeLimit = "LIMIT"
const OrderTypeSTOP = "STOP"
const OrderValidityGoodTilDate = "GTD"
const OrderValidityGoodForDay = "GFD"

// NewOrderListFromJSON - creates a slice of OrderDetails from a JSON string
func NewOrderListFromJSON(jsonStr []byte) *[]OrderDetails {
	var o []OrderDetails
	err := json.Unmarshal(jsonStr, &o)
	if err != nil {
		return nil
	}
	return &o
}

// NewOrderResponseFromJSON - creates a slice of OrderResponse from a JSON string
func NewOrderResponseFromJSON(jsonStr []byte) *OrderResponse {
	var o OrderResponse
	err := json.Unmarshal(jsonStr, &o)
	if err != nil {
		return nil
	}
	return &o
}

// NewBuyOrder - create a new buy order
func NewBuyOrder() *Order {
	var o Order
	o.Side = OrderBUY
	o.Type = OrderTypeLimit
	o.Units = 0
	o.Price = 0.0
	o.Validity = OrderValidityGoodTilDate
	o.ValidityDate = time.Now().AddDate(0, 1, 0).Format("2006-01-02")
	o.AllowAwaitingTrigger = true
	return &o
}

// NewSellOrder - create a new sell order
func NewSellOrder() *Order {
	var o Order
	o.Side = OrderSELL
	o.Type = OrderTypeLimit
	o.Units = 0
	o.Price = 0.0
	o.Validity = OrderValidityGoodTilDate
	o.ValidityDate = time.Now().AddDate(0, 1, 0).Format("2006-01-02")
	o.AllowAwaitingTrigger = true
	return &o
}

// Order information
type Order struct {
	Side                 string  `json:"side"`
	Type                 string  `json:"type"`
	Units                int     `json:"units"`
	Price                float64 `json:"price,string"`
	Validity             string  `json:"validity"`
	ValidityDate         string  `json:"validityDate"`
	InstrumentCode       string  `json:"instrumentCode"`
	AllowAwaitingTrigger bool    `json:"allowAwaitingTrigger"`
}

// AsJSON - convert to json
func (o *Order) AsJSON() []byte {
	j, err := json.Marshal(o)
	if err != nil {
		return []byte("{}")
	}
	return j
}

// OrderDetails - holds information about existing or created orders
type OrderDetails struct {
	ID                         string      `json:"id"`
	Broker                     string      `json:"broker"`
	BrokerOrderID              interface{} `json:"brokerOrderId"`
	BrokerOrderVersionID       interface{} `json:"brokerOrderVersionId"`
	BrokerInstructionID        int         `json:"brokerInstructionId"`
	BrokerInstructionVersionID int         `json:"brokerInstructionVersionId"`
	UserID                     string      `json:"userId"`
	InstrumentID               interface{} `json:"instrumentId"`
	InstrumentCode             string      `json:"instrumentCode"`
	Side                       string      `json:"side"`
	LimitPrice                 float64     `json:"limitPrice"`
	TriggerPrice               float64     `json:"triggerPrice"`
	TrailingPercentage         interface{} `json:"trailingPercentage"`
	Validity                   string      `json:"validity"`
	ValidityDate               string      `json:"validityDate"`
	Type                       string      `json:"type"`
	PlacedTimestamp            string      `json:"placedTimestamp"`
	CompletedTimestamp         interface{} `json:"completedTimestamp"`
	ExpiresAt                  string      `json:"expiresAt"`
	OrderStatus                string      `json:"orderStatus"`
	OrderCompletionType        string      `json:"orderCompletionType"`
	FilledUnits                int         `json:"filledUnits"`
	AveragePrice               interface{} `json:"averagePrice"`
	UnitsRemaining             int         `json:"unitsRemaining"`
	UnitsRequested             int         `json:"unitsRequested"`
	EstimatedBrokerage         float64     `json:"estimatedBrokerage"`
	EstimatedExchangeFees      float64     `json:"estimatedExchangeFees"`
	CancellationEventSent      bool        `json:"cancellationEventSent"`
	BrokerageDiscount          float64     `json:"brokerageDiscount"`
	PendingBrokerage           float64     `json:"pendingBrokerage"`
	ChargedBrokerage           float64     `json:"chargedBrokerage"`
	FcPlacementAttempts        interface{} `json:"fcPlacementAttempts"`
	AllowAwaitingTrigger       bool        `json:"allowAwaitingTrigger"`
	CancellationReason         string      `json:"cancellationReason"`
	CancellationDetail         interface{} `json:"cancellationDetail"`
	CurrentExecutionPrice      interface{} `json:"currentExecutionPrice"`
	AnchorPrice                interface{} `json:"anchorPrice"`
	StakeManaged               bool        `json:"stakeManaged"`
}

// OrderResponse - holds full response when creating/deleting an order
type OrderResponse struct {
	Order OrderDetails `json:"order"`
}
