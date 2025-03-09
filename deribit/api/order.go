package api

import (
	"fmt"
	"strings"
)

type OrderService struct {
	client *Client
}

const (
	urlPathBuy                         = "/private/buy"
	urlPathSell                        = "/private/sell"
	urlPathCancelOneOrder              = "/private/cancel"
	urlPathCancelAllOrder              = "/private/cancel_all"
	urlPathCancelAllByInstrument       = "/private/cancel_all_by_instrument"
	urlPathGetOrderState               = "/private/get_order_state"
	urlPathGetOrderStateByLabel        = "/private/get_order_state_by_label"
	urlPathGetOpenOrders               = "/private/get_open_orders"
	urlPathGetOpenOrdersByInstrument   = "/private/get_open_orders_by_instrument"
	urlPathGetOrderHistoryByCurrency   = "/private/get_order_history_by_currency"
	urlPathGetOrderHistoryByInstrument = "/private/get_order_history_by_instrument"
	urlPathGetTriggerOrderHistory      = "/private/get_trigger_order_history"
	urlPathGetMargins                  = "/private/get_margins"
)

type OTOCOConfig struct {
	Amount         float64 `json:"amount,omitempty"`
	Direction      string  `json:"direction"`
	Type           string  `json:"type,omitempty"`
	Label          string  `json:"label,omitempty"`
	Price          float64 `json:"price,omitempty"`
	ReduceOnly     bool    `json:"reduce_only,omitempty"`
	TimeInForce    string  `json:"time_in_force,omitempty"`
	PostOnly       bool    `json:"post_only,omitempty"`
	RejectPostOnly bool    `json:"reject_post_only,omitempty"`
	TriggerPrice   float64 `json:"trigger_price,omitempty"`
	TriggerOffset  float64 `json:"trigger_offset,omitempty"`
	Trigger        string  `json:"trigger,omitempty"`
}

type OrderRequest struct {
	InstrumentName       string        `json:"instrument_name"`
	Amount               float64       `json:"amount,omitempty"`
	Contracts            int64         `json:"contracts,omitempty"`
	Type                 string        `json:"type,omitempty"`
	Label                string        `json:"label,omitempty"`
	Price                float64       `json:"price,omitempty"`
	TimeInForce          string        `json:"time_in_force,omitempty"`
	MaxShow              int64         `json:"max_show,omitempty"`
	PostOnly             bool          `json:"post_only,omitempty"`
	RejectPostOnly       bool          `json:"reject_post_only,omitempty"`
	ReduceOnly           bool          `json:"reduce_only,omitempty"`
	TriggerPrice         float64       `json:"trigger_price,omitempty"`
	TriggerOffset        float64       `json:"trigger_offset,omitempty"`
	Trigger              string        `json:"trigger,omitempty"`
	Advanced             string        `json:"advanced,omitempty"`
	MMP                  bool          `json:"mmp,omitempty"`
	ValidUntil           int64         `json:"valid_until,omitempty"`
	LinkedOrderType      string        `json:"linked_order_type,omitempty"`
	TriggerFillCondition string        `json:"trigger_fill_condition,omitempty"`
	OTOCOConfig          []OTOCOConfig `json:"otoco_config,omitempty"`
}

