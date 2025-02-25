package ws

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

type OpenOrder struct {
	Quote                 bool        `json:"quote"`
	Triggered             bool        `json:"triggered"`
	Mobile                bool        `json:"mobile"`
	AppName               string      `json:"app_name"`
	Implv                 float64     `json:"implv"`
	USD                   float64     `json:"usd"`
	OtoOrderIDs           []string    `json:"oto_order_ids"`
	API                   bool        `json:"api"`
	AveragePrice          float64     `json:"average_price"`
	Advanced              string      `json:"advanced"`
	OrderID               string      `json:"order_id"`
	PostOnly              bool        `json:"post_only"`
	FilledAmount          float64     `json:"filled_amount"`
	Trigger               string      `json:"trigger"`
	TriggerOrderID        string      `json:"trigger_order_id"`
	Direction             string      `json:"direction"`
	Contracts             float64     `json:"contracts"`
	IsSecondaryOto        bool        `json:"is_secondary_oto"`
	Replaced              bool        `json:"replaced"`
	MmpGroup              string      `json:"mmp_group"`
	Mmp                   bool        `json:"mmp"`
	LastUpdateTimestamp   int64       `json:"last_update_timestamp"`
	CreationTimestamp     int64       `json:"creation_timestamp"`
	CancelReason          string      `json:"cancel_reason"`
	MmpCancelled          bool        `json:"mmp_cancelled"`
	QuoteID               string      `json:"quote_id"`
	OrderState            string      `json:"order_state"`
	IsRebalance           bool        `json:"is_rebalance"`
	RejectPostOnly        bool        `json:"reject_post_only"`
	Label                 string      `json:"label"`
	IsLiquidation         bool        `json:"is_liquidation"`
	Price                 interface{} `json:"price"`
	Web                   bool        `json:"web"`
	TimeInForce           string      `json:"time_in_force"`
	TriggerReferencePrice float64     `json:"trigger_reference_price"`
	OrderType             string      `json:"order_type"`
	IsPrimaryOtoCo        bool        `json:"is_primary_otoco"`
	OriginalOrderType     string      `json:"original_order_type"`
	BlockTrade            bool        `json:"block_trade"`
	TriggerPrice          float64     `json:"trigger_price"`
	OcoRef                string      `json:"oco_ref"`
	TriggerOffset         float64     `json:"trigger_offset"`
	QuoteSetID            string      `json:"quote_set_id"`
	AutoReplaced          bool        `json:"auto_replaced"`
	ReduceOnly            bool        `json:"reduce_only"`
	MaxShow               float64     `json:"max_show"`
	Amount                float64     `json:"amount"`
	RiskReducing          bool        `json:"risk_reducing"`
	InstrumentName        string      `json:"instrument_name"`
	TriggerFillCondition  string      `json:"trigger_fill_condition"`
	PrimaryOrderID        string      `json:"primary_order_id"`
}

type Position struct {
	AveragePrice              float64 `json:"average_price"`
	AveragePriceUSD           float64 `json:"average_price_usd"`
	Delta                     float64 `json:"delta"`
	Direction                 string  `json:"direction"`
	EstimatedLiquidationPrice float64 `json:"estimated_liquidation_price"`
	FloatingProfitLoss        float64 `json:"floating_profit_loss"`
	FloatingProfitLossUSD     float64 `json:"floating_profit_loss_usd"`
	Gamma                     float64 `json:"gamma"`
	IndexPrice                float64 `json:"index_price"`
	InitialMargin             float64 `json:"initial_margin"`
	InstrumentName            string  `json:"instrument_name"`
	InterestValue             float64 `json:"interest_value"`
	Kind                      string  `json:"kind"`
	Leverage                  int     `json:"leverage"`
	MaintenanceMargin         float64 `json:"maintenance_margin"`
	MarkPrice                 float64 `json:"mark_price"`
	OpenOrdersMargin          float64 `json:"open_orders_margin"`
	RealizedFunding           float64 `json:"realized_funding"`
	RealizedProfitLoss        float64 `json:"realized_profit_loss"`
	SettlementPrice           float64 `json:"settlement_price"`
	Size                      float64 `json:"size"`
	SizeCurrency              float64 `json:"size_currency"`
	Theta                     float64 `json:"theta"`
	TotalProfitLoss           float64 `json:"total_profit_loss"`
	Vega                      float64 `json:"vega"`
}

type PositionsResponse struct {
	ID      int        `json:"id"`
	JSONRPC string     `json:"jsonrpc"`
	Result  []Position `json:"result"`
}

type PositionResponse struct {
	ID      int      `json:"id"`
	JSONRPC string   `json:"jsonrpc"`
	Result  Position `json:"result"`
}

// ## Get All positions list in account
/**
* currency - BTC/ETH/SOL/USDC/USDT
* kind - future/option/spot/future_combo/option_combo
* subaccountId - account id in integer
 */
func GetPositions(client *DeribitClient, currency string, kind string) error {
	// Create the request message
	msg := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "private/get_positions",
		"params": map[string]interface{}{
			"currency": currency,
			"kind":     kind,
		},
	}

	// Marshal the request message to JSON
	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal request GetPositions message: %w", err)
	}

	// Send the request over the WebSocket
	err = client.conn.WriteMessage(websocket.TextMessage, jsonMsg)
	if err != nil {
		return fmt.Errorf("failed to send GetPositions request: %w", err)
	}

	return nil
}

// ## Get One position list in account
/**
* instrument_name - BTC_USD/BTC-PERPEPTUAL
 */
func GetPosition(client *DeribitClient, instrument_name string) error {
	// Create the request message
	msg := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "private/get_position",
		"params": map[string]interface{}{
			"instrument_name": instrument_name,
		},
	}

	// Marshal the request message to JSON
	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal request GetPosition message: %w", err)
	}

	// Send the request over the WebSocket
	err = client.conn.WriteMessage(websocket.TextMessage, jsonMsg)
	if err != nil {
		return fmt.Errorf("failed to send GetPosition request: %w", err)
	}

	return nil
}
