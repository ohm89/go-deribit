package ws

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

type Limits struct {
	Max_position_size  float64 `json:"max_position_size"`
	Max_leverage       float64 `json:"max_leverage"`
	Max_open_orders    int     `json:"max_open_orders"`
	Max_position_value float64 `json:"max_position_value"`
	Max_order_size     float64 `json:"max_order_size"`
	Max_order_value    float64 `json:"max_order_value"`
}

type Fee struct {
	Currency        string  `json:"currency"`
	Fee_type        string  `json:"fee_type"`
	Instrument_type string  `json:"instrument_type"`
	Maker_fee       float64 `json:"maker_fee"`
	Taker_fee       float64 `json:"taker_fee"`
}

type AccountSummary struct {
	Options_pl                   float64            `json:"options_pl"`
	Projected_delta_total        float64            `json:"projected_delta_total"`
	Options_theta_map            map[string]float64 `json:"options_theta_map"`
	Has_non_block_chain_equity   bool               `json:"has_non_block_chain_equity"`
	Total_margin_balance_usd     float64            `json:"total_margin_balance_usd"`
	Limits                       Limits             `json:"limits"`
	Total_delta_total_usd        float64            `json:"total_delta_total_usd"`
	Available_withdrawal_funds   float64            `json:"available_withdrawal_funds"`
	Options_session_rpl          float64            `json:"options_session_rpl"`
	Futures_session_rpl          float64            `json:"futures_session_rpl"`
	Total_pl                     float64            `json:"total_pl"`
	Spot_reserve                 float64            `json:"spot_reserve"`
	Fees                         []Fee              `json:"fees"`
	Additional_reserve           float64            `json:"additional_reserve"`
	Options_session_upl          float64            `json:"options_session_upl"`
	Cross_collateral_enabled     bool               `json:"cross_collateral_enabled"`
	Options_value                float64            `json:"options_value"`
	Options_vega_map             map[string]float64 `json:"options_vega_map"`
	Maintenance_margin           float64            `json:"maintenance_margin"`
	Futures_session_upl          float64            `json:"futures_session_upl"`
	Portfolio_margining_enabled  bool               `json:"portfolio_margining_enabled"`
	Futures_pl                   float64            `json:"futures_pl"`
	Options_gamma_map            map[string]float64 `json:"options_gamma_map"`
	Currency                     string             `json:"currency"`
	Options_delta                float64            `json:"options_delta"`
	Initial_margin               float64            `json:"initial_margin"`
	Projected_maintenance_margin float64            `json:"projected_maintenance_margin"`
	Available_funds              float64            `json:"available_funds"`
	Equity                       float64            `json:"equity"`
	Margin_model                 string             `json:"margin_model"`
	Balance                      float64            `json:"balance"`
	Session_upl                  float64            `json:"session_upl"`
	Margin_balance               float64            `json:"margin_balance"`
	Deposit_address              string             `json:"deposit_address"`
	Options_theta                float64            `json:"options_theta"`
	Total_initial_margin_usd     float64            `json:"total_initial_margin_usd"`
	Estimated_liquidation_ratio  float64            `json:"estimated_liquidation_ratio"`
	Session_rpl                  float64            `json:"session_rpl"`
	Fee_balance                  float64            `json:"fee_balance"`
	Total_maintenance_margin_usd float64            `json:"total_maintenance_margin_usd"`
	Options_vega                 float64            `json:"options_vega"`
	Projected_initial_margin     float64            `json:"projected_initial_margin"`
	Options_gamma                float64            `json:"options_gamma"`
	Total_equity_usd             float64            `json:"total_equity_usd"`
	Delta_total                  float64            `json:"delta_total"`
}

type AccountSummariesResult struct {
	Creation_timestamp                   int64            `json:"creation_timestamp"`
	Email                                string           `json:"email"`
	ID                                   int              `json:"id"`
	Interuser_transfers_enabled          bool             `json:"interuser_transfers_enabled"`
	Login_enabled                        bool             `json:"login_enabled"`
	Mmp_enabled                          bool             `json:"mmp_enabled"`
	Referrer_id                          string           `json:"referrer_id"`
	Security_keys_enabled                bool             `json:"security_keys_enabled"`
	Self_trading_extended_to_subaccounts string           `json:"self_trading_extended_to_subaccounts"`
	Self_trading_reject_mode             string           `json:"self_trading_reject_mode"`
	Summaries                            []AccountSummary `json:"summaries"`
	System_name                          string           `json:"system_name"`
	Type                                 string           `json:"type"`
	Username                             string           `json:"username"`
}

type AccountSummariesResponse struct {
	ID      int                    `json:"id"`
	JsonRPC string                 `json:"jsonrpc"`
	Result  AccountSummariesResult `json:"result"`
}

type AccountSummaryResponse struct {
	ID      int            `json:"id"`
	JsonRPC string         `json:"jsonrpc"`
	Result  AccountSummary `json:"result"`
}

// ## Get account summaries list of all account
func GetAccountSummaries(client *DeribitClient, extended bool) error {
	// Create the request message
	msg := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "private/get_account_summaries",
		"params": map[string]interface{}{
			"extended": extended,
		},
	}

	// Marshal the request message to JSON
	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal request message: %w", err)
	}

	// Send the request over the WebSocket
	err = client.conn.WriteMessage(websocket.TextMessage, jsonMsg)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}

	return nil
}

// ## Get one account summary in one currency [BTC/ETH/USDC/USDT/SOL/BNB]
func GetAccountSummary(client *DeribitClient, currency string, extended bool) error {
	// Create the request message
	msg := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "private/get_account_summary",
		"params": map[string]interface{}{
			"currency": currency,
			"extended": extended,
		},
	}

	// Marshal the request message to JSON
	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal request message: %w", err)
	}

	// Send the request over the WebSocket
	err = client.conn.WriteMessage(websocket.TextMessage, jsonMsg)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}

	return nil
}
