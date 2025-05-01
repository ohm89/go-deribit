package api

import (
	"fmt"
)

type TradeService struct {
	client *Client
}

const (
	urlPathGetUserTradesByInstrument = "/private/get_user_trades_by_instrument"
)

// Trade represents a trade from the Deribit API
type Trade struct {
	TradeID          string  `json:"trade_id"`
	TradeSeq         int64   `json:"trade_seq"`
	Timestamp        int64   `json:"timestamp"`
	TickDirection    int     `json:"tick_direction"`
	State            string  `json:"state"`
	SelfTrade        bool    `json:"self_trade"`
	ReduceOnly       bool    `json:"reduce_only"`
	PriceIndex       float64 `json:"price_index,omitempty"`
	PostOnly         bool    `json:"post_only"`
	OrderType        string  `json:"order_type"`
	OrderID          string  `json:"order_id"`
	MatchingID       string  `json:"matching_id"`
	MarkPrice        float64 `json:"mark_price,omitempty"`
	Liquidation      string  `json:"liquidation,omitempty"`
	IV               float64 `json:"iv,omitempty"`
	InstrumentName   string  `json:"instrument_name"`
	IndexPrice       float64 `json:"index_price,omitempty"`
	Fee              float64 `json:"fee"`
	Direction        string  `json:"direction"`
	Amount           float64 `json:"amount"`
	Price            float64 `json:"price"`
	UnderlyingPrice  float64 `json:"underlying_price,omitempty"`
	UnderlyingIndex  string  `json:"underlying_index,omitempty"`
	Label            string  `json:"label,omitempty"`
	FeeCurrency      string  `json:"fee_currency"`
	BlockTradeID     string  `json:"block_trade_id,omitempty"`
	API              bool    `json:"api"`
	AdvancedOptions  string  `json:"advanced,omitempty"`
	LiquidityRole    string  `json:"liquidity_role,omitempty"`
	ProfitLoss       float64 `json:"profit_loss,omitempty"`
	Triggered        bool    `json:"triggered,omitempty"`
	TriggerPrice     float64 `json:"trigger_price,omitempty"`
	TriggerDirection string  `json:"trigger_direction,omitempty"`
	ImpliedVolatility float64 `json:"implied_volatility,omitempty"`
}

type GetUserTradesByInstrumentParams struct {
	InstrumentName    string `json:"instrument_name"`
	StartSeq          int64  `json:"start_seq,omitempty"`
	EndSeq            int64  `json:"end_seq,omitempty"`
	Count             int    `json:"count,omitempty"`
	Historical        bool   `json:"historical,omitempty"`
	Sorting           string `json:"sorting,omitempty"`
	StartTimestamp    int64  `json:"start_timestamp,omitempty"`
	EndTimestamp      int64  `json:"end_timestamp,omitempty"`
}

type GetUserTradesByInstrumentResponse struct {
	ID      uint64  `json:"id"`
	JSONRPC string  `json:"jsonrpc"`
	Result  struct {
		Trades []Trade `json:"trades"`
		HasMore bool   `json:"has_more"`
	} `json:"result"`
	Error *ResponseError `json:"error,omitempty"`
}

// GetUserTradesByInstrument retrieves the user's trades by instrument
func (s *TradeService) GetUserTradesByInstrument(params GetUserTradesByInstrumentParams) (*GetUserTradesByInstrumentResponse, error) {
	var resp GetUserTradesByInstrumentResponse
	
	// Prepare request body
	requestBody := make(map[string]interface{})
	requestBody["instrument_name"] = params.InstrumentName
	
	if params.StartSeq != 0 {
		requestBody["start_seq"] = params.StartSeq
	}
	
	if params.EndSeq != 0 {
		requestBody["end_seq"] = params.EndSeq
	}
	
	if params.Count != 0 {
		requestBody["count"] = params.Count
	}
	
	requestBody["historical"] = params.Historical
	
	if params.Sorting != "" {
		requestBody["sorting"] = params.Sorting
	}
	
	if params.StartTimestamp != 0 {
		requestBody["start_timestamp"] = params.StartTimestamp
	}
	
	if params.EndTimestamp != 0 {
		requestBody["end_timestamp"] = params.EndTimestamp
	}
	
	uri := fmt.Sprintf("%s%s%s", 
		s.client.baseURL, 
		defaultAPIURL, 
		urlPathGetUserTradesByInstrument,
	)
	
	err := s.client.DoPrivateRPC(
		uri, 
		urlPathGetUserTradesByInstrument, 
		requestBody, 
		&resp,
	)
	
	if err != nil {
		return nil, err
	}
	
	
	return &resp, nil
}