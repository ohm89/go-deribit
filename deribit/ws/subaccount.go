package ws

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

type PortfolioItem struct {
	AdditionalReserve        float64 `json:"additional_reserve"`
	AvailableFunds           float64 `json:"available_funds"`
	AvailableWithdrawalFunds float64 `json:"available_withdrawal_funds"`
	Balance                  float64 `json:"balance"`
	Currency                 string  `json:"currency"`
	Equity                   float64 `json:"equity"`
	InitialMargin            float64 `json:"initial_margin"`
	MaintenanceMargin        float64 `json:"maintenance_margin"`
	MarginBalance            float64 `json:"margin_balance"`
	SpotReserve              float64 `json:"spot_reserve"`
}

type Portfolio struct {
	BTC PortfolioItem `json:"btc"`
	ETH PortfolioItem `json:"eth"`
}

type SubAccount struct {
	Email                   string    `json:"email"`
	ID                      int       `json:"id"`
	IsPassword              bool      `json:"is_password"`
	LoginEnabled            bool      `json:"login_enabled"`
	MarginModel             string    `json:"margin_model"`
	NotConfirmedEmail       string    `json:"not_confirmed_email"`
	Portfolio               Portfolio `json:"portfolio"`
	ProofID                 string    `json:"proof_id"`
	ProofIDSignature        string    `json:"proof_id_signature"`
	ReceiveNotifications    bool      `json:"receive_notifications"`
	SecurityKeysAssignments []string  `json:"security_keys_assignments"`
	SecurityKeysEnabled     bool      `json:"security_keys_enabled"`
	SystemName              string    `json:"system_name"`
	Type                    string    `json:"type"`
	Username                string    `json:"username"`
}

type SubAccountsResponse struct {
	ID      int          `json:"id"`
	JSONRPC string       `json:"jsonrpc"`
	Result  []SubAccount `json:"result"`
}

type SubAccountsDetail struct {
	OpenOrders []OpenOrder `json:"open_orders"`
	Positions  []Position  `json:"positions"`
	UID        int         `json:"uid"`
}

type SubAccountsDetailsResponse struct {
	ID      int               `json:"id"`
	JSONRPC string            `json:"jsonrpc"`
	Result  SubAccountsDetail `json:"result"`
}

// ## Get All Subaccounts Details
func GetSubAccounts(client *DeribitClient, withPortfolio bool) error {

	// Create the request message
	msg := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "private/get_subaccounts",
		"params": map[string]interface{}{
			"with_portfolio": withPortfolio,
		},
	}

	// Marshal the request message to JSON
	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal request GetSubAccounts message: %w", err)
	}

	// Send the request over the WebSocket
	err = client.conn.WriteMessage(websocket.TextMessage, jsonMsg)
	if err != nil {
		return fmt.Errorf("failed to send GetSubAccounts request: %w", err)
	}

	return nil
}

// ## Get subaccounts positions
func GetSubAccountsDetails(client *DeribitClient, currency string, withOpenOrders bool) error {

	// Create the request message
	msg := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "private/get_subaccounts_details",
		"params": map[string]interface{}{
			"currency":         currency,
			"with_open_orders": withOpenOrders,
		},
	}

	// Marshal the request message to JSON
	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal request GetSubAccountsDetails message: %w", err)
	}

	// Send the request over the WebSocket
	err = client.conn.WriteMessage(websocket.TextMessage, jsonMsg)
	if err != nil {
		return fmt.Errorf("failed to send GetSubAccountsDetails request: %w", err)
	}

	return nil
}