type OrderResultOrderResponse struct {
	Quote                 bool     `json:"quote"`
	Triggered             bool     `json:"triggered"`
	Mobile                bool     `json:"mobile,omitempty"`
	AppName               string   `json:"app_name,omitempty"`
	Implv                 float64  `json:"implv,omitempty"`
	USD                   float64  `json:"usd,omitempty"`
	OtoOrderIds           []string `json:"oto_order_ids"`
	API                   bool     `json:"api"`
	AveragePrice          float64  `json:"average_price"`
	Advanced              string   `json:"advanced,omitempty"`
	OrderID               string   `json:"order_id"`
	PostOnly              bool     `json:"post_only"`
	FilledAmount          float64  `json:"filled_amount"`
	Trigger               string   `json:"trigger,omitempty"`
	TriggerOrderID        string   `json:"trigger_order_id,omitempty"`
	Direction             string   `json:"direction"`
	Contracts             float64  `json:"contracts,omitempty"`
	IsSecondaryOto        bool     `json:"is_secondary_oto,omitempty"`
	Replaced              bool     `json:"replaced"`
	MMPGroup              string   `json:"mmp_group,omitempty"`
	MMP                   bool     `json:"mmp"`
	LastUpdateTimestamp   int64    `json:"last_update_timestamp"`
	CreationTimestamp     int64    `json:"creation_timestamp"`
	CancelReason          string   `json:"cancel_reason,omitempty"`
	MMPCancelled          bool     `json:"mmp_cancelled,omitempty"`
	QuoteID               string   `json:"quote_id,omitempty"`
	OrderState            string   `json:"order_state"`
	IsRebalance           bool     `json:"is_rebalance,omitempty"`
	RejectPostOnly        bool     `json:"reject_post_only,omitempty"`
	Label                 string   `json:"label,omitempty"`
	IsLiquidation         bool     `json:"is_liquidation,omitempty"`
	Price                 float64  `json:"price"`
	Web                   bool     `json:"web,omitempty"`
	TimeInForce           string   `json:"time_in_force"`
	TriggerReferencePrice float64  `json:"trigger_reference_price,omitempty"`
	OrderType             string   `json:"order_type"`
	IsPrimaryOtoco        bool     `json:"is_primary_otoco,omitempty"`
	OriginalOrderType     string   `json:"original_order_type,omitempty"`
	BlockTrade            bool     `json:"block_trade,omitempty"`
	TriggerPrice          float64  `json:"trigger_price,omitempty"`
	OcoRef                string   `json:"oco_ref,omitempty"`
	TriggerOffset         float64  `json:"trigger_offset,omitempty"`
	QuoteSetID            string   `json:"quote_set_id,omitempty"`
	AutoReplaced          bool     `json:"auto_replaced,omitempty"`
	ReduceOnly            bool     `json:"reduce_only,omitempty"`
	MaxShow               float64  `json:"max_show,omitempty"`
	Amount                float64  `json:"amount"`
	RiskReducing          bool     `json:"risk_reducing,omitempty"`
	InstrumentName        string   `json:"instrument_name"`
	TriggerFillCondition  string   `json:"trigger_fill_condition,omitempty"`
	PrimaryOrderID        string   `json:"primary_order_id,omitempty"`
}

type OrderResultTradeResponse struct {
	TradeID         string                     `json:"trade_id"`
	TickDirection   int                        `json:"tick_direction"`
	FeeCurrency     string                     `json:"fee_currency"`
	API             bool                       `json:"api"`
	Advanced        string                     `json:"advanced,omitempty"`
	OrderID         string                     `json:"order_id"`
	Liquidity       string                     `json:"liquidity"`
	PostOnly        bool                       `json:"post_only"`
	Direction       string                     `json:"direction"`
	Contracts       float64                    `json:"contracts,omitempty"`
	MMP             bool                       `json:"mmp"`
	Fee             float64                    `json:"fee"`
	QuoteID         string                     `json:"quote_id,omitempty"`
	IndexPrice      float64                    `json:"index_price"`
	Label           string                     `json:"label,omitempty"`
	BlockTradeID    string                     `json:"block_trade_id,omitempty"`
	Price           float64                    `json:"price"`
	ComboID         string                     `json:"combo_id,omitempty"`
	MatchingID      string                     `json:"matching_id"`
	OrderType       string                     `json:"order_type"`
	ProfitLoss      float64                    `json:"profit_loss"`
	Timestamp       int64                      `json:"timestamp"`
	IV              float64                    `json:"iv,omitempty"`
	State           string                     `json:"state"`
	UnderlyingPrice float64                    `json:"underlying_price,omitempty"`
	QuoteSetID      string                     `json:"quote_set_id,omitempty"`
	MarkPrice       float64                    `json:"mark_price"`
	BlockRFQID      int                        `json:"block_rfq_id,omitempty"`
	ComboTradeID    int                        `json:"combo_trade_id,omitempty"`
	ReduceOnly      bool                       `json:"reduce_only"`
	Amount          float64                    `json:"amount"`
	Liquidation     string                     `json:"liquidation,omitempty"`
	TradeSeq        int                        `json:"trade_seq"`
	RiskReducing    bool                       `json:"risk_reducing"`
	InstrumentName  string                     `json:"instrument_name"`
	Legs            []OrderResultTradeResponse `json:"legs,omitempty"`
}

type OrderResultResponse struct {
	Order  OrderResultOrderResponse   `json:"order"`
	Trades []OrderResultTradeResponse `json:"trades"`
}

type OrderResponse struct {
	Id      uint64              `json:"id"`
	Jsonrpc string              `json:"jsonrpc"`
	Result  OrderResultResponse `json:"result"`
}

type CancelAllResponse struct {
	ID      uint64 `json:"id"`
	JSONRPC string `json:"jsonrpc"`
	Result  int    `json:"result"`
}

