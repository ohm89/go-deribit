package api

import "fmt"

type MarketService struct {
	client *Client
}

const (
	urlPathGetFundingChartData             = "/public/get_funding_chart_data"
	urlPathGetFundingRateHistory           = "/public/get_funding_rate_history"
	urlPathGetFundingRateValue             = "/public/get_funding_rate_value"
	urlPathGetHistoricalVolatility         = "/public/get_historical_volatility"
	urlPathGetIndexPrice                   = "/public/get_index_price"
	urlPathGetIndexPriceNames              = "/public/get_index_price_names"
	urlPathGetInstrument                   = "/public/get_instrument"
	urlPathGetInstruments                  = "/public/get_instruments"
	urlPathGetLastSettlementsByInstrument  = "/public/get_last_settlements_by_instrument"
	urlPathGetLastTradeByCurrencyAndTime   = "/public/get_last_trades_by_currency_and_time"
	urlPathGetLastTradeByInstrument        = "/public/get_last_trades_by_instrument"
	urlPathGetLastTradeByInstrumentAndTime = "/public/get_last_trades_by_instrument_and_time"
	urlPathGetMarkPriceHistory             = "/public/get_mark_price_history"
	urlPathGetOrderBook                    = "/public/get_order_book"
	urlPathGetOrderBookByInstrumentId      = "/public/get_order_book_by_instrument_id"
	urlPathGetTradeVolumes                 = "/public/get_trade_volumes"
	// OHLCV
	urlPathGetTradingViewChartData = "/public/get_tradingview_chart_data"
	urlPathGetVolatilityIndexData  = "/public/get_volatility_index_data"
	urlPathGetTicker               = "/public/ticker"
)

// ## ------------------------------------------------------------------------

// FundingChartDataRequest represents the request parameters for the GetFundingChartData function.
type FundingChartDataRequest struct {
	// InstrumentName is the name of the instrument.
	InstrumentName string `json:"instrument_name"`
	// Length is the time period for the chart data. Possible values are "8h", "24h", and "1m".
	Length string `json:"length"`
}

// FundingChartDataResult represents the result section of the response for the GetFundingChartData function.
type FundingChartDataResult struct {
	// CurrentInterest is the current funding interest rate.
	CurrentInterest float64 `json:"current_interest"`
	// Data is an array of historical funding interest rate data.
	Data []struct {
		// IndexPrice is the current index price.
		IndexPrice float64 `json:"index_price"`
		// Interest8h is the historical funding interest rate for the past 8 hours.
		Interest8h float64 `json:"interest_8h"`
		// Timestamp is the timestamp (in milliseconds since the Unix epoch) of the data point.
		Timestamp int64 `json:"timestamp"`
	} `json:"data"`
	// Interest8h is the current 8-hour funding interest rate.
	Interest8h float64 `json:"interest_8h"`
}

// FundingChartDataResponse represents the response structure for the GetFundingChartData function.
type FundingChartDataResponse struct {
	JSONRPC string                 `json:"jsonrpc"`
	Result  FundingChartDataResult `json:"result"`
}

