package ws

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type DeribitClient struct {
	conn         *websocket.Conn
	channels     []string
	mu           sync.Mutex
	clientID     string
	clientSecret string
	accessToken  string
	refreshToken string
}

type heartBeatResponse struct {
	params  string          `json:"type"`
	Channel string          `json:"channel,omitempty"`
	Market  string          `json:"market,omitempty"`
	Data    json.RawMessage `json:"data,omitempty"`
	Code    int             `json:"code,omitempty"`
	Msg     string          `json:"msg,omitempty"`
}

// NewDeribitClient is an exported function that creates a new Deribit WebSocket client
func NewDeribitClient(clientID, clientSecret string) *DeribitClient {
	return &DeribitClient{
		clientID:     clientID,
		clientSecret: clientSecret,
	}
}

func (c *DeribitClient) Connect(websocketUrl string) error {

	// ## DEBUG
	fmt.Printf("Web Socket URL: %s \n", websocketUrl)

	// WebSocket connection URL
	u := url.URL{Scheme: "wss", Host: websocketUrl, Path: "/ws/api/v2"}

	// Connect to the WebSocket
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return fmt.Errorf("failed to connect to WebSocket: %w", err)
	}

	c.conn = conn

	// Authenticate the WebSocket connection
	err = c.authenticate()
	if err != nil {
		c.conn.Close()
		return fmt.Errorf("failed to authenticate: %w", err)
	}

	return nil
}

func (c *DeribitClient) GetConn() *websocket.Conn {
	return c.conn
}

func (c *DeribitClient) authenticate() error {
	msg := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "public/auth",
		"params": map[string]string{
			"grant_type":    "client_credentials",
			"client_id":     c.clientID,
			"client_secret": c.clientSecret,
		},
	}

	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal authentication message: %w", err)
	}

	err = c.conn.WriteMessage(websocket.TextMessage, jsonMsg)
	if err != nil {
		return fmt.Errorf("failed to send authentication message: %w", err)
	}

	_, message, err := c.conn.ReadMessage()
	if err != nil {
		return fmt.Errorf("failed to read authentication response: %w", err)
	}

	log.Printf("Received authentication response: %s", message)
	fmt.Printf("Received authentication response: %s", message)

	// Parse the authentication response and save the access_token
	var authResponse map[string]interface{}
	err = json.Unmarshal(message, &authResponse)
	if err != nil {
		return fmt.Errorf("failed to unmarshal authentication response: %w", err)
	}

	result, ok := authResponse["result"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid authentication response format")
	}

	accessToken, ok := result["access_token"].(string)
	if !ok {
		return fmt.Errorf("failed to extract access_token from authentication response")
	}

	c.accessToken = accessToken

	refreshToken, ok := result["refresh_token"].(string)
	if !ok {
		return fmt.Errorf("failed to extract refresh_token from authentication response")
	}

	c.refreshToken = refreshToken

	// ## [DEBUG]
	// fmt.Printf("\n\n Access Token: %s \n\n", c.accessToken)
	// fmt.Printf("\n\n Refresh Token: %s \n\n", c.refreshToken)

	return nil
}

func (c *DeribitClient) Subscribe(channels ...string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.channels = append(c.channels, channels...)

	msg := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  "public/subscribe",
		"params": map[string]interface{}{
			"channels": c.channels,
		},
	}

	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		log.Fatalf("failed to marshal subscription message: %v", err)
	}

	err = c.conn.WriteMessage(websocket.TextMessage, jsonMsg)
	if err != nil {
		log.Fatalf("failed to send subscription message: %v", err)
	}
}

func (c *DeribitClient) ReceiveRaw() ([]byte, error) {
	_, message, err := c.conn.ReadMessage()
	if err != nil {
		log.Fatalf("failed to read message: %v", err)
	}
	fmt.Printf("Received from server: %s\n", message)

	return message, err
}

func (c *DeribitClient) Ping() error {
	msg := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "public/test",
		"params":  map[string]interface{}{},
	}

	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		log.Fatalf("failed to marshal test(ping) message: %v", err)
	}

	err = c.conn.WriteMessage(websocket.TextMessage, jsonMsg)
	if err != nil {
		log.Fatalf("failed to send test(ping) message: %v", err)
	}

	return err
}

func (c *DeribitClient) PingRegular(ctx context.Context, duration time.Duration) {
	go func() {
		t := time.NewTicker(duration)
		for {
			select {
			case <-ctx.Done():
				return
			case <-t.C:
				if err := c.Ping(); err != nil {
					return
				}
			}
		}
	}()
}

func (c *DeribitClient) SetHeartBeat() error {
	msg := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "public/set_heartbeat",
		"params": map[string]interface{}{
			"interval": 20,
		},
	}

	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		log.Fatalf("failed to marshal SetHeartBeat message: %v", err)
	}

	err = c.conn.WriteMessage(websocket.TextMessage, jsonMsg)
	if err != nil {
		log.Fatalf("failed to send SetHeartBeat message: %v", err)
	}

	return err
}

func (c *DeribitClient) handleHeartbeat() error {
	log.Println("Received heartbeat message")
	fmt.Println("Received heartbeat message")

	// Prepare the heartbeat response
	err := c.Ping()

	return err
}

func (c *DeribitClient) HandleMessage(message []byte) error {
	var msg map[string]interface{}
	err := json.Unmarshal(message, &msg)
	if err != nil {
		return fmt.Errorf("failed to unmarshal message: %w", err)
	}

	method, _ := msg["method"].(string)
	// if !ok {
	// 	return fmt.Errorf("message does not have a 'method' field")
	// }

	switch method {
	case "heartbeat":
		return c.handleHeartbeat()
	default:
		fmt.Printf("Received message: %s\n", message)
		return nil
	}
}

// ## Hello to set program for deribit to known software
func (c *DeribitClient) Hello(softwareClientName string, softwareClientVersion string) error {
	msg := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "public/hello",
		"params": map[string]interface{}{
			"client_name":    softwareClientName,
			"client_version": softwareClientVersion,
		},
	}

	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		log.Fatalf("failed to marshal Hello message: %v", err)
	}

	err = c.conn.WriteMessage(websocket.TextMessage, jsonMsg)
	if err != nil {
		log.Fatalf("failed to send Hello message: %v", err)
	}

	return err
}

// ## Main Run For Testing client
func (c *DeribitClient) Run() {
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			log.Fatalf("failed to read message: %v", err)
		}
		fmt.Printf("Received from server: %s\n", message)
	}
}
