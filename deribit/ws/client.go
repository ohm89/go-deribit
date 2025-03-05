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
	websocketUrl string
	conn         *websocket.Conn
	channels     []string
	mu           sync.Mutex
	clientID     string
	clientSecret string
	accessToken  string
	refreshToken string
	isPrivate    bool
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
	fmt.Printf("Web Socket URL: %s \n\n", websocketUrl)

	// WebSocket connection URL
	u := url.URL{Scheme: "wss", Host: websocketUrl, Path: "/ws/api/v2"}

	// Connect to the WebSocket
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return fmt.Errorf("failed to connect to WebSocket: %w", err)
	}

	c.conn = conn
	c.websocketUrl = websocketUrl

	return nil
}

func (c *DeribitClient) GetConn() *websocket.Conn {
	return c.conn
}

func (c *DeribitClient) Subscribe(channels ...string) error {
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
		// log.Fatalf("failed to marshal subscription message: %v", err)
		return err
	}

	err = c.conn.WriteMessage(websocket.TextMessage, jsonMsg)
	if err != nil {
		// log.Fatalf("failed to send subscription message: %v", err)
		return err
	}

	return nil
}

func (c *DeribitClient) PrivateSubscribe(channels ...string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.channels = append(c.channels, channels...)

	msg := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  "private/subscribe",
		"params": map[string]interface{}{
			"channels": c.channels,
		},
	}

	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		// log.Fatalf("failed to marshal subscription message: %v", err)
		return err
	}

	err = c.conn.WriteMessage(websocket.TextMessage, jsonMsg)
	if err != nil {
		// log.Fatalf("failed to send subscription message: %v", err)
		return err
	}

	return nil
}

func (c *DeribitClient) Unsubscribe(channels ...string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Remove the channels from the list of subscribed channels
	newChannels := make([]string, 0, len(c.channels))
	for _, ch := range c.channels {
		found := false
		for _, unsubCh := range channels {
			if ch == unsubCh {
				found = true
				break
			}
		}
		if !found {
			newChannels = append(newChannels, ch)
		}
	}
	c.channels = newChannels

	msg := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  "public/unsubscribe",
		"params": map[string]interface{}{
			"channels": channels,
		},
	}

	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	err = c.conn.WriteMessage(websocket.TextMessage, jsonMsg)
	if err != nil {
		return err
	}

	return nil
}

func (c *DeribitClient) PrivateUnsubscribe(channels ...string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Remove the channels from the list of subscribed channels
	newChannels := make([]string, 0, len(c.channels))
	for _, ch := range c.channels {
		found := false
		for _, unsubCh := range channels {
			if ch == unsubCh {
				found = true
				break
			}
		}
		if !found {
			newChannels = append(newChannels, ch)
		}
	}
	c.channels = newChannels

	msg := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  "private/unsubscribe",
		"params": map[string]interface{}{
			"channels": channels,
		},
	}

	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	err = c.conn.WriteMessage(websocket.TextMessage, jsonMsg)
	if err != nil {
		return err
	}

	return nil
}

func (c *DeribitClient) UnsubscribeAll() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	msg := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  "public/unsubscribe_all",
		"params":  map[string]interface{}{},
	}

	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	err = c.conn.WriteMessage(websocket.TextMessage, jsonMsg)
	if err != nil {
		return err
	}

	return nil
}

func (c *DeribitClient) PrivateUnsubscribeAll() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	msg := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  "private/unsubscribe_all",
		"params":  map[string]interface{}{},
	}

	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	err = c.conn.WriteMessage(websocket.TextMessage, jsonMsg)
	if err != nil {
		return err
	}

	return nil
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

func (c *DeribitClient) SetHeartBeat(interval int) error {
	msg := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "public/set_heartbeat",
		"params": map[string]interface{}{
			"interval": interval, // ## In second
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

	// Prepare the heartbeat response
	err := c.Ping()

	return err
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
func (c *DeribitClient) Close() {
	c.conn.Close()
}

func (c *DeribitClient) handleTextMessage(message []byte) {
	var msg map[string]interface{}
	err := json.Unmarshal(message, &msg)
	if err != nil {
		log.Printf("failed to unmarshal message: %v", err)
		return
	}

	method, _ := msg["method"].(string)
	switch method {
	case "heartbeat":
		c.handleHeartbeat()
	default:
		fmt.Printf("Received message [Private: %t]: %s\n\n", c.isPrivate, message)
	}
}

// ## Reconnect
func (c *DeribitClient) reconnect() error {
	// Close the existing connection
	c.conn.Close()

	// Reconnect to the WebSocket
	err := c.Connect(c.websocketUrl)
	if err != nil {
		return err
	}

	// Resubscribe to the channels
	if c.isPrivate {
		err = c.PrivateSubscribe(c.channels...)
	} else {
		err = c.Subscribe(c.channels...)
	}
	if err != nil {
		return err
	}

	return nil
}

// ## Main Run Client in loop
func (c *DeribitClient) Run() {
	for {
		messageType, message, err := c.conn.ReadMessage()
		if err != nil {
			// Check if the error is a WebSocket error
			if _, ok := err.(*websocket.CloseError); ok {
				// Reconnect the WebSocket
				err = c.reconnect()
				if err != nil {
					log.Printf("failed to reconnect: %v", err)
					return
				}
				continue
			}
			log.Printf("failed to read message: %v", err)
			continue
		}

		switch messageType {
		case websocket.TextMessage:
			c.handleTextMessage(message)
		case websocket.BinaryMessage:
			fmt.Printf("Received binary message from server: %v\n", message)
		default:
			log.Printf("unexpected message type: %d", messageType)
		}
	}
}

// ## ----------------- Event --------------

type WebSocketResponse struct {
	JSONRPC string          `json:"jsonrpc"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params"`
}

type ChannelInfo struct {
	Channel string          `json:"channel"`
	Data    json.RawMessage `json:"data"`
}

func (c *DeribitClient) Receive() (*WebSocketResponse, error) {
	// Read a message from the WebSocket connection
	_, message, err := c.conn.ReadMessage()
	if err != nil {
		return nil, fmt.Errorf("failed to read message: %w", err)
	}

	// Parse the base WebSocket response
	var wsResponse WebSocketResponse
	if err := json.Unmarshal(message, &wsResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal WebSocket response: %w", err)
	}

	// If the response is not a subscription, check is heartbeat, if not return as is
	if wsResponse.Method != "subscription" {
		if wsResponse.Method == "heartbeat" {
			c.handleHeartbeat()
		} else {
			return &wsResponse, nil
		}
	}

	// Extract channel information
	var channelInfo ChannelInfo
	if err := json.Unmarshal(wsResponse.Params, &channelInfo); err != nil {
		return nil, fmt.Errorf("failed to unmarshal channel info: %w", err)
	}

	return &wsResponse, nil
}
