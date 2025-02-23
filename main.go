package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"bitbucket.org/ohm89/go-deribit/deribit/ws"
)

type Config struct {
	API_URL       string `json:"API_URL"`
	WS_URL        string `json:"WS_URL"`
	CLIENT_ID     string `json:"CLIENT_ID"`
	CLIENT_SECRET string `json:"CLIENT_SECRET`
	SUBACCOUNT    string `json:"SUBACCOUNT"`
}

const (
	pathConfig = "./config/development.json"
)

func main() {
	fmt.Println("********* Start Program ***********")

	// ## Get Config File
	configFile, _ := os.ReadFile(pathConfig)

	config := Config{}

	_ = json.Unmarshal([]byte(configFile), &config)

	client := ws.NewDeribitClient(config.CLIENT_ID, config.CLIENT_SECRET)

	err := client.Connect(config.WS_URL)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	err3 := client.Hello("test-go-deribit", "0.0.1")
	if err3 != nil {
		log.Fatalf("failed to hello: %v", err)
	}

	err2 := client.SetHeartBeat(60)
	if err2 != nil {
		log.Fatalf("failed to set heart beat: %v", err)
	}

	// Subscribe to multiple channels
	client.Subscribe(
		"deribit_price_index.btc_usd",
		"deribit_price_index.btc_usdc",
		"deribit_price_index.btc_usdt",
	)

	// Start listening for messages
	// client.Run()

	// Start listening for messages
	conn := client.GetConn()

	for {
		_, message, err := conn.ReadMessage()

		if err != nil {
			log.Fatalf("failed to read message: %v", err)
		}

		err2 := client.HandleMessage(message)
		if err2 != nil {
			log.Fatalf("failed to handle message: %v", err)
		}

	}
}
