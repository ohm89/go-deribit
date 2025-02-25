package api

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/valyala/fasthttp"
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

type AuthError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type AuthResponse struct {
	ID      int        `json:"id"`
	JSONRPC string     `json:"jsonrpc"`
	Result  AuthResult `json:"result,omitempty"`
	Error   *AuthError `json:"error,omitempty"`
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

func Authenticate(c *Client) (*AuthResponse, error) {
	if c.clientID == "" || len(c.clientSecret) == 0 {
		return nil, errors.New("API clientID and clientSecret not configured")
	}

	req, resp := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()

	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}()

	authRequest := &AuthRequest{
		GrantType:    "client_credentials",
		ClientID:     c.clientID,
		ClientSecret: c.clientSecret,
	}

	uri := fmt.Sprintf(
		"%s%s/public/auth?grant_type=%s&client_id=%s&client_secret=%s",
		c.baseURL,
		defaultAPIURL,
		authRequest.GrantType,
		authRequest.ClientID,
		authRequest.ClientSecret,
	)

	req.SetRequestURI(uri)
	req.Header.SetMethod("GET")
	req.Header.Set("Content-Type", "application/json")

	if err := c.client.Do(req, resp); err != nil {
		return nil, err
	}

	var data AuthResponse
	if err := json.Unmarshal(resp.Body(), &data); err != nil {
		return nil, err
	}

	if data.Error != nil {
		return nil, fmt.Errorf("authentication failed: code: %d, message: %s", data.Error.Code, data.Error.Message)
	}

	c.accessToken = data.Result.AccessToken
	c.refreshToken = data.Result.RefreshToken

	return &data, nil
}