type CancelAllByInstrumentResponse struct {
	ID      uint64 `json:"id"`
	JSONRPC string `json:"jsonrpc"`
	Result  int    `json:"result"`
}

type OrderState struct {
	Quote                 bool     `json:"quote"`
	Triggered             bool     `json:"triggered"`
	Mobile                bool     `json:"mobile,omitempty"`
	AppName               string   `json:"app_name,omitempty"`
	Implv                 float64  `json:"implv,omitempty"`
	USD                   float64  `json:"usd,omitempty"`
	OtoOrderIDs           []string `json:"oto_order_ids"`
	API                   bool     `json:"api"`
	AveragePrice          float64  `json:"average_price"`
	Advanced              string   `json:"advanced,omitempty"`
	OrderID               string   `json:"order_id"`
	PostOnly              bool     `json:"post_only"`
	FilledAmount          float64  `json:"filled_amount"`
	Trigger               string   `json:"trigger,omitempty"`
	TriggerOrderID        string   `json:"trigger_order_id,omitempty"`
	Direction             string   `json:"direction"`
	Contracts             float64  `json:"contracts,omitempty"`
	IsSecondaryOto        bool     `json:"is_secondary_oto,omitempty"`
	Replaced              bool     `json:"replaced"`
	MMPGroup              string   `json:"mmp_group,omitempty"`
	MMP                   bool     `json:"mmp"`
	LastUpdateTimestamp   int64    `json:"last_update_timestamp"`
	CreationTimestamp     int64    `json:"creation_timestamp"`
	CancelReason          string   `json:"cancel_reason,omitempty"`
	MMPCancelled          bool     `json:"mmp_cancelled,omitempty"`
	QuoteID               string   `json:"quote_id,omitempty"`
	OrderState            string   `json:"order_state"`
	IsRebalance           bool     `json:"is_rebalance,omitempty"`
	RejectPostOnly        bool     `json:"reject_post_only,omitempty"`
	Label                 string   `json:"label,omitempty"`
	IsLiquidation         bool     `json:"is_liquidation,omitempty"`
	Price                 float64  `json:"price"`
	Web                   bool     `json:"web,omitempty"`
	TimeInForce           string   `json:"time_in_force"`
	TriggerReferencePrice float64  `json:"trigger_reference_price,omitempty"`
	OrderType             string   `json:"order_type"`
	IsPrimaryOtoco        bool     `json:"is_primary_otoco,omitempty"`
	OriginalOrderType     string   `json:"original_order_type,omitempty"`
	BlockTrade            bool     `json:"block_trade,omitempty"`
	TriggerPrice          float64  `json:"trigger_price,omitempty"`
	OcoRef                string   `json:"oco_ref,omitempty"`
	TriggerOffset         float64  `json:"trigger_offset,omitempty"`
	QuoteSetID            string   `json:"quote_set_id,omitempty"`
	AutoReplaced          bool     `json:"auto_replaced,omitempty"`
	ReduceOnly            bool     `json:"reduce_only,omitempty"`
	MaxShow               float64  `json:"max_show,omitempty"`
	Amount                float64  `json:"amount"`
	RiskReducing          bool     `json:"risk_reducing,omitempty"`
	InstrumentName        string   `json:"instrument_name"`
	TriggerFillCondition  string   `json:"trigger_fill_condition,omitempty"`
	PrimaryOrderID        string   `json:"primary_order_id,omitempty"`
}

type GetOrderStateResponse struct {
	ID      uint64     `json:"id"`
	JSONRPC string     `json:"jsonrpc"`
	Result  OrderState `json:"result"`
}

type GetOrderStateByLabelResponse struct {
	ID      uint64       `json:"id"`
	JSONRPC string       `json:"jsonrpc"`
	Result  []OrderState `json:"result"`
}

type GetOpenOrdersResponse struct {
	ID      uint64       `json:"id"`
	JSONRPC string       `json:"jsonrpc"`
	Result  []OrderState `json:"result"`
}

type GetOpenOrdersByInstrumentResponse struct {
	ID      uint64       `json:"id"`
	JSONRPC string       `json:"jsonrpc"`
	Result  []OrderState `json:"result"`
}

type GetOrderHistoryByCurrencyResponse struct {
	ID      uint64       `json:"id"`
	JSONRPC string       `json:"jsonrpc"`
	Result  []OrderState `json:"result"`
}

