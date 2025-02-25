package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"bitbucket.org/ohm89/go-deribit/deribit/api"
	"bitbucket.org/ohm89/go-deribit/deribit/ws"
)

type Config struct {
	API_URL       string `json:"API_URL,omitempty"`
	WS_URL        string `json:"WS_URL,omitempty"`
	CLIENT_ID     string `json:"CLIENT_ID,omitempty"`
	CLIENT_SECRET string `json:"CLIENT_SECRET,omitempty"`
	SUBACCOUNT    string `json:"SUBACCOUNT,omitempty"`
	USER_ID       uint64 `json:"USER_ID,omitempty"`
	NAME          string `json:"NAME,omitempty"`
	VERSION       string `json:"VERSION,omitempty"`
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

	isUseWebSocket := false
	isUseAPI := true

	// ## ------ Websocket Testing and usage --------------
	if isUseWebSocket {

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

		// ## -------------- Test Create Buy Order ---------------------
		// // ## Create a new buy order request
		// orderRequest := &ws.OrderRequest{
		// 	InstrumentName: "BTC_USDC",
		// 	Amount:         0.0001,
		// 	Price:          18000,
		// 	Type:           "limit",
		// 	Label:          "limit0000243",
		// }

		// errPrivate = ws.CreateBuyOrder(privateClient, orderRequest)
		// if errPrivate != nil {
		// 	privateClient.Close()
		// 	log.Fatalf("failed to create buy order: %v", errPrivate)
		// }

		// orderRequest2 := &ws.OrderRequest{
		// 	InstrumentName: "BTC_USDC",
		// 	Amount:         0.0001,
		// 	Price:          18000,
		// 	Type:           "limit",
		// 	Label:          "limit0000244",
		// }

		// errPrivate = ws.CreateBuyOrder(privateClient, orderRequest2)
		// if errPrivate != nil {
		// 	privateClient.Close()
		// 	log.Fatalf("failed to create buy order: %v", errPrivate)
		// }

		// ## -------------- Test Create Sell Order ---------------------
		// // ## Create a new buy order request
		// orderRequest := &ws.OrderRequest{
		// 	InstrumentName: "BTC_USDC",
		// 	Amount:         0.0001,
		// 	Price:          200000,
		// 	Type:           "limit",
		// 	Label:          "limitSell0000243",
		// }

		// errPrivate = ws.CreateSellOrder(privateClient, orderRequest)
		// if errPrivate != nil {
		// 	privateClient.Close()
		// 	log.Fatalf("failed to create buy order: %v", errPrivate)
		// }

		// orderRequest2 := &ws.OrderRequest{
		// 	InstrumentName: "BTC_USDC",
		// 	Amount:         0.0001,
		// 	Price:          380000,
		// 	Type:           "limit",
		// 	Label:          "limitSell0000244",
		// }

		// errPrivate = ws.CreateSellOrder(privateClient, orderRequest2)
		// if errPrivate != nil {
		// 	privateClient.Close()
		// 	log.Fatalf("failed to create buy order: %v", errPrivate)
		// }

		// ## -------------- Test Cancel Order ---------------------

		// errPrivate = ws.CancelOneOrder(privateClient, "BTC_USDC-2731791644")
		// if errPrivate != nil {
		// 	privateClient.Close()
		// 	log.Fatalf("failed to create buy order: %v", errPrivate)
		// }

		// ## -------------- Test Cancel All ---------------------

		// errPrivate = ws.CancelAllOrders(privateClient)
		// if errPrivate != nil {
		// 	privateClient.Close()
		// 	log.Fatalf("failed to create buy order: %v", errPrivate)
		// }

		// ## -------------- Test Get Account ---------------------

		// // ## GEt All Currencies in account
		// errPrivate = ws.GetAccountSummaries(privateClient, false)
		// if errPrivate != nil {
		// 	privateClient.Close()
		// 	log.Fatalf("failed to GetAccountSummaries: %v", errPrivate)
		// }

		// // ## Get One Currencies in account
		// errPrivate = ws.GetAccountSummary(privateClient, "USDC", false)
		// if errPrivate != nil {
		// 	privateClient.Close()
		// 	log.Fatalf("failed to GetAccountSummary: %v", errPrivate)
		// }

		// errPrivate = ws.GetAccountSummary(privateClient, "BTC", false)
		// if errPrivate != nil {
		// 	privateClient.Close()
		// 	log.Fatalf("failed to GetAccountSummary: %v", errPrivate)
		// }

		// ## -------------- Test Get SubAccount ---------------------

		// errPrivate = ws.GetSubAccounts(privateClient, false)
		// if errPrivate != nil {
		// 	privateClient.Close()
		// 	log.Fatalf("failed to GetSubAccounts: %v", errPrivate)
		// }

		// errPrivate = ws.GetSubAccountsDetails(privateClient, "USDC", true)
		// if errPrivate != nil {
		// 	privateClient.Close()
		// 	log.Fatalf("failed to GetSubAccountsDetails: %v", errPrivate)
		// }

		// ## -------------- Test Get Position ---------------------
		// errPrivate = ws.GetPositions(privateClient, "USDC", "future")
		// if errPrivate != nil {
		// 	privateClient.Close()
		// 	log.Fatalf("failed to GetPositions: %v", errPrivate)
		// }

		// errPrivate = ws.GetPosition(privateClient, "BTC-PERPETUAL")
		// if errPrivate != nil {
		// 	privateClient.Close()
		// 	log.Fatalf("failed to GetPositions: %v", errPrivate)
		// }

		// ## -------------- Main Loop (Concurrent GO) ---------------------
		// Start concurrent tasks for HandleReadMessage and HandleHeartBeatMessage
		// Start concurrent tasks for HandleReadMessage and HandleHeartBeatMessage
		// go client.Run()
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

	// ## ------ API Testing and usage --------------
	if isUseAPI {

		// ## Create New http Client
		apiClient := api.New(
			"https://"+config.API_URL,
			config.CLIENT_ID,
			config.CLIENT_SECRET,
		)

		// ## ------- Test [Buy] Order ----------
		orderBuyResponse, err := apiClient.Orders.Buy(
			"BTC_USDC",
			0.0001,
			0,
			"limit",
			"my-btc-usdc-order",
			18000.0,
			"good_til_cancelled",
			0,
			false,
			false,
			false,
			0,
			0,
			"",
			"",
			false,
			0,
			"",
			"",
			nil,
		)
		if err != nil {
			log.Fatalf("failed [Buy] API: %+v", err)
		}

		fmt.Println("[1] Buy Resp: ")
		fmt.Printf("%#v", orderBuyResponse)
		fmt.Printf("\n\n")

		// ## ------- Test [CancelOneOrder] Order ----------
		if orderBuyResponse.Result.Order.OrderID != "" {
			orderCancelOneOrderResponse, err := apiClient.Orders.Cancel(
				orderBuyResponse.Result.Order.OrderID,
			)
			if err != nil {
				log.Fatalf("failed [Cancel] API: %+v", err)
			}

			fmt.Println("[3] Cancel Resp: ")
			fmt.Printf("%#v", orderCancelOneOrderResponse)
			fmt.Printf("\n\n")
		}

		// ## ------- Test [Buy] Order 2 ----------
		orderBuyResponse2, err := apiClient.Orders.Buy(
			"BTC_USDC",
			0.0001,
			0,
			"limit",
			"my-btc-usdc-order",
			19000.0,
			"good_til_cancelled",
			0,
			false,
			false,
			false,
			0,
			0,
			"",
			"",
			false,
			0,
			"",
			"",
			nil,
		)
		if err != nil {
			log.Fatalf("failed [Buy] API 2: %+v", err)
		}

		fmt.Println("[1] Buy Resp2: ")
		fmt.Printf("%#v", orderBuyResponse2)
		fmt.Printf("\n\n")

		// ## ------- Test [CancelAll] Order ----------

		orderCancelAllOrderResponse, err := apiClient.Orders.CancelAll()
		if err != nil {
			log.Fatalf("failed [CancelAll] API: %+v", err)
		}

		fmt.Println("[4] CancelAll Resp: ")
		fmt.Printf("%#v", orderCancelAllOrderResponse)
		fmt.Printf("\n\n")

		// ## ------- Test [Buy] Order 3 ----------
		orderBuyResponse3, err := apiClient.Orders.Buy(
			"BTC_USDC",
			0.0001,
			0,
			"limit",
			"my-btc-usdc-order",
			20000.0,
			"good_til_cancelled",
			0,
			false,
			false,
			false,
			0,
			0,
			"",
			"",
			false,
			0,
			"",
			"",
			nil,
		)
		if err != nil {
			log.Fatalf("failed [Buy] API 2: %+v", err)
		}

		fmt.Println("[33] Buy Resp3: ")
		fmt.Printf("%#v", orderBuyResponse3)
		fmt.Printf("\n\n")

		// ## ------- Test [Buy] Order 4 ----------
		// orderBuyResponse4, err := apiClient.Orders.Buy(
		// 	"SOL_USDC",
		// 	1,
		// 	0,
		// 	"limit",
		// 	"my-btc-usdc-order",
		// 	10.0,
		// 	"good_til_cancelled",
		// 	0,
		// 	false,
		// 	false,
		// 	false,
		// 	0,
		// 	0,
		// 	"",
		// 	"",
		// 	false,
		// 	0,
		// 	"",
		// 	"",
		// 	nil,
		// )
		// if err != nil {
		// 	log.Fatalf("failed [Buy] API 4: %+v", err)
		// }

		// fmt.Println("[1] Buy Resp4: ")
		// fmt.Printf("%#v", orderBuyResponse4)
		// fmt.Printf("\n\n")

		// ## ------- Test [CancelByInstrument] Order ----------

		orderCancelByInstrumentResponse, err := apiClient.Orders.CancelAllByInstrument(
			"BTC_USDC",
			"all",
			false,
			false,
			false,
		)
		if err != nil {
			log.Fatalf("failed [CancelByInstrument] API: %+v", err)
		}

		fmt.Println("[55] CancelByInstrument Resp: ")
		fmt.Printf("%#v", orderCancelByInstrumentResponse)
		fmt.Printf("\n\n")

		// ## ------- Test [Buy] Order 6 ----------
		orderBuyResponse6, err := apiClient.Orders.Buy(
			"SOL_USDC",
			1,
			0,
			"limit",
			"my-sol-6-usdc-orde-001",
			10.0,
			"good_til_cancelled",
			0,
			false,
			false,
			false,
			0,
			0,
			"",
			"",
			false,
			0,
			"",
			"",
			nil,
		)
		if err != nil {
			log.Fatalf("failed [Buy] API 6: %+v", err)
		}

		fmt.Println("[1] Buy Resp6: ")
		fmt.Printf("%#v", orderBuyResponse6)
		fmt.Printf("\n\n")

		// ## ------- Test [GetOrderState] Order 6 ----------
		if orderBuyResponse6.Result.Order.OrderID != "" {
			orderState6, err := apiClient.Orders.GetOrderState(
				orderBuyResponse6.Result.Order.OrderID,
			)
			if err != nil {
				log.Fatalf("failed [GetOrderState] API: %+v", err)
			}

			fmt.Println("[66] GetOrderState Resp: ")
			fmt.Printf("%#v", orderState6)
			fmt.Printf("\n\n")
		}

		// ## ------- Test [GetOrderStateByLabel] Order 6 ----------
		if orderBuyResponse6.Result.Order.OrderID != "" {
			orderState67, err := apiClient.Orders.GetOrderStateByLabel(
				"USDC",
				orderBuyResponse6.Result.Order.Label,
			)
			if err != nil {
				log.Fatalf("failed [GetOrderStateByLabel] API: %+v", err)
			}

			fmt.Println("[67] GetOrderStateByLabel Resp: ")
			fmt.Printf("%#v", orderState67)
			fmt.Printf("\n\n")
		}

		// ## ------- Test [GetOpenOrders] Order 6 ----------

		openSpotOrderList, err := apiClient.Orders.GetOpenOrders(
			"spot",
			"all",
		)
		if err != nil {
			log.Fatalf("failed [GetOpenOrders] API: %+v", err)
		}

		fmt.Println("[68] GetOpenOrders Resp: ")
		fmt.Printf("%#v", openSpotOrderList)
		fmt.Printf("\n\n")

		// ## ------- Test [GetOpenOrdersByInstrument] Order 6 ----------

		openSpotOrderList2, err := apiClient.Orders.GetOpenOrdersByInstrument(
			"SOL_USDC",
			"all",
		)
		if err != nil {
			log.Fatalf("failed [GetOpenOrdersByInstrument] API: %+v", err)
		}

		fmt.Println("[68] GetOpenOrdersByInstrument Resp: ")
		fmt.Printf("%#v", openSpotOrderList2)
		fmt.Printf("\n\n")

		// ## ------- Test [GetOrderHistoryByCurrency] Order 6 ----------

		orderHistoryList1, err := apiClient.Orders.GetOrderHistoryByCurrency(
			"USDC",
			"spot",
			20,
			0,
			false,
			false,
		)
		if err != nil {
			log.Fatalf("failed [GetOrderHistoryByCurrency] API: %+v", err)
		}

		fmt.Println("[68] GetOrderHistoryByCurrency Resp: ")
		fmt.Printf("%#v", orderHistoryList1)
		fmt.Printf("\n\n")

		// ## ------- Test [GetOrderHistoryByInstrument] Order 6 ----------

		orderHistoryList2, err := apiClient.Orders.GetOrderHistoryByInstrument(
			"SOL_USDC",
			20,
			0,
			false,
			true,
		)
		if err != nil {
			log.Fatalf("failed [GetOrderHistoryByInstrument] API: %+v", err)
		}

		fmt.Println("[69] GetOrderHistoryByInstrument Resp: ")
		fmt.Printf("%#v", orderHistoryList2)
		fmt.Printf("\n\n")

		// ## ------- Test [GetTriggerOrderHistory] Order 6 ----------

		triggerOrderHistoryList1, err := apiClient.Orders.GetTriggerOrderHistory(
			"USDC",
			"",
			20,
			"",
		)
		if err != nil {
			log.Fatalf("failed [GetTriggerOrderHistory] API: %+v", err)
		}

		fmt.Println("[70] GetTriggerOrderHistory Resp: ")
		fmt.Printf("%#v", triggerOrderHistoryList1)
		fmt.Printf("\n\n")

		// ## ------- Test [CancelAll] Order ----------

		orderCancelAllLast, err := apiClient.Orders.CancelAll()
		if err != nil {
			log.Fatalf("failed [CancelAll] API: %+v", err)
		}

		fmt.Println("[9999] CancelAll Resp: ")
		fmt.Printf("%#v", orderCancelAllLast)
		fmt.Printf("\n\n")

		// ## ------- Test [Sell] Order ----------
		// orderResponse2, err := apiClient.Orders.Sell(
		// 	"BTC_USDC",
		// 	0.0001,
		// 	0,
		// 	"limit",
		// 	"my-btc-usdc-order",
		// 	180000.0,
		// 	"good_til_cancelled",
		// 	0,
		// 	false,
		// 	false,
		// 	false,
		// 	0,
		// 	0,
		// 	"",
		// 	"",
		// 	false,
		// 	0,
		// 	"",
		// 	"",
		// 	nil,
		// )
		// if err != nil {
		// 	log.Fatalf("failed [Sell] API: %+v", err)
		// }

		// fmt.Println("[2] Sell Resp: ")
		// fmt.Printf("%#v", orderResponse2)
		// fmt.Printf("\n\n")

		// ## ------- Test [GetAccountSummaries] ----------
		// dataResponse, err := apiClient.Accounts.GetAccountSummaries(true)
		// if err != nil {
		// 	log.Fatalf("failed [GetAccountSummaries] API: %v", err)
		// }

		// fmt.Println("[1] GetAccountSummaries Resp: ")
		// fmt.Printf("%#v", dataResponse)
		// fmt.Printf("\n\n")

		// ## ------- Test [GetAccountSummary] ----------
		// dataResponse2, err := apiClient.Accounts.GetAccountSummary("USDC", true)
		// if err != nil {
		// 	log.Fatalf("failed [GetAccountSummary] API: %v", err)
		// }

		// fmt.Println("[2] GetAccountSummary Resp: ")
		// fmt.Printf("%#v", dataResponse2)
		// fmt.Printf("\n\n")

	}
}
