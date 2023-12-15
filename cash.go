package stakego

import (
    "encoding/json"
)

// NewCashFromJSON - create a Cash item from a json byte slice
func NewCashFromJSON(jsonStr []byte) *Cash {
    var c Cash
    err := json.Unmarshal(jsonStr, &c)
    if err != nil {
        return nil
    }
    return &c
}

// Cash - stores the result of from a cash request
type Cash struct {
    SettledCash                    float64 `json:"settledCash"`
    PostedBalance                  float64 `json:"postedBalance"`
    TradeSettlement                float64 `json:"tradeSettlement"`
    BuyingPower                    float64 `json:"buyingPower"`
    PendingBuys                    float64 `json:"pendingBuys"`
    PendingBids                    float64 `json:"pendingBids"`
    SettlementHold                 float64 `json:"settlementHold"`
    PendingWithdrawals             float64 `json:"pendingWithdrawals"`
    CashAvailableForWithdrawal     float64 `json:"cashAvailableForWithdrawal"`
    CashAvailableForWithdrawalRaw  float64 `json:"cashAvailableForWithdrawalRaw"`
    CashAvailableForTransfer       float64 `json:"cashAvailableForTransfer"`
    CashAvailableForWithdrawalHold float64 `json:"cashAvailableForWithdrawalHold"`
    ClearingCash                   float64 `json:"clearingCash"`
}