type GetOrderHistoryByInstrumentResponse struct {
	ID      uint64       `json:"id"`
	JSONRPC string       `json:"jsonrpc"`
	Result  []OrderState `json:"result"`
}

type TriggerOrderHistoryEntry struct {
	Amount              float64 `json:"amount"`
	Direction           string  `json:"direction"`
	InstrumentName      string  `json:"instrument_name"`
	IsSecondaryOto      bool    `json:"is_secondary_oto"`
	Label               string  `json:"label"`
	LastUpdateTimestamp int64   `json:"last_update_timestamp"`
	OcoRef              string  `json:"oco_ref"`
	OrderID             string  `json:"order_id"`
	OrderState          string  `json:"order_state"`
	OrderType           string  `json:"order_type"`
	PostOnly            bool    `json:"post_only"`
	Price               float64 `json:"price"`
	ReduceOnly          bool    `json:"reduce_only"`
	Request             string  `json:"request"`
	Source              string  `json:"source"`
	Timestamp           int64   `json:"timestamp"`
	Trigger             string  `json:"trigger"`
	TriggerOffset       float64 `json:"trigger_offset"`
	TriggerOrderID      string  `json:"trigger_order_id"`
	TriggerPrice        float64 `json:"trigger_price"`
}

type TriggerOrderHistoryResult struct {
	Continuation string                     `json:"continuation"`
	Entries      []TriggerOrderHistoryEntry `json:"entries"`
}

type GetTriggerOrderHistoryResponse struct {
	ID      uint64                     `json:"id"`
	JSONRPC string                     `json:"jsonrpc"`
	Result  *TriggerOrderHistoryResult `json:"result"`
}

// ## --------------------------------------------------------------------------

