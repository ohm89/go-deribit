package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"bitbucket.org/ohm89/go-deribit/deribit/ws"
)

type Config struct {
	API_URL       string `json:"API_URL"`
	WS_URL        string `json:"WS_URL"`
	CLIENT_ID     string `json:"CLIENT_ID"`
	CLIENT_SECRET string `json:"CLIENT_SECRET`
	SUBACCOUNT    string `json:"SUBACCOUNT"`
	NAME          string `json:"NAME"`
	VERSION       string `json:"VERSION"`
}

const (
	pathConfig = "./config/development.json"
)

func main() {

	// ## Get Config File
	configFile, _ := os.ReadFile(pathConfig)

	config := Config{}

	_ = json.Unmarshal([]byte(configFile), &config)

	fmt.Printf("\n\n********* Start Program %s *********** \n\n", config.NAME)

	// ## ------ Create Market Client connection --------------
	client := ws.NewDeribitClient(config.CLIENT_ID, config.CLIENT_SECRET)

	err := client.Connect(config.WS_URL)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	// ## Send Hello Software the WebSocket connection
	err = client.Hello(config.NAME, config.VERSION)
	if err != nil {
		client.Close()
		log.Fatalf("failed to hello: %v", err)
	}

	// ## Set WebSocket HeartBeat Interval
	err = client.SetHeartBeat(60)
	if err != nil {
		client.Close()
		log.Fatalf("failed to set heart beat: %v", err)
	}

	// ## Subscribe to multiple channels
	err = client.Subscribe(
		"deribit_price_index.btc_usd",
		"deribit_price_index.btc_usdc",
		"deribit_price_index.btc_usdt",
	)
	if err != nil {
		client.Close()
		log.Fatalf("failed to subscribe : %v", err)
	}

	// ## ------ Create Private Client connection for trade --------------

	privateClient := ws.NewDeribitClient(config.CLIENT_ID, config.CLIENT_SECRET)

	errPrivate := privateClient.Connect(config.WS_URL)
	if errPrivate != nil {
		log.Fatalf("failed to connect: %v", errPrivate)
	}

	// ## Send Hello Software the WebSocket connection
	errPrivate = privateClient.Hello(config.NAME, config.VERSION)
	if errPrivate != nil {
		privateClient.Close()
		log.Fatalf("failed to hello: %v", errPrivate)
	}

	// ## Set WebSocket HeartBeat Interval
	errPrivate = client.SetHeartBeat(60)
	if errPrivate != nil {
		privateClient.Close()
		log.Fatalf("failed to set heart beat: %v", errPrivate)
	}

	// ## Authenticate the WebSocket connection
	_, err4 := ws.Authenticate(privateClient)
	if err4 != nil {
		client.Close()
		log.Fatalf("failed to authenticate: %v", err4)
	}

	// ## -------------- Main Loop (Concurrent GO) ---------------------
	// Start concurrent tasks for HandleReadMessage and HandleHeartBeatMessage
	// Start concurrent tasks for HandleReadMessage and HandleHeartBeatMessage
	go client.Run()
	go privateClient.Run()

	// ## -------------- Termination ---------------------
	// ## Wait for termination signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan

	// ## Gracefully shut down the client
	client.Close()
	privateClient.Close()
	log.Println("Shutting down...")

}