// GetFundingChartData retrieves the funding chart data for the specified instrument and time period.
func (s *MarketService) GetFundingChartData(request *FundingChartDataRequest) (*FundingChartDataResponse, error) {
	var resp FundingChartDataResponse
	uri := fmt.Sprintf("%s%s%s?instrument_name=%s&length=%s",
		s.client.baseURL,
		defaultAPIURL,
		urlPathGetFundingChartData,
		request.InstrumentName,
		request.Length,
	)
	err := s.client.DoPublic(uri, "GET", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ## ------------------------------------------------------------------------

// FundingRateHistoryEntry represents a single entry in the funding rate history.
type FundingRateHistoryEntry struct {
	// Timestamp is the timestamp (in milliseconds since the Unix epoch) of the data point.
	Timestamp int64 `json:"timestamp"`
	// IndexPrice is the price in the base currency.
	IndexPrice float64 `json:"index_price"`
	// PrevIndexPrice is the previous price in the base currency.
	PrevIndexPrice float64 `json:"prev_index_price"`
	// Interest1h is the 1-hour funding interest rate.
	Interest1h float64 `json:"interest_1h"`
	// Interest8h is the 8-hour funding interest rate.
	Interest8h float64 `json:"interest_8h"`
}

// FundingRateHistoryResponse represents the response structure for the GetFundingRateHistory function.
type FundingRateHistoryResponse struct {
	JSONRPC string                    `json:"jsonrpc"`
	ID      uint64                    `json:"id"`
	Result  []FundingRateHistoryEntry `json:"result"`
}

// GetFundingRateHistory retrieves the funding rate history for the specified instrument.
func (s *MarketService) GetFundingRateHistory(
	instrumentName string,
	startTimestamp int64,
	endTimestamp int64,
) (*FundingRateHistoryResponse, error) {
	var resp FundingRateHistoryResponse
	uri := fmt.Sprintf(
		"%s%s%s?instrument_name=%s&start_timestamp=%d&end_timestamp=%d",
		s.client.baseURL,
		defaultAPIURL,
		urlPathGetFundingRateHistory,
		instrumentName,
		startTimestamp,
		endTimestamp,
	)
	err := s.client.DoPublic(uri, "GET", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ## ------------------------------------------------------------------------

// FundingRateValueResponse represents the response structure for the GetFundingRateValue function.
type FundingRateValueResponse struct {
	JSONRPC string  `json:"jsonrpc"`
	ID      uint64  `json:"id"`
	Result  float64 `json:"result"`
}

// GetFundingRateValue retrieves the current funding rate for the specified instrument.
func (s *MarketService) GetFundingRateValue(
	instrumentName string,
	startTimestamp int64,
	endTimestamp int64,
) (*FundingRateValueResponse, error) {
	var resp FundingRateValueResponse
	uri := fmt.Sprintf("%s%s%s?instrument_name=%s&start_timestamp=%d&end_timestamp=%d",
		s.client.baseURL,
		defaultAPIURL,
		urlPathGetFundingRateValue,
		instrumentName,
		startTimestamp,
		endTimestamp,
	)
	err := s.client.DoPublic(uri, "GET", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ## ------------------------------------------------------------------------

// HistoricalVolatilityResponse represents the response structure for the GetHistoricalVolatility function.
type HistoricalVolatilityResponse struct {
	JSONRPC string       `json:"jsonrpc"`
	ID      uint64       `json:"id"`
	Result  [][2]float64 `json:"result"`
}

// GetHistoricalVolatility retrieves the historical volatility for the specified currency.
func (s *MarketService) GetHistoricalVolatility(currency string) (*HistoricalVolatilityResponse, error) {
	var resp HistoricalVolatilityResponse
	uri := fmt.Sprintf("%s%s%s?currency=%s",
		s.client.baseURL,
		defaultAPIURL,
		urlPathGetHistoricalVolatility,
		currency,
	)
	err := s.client.DoPublic(uri, "GET", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ## ------------------------------------------------------------------------

type IndexPriceResult struct {
	EstimatedDeliveryPrice float64 `json:"estimated_delivery_price"`
	IndexPrice             float64 `json:"index_price"`
}

// IndexPriceResponse represents the response structure for the GetIndexPrice function.
type IndexPriceResponse struct {
	JSONRPC string           `json:"jsonrpc"`
	Result  IndexPriceResult `json:"result"`
}

// GetIndexPrice retrieves the index price for the specified index name.
func (s *MarketService) GetIndexPrice(indexName string) (*IndexPriceResponse, error) {
	var resp IndexPriceResponse
	uri := fmt.Sprintf("%s%s%s?index_name=%s",
		s.client.baseURL,
		defaultAPIURL,
		urlPathGetIndexPrice,
		indexName,
	)
	err := s.client.DoPublic(uri, "GET", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ## ------------------------------------------------------------------------
// IndexPriceNamesResponse represents the response structure for the GetIndexPriceNames function.
type IndexPriceNamesResponse struct {
	JSONRPC string   `json:"jsonrpc"`
	ID      uint64   `json:"id"`
	Result  []string `json:"result"`
}

// GetIndexPriceNames retrieves the list of available index price names.
func (s *MarketService) GetIndexPriceNames() (*IndexPriceNamesResponse, error) {
	var resp IndexPriceNamesResponse
	uri := fmt.Sprintf("%s%s%s",
		s.client.baseURL,
		defaultAPIURL,
		urlPathGetIndexPriceNames,
	)
	err := s.client.DoPublic(uri, "GET", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ## ------------------------------------------------------------------------
// InstrumentResponse represents the response structure for the GetInstrument function.
type InstrumentResult struct {
	BaseCurrency             string  `json:"base_currency"`
	BlockTradeCommission     float64 `json:"block_trade_commission"`
	BlockTradeMinTradeAmount float64 `json:"block_trade_min_trade_amount"`
	BlockTradeTickSize       float64 `json:"block_trade_tick_size"`
	ContractSize             float64 `json:"contract_size"`
	CounterCurrency          string  `json:"counter_currency"`
	CreationTimestamp        int64   `json:"creation_timestamp"`
	ExpirationTimestamp      int64   `json:"expiration_timestamp"`
	InstrumentID             int64   `json:"instrument_id"`
	InstrumentName           string  `json:"instrument_name"`
	InstrumentType           string  `json:"instrument_type"`
	IsActive                 bool    `json:"is_active"`
	Kind                     string  `json:"kind"`
	MakerCommission          float64 `json:"maker_commission"`
	MinTradeAmount           float64 `json:"min_trade_amount"`
	OptionType               string  `json:"option_type"`
	PriceIndex               string  `json:"price_index"`
	QuoteCurrency            string  `json:"quote_currency"`
	RFQ                      bool    `json:"rfq"`
	SettlementCurrency       string  `json:"settlement_currency"`
	SettlementPeriod         string  `json:"settlement_period"`
	Strike                   float64 `json:"strike"`
	TakerCommission          float64 `json:"taker_commission"`
	TickSize                 float64 `json:"tick_size"`
	TickSizeSteps            []struct {
		AbovePrice float64 `json:"above_price"`
		TickSize   float64 `json:"tick_size"`
	} `json:"tick_size_steps"`
}

type InstrumentResponse struct {
	JSONRPC string           `json:"jsonrpc"`
	ID      uint64           `json:"id"`
	Result  InstrumentResult `json:"result"`
}

// GetInstrument retrieves the details of the specified instrument.
func (s *MarketService) GetInstrument(instrumentName string) (*InstrumentResponse, error) {
	var resp InstrumentResponse
	uri := fmt.Sprintf("%s%s%s?instrument_name=%s",
		s.client.baseURL,
		defaultAPIURL,
		urlPathGetInstrument,
		instrumentName,
	)
	err := s.client.DoPublic(uri, "GET", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ## ------------------------------------------------------------------------

// InstrumentsResponse represents the response structure for the GetInstruments function.
type InstrumentsResponse struct {
	JSONRPC string             `json:"jsonrpc"`
	ID      uint64             `json:"id"`
	Result  []InstrumentResult `json:"result"`
}

// GetInstruments retrieves a list of available instruments.
func (s *MarketService) GetInstruments(currency string, kind string, expired bool) (*InstrumentsResponse, error) {
	var resp InstrumentsResponse
	uri := fmt.Sprintf("%s%s%s?currency=%s&kind=%s&expired=%t",
		s.client.baseURL,
		defaultAPIURL,
		urlPathGetInstruments,
		currency,
		kind,
		expired,
	)
	err := s.client.DoPublic(uri, "GET", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ## ------------------------------------------------------------------------
// LastSettlementEntry represents a single settlement event.
type LastSettlementEntry struct {
	Funded            float64 `json:"funded"`
	Funding           float64 `json:"funding"`
	IndexPrice        float64 `json:"index_price"`
	InstrumentName    string  `json:"instrument_name"`
	MarkPrice         float64 `json:"mark_price"`
	Position          float64 `json:"position"`
	ProfitLoss        float64 `json:"profit_loss"`
	SessionBankruptcy float64 `json:"session_bankruptcy"`
	SessionProfitLoss float64 `json:"session_profit_loss"`
	SessionTax        float64 `json:"session_tax"`
	SessionTaxRate    float64 `json:"session_tax_rate"`
	Socialized        float64 `json:"socialized"`
	Timestamp         int64   `json:"timestamp"`
	Type              string  `json:"type"`
}

// LastSettlementsResponse represents the response structure for the GetLastSettlementsByInstrument function.
type LastSettlementsResponse struct {
	JSONRPC string `json:"jsonrpc"`
	ID      uint64 `json:"id"`
	Result  struct {
		Continuation string                `json:"continuation"`
		Settlements  []LastSettlementEntry `json:"settlements"`
	} `json:"result"`
}

// GetLastSettlementsByInstrument retrieves the last settlements for the specified instrument.
func (s *MarketService) GetLastSettlementsByInstrument(
	instrumentName string,
	settlementType string,
	count int,
	continuation string,
	searchStartTimestamp int64,
) (*LastSettlementsResponse, error) {
	var resp LastSettlementsResponse
	uri := fmt.Sprintf("%s%s%s?instrument_name=%s&type=%s&count=%d&continuation=%s&search_start_timestamp=%d",
		s.client.baseURL,
		defaultAPIURL,
		urlPathGetLastSettlementsByInstrument,
		instrumentName,
		settlementType,
		count,
		continuation,
		searchStartTimestamp,
	)
	err := s.client.DoPublic(uri, "GET", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ## ------------------------------------------------------------------------
// LastTradeResponse represents a single trade.
type LastTradeResponse struct {
	Amount             float64 `json:"amount"`
	BlockRFQID         int     `json:"block_rfq_id"`
	BlockTradeID       string  `json:"block_trade_id"`
	BlockTradeLegCount int     `json:"block_trade_leg_count"`
	ComboID            string  `json:"combo_id"`
	ComboTradeID       float64 `json:"combo_trade_id"`
	Contracts          float64 `json:"contracts"`
	Direction          string  `json:"direction"`
	IndexPrice         float64 `json:"index_price"`
	InstrumentName     string  `json:"instrument_name"`
	IV                 float64 `json:"iv"`
	Liquidation        string  `json:"liquidation"`
	MarkPrice          float64 `json:"mark_price"`
	Price              float64 `json:"price"`
	TickDirection      int     `json:"tick_direction"`
	Timestamp          int64   `json:"timestamp"`
	TradeID            string  `json:"trade_id"`
	TradeSeq           int     `json:"trade_seq"`
}

// LastTradesByCurrencyAndTimeResponse represents the response structure for the GetLastTradesByCurrencyAndTime function.
type LastTradesByCurrencyAndTimeResponse struct {
	JSONRPC string `json:"jsonrpc"`
	ID      uint64 `json:"id"`
	Result  struct {
		HasMore bool                `json:"has_more"`
		Trades  []LastTradeResponse `json:"trades"`
	} `json:"result"`
}

// GetLastTradesByCurrencyAndTime retrieves the last trades for the specified currency and time range.
func (s *MarketService) GetLastTradesByCurrencyAndTime(
	currency string,
	startTimestamp int64,
	endTimestamp int64,
	count int,
) (*LastTradesByCurrencyAndTimeResponse, error) {
	var resp LastTradesByCurrencyAndTimeResponse
	uri := fmt.Sprintf("%s%s%s?currency=%s&start_timestamp=%d&end_timestamp=%d&count=%d",
		s.client.baseURL,
		defaultAPIURL,
		urlPathGetLastTradeByCurrencyAndTime,
		currency,
		startTimestamp,
		endTimestamp,
		count,
	)
	err := s.client.DoPublic(uri, "GET", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ## ------------------------------------------------------------------------

// LastTradesByInstrumentResponse represents the response structure for the GetLastTradesByInstrument function.
type LastTradesByInstrumentResponse struct {
	JSONRPC string `json:"jsonrpc"`
	ID      uint64 `json:"id"`
	Result  struct {
		HasMore bool                `json:"has_more"`
		Trades  []LastTradeResponse `json:"trades"`
	} `json:"result"`
}

// GetLastTradesByInstrument retrieves the last trades for the specified instrument.
func (s *MarketService) GetLastTradesByInstrument(
	instrumentName string,
	startSeq int,
	endSeq int,
	startTimestamp int64,
	endTimestamp int64,
	count int,
	sorting string,
) (*LastTradesByInstrumentResponse, error) {
	var resp LastTradesByInstrumentResponse
	uri := fmt.Sprintf(
		"%s%s%s?instrument_name=%s",
		s.client.baseURL,
		defaultAPIURL,
		urlPathGetLastTradeByInstrument,
		instrumentName,
	)

	if startSeq != 0 {
		uri += fmt.Sprintf("&start_seq=%d", startSeq)
	}
	if endSeq != 0 {
		uri += fmt.Sprintf("&end_seq=%d", endSeq)
	}
	if startTimestamp != 0 {
		uri += fmt.Sprintf("&start_timestamp=%d", startTimestamp)
	}
	if endTimestamp != 0 {
		uri += fmt.Sprintf("&end_timestamp=%d", endTimestamp)
	}
	if count != 0 {
		uri += fmt.Sprintf("&count=%d", count)
	}
	if sorting != "" {
		uri += fmt.Sprintf("&sorting=%s", sorting)
	}

	err := s.client.DoPublic(uri, "GET", nil, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// ## ------------------------------------------------------------------------

// GetLastTradesByInstrumentAndTimeResponse represents the response structure for the GetLastTradesByInstrumentAndTime function.
type GetLastTradesByInstrumentAndTimeResponse struct {
	JSONRPC string `json:"jsonrpc"`
	ID      uint64 `json:"id"`
	Result  struct {
		HasMore bool                `json:"has_more"`
		Trades  []LastTradeResponse `json:"trades"`
	} `json:"result"`
}

// GetLastTradesByInstrumentAndTime retrieves the last trades for the specified instrument and time range.
func (s *MarketService) GetLastTradesByInstrumentAndTime(
	instrumentName string,
	startTimestamp int64,
	endTimestamp int64,
	count int,
	sorting string,
) (*GetLastTradesByInstrumentAndTimeResponse, error) {
	var resp GetLastTradesByInstrumentAndTimeResponse
	uri := fmt.Sprintf(
		"%s%s%s?instrument_name=%s&start_timestamp=%d&end_timestamp=%d",
		s.client.baseURL,
		defaultAPIURL,
		urlPathGetLastTradeByInstrumentAndTime,
		instrumentName,
		startTimestamp,
		endTimestamp,
	)

	if count != 0 {
		uri += fmt.Sprintf("&count=%d", count)
	}
	if sorting != "" {
		uri += fmt.Sprintf("&sorting=%s", sorting)
	}

	err := s.client.DoPublic(uri, "GET", nil, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// ## ------------------------------------------------------------------------

// MarkPriceHistoryResponse represents the response structure for the GetMarkPriceHistory function.
type MarkPriceHistoryResponse struct {
	JSONRPC string       `json:"jsonrpc"`
	ID      uint64       `json:"id"`
	Result  [][2]float64 `json:"result"`
}

// GetMarkPriceHistory retrieves the mark price history for the specified instrument.
func (s *MarketService) GetMarkPriceHistory(
	instrumentName string,
	startTimestamp int64,
	endTimestamp int64,

) (*MarkPriceHistoryResponse, error) {
	var resp MarkPriceHistoryResponse
	uri := fmt.Sprintf(
		"%s%s%s?instrument_name=%s&start_timestamp=%d&end_timestamp=%d",
		s.client.baseURL,
		defaultAPIURL,
		urlPathGetMarkPriceHistory,
		instrumentName,
		startTimestamp,
		endTimestamp,
	)

	err := s.client.DoPublic(uri, "GET", nil, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// ## ------------------------------------------------------------------------

// OrderBookResult represents the result section of the response for the GetOrderBook function.
type OrderBookResult struct {
	AskIV          float64     `json:"ask_iv"`
	Asks           [][]float64 `json:"asks"`
	BestAskAmount  float64     `json:"best_ask_amount"`
	BestAskPrice   float64     `json:"best_ask_price"`
	BestBidAmount  float64     `json:"best_bid_amount"`
	BestBidPrice   float64     `json:"best_bid_price"`
	BidIV          float64     `json:"bid_iv"`
	Bids           [][]float64 `json:"bids"`
	CurrentFunding float64     `json:"current_funding"`
	DeliveryPrice  float64     `json:"delivery_price"`
	Funding8h      float64     `json:"funding_8h"`
	Greeks         struct {
		Delta float64 `json:"delta"`
		Gamma float64 `json:"gamma"`
		Rho   float64 `json:"rho"`
		Theta float64 `json:"theta"`
		Vega  float64 `json:"vega"`
	} `json:"greeks"`
	IndexPrice      float64 `json:"index_price"`
	InstrumentName  string  `json:"instrument_name"`
	InterestRate    float64 `json:"interest_rate"`
	LastPrice       float64 `json:"last_price"`
	MarkIV          float64 `json:"mark_iv"`
	MarkPrice       float64 `json:"mark_price"`
	MaxPrice        float64 `json:"max_price"`
	MinPrice        float64 `json:"min_price"`
	OpenInterest    float64 `json:"open_interest"`
	SettlementPrice float64 `json:"settlement_price"`
	State           string  `json:"state"`
	Stats           struct {
		High        float64 `json:"high"`
		Low         float64 `json:"low"`
		PriceChange float64 `json:"price_change"`
		Volume      float64 `json:"volume"`
		VolumeUSD   float64 `json:"volume_usd"`
	} `json:"stats"`
	Timestamp       int64   `json:"timestamp"`
	UnderlyingIndex string  `json:"underlying_index"`
	UnderlyingPrice float64 `json:"underlying_price"`
}

// OrderBookResponse represents the response structure for the GetOrderBook function.
type OrderBookResponse struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      uint64          `json:"id"`
	Result  OrderBookResult `json:"result"`
}

// GetOrderBook retrieves the order book for the specified instrument.
func (s *MarketService) GetOrderBook(
	instrumentName string,
	depth int,
) (*OrderBookResponse, error) {
	var resp OrderBookResponse
	uri := fmt.Sprintf(
		"%s%s%s?instrument_name=%s",
		s.client.baseURL,
		defaultAPIURL,
		urlPathGetOrderBook,
		instrumentName,
	)

	if depth > 0 {
		uri += fmt.Sprintf("&depth=%d", depth)
	}

	err := s.client.DoPublic(uri, "GET", nil, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// ## ------------------------------------------------------------------------

// GetOrderBookByInstrumentResponse represents the response structure for the GetOrderBookByInstrumentId function.
type GetOrderBookByInstrumentResponse struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      uint64          `json:"id"`
	Result  OrderBookResult `json:"result"`
}

// GetOrderBookByInstrumentId retrieves the order book for the specified instrument ID.
func (s *MarketService) GetOrderBookByInstrumentId(
	instrumentID int,
	depth int,
) (*GetOrderBookByInstrumentResponse, error) {
	var resp GetOrderBookByInstrumentResponse
	uri := fmt.Sprintf(
		"%s%s%s?instrument_id=%d",
		s.client.baseURL,
		defaultAPIURL,
		urlPathGetOrderBookByInstrumentId,
		instrumentID,
	)

	if depth > 0 {
		uri += fmt.Sprintf("&depth=%d", depth)
	}

	err := s.client.DoPublic(uri, "GET", nil, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// ## ------------------------------------------------------------------------

// TradeVolumeResponse represents the response structure for the GetTradeVolumes function.
type TradeVolumeResponse struct {
	CallsVolume      float64 `json:"calls_volume"`
	CallsVolume30d   float64 `json:"calls_volume_30d"`
	CallsVolume7d    float64 `json:"calls_volume_7d"`
	Currency         string  `json:"currency"`
	FuturesVolume    float64 `json:"futures_volume"`
	FuturesVolume30d float64 `json:"futures_volume_30d"`
	FuturesVolume7d  float64 `json:"futures_volume_7d"`
	PutsVolume       float64 `json:"puts_volume"`
	PutsVolume30d    float64 `json:"puts_volume_30d"`
	PutsVolume7d     float64 `json:"puts_volume_7d"`
	SpotVolume       float64 `json:"spot_volume"`
	SpotVolume30d    float64 `json:"spot_volume_30d"`
	SpotVolume7d     float64 `json:"spot_volume_7d"`
}

// GetTradeVolumesResponse represents the response structure for the GetTradeVolumes function.
type GetTradeVolumesResponse struct {
	JSONRPC string                `json:"jsonrpc"`
	ID      uint64                `json:"id"`
	Result  []TradeVolumeResponse `json:"result"`
}

// GetTradeVolumes retrieves the trade volumes for the specified currency.
func (s *MarketService) GetTradeVolumes(
	extended bool,
) (*GetTradeVolumesResponse, error) {
	var resp GetTradeVolumesResponse
	uri := fmt.Sprintf(
		"%s%s%s",
		s.client.baseURL,
		defaultAPIURL,
		urlPathGetTradeVolumes,
	)

	if extended {
		uri += "?extended=true"
	}

	err := s.client.DoPublic(uri, "GET", nil, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// ## ------------------------------------------------------------------------

type OHLCVResult struct {
	Close  []float64 `json:"close"`
	Cost   []float64 `json:"cost"`
	High   []float64 `json:"high"`
	Low    []float64 `json:"low"`
	Open   []float64 `json:"open"`
	Status string    `json:"status"`
	Ticks  []int64   `json:"ticks"`
	Volume []float64 `json:"volume"`
}

// TradingViewChartDataResponse represents the response structure for the GetTradingViewChartData function.
type TradingViewChartDataResponse struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      uint64      `json:"id"`
	Result  OHLCVResult `json:"result"`
	UsIn    int64       `json:"usIn"`
	UsOut   int64       `json:"usOut"`
	UsDiff  int64       `json:"usDiff"`
	Testnet bool        `json:"testnet"`
}

// GetTradingViewChartData retrieves the trading view chart data for the specified instrument.
func (s *MarketService) GetTradingViewChartData(
	instrumentName string,
	startTimestamp int64,
	endTimestamp int64,
	resolution string,
) (*TradingViewChartDataResponse, error) {
	var resp TradingViewChartDataResponse
	uri := fmt.Sprintf(
		"%s%s%s?instrument_name=%s&start_timestamp=%d&end_timestamp=%d&resolution=%s",
		s.client.baseURL,
		defaultAPIURL,
		urlPathGetTradingViewChartData,
		instrumentName,
		startTimestamp,
		endTimestamp,
		resolution,
	)

	err := s.client.DoPublic(uri, "GET", nil, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// ## ------------------------------------------------------------------------

// VolatilityIndexDataEntry represents a single entry in the volatility index data.
type VolatilityIndexDataEntry struct {
	Timestamp int64   `json:"timestamp"`
	Open      float64 `json:"open"`
	High      float64 `json:"high"`
	Low       float64 `json:"low"`
	Close     float64 `json:"close"`
}

// VolatilityIndexDataResponse represents the response structure for the GetVolatilityIndexData function.
type VolatilityIndexDataResponse struct {
	JSONRPC string `json:"jsonrpc"`
	ID      uint64 `json:"id"`
	Result  struct {
		Continuation string      `json:"continuation"`
		Data         [][]float64 `json:"data"`
	} `json:"result"`
}

// GetVolatilityIndexData retrieves the volatility index data for the specified instrument.
func (s *MarketService) GetVolatilityIndexData(
	currency string,
	startTimestamp int64,
	endTimestamp int64,
	resolution string,
) (*VolatilityIndexDataResponse, error) {
	var resp VolatilityIndexDataResponse
	uri := fmt.Sprintf(
		"%s%s%s?currency=%s&start_timestamp=%d&end_timestamp=%d&resolution=%s",
		s.client.baseURL,
		defaultAPIURL,
		urlPathGetVolatilityIndexData,
		currency,
		startTimestamp,
		endTimestamp,
		resolution,
	)

	err := s.client.DoPublic(uri, "GET", nil, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// ## ------------------------------------------------------------------------

// TickerResult represents the response structure for the GetTicker function.
type TickerResult struct {
	BestAskAmount          float64 `json:"best_ask_amount"`
	BestAskPrice           float64 `json:"best_ask_price"`
	BestBidAmount          float64 `json:"best_bid_amount"`
	BestBidPrice           float64 `json:"best_bid_price"`
	CurrentFunding         float64 `json:"current_funding"`
	EstimatedDeliveryPrice float64 `json:"estimated_delivery_price"`
	Funding8h              float64 `json:"funding_8h"`
	IndexPrice             float64 `json:"index_price"`
	InstrumentName         string  `json:"instrument_name"`
	InterestValue          float64 `json:"interest_value"`
	LastPrice              float64 `json:"last_price"`
	MarkPrice              float64 `json:"mark_price"`
	MaxPrice               float64 `json:"max_price"`
	MinPrice               float64 `json:"min_price"`
	OpenInterest           float64 `json:"open_interest"`
	SettlementPrice        float64 `json:"settlement_price"`
	State                  string  `json:"state"`
	Stats                  struct {
		High        float64 `json:"high"`
		Low         float64 `json:"low"`
		PriceChange float64 `json:"price_change"`
		Volume      float64 `json:"volume"`
		VolumeUSD   float64 `json:"volume_usd"`
	} `json:"stats"`
	Timestamp int64 `json:"timestamp"`
}

type TickerResponse struct {
	JSONRPC string       `json:"jsonrpc"`
	ID      uint64       `json:"id"`
	Result  TickerResult `json:"result"`
}

// GetTicker retrieves the ticker data for the specified instrument.
func (s *MarketService) GetTicker(
	instrumentName string,
) (*TickerResponse, error) {
	var resp TickerResponse
	uri := fmt.Sprintf(
		"%s%s%s?instrument_name=%s",
		s.client.baseURL,
		defaultAPIURL,
		urlPathGetTicker,
		instrumentName,
	)

	err := s.client.DoPublic(uri, "GET", nil, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