// ## Create Buy Order
func (s *OrderService) Buy(
	instrumentName string,
	amount float64,
	contracts int64,
	orderType string,
	label string,
	price float64,
	timeInForce string,
	maxShow int64,
	postOnly bool,
	rejectPostOnly bool,
	reduceOnly bool,
	triggerPrice float64,
	triggerOffset float64,
	trigger string,
	advanced string,
	mmp bool,
	validUntil int64,
	linkedOrderType string,
	triggerFillCondition string,
	otocoConfig []OTOCOConfig,
) (*OrderResponse, error) {
	var resp OrderResponse
	uri := fmt.Sprintf("%s%s%s", s.client.baseURL, defaultAPIURL, urlPathBuy)
	queryParams := make([]string, 0, 20)

	if instrumentName != "" {
		queryParams = append(queryParams, fmt.Sprintf("instrument_name=%s", instrumentName))
	}
	if amount != 0 {
		queryParams = append(queryParams, fmt.Sprintf("amount=%f", amount))
	}
	if contracts != 0 {
		queryParams = append(queryParams, fmt.Sprintf("contracts=%d", contracts))
	}
	if orderType != "" {
		queryParams = append(queryParams, fmt.Sprintf("type=%s", orderType))
	}
	if label != "" {
		queryParams = append(queryParams, fmt.Sprintf("label=%s", label))
	}
	if price != 0 {
		queryParams = append(queryParams, fmt.Sprintf("price=%f", price))
	}
	if timeInForce != "" {
		queryParams = append(queryParams, fmt.Sprintf("time_in_force=%s", timeInForce))
	}
	if maxShow != 0 {
		queryParams = append(queryParams, fmt.Sprintf("max_show=%d", maxShow))
	}
	if postOnly {
		queryParams = append(queryParams, "post_only=true")
	}
	if rejectPostOnly {
		queryParams = append(queryParams, "reject_post_only=true")
	}
	if reduceOnly {
		queryParams = append(queryParams, "reduce_only=true")
	}
	if triggerPrice != 0 {
		queryParams = append(queryParams, fmt.Sprintf("trigger_price=%f", triggerPrice))
	}
	if triggerOffset != 0 {
		queryParams = append(queryParams, fmt.Sprintf("trigger_offset=%f", triggerOffset))
	}
	if trigger != "" {
		queryParams = append(queryParams, fmt.Sprintf("trigger=%s", trigger))
	}
	if advanced != "" {
		queryParams = append(queryParams, fmt.Sprintf("advanced=%s", advanced))
	}
	if mmp {
		queryParams = append(queryParams, "mmp=true")
	}
	if validUntil != 0 {
		queryParams = append(queryParams, fmt.Sprintf("valid_until=%d", validUntil))
	}
	if linkedOrderType != "" {
		queryParams = append(queryParams, fmt.Sprintf("linked_order_type=%s", linkedOrderType))
	}
	if triggerFillCondition != "" {
		queryParams = append(queryParams, fmt.Sprintf("trigger_fill_condition=%s", triggerFillCondition))
	}

	if len(otocoConfig) > 0 {
		for _, config := range otocoConfig {
			if config.Amount != 0 {
				queryParams = append(queryParams, fmt.Sprintf("otoco_config[]=amount:%f", config.Amount))
			}
			if config.Direction != "" {
				queryParams = append(queryParams, fmt.Sprintf("otoco_config[]=direction:%s", config.Direction))
			}
			if config.Type != "" {
				queryParams = append(queryParams, fmt.Sprintf("otoco_config[]=type:%s", config.Type))
			}
			if config.Label != "" {
				queryParams = append(queryParams, fmt.Sprintf("otoco_config[]=label:%s", config.Label))
			}
			if config.Price != 0 {
				queryParams = append(queryParams, fmt.Sprintf("otoco_config[]=price:%f", config.Price))
			}
			if config.ReduceOnly {
				queryParams = append(queryParams, "otoco_config[]=reduce_only:true")
			}
			if config.TimeInForce != "" {
				queryParams = append(queryParams, fmt.Sprintf("otoco_config[]=time_in_force:%s", config.TimeInForce))
			}
			if config.PostOnly {
				queryParams = append(queryParams, "otoco_config[]=post_only:true")
			}
			if config.RejectPostOnly {
				queryParams = append(queryParams, "otoco_config[]=reject_post_only:true")
			}
			if config.TriggerPrice != 0 {
				queryParams = append(queryParams, fmt.Sprintf("otoco_config[]=trigger_price:%f", config.TriggerPrice))
			}
			if config.TriggerOffset != 0 {
				queryParams = append(queryParams, fmt.Sprintf("otoco_config[]=trigger_offset:%f", config.TriggerOffset))
			}
			if config.Trigger != "" {
				queryParams = append(queryParams, fmt.Sprintf("otoco_config[]=trigger:%s", config.Trigger))
			}
		}
	}

	uri += "?" + strings.Join(queryParams, "&")

	// ## [DEBUG]
	// fmt.Printf("buy uri: %s \n", uri)

	err := s.client.DoPrivate(uri, "GET", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ## Create PostBuy Order
func (s *OrderService) PostBuy(
	instrumentName string,
	amount float64,
	contracts int64,
	orderType string,
	label string,
	price float64,
	timeInForce string,
	maxShow int64,
	postOnly bool,
	rejectPostOnly bool,
	reduceOnly bool,
	triggerPrice float64,
	triggerOffset float64,
	trigger string,
	advanced string,
	mmp bool,
	validUntil int64,
	linkedOrderType string,
	triggerFillCondition string,
	otocoConfig []OTOCOConfig,
) (*OrderResponse, error) {
	var resp OrderResponse
	uri := fmt.Sprintf("%s%s%s", s.client.baseURL, defaultAPIURL, urlPathBuy)

	orderRequest := OrderRequest{
		InstrumentName:       instrumentName,
		Amount:               amount,
		Contracts:            contracts,
		Type:                 orderType,
		Label:                label,
		Price:                price,
		TimeInForce:          timeInForce,
		MaxShow:              maxShow,
		PostOnly:             postOnly,
		RejectPostOnly:       rejectPostOnly,
		ReduceOnly:           reduceOnly,
		TriggerPrice:         triggerPrice,
		TriggerOffset:        triggerOffset,
		Trigger:              trigger,
		Advanced:             advanced,
		MMP:                  mmp,
		ValidUntil:           validUntil,
		LinkedOrderType:      linkedOrderType,
		TriggerFillCondition: triggerFillCondition,
		OTOCOConfig:          otocoConfig,
	}

	err := s.client.DoPrivate(uri, "POST", orderRequest, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ## Create Sell Order
func (s *OrderService) Sell(
	instrumentName string,
	amount float64,
	contracts int64,
	orderType string,
	label string,
	price float64,
	timeInForce string,
	maxShow int64,
	postOnly bool,
	rejectPostOnly bool,
	reduceOnly bool,
	triggerPrice float64,
	triggerOffset float64,
	trigger string,
	advanced string,
	mmp bool,
	validUntil int64,
	linkedOrderType string,
	triggerFillCondition string,
	otocoConfig []OTOCOConfig,
) (*OrderResponse, error) {
	var resp OrderResponse
	uri := fmt.Sprintf("%s%s%s", s.client.baseURL, defaultAPIURL, urlPathSell)
	queryParams := make([]string, 0, 20)

	if instrumentName != "" {
		queryParams = append(queryParams, fmt.Sprintf("instrument_name=%s", instrumentName))
	}
	if amount != 0 {
		queryParams = append(queryParams, fmt.Sprintf("amount=%f", amount))
	}
	if contracts != 0 {
		queryParams = append(queryParams, fmt.Sprintf("contracts=%d", contracts))
	}
	if orderType != "" {
		queryParams = append(queryParams, fmt.Sprintf("type=%s", orderType))
	}
	if label != "" {
		queryParams = append(queryParams, fmt.Sprintf("label=%s", label))
	}
	if price != 0 {
		queryParams = append(queryParams, fmt.Sprintf("price=%f", price))
	}
	if timeInForce != "" {
		queryParams = append(queryParams, fmt.Sprintf("time_in_force=%s", timeInForce))
	}
	if maxShow != 0 {
		queryParams = append(queryParams, fmt.Sprintf("max_show=%d", maxShow))
	}
	if postOnly {
		queryParams = append(queryParams, "post_only=true")
	}
	if rejectPostOnly {
		queryParams = append(queryParams, "reject_post_only=true")
	}
	if reduceOnly {
		queryParams = append(queryParams, "reduce_only=true")
	}
	if triggerPrice != 0 {
		queryParams = append(queryParams, fmt.Sprintf("trigger_price=%f", triggerPrice))
	}
	if triggerOffset != 0 {
		queryParams = append(queryParams, fmt.Sprintf("trigger_offset=%f", triggerOffset))
	}
	if trigger != "" {
		queryParams = append(queryParams, fmt.Sprintf("trigger=%s", trigger))
	}
	if advanced != "" {
		queryParams = append(queryParams, fmt.Sprintf("advanced=%s", advanced))
	}
	if mmp {
		queryParams = append(queryParams, "mmp=true")
	}
	if validUntil != 0 {
		queryParams = append(queryParams, fmt.Sprintf("valid_until=%d", validUntil))
	}
	if linkedOrderType != "" {
		queryParams = append(queryParams, fmt.Sprintf("linked_order_type=%s", linkedOrderType))
	}
	if triggerFillCondition != "" {
		queryParams = append(queryParams, fmt.Sprintf("trigger_fill_condition=%s", triggerFillCondition))
	}

	if len(otocoConfig) > 0 {
		for _, config := range otocoConfig {
			if config.Amount != 0 {
				queryParams = append(queryParams, fmt.Sprintf("otoco_config[]=amount:%f", config.Amount))
			}
			if config.Direction != "" {
				queryParams = append(queryParams, fmt.Sprintf("otoco_config[]=direction:%s", config.Direction))
			}
			if config.Type != "" {
				queryParams = append(queryParams, fmt.Sprintf("otoco_config[]=type:%s", config.Type))
			}
			if config.Label != "" {
				queryParams = append(queryParams, fmt.Sprintf("otoco_config[]=label:%s", config.Label))
			}
			if config.Price != 0 {
				queryParams = append(queryParams, fmt.Sprintf("otoco_config[]=price:%f", config.Price))
			}
			if config.ReduceOnly {
				queryParams = append(queryParams, "otoco_config[]=reduce_only:true")
			}
			if config.TimeInForce != "" {
				queryParams = append(queryParams, fmt.Sprintf("otoco_config[]=time_in_force:%s", config.TimeInForce))
			}
			if config.PostOnly {
				queryParams = append(queryParams, "otoco_config[]=post_only:true")
			}
			if config.RejectPostOnly {
				queryParams = append(queryParams, "otoco_config[]=reject_post_only:true")
			}
			if config.TriggerPrice != 0 {
				queryParams = append(queryParams, fmt.Sprintf("otoco_config[]=trigger_price:%f", config.TriggerPrice))
			}
			if config.TriggerOffset != 0 {
				queryParams = append(queryParams, fmt.Sprintf("otoco_config[]=trigger_offset:%f", config.TriggerOffset))
			}
			if config.Trigger != "" {
				queryParams = append(queryParams, fmt.Sprintf("otoco_config[]=trigger:%s", config.Trigger))
			}
		}
	}

	uri += "?" + strings.Join(queryParams, "&")

	// ## [DEBUG]
	// fmt.Printf("Sell uri: %s \n", uri)

	err := s.client.DoPrivate(uri, "GET", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ## Create PostSell Order
func (s *OrderService) PostSell(
	instrumentName string,
	amount float64,
	contracts int64,
	orderType string,
	label string,
	price float64,
	timeInForce string,
	maxShow int64,
	postOnly bool,
	rejectPostOnly bool,
	reduceOnly bool,
	triggerPrice float64,
	triggerOffset float64,
	trigger string,
	advanced string,
	mmp bool,
	validUntil int64,
	linkedOrderType string,
	triggerFillCondition string,
	otocoConfig []OTOCOConfig,
) (*OrderResponse, error) {
	var resp OrderResponse
	uri := fmt.Sprintf("%s%s%s", s.client.baseURL, defaultAPIURL, urlPathSell)

	orderRequest := OrderRequest{
		InstrumentName:       instrumentName,
		Amount:               amount,
		Contracts:            contracts,
		Type:                 orderType,
		Label:                label,
		Price:                price,
		TimeInForce:          timeInForce,
		MaxShow:              maxShow,
		PostOnly:             postOnly,
		RejectPostOnly:       rejectPostOnly,
		ReduceOnly:           reduceOnly,
		TriggerPrice:         triggerPrice,
		TriggerOffset:        triggerOffset,
		Trigger:              trigger,
		Advanced:             advanced,
		MMP:                  mmp,
		ValidUntil:           validUntil,
		LinkedOrderType:      linkedOrderType,
		TriggerFillCondition: triggerFillCondition,
		OTOCOConfig:          otocoConfig,
	}

	err := s.client.DoPrivate(uri, "POST", orderRequest, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ## Cancel One Order By ID
func (s *OrderService) Cancel(orderID string) (*OrderResponse, error) {
	var resp OrderResponse
	uri := fmt.Sprintf("%s%s%s?order_id=%s",
		s.client.baseURL,
		defaultAPIURL,
		urlPathCancelOneOrder,
		orderID,
	)
	err := s.client.DoPrivate(uri, "GET", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ## Cancel All Open Order
func (s *OrderService) CancelAll() (*CancelAllResponse, error) {
	var resp CancelAllResponse
	uri := fmt.Sprintf("%s%s%s", s.client.baseURL, defaultAPIURL, urlPathCancelAllOrder)
	err := s.client.DoPrivate(uri, "GET", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ## Cancel All By Instrument
func (s *OrderService) CancelAllByInstrument(
	instrumentName string,
	orderType string,
	detailed bool,
	includeCombos bool,
	freezeQuotes bool,
) (*CancelAllByInstrumentResponse, error) {
	var resp CancelAllByInstrumentResponse
	uri := fmt.Sprintf(
		"%s%s%s?instrument_name=%s",
		s.client.baseURL,
		defaultAPIURL,
		urlPathCancelAllByInstrument,
		instrumentName,
	)

	if orderType != "" {
		uri += fmt.Sprintf("&type=%s", orderType)
	}
	if detailed {
		uri += "&detailed=true"
	}
	if includeCombos {
		uri += "&include_combos=true"
	}
	if freezeQuotes {
		uri += "&freeze_quotes=true"
	}

	err := s.client.DoPrivate(uri, "GET", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ## Get Order State by order_id
func (s *OrderService) GetOrderState(orderID string) (*GetOrderStateResponse, error) {
	var resp GetOrderStateResponse
	uri := fmt.Sprintf(
		"%s%s%s?order_id=%s",
		s.client.baseURL,
		defaultAPIURL,
		urlPathGetOrderState,
		orderID,
	)
	err := s.client.DoPrivate(uri, "GET", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ## Get Order State by order label from api
func (s *OrderService) GetOrderStateByLabel(currency, label string) (*GetOrderStateByLabelResponse, error) {
	var resp GetOrderStateByLabelResponse
	uri := fmt.Sprintf(
		"%s%s%s?currency=%s&label=%s",
		s.client.baseURL,
		defaultAPIURL,
		urlPathGetOrderStateByLabel,
		currency,
		label,
	)
	err := s.client.DoPrivate(uri, "GET", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ## Get Open Orders
func (s *OrderService) GetOpenOrders(kind, orderType string) (*GetOpenOrdersResponse, error) {
	var resp GetOpenOrdersResponse
	uri := fmt.Sprintf("%s%s%s",
		s.client.baseURL,
		defaultAPIURL,
		urlPathGetOpenOrders,
	)

	queryParams := make([]string, 0, 2)
	if kind != "" {
		queryParams = append(queryParams, fmt.Sprintf("kind=%s", kind))
	}
	if orderType != "" {
		queryParams = append(queryParams, fmt.Sprintf("type=%s", orderType))
	}

	if len(queryParams) > 0 {
		uri += "?" + strings.Join(queryParams, "&")
	}

	err := s.client.DoPrivate(uri, "GET", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ## Get Open Orders by Instrument
func (s *OrderService) GetOpenOrdersByInstrument(instrumentName, orderType string) (*GetOpenOrdersByInstrumentResponse, error) {
	var resp GetOpenOrdersByInstrumentResponse
	uri := fmt.Sprintf(
		"%s%s%s?instrument_name=%s",
		s.client.baseURL,
		defaultAPIURL,
		urlPathGetOpenOrdersByInstrument,
		instrumentName,
	)

	if orderType != "" {
		uri += fmt.Sprintf("&type=%s", orderType)
	}

	err := s.client.DoPrivate(uri, "GET", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ## Get History Orders By Currency
func (s *OrderService) GetOrderHistoryByCurrency(
	currency string,
	kind string,
	count int,
	offset int,
	includeOld bool,
	includeUnfilled bool,
) (*GetOrderHistoryByCurrencyResponse, error) {
	var resp GetOrderHistoryByCurrencyResponse
	uri := fmt.Sprintf(
		"%s%s%s?currency=%s",
		s.client.baseURL,
		defaultAPIURL,
		urlPathGetOrderHistoryByCurrency,
		currency,
	)

	if kind != "" {
		uri += fmt.Sprintf("&kind=%s", kind)
	}
	if count > 0 {
		uri += fmt.Sprintf("&count=%d", count)
	}
	if offset > 0 {
		uri += fmt.Sprintf("&offset=%d", offset)
	}
	if includeOld {
		uri += "&include_old=true"
	}
	if includeUnfilled {
		uri += "&include_unfilled=true"
	}

	err := s.client.DoPrivate(uri, "GET", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ## Get Order History By Instrument
func (s *OrderService) GetOrderHistoryByInstrument(
	instrumentName string,
	count int,
	offset int,
	includeOld bool,
	includeUnfilled bool,
) (*GetOrderHistoryByInstrumentResponse, error) {
	var resp GetOrderHistoryByInstrumentResponse
	uri := fmt.Sprintf(
		"%s%s%s?instrument_name=%s",
		s.client.baseURL,
		defaultAPIURL,
		urlPathGetOrderHistoryByInstrument,
		instrumentName,
	)

	if count > 0 {
		uri += fmt.Sprintf("&count=%d", count)
	}
	if offset > 0 {
		uri += fmt.Sprintf("&offset=%d", offset)
	}
	if includeOld {
		uri += "&include_old=true"
	}
	if includeUnfilled {
		uri += "&include_unfilled=true"
	}

	err := s.client.DoPrivate(uri, "GET", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ## Get Trigger Order History
func (s *OrderService) GetTriggerOrderHistory(
	currency string,
	instrumentName string,
	count int,
	continuation string,
) (*GetTriggerOrderHistoryResponse, error) {
	var resp GetTriggerOrderHistoryResponse
	uri := fmt.Sprintf(
		"%s%s%s?currency=%s",
		s.client.baseURL,
		defaultAPIURL,
		urlPathGetTriggerOrderHistory,
		currency,
	)

	if instrumentName != "" {
		uri += fmt.Sprintf("&instrument_name=%s", instrumentName)
	}
	if count > 0 {
		uri += fmt.Sprintf("&count=%d", count)
	}
	if continuation != "" {
		uri += fmt.Sprintf("&continuation=%s", continuation)
	}

	err := s.client.DoPrivate(uri, "GET", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type MarginResult struct {
	Buy      float64 `json:"buy"`
	MaxPrice float64 `json:"max_price"`
	MinPrice float64 `json:"min_price"`
	Sell     float64 `json:"sell"`
}

type GetMarginsResponse struct {
	ID      uint64       `json:"id"`
	JSONRPC string       `json:"jsonrpc"`
	Result  MarginResult `json:"result"`
}

// ## Get Margins
func (s *OrderService) GetMargins(
	instrumentName string,
	amount float64,
	price float64,
) (*GetMarginsResponse, error) {
	var resp GetMarginsResponse
	uri := fmt.Sprintf(
		"%s%s%s?instrument_name=%s&amount=%f&price=%f",
		s.client.baseURL,
		defaultAPIURL,
		urlPathGetMargins,
		instrumentName,
		amount,
		price,
	)

	err := s.client.DoPrivate(uri, "GET", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
