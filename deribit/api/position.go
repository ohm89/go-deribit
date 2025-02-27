package api

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

type PositionService struct {
	client *Client
}

const (
	urlPathGetPosition        = "/private/get_position"
	urlPathGetPositions       = "/private/get_positions"
	urlPathGetSimulateMargins = "/private/simulate_portfolio"
	urlPathClosePosition      = "/private/close_position"
)

type GetPositionDetailsResponse struct {
	ID      uint64   `json:"id"`
	JSONRPC string   `json:"jsonrpc"`
	Result  Position `json:"result"`
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

// ## Get Position
func (s *PositionService) GetPosition(
	instrumentName string,
) (*GetPositionDetailsResponse, error) {
	var resp GetPositionDetailsResponse
	uri := fmt.Sprintf(
		"%s%s%s?instrument_name=%s",
		s.client.baseURL,
		defaultAPIURL,
		urlPathGetPosition,
		instrumentName,
	)

	err := s.client.DoPrivate(uri, "GET", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetPositionsResponse struct {
	ID      uint64     `json:"id"`
	JSONRPC string     `json:"jsonrpc"`
	Result  []Position `json:"result"`
}

func (s *PositionService) GetPositions(
	currency string,
	kind string,
	subaccountID int,
) (*GetPositionsResponse, error) {
	var resp GetPositionsResponse
	uri := fmt.Sprintf("%s%s%s", s.client.baseURL, defaultAPIURL, urlPathGetPositions)
	queryParams := make([]string, 0, 3)

	if currency != "" {
		queryParams = append(queryParams, fmt.Sprintf("currency=%s", currency))
	}
	if kind != "" {
		queryParams = append(queryParams, fmt.Sprintf("kind=%s", kind))
	}
	if subaccountID > 0 {
		queryParams = append(queryParams, fmt.Sprintf("subaccount_id=%d", subaccountID))
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

type FTUSummary struct {
	ShortTotalCost float64   `json:"short_total_cost"`
	PlVec          []float64 `json:"pl_vec"`
	LongTotalCost  float64   `json:"long_total_cost"`
	ExpTimestamp   int64     `json:"exp_tstamp"`
}

type FTUEntry struct {
	TotalCost      float64   `json:"total_cost"`
	Size           float64   `json:"size"`
	PlVec          []float64 `json:"pl_vec"`
	MarkPrice      float64   `json:"mark_price"`
	InstrumentName string    `json:"instrument_name"`
	ExpTimestamp   int64     `json:"exp_tstamp"`
}

type PortfolioMarginsResult struct {
	VolumeRange          []float64          `json:"vol_range"`
	VegaPow2             float64            `json:"vega_pow2"`
	VegaPow1             float64            `json:"vega_pow1"`
	Skew                 float64            `json:"skew"`
	PriceRange           float64            `json:"price_range"`
	OptSumContinguency   float64            `json:"opt_sum_continguency"`
	OptContinguency      float64            `json:"opt_continguency"`
	Kurtosis             float64            `json:"kurtosis"`
	InterestRate         float64            `json:"int_rate"`
	InitialMarginFactor  float64            `json:"initial_margin_factor"`
	FtuContinguency      float64            `json:"ftu_continguency"`
	AtmRange             float64            `json:"atm_range"`
	ProjectedMarginPos   float64            `json:"projected_margin_pos"`
	ProjectedMargin      float64            `json:"projected_margin"`
	PositionSizes        map[string]float64 `json:"position_sizes"`
	Pls                  []float64          `json:"pls"`
	PcoOpt               float64            `json:"pco_opt"`
	PcoFtu               float64            `json:"pco_ftu"`
	OptSummary           []interface{}      `json:"opt_summary"`
	OptPls               []float64          `json:"opt_pls"`
	OptEntries           []interface{}      `json:"opt_entries"`
	MarginPos            float64            `json:"margin_pos"`
	Margin               float64            `json:"margin"`
	FtuSummary           []FTUSummary       `json:"ftu_summary"`
	FtuPls               []float64          `json:"ftu_pls"`
	FtuEntries           []FTUEntry         `json:"ftu_entries"`
	CoOpt                float64            `json:"co_opt"`
	CoFtu                float64            `json:"co_ftu"`
	CalculationTimestamp int64              `json:"calculation_timestamp"`
}

type SimulatePortfolioResponse struct {
	ID         uint64                 `json:"id"`
	JSONRPC    string                 `json:"jsonrpc"`
	Result     PortfolioMarginsResult `json:"result"`
	UsIn       int64                  `json:"usIn"`
	UsOut      int64                  `json:"usOut"`
	UsDiff     int64                  `json:"usDiff"`
	TestnetApp bool                   `json:"testnet"`
}

// ## Get Simulated Margins
func (s *PositionService) GetSimulateMargins(
	currency string,
	addPositions bool,
	simulatedPositions map[string]float64,
) (*SimulatePortfolioResponse, error) {
	var resp SimulatePortfolioResponse
	uri := fmt.Sprintf("%s%s%s", s.client.baseURL, defaultAPIURL, urlPathGetSimulateMargins)
	queryParams := make([]string, 0, 3)

	queryParams = append(queryParams, fmt.Sprintf("currency=%s", currency))
	if addPositions {
		queryParams = append(queryParams, "add_positions=true")
	} else {
		queryParams = append(queryParams, "add_positions=false")
	}

	if len(simulatedPositions) > 0 {
		// Encode the simulated positions map as a JSON string
		simulatedPositionsJSON, err := json.Marshal(simulatedPositions)
		if err != nil {
			return nil, err
		}
		queryParams = append(queryParams, fmt.Sprintf("simulated_positions=%s", url.QueryEscape(string(simulatedPositionsJSON))))
	}

	uri += "?" + strings.Join(queryParams, "&")

	err := s.client.DoPrivate(uri, "GET", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ## ---------------------------------------
type ClosePositionRequest struct {
	InstrumentName string  `json:"instrument_name"`
	Type           string  `json:"type"`
	Price          float64 `json:"price,omitempty"`
}

type ClosePositionResponse struct {
	ID      uint64        `json:"id"`
	JSONRPC string        `json:"jsonrpc"`
	Result  ClosePosition `json:"result"`
}

type ClosePosition struct {
	Order  OrderResultOrderResponse   `json:"order"`
	Trades []OrderResultTradeResponse `json:"trades"`
}

// ## Close Position
func (s *PositionService) ClosePosition(
	instrumentName string,
	orderType string,
	price float64,
) (*ClosePositionResponse, error) {
	var resp ClosePositionResponse
	uri := fmt.Sprintf(
		"%s%s%s?instrument_name=%s&type=%s&price=%f",
		s.client.baseURL,
		defaultAPIURL,
		urlPathClosePosition,
		instrumentName,
		orderType,
		price,
	)
	err := s.client.DoPrivate(uri, "GET", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
