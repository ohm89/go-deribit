package api

import (
	"fmt"
)

type AccountService struct {
	client *Client
}

const (
	urlPathAccountSummaries = "/private/get_account_summaries"
	urlPathAccountSummary   = "/private/get_account_summary"
)

type AccountSummary struct {
	Currency                     string                 `json:"currency"`
	DeltaTotalMap                map[string]float64     `json:"delta_total_map"`
	MarginBalance                float64                `json:"margin_balance"`
	FuturesSessionRPL            float64                `json:"futures_session_rpl"`
	OptionsSessionRPL            float64                `json:"options_session_rpl"`
	EstimatedLiquidationRatioMap map[string]float64     `json:"estimated_liquidation_ratio_map"`
	SessionUPL                   float64                `json:"session_upl"`
	EstimatedLiquidationRatio    float64                `json:"estimated_liquidation_ratio"`
	OptionsGammaMap              map[string]float64     `json:"options_gamma_map"`
	OptionsVega                  float64                `json:"options_vega"`
	OptionsValue                 float64                `json:"options_value"`
	AvailableWithdrawalFunds     float64                `json:"available_withdrawal_funds"`
	ProjectedDeltaTotal          float64                `json:"projected_delta_total"`
	MaintenanceMargin            float64                `json:"maintenance_margin"`
	TotalPL                      float64                `json:"total_pl"`
	Limits                       map[string]interface{} `json:"limits"`
	ProjectedMaintenanceMargin   float64                `json:"projected_maintenance_margin"`
	AvailableFunds               float64                `json:"available_funds"`
	OptionsDelta                 float64                `json:"options_delta"`
	Balance                      float64                `json:"balance"`
	Equity                       float64                `json:"equity"`
	FuturesSessionUPL            float64                `json:"futures_session_upl"`
	FeeBalance                   float64                `json:"fee_balance"`
	OptionsSessionUPL            float64                `json:"options_session_upl"`
	ProjectedInitialMargin       float64                `json:"projected_initial_margin"`
	OptionsTheta                 float64                `json:"options_theta"`
	PortfolioMarginingEnabled    bool                   `json:"portfolio_margining_enabled"`
	CrossCollateralEnabled       bool                   `json:"cross_collateral_enabled"`
	MarginModel                  string                 `json:"margin_model"`
	OptionsVegaMap               map[string]float64     `json:"options_vega_map"`
	FuturesPL                    float64                `json:"futures_pl"`
	OptionsPL                    float64                `json:"options_pl"`
	InitialMargin                float64                `json:"initial_margin"`
	SpotReserve                  float64                `json:"spot_reserve"`
	DeltaTotal                   float64                `json:"delta_total"`
	OptionsGamma                 float64                `json:"options_gamma"`
	SessionRPL                   float64                `json:"session_rpl"`
}

type AccountSummaries struct {
	ID                               uint64           `json:"id"`
	Email                            string           `json:"email"`
	SystemName                       string           `json:"system_name"`
	Username                         string           `json:"username"`
	CreationTimestamp                int64            `json:"creation_timestamp"`
	Type                             string           `json:"type"`
	ReferrerID                       interface{}      `json:"referrer_id"`
	LoginEnabled                     bool             `json:"login_enabled"`
	SecurityKeysEnabled              bool             `json:"security_keys_enabled"`
	MMPEnabled                       bool             `json:"mmp_enabled"`
	InteruserTransfersEnabled        bool             `json:"interuser_transfers_enabled"`
	SelfTradingRejectMode            string           `json:"self_trading_reject_mode"`
	SelfTradingExtendedToSubaccounts bool             `json:"self_trading_extended_to_subaccounts"`
	Summaries                        []AccountSummary `json:"summaries"`
}

type AccountSummariesResponse struct {
	ID      uint64           `json:"id"`
	JSONRPC string           `json:"jsonrpc"`
	Result  AccountSummaries `json:"result"`
	Error   *ResponseError   `json:"error,omitempty"`
}

type AccountSummaryResponse struct {
	ID      uint64         `json:"id"`
	JSONRPC string         `json:"jsonrpc"`
	Result  AccountSummary `json:"result"`
	Error   *ResponseError `json:"error,omitempty"`
}

// ## Get All Asset (Currency) in Account
func (s *AccountService) GetAccountSummaries(extended bool) (*AccountSummariesResponse, error) {
	var resp AccountSummariesResponse
	uri := fmt.Sprintf(
		"%s%s%s?extended=%t",
		s.client.baseURL,
		defaultAPIURL,
		urlPathAccountSummaries,
		extended,
	)

	err := s.client.DoPrivate(uri, "GET", nil, &resp)

	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, fmt.Errorf("API request failed: code=%d, message=%s", resp.Error.Code, resp.Error.Message)
	}

	return &resp, nil
}

func (s *AccountService) GetAccountSummary(currency string, extended bool) (*AccountSummaryResponse, error) {
	var resp AccountSummaryResponse
	uri := fmt.Sprintf(
		"%s%s%s?currency=%s&extended=%t",
		s.client.baseURL,
		defaultAPIURL,
		urlPathAccountSummary,
		currency,
		extended,
	)
	err := s.client.DoPrivate(uri, "GET", nil, &resp)
	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, fmt.Errorf("API request failed: code=%d, message=%s", resp.Error.Code, resp.Error.Message)
	}

	return &resp, nil
}
