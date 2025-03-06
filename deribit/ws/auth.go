package ws

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type AuthResult struct {
	AccessToken        string   `json:"access_token"`
	EnabledFeatures    []string `json:"enabled_features"`
	ExpiresIn          int      `json:"expires_in"`
	GoogleLogin        bool     `json:"google_login"`
	MandatoryTFAStatus string   `json:"mandatory_tfa_status"`
	RefreshToken       string   `json:"refresh_token"`
	Scope              string   `json:"scope"`
	SID                string   `json:"sid,omitempty"`
	State              string   `json:"state,omitempty"`
	TokenType          string   `json:"token_type"`
}

type AuthResponse struct {
	ID      int        `json:"id"`
	JSONRPC string     `json:"jsonrpc"`
	Result  AuthResult `json:"result"`
}

type AuthRequest struct {
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RefreshToken string `json:"refresh_token,omitempty"`
	Timestamp    int64  `json:"timestamp,omitempty"`
	Signature    string `json:"signature,omitempty"`
	Nonce        string `json:"nonce,omitempty"`
	Data         string `json:"data,omitempty"`
	State        string `json:"state,omitempty"`
	Scope        string `json:"scope,omitempty"`
}

func Authenticate(c *DeribitClient) (*AuthResponse, error) {

	authRequest := &AuthRequest{
		GrantType:    "client_credentials",
		ClientID:     c.clientID,
		ClientSecret: c.clientSecret,
	}

	msg := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "public/auth",
		"params":  authRequest,
	}

	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal authentication message: %w", err)
	}

	err = c.conn.WriteMessage(websocket.TextMessage, jsonMsg)
	if err != nil {
		return nil, fmt.Errorf("failed to send authentication message: %w", err)
	}

	_, message, err := c.conn.ReadMessage()
	if err != nil {
		return nil, fmt.Errorf("failed to read authentication response: %w", err)
	}

	log.Printf("Received authentication response: %s \n\n", message)

	// Parse the authentication response and save the access_token
	var authResponse AuthResponse
	err = json.Unmarshal(message, &authResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal authentication response: %w", err)
	}

	result := authResponse.Result

	accessToken := result.AccessToken
	refreshToken := result.RefreshToken

	c.accessToken = accessToken
	c.refreshToken = refreshToken

	c.isPrivate = true

	// ## [DEBUG]
	// fmt.Printf("\n\n Access Token: %s \n\n", c.accessToken)
	// fmt.Printf("\n\n Refresh Token: %s \n\n", c.refreshToken)

	return &authResponse, nil
}

func RefreshAuth(c *DeribitClient) (*AuthResponse, error) {

	authRequest := &AuthRequest{
		GrantType:    "refresh_token",
		RefreshToken: c.refreshToken,
	}

	msg := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "public/auth",
		"params":  authRequest,
	}

	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal authentication message: %w", err)
	}

	err = c.conn.WriteMessage(websocket.TextMessage, jsonMsg)
	if err != nil {
		return nil, fmt.Errorf("failed to send authentication message: %w", err)
	}

	_, message, err := c.conn.ReadMessage()
	if err != nil {
		return nil, fmt.Errorf("failed to read authentication response: %w", err)
	}

	log.Printf("Received authentication response: %s \n\n", message)

	// Parse the authentication response and save the access_token
	var authResponse AuthResponse
	err = json.Unmarshal(message, &authResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal authentication response: %w", err)
	}

	result := authResponse.Result

	accessToken := result.AccessToken
	refreshToken := result.RefreshToken

	c.accessToken = accessToken
	c.refreshToken = refreshToken

	c.isPrivate = true

	// ## [DEBUG]
	// fmt.Printf("\n\n Access Token: %s \n\n", c.accessToken)
	// fmt.Printf("\n\n Refresh Token: %s \n\n", c.refreshToken)

	return &authResponse, nil
}
