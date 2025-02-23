package api

type OTOCOConfig struct {
	Amount         float64 `json:"amount"`
	Direction      string  `json:"direction"`
	Type           string  `json:"type"`
	Label          string  `json:"label"`
	Price          float64 `json:"price"`
	ReduceOnly     bool    `json:reduce_only`
	TimeInForce    string  `json:time_in_force`
	PostOnly       bool    `json:post_only`
	RejectPostOnly bool    `json:reject_post_only`
	TriggerPrice   float64 `json:trigger_price`
	TriggerOffset  float64 `json:trigger_offset`
	Trigger        string  `json:trigger`
}

type OrderRequest struct {
	InstrumentName       string        `json:instrument_name`
	Amount               float64       `json:"amount"`
	Contracts            float64       `json:contracts`
	Type                 string        `json:type`
	Label                string        `json:label`
	Price                float64       `json:price`
	TimeInForce          string        `json:time_in_force`
	MaxShow              float64       `json:max_show`
	PostOnly             bool          `json:post_only`
	RejectPostOnly       bool          `json:reject_post_only`
	ReduceOnly           bool          `json:reduce_only`
	TriggerPrice         float64       `json:trigger_price`
	TriggerOffset        float64       `json:trigger_offset`
	Trigger              string        `json:trigger`
	Advanced             string        `json:advanced`
	Mmp                  bool          `json:mmp`
	ValidUntil           bool          `json:valid_until`
	LinkedOrderType      string        `json:linked_order_type`
	TriggerFillCondition string        `json:trigger_fill_condition`
	OtocoConfig          []OTOCOConfig `json:otoco_config`
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
	PostOnly        string                     `json:"post_only"`
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
	ReduceOnly      string                     `json:"reduce_only"`
	Amount          float64                    `json:"amount"`
	Liquidation     string                     `json:"liquidation,omitempty"`
	TradeSeq        int                        `json:"trade_seq"`
	RiskReducing    bool                       `json:"risk_reducing"`
	InstrumentName  string                     `json:"instrument_name"`
	Legs            []OrderResultTradeResponse `json:"legs,omitempty"`
}

type OrderResultResponse struct {
	Order OrderResultOrderResponse `json:order`
	Trade OrderResultTradeResponse `json:trade`
}

type OrderResponse struct {
	Id      uint64              `json:"id"`
	Jsonrpc string              `json:jsonrpc`
	Result  OrderResultResponse `json:result`
}
