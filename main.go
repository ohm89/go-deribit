package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	isUsePrivateWebSocket := true
	isUseOrderAPI := false
	isUseMarketAPI := false

	isUsePublicWebSocket := false
	isUsePositionAPI := false

	// ## ------ Websocket Testing and usage --------------
	if isUsePositionAPI {
		// ## Create New http Client
		apiClient := api.New(
			"https://"+config.API_URL,
			config.CLIENT_ID,
			config.CLIENT_SECRET,
		)

		// ## ------- [Pre-Condition] ----
		// orderBuyResponse, err := apiClient.Orders.Buy(
		// 	"SOL_USDC-PERPETUAL",
		// 	0.1,
		// 	0,
		// 	"market",
		// 	"my-sol-usdc-future-perp-test",
		// 	0,
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
		// 	log.Fatalf("failed [Buy-Future] API: %+v", err)
		// }

		// fmt.Println("[1] Buy Resp: ")
		// fmt.Printf("%#v", orderBuyResponse)
		// fmt.Printf("\n\n")

		// ## ------- [GetPosition] ----
		positionResponse, err := apiClient.Positions.GetPosition("SOL_USDC-PERPETUAL")
		if err != nil {
			// Handle the error
			log.Fatalf("failed [GetPosition] API: %+v", err)
		}

		// Access the position details from the positionResponse.Result field
		position := positionResponse.Result
		fmt.Println("Position size:", position.Size)
		fmt.Println("Mark price:", position.MarkPrice)
		fmt.Println(" ")
		fmt.Println(" ")

		// ## ------- [GetPositions] ----
		positionsResponse, err := apiClient.Positions.GetPositions("USDC", "future", 0)
		if err != nil {
			// Handle the error
			log.Fatalf("failed [GetPositions] API: %+v", err)
		}

		// Access the positions from the positionsResponse.Result field
		for _, position := range positionsResponse.Result {
			fmt.Println("Instrument:", position.InstrumentName)
			fmt.Println("Size:", position.Size)
			fmt.Println("Mark Price:", position.MarkPrice)
			// Access other position details as needed
		}

		fmt.Println(" ")
		fmt.Println(" ")

		// ## ------- [GetMargins] ----

		marginsResponse, err := apiClient.Orders.GetMargins("SOL_USDC-PERPETUAL", 0.1, 139)
		if err != nil {
			// Handle the error
			log.Fatalf("failed [GetMargins] API: %+v", err)
		}

		fmt.Println("Buy Margin:", marginsResponse.Result.Buy)
		fmt.Println("Sell Margin:", marginsResponse.Result.Sell)
		fmt.Println("Max Price:", marginsResponse.Result.MaxPrice)
		fmt.Println("Min Price:", marginsResponse.Result.MinPrice)

		fmt.Println(" ")
		fmt.Println(" ")

		// ## ------- [GetSimulateMargins] ----

		simulatedPositions := map[string]float64{
			"SOL_USDC-PERPETUAL": 100,
		}
		simulatePortfolioResponse, err := apiClient.Positions.GetSimulateMargins(
			"USDC",
			true,
			simulatedPositions,
		)
		if err != nil {
			// Handle the error
			log.Fatalf("failed [GetSimulateMargins] API: %+v", err)
		}

		fmt.Printf("Success [GetSimulateMargins] API: %+v", simulatePortfolioResponse)
		// Access the simulated portfolio margin information from the simulatePortfolioResponse.Result field
		// fmt.Println("Projected Margin:", simulatePortfolioResponse.Result.ProjectedMargin)
		// fmt.Println("Margin:", simulatePortfolioResponse.Result.Margin)

		fmt.Println(" ")
		fmt.Println(" ")

		// ## ------- [GetSimulateMargins] ----
		closePositionResponse, err := apiClient.Positions.ClosePosition(
			"SOL_USDC-PERPETUAL",
			"market",
			140,
		)
		if err != nil {
			// Handle the error
			log.Fatalf("failed [ClosePosition] API: %+v", err)
		}

		fmt.Println("ClosePosition: ")
		fmt.Println("Order ID:", closePositionResponse.Result.Order.OrderID)
		fmt.Println("Trades:")
		for _, trade := range closePositionResponse.Result.Trades {
			fmt.Println("  - Trade ID:", trade.TradeID)
			fmt.Println("    Price:", trade.Price)
			fmt.Println("    Direction:", trade.Direction)
			// Access other trade details as needed
		}
	}

	// ## ------ Websocket Testing and usage --------------
	if isUsePrivateWebSocket {

		// // ## ------ Create Market Client connection --------------
		// client := ws.NewDeribitClient(config.CLIENT_ID, config.CLIENT_SECRET)

		// err := client.Connect(config.WS_URL)
		// if err != nil {
		// 	log.Fatalf("failed to connect: %v", err)
		// }

		// // ## Send Hello Software the WebSocket connection
		// err = client.Hello(config.NAME, config.VERSION)
		// if err != nil {
		// 	client.Close()
		// 	log.Fatalf("failed to hello: %v", err)
		// }

		// // ## Set WebSocket HeartBeat Interval
		// err = client.SetHeartBeat(60)
		// if err != nil {
		// 	client.Close()
		// 	log.Fatalf("failed to set heart beat: %v", err)
		// }

		// // ## Subscribe to multiple channels
		// err = client.Subscribe(
		// 	"deribit_price_index.btc_usd",
		// 	"deribit_price_index.btc_usdc",
		// 	"deribit_price_index.btc_usdt",
		// )
		// if err != nil {
		// 	client.Close()
		// 	log.Fatalf("failed to subscribe : %v", err)
		// }

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
		errPrivate = privateClient.SetHeartBeat(60)
		if errPrivate != nil {
			privateClient.Close()
			log.Fatalf("failed to set heart beat: %v", errPrivate)
		}

		// ## Authenticate the WebSocket connection
		_, err4 := ws.Authenticate(privateClient)
		if err4 != nil {
			privateClient.Close()
			log.Fatalf("failed to authenticate: %v", err4)
		}

		// ## Subscribe to multiple channels
		err := privateClient.PrivateSubscribe(
			// "deribit_price_index.sol_usdc", // public market index price data
			// "trades.sol_usdc.raw",          // Trade Signal from our account
			// "ticker.SOL_USDC.raw",          // ## Ticker of btc_usdc pairs [raw/100ms]
			// "chart.trades.SOL_USDC.1",      // ## OHLCV of chart data [1/3/60 -> in minutes except 1D = 1 day]
			// "quote.SOL_USDC",               // ## Quote of btc_usdc pairs [raw/100ms]
			// "perpetual.SOL-PERPETUAL.raw",  // ## Perpetual of BTC-PERPETUAL pairs [raw/100ms

			"trades.btc_usdc.raw",
		)
		if err != nil {
			privateClient.Close()
			log.Fatalf("failed to PrivateSubscribe channel : %v", err)
		}

		// // ## UnSubscribe to all channels
		// err = privateClient.PrivateUnsubscribeAll()
		// if err != nil {
		// 	privateClient.Close()
		// 	log.Fatalf("failed to PrivateUnsubscribeAll : %v", err)
		// }

		// // ## -------------- Test Create Buy Order 2 ---------------------

		// // ## Create a new buy order request
		// orderRequest := &ws.OrderRequest{
		// 	InstrumentName: "SOL_USDC",
		// 	Amount:         1,
		// 	Price:          10,
		// 	Type:           "limit",
		// 	Label:          "limit0000243",
		// }

		// errPrivate = ws.CreateBuyOrder(privateClient, orderRequest)
		// if errPrivate != nil {
		// 	privateClient.Close()
		// 	log.Fatalf("failed to create buy order: %v", errPrivate)
		// }

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
		// client.Close()
		privateClient.Close()
		log.Println("Shutting down...")
	}

	// ## ------ API Testing and usage order.go --------------
	if isUseOrderAPI {

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

	// ## ------ API Testing and usage market.go --------------
	if isUseMarketAPI {

		// ## Create New http Client
		apiClient := api.New(
			"https://"+config.API_URL,
			config.CLIENT_ID,
			config.CLIENT_SECRET,
		)

		// @@ ------------ [1] GetFundingChartData --------

		requestFundingChartData := &api.FundingChartDataRequest{
			InstrumentName: "BTC-PERPETUAL",
			Length:         "8h", // Can be "8h", "24h", or "1m"
		}

		fundingChartDataResponse, err := apiClient.Markets.GetFundingChartData(requestFundingChartData)
		if err != nil {
			// Handle error
			log.Fatalf("failed [GetFundingChartData]: %+v", err)
		}

		fmt.Printf("Current interest: %f\n", fundingChartDataResponse.Result.CurrentInterest)
		fmt.Printf("Current interest_8h: %v\n", fundingChartDataResponse.Result.Interest8h)
		fmt.Println("")
		// fmt.Printf("Current funding chart data: %v\n", fundingChartDataResponse.Result.Data)
		fmt.Println("")
		fmt.Println("")

		// @@ ------------ [2] GetFundingRateHistory --------
		now := time.Now()
		oneWeekAgo := now.Add(-7 * 24 * time.Hour) // 7 days ago
		startTimestamp := oneWeekAgo.Unix() * 1000 // Convert to milliseconds
		// Get current timestamp
		endTimestamp := now.Unix() * 1000 // Current timestamp in milliseconds

		instrumentName := "BTC-PERPETUAL" // September 27, 2019 00:00:00 UTC

		fundingRateHistoryResponse, err := apiClient.Markets.GetFundingRateHistory(
			instrumentName,
			startTimestamp,
			endTimestamp,
		)
		if err != nil {
			// Handle error
			log.Fatalf("failed [GetFundingRateHistory]: %+v", err)
		}

		fmt.Println("Funding Rate History:")
		fmt.Println("")
		fmt.Printf("%+v\n", fundingRateHistoryResponse.Result)
		fmt.Println("")

		fmt.Println("")
		fmt.Println("")

		// @@ ------------ [3] GetFundingRateValue --------
		now = time.Now()
		oneWeekAgo = now.Add(-7 * 24 * time.Hour) // 7 days ago
		startTimestamp = oneWeekAgo.Unix() * 1000 // Convert to milliseconds
		// Get current timestamp
		endTimestamp = now.Unix() * 1000 // Current timestamp in milliseconds

		instrumentName = "BTC-PERPETUAL" // September 27, 2019 00:00:00 UTC

		fundingRateValueResponse, err := apiClient.Markets.GetFundingRateValue(
			instrumentName,
			startTimestamp,
			endTimestamp,
		)
		if err != nil {
			// Handle error
			log.Fatalf("failed [GetFundingRateValue]: %+v", err)
		}

		fmt.Println("Funding Rate Value:")
		fmt.Println("")
		fmt.Printf("%+v\n", fundingRateValueResponse.Result)
		fmt.Println("")

		fmt.Println("")
		fmt.Println("")

		// @@ ------------ [4] GetHistoricalVolatility --------
		currency := "BTC"

		historicalVolatilityResponse, err := apiClient.Markets.GetHistoricalVolatility(currency)
		if err != nil {
			// Handle error
			log.Fatalf("failed [GetHistoricalVolatility]: %+v", err)
		}

		for _, entry := range historicalVolatilityResponse.Result {
			timestamp := entry[0]
			volatility := entry[1]
			entryTime := time.Unix(int64(timestamp/1000), 0).Format("2006-01-02 15:04:05")
			// fmt.Printf("Timestamp: %d, Volatility: %f\n", int64(timestamp), volatility)
			fmt.Printf("Timestamp: %s, Volatility: %f\n", entryTime, volatility)
		}

		fmt.Println("")
		fmt.Println("")

		// @@ ------------ [5] GetIndexPrice --------

		// Example usage of GetIndexPrice
		indexName := "btc_usd"

		indexPriceResponse, err := apiClient.Markets.GetIndexPrice(indexName)
		if err != nil {
			// Handle error
			log.Fatalf("failed [GetIndexPrice]: %+v", err)
		}

		fmt.Printf("Index Price: %f\n", indexPriceResponse.Result.IndexPrice)
		fmt.Printf("Estimated Delivery Price: %f\n", indexPriceResponse.Result.EstimatedDeliveryPrice)
		fmt.Println("")
		fmt.Println("")

		// @@ ------------ [6] GetIndexPriceNames --------
		indexPriceNamesResponse, err := apiClient.Markets.GetIndexPriceNames()
		if err != nil {
			// Handle error
			log.Fatalf("failed [GetIndexPriceNames]: %+v", err)
		}

		fmt.Println("Available Index Price Names:")
		for _, indexName := range indexPriceNamesResponse.Result {
			fmt.Println(indexName)
		}

		fmt.Println("")
		fmt.Println("")

		// @@ ------------ [7] GetInstrument --------
		instrumentName = "BTC-PERPETUAL"

		instrumentResponse, err := apiClient.Markets.GetInstrument(instrumentName)
		if err != nil {
			// Handle error
			log.Fatalf("failed [GetInstrument]: %+v", err)
		}

		fmt.Println("GetInstrument:")
		fmt.Printf("%+v\n", instrumentResponse.Result)
		fmt.Println("")

		fmt.Println("")
		fmt.Println("")

		// @@ ------------ [8] GetInstruments --------
		currency = "USDC"
		kind := "spot"
		expired := false

		instrumentsResponse, err := apiClient.Markets.GetInstruments(currency, kind, expired)
		if err != nil {
			// Handle error
			log.Fatalf("failed [GetInstruments]: %+v", err)
		}

		fmt.Println("Available Instruments:")
		for _, instrument := range instrumentsResponse.Result {
			fmt.Printf("Instrument Name: %s\n", instrument.InstrumentName)
			fmt.Printf("Instrument ID: %d\n", instrument.InstrumentID)
			fmt.Printf("Kind: %s\n", instrument.Kind)
			fmt.Printf("Base Currency: %s\n", instrument.BaseCurrency)
			fmt.Printf("Quote Currency: %s\n", instrument.QuoteCurrency)
			fmt.Printf("Settlement Currency: %s\n", instrument.SettlementCurrency)
			fmt.Printf("Settlement Period: %s\n", instrument.SettlementPeriod)
			fmt.Printf("Expiration Timestamp: %d\n", instrument.ExpirationTimestamp)
			fmt.Printf("Creation Timestamp: %d\n", instrument.CreationTimestamp)
			fmt.Printf("Contract Size: %f\n", instrument.ContractSize)
			fmt.Printf("Tick Size: %f\n", instrument.TickSize)
			fmt.Printf("Maker Commission: %f\n", instrument.MakerCommission)
			fmt.Printf("Taker Commission: %f\n", instrument.TakerCommission)
			fmt.Printf("Block Trade Tick Size: %f\n", instrument.BlockTradeTickSize)
			fmt.Printf("Block Trade Min Trade Amount: %f\n", instrument.BlockTradeMinTradeAmount)
			fmt.Printf("Block Trade Commission: %f\n", instrument.BlockTradeCommission)
			fmt.Printf("Instrument Type: %s\n", instrument.InstrumentType)
			// fmt.Printf("Max Leverage: %d\n", int(instrument.MaxLeverage))
			// fmt.Printf("Max Liquidation Commission: %f\n", instrument.MaxLiquidationCommission)
			fmt.Printf("Minimum Trade Amount: %f\n", instrument.MinTradeAmount)
			fmt.Printf("Is Active: %t\n", instrument.IsActive)
			fmt.Println("")
		}

		fmt.Println("")
		fmt.Println("")

		// @@ ------------ [9] GetLastSettlementsByInstrument --------
		instrumentName = "BTC-PERPETUAL"
		settlementType := "settlement"
		count := 1
		continuation := ""
		searchStartTimestamp := time.Now().Add(-7*24*time.Hour).Unix() * 1000 // 7 days ago

		lastSettlementsResponse, err := apiClient.Markets.GetLastSettlementsByInstrument(instrumentName, settlementType, count, continuation, searchStartTimestamp)
		if err != nil {
			// Handle error
			log.Fatalf("failed [GetLastSettlementsByInstrument]: %+v", err)
		}

		fmt.Println("Last Settlements:")
		for _, settlement := range lastSettlementsResponse.Result.Settlements {
			fmt.Printf("Timestamp: %d, Type: %s, Instrument Name: %s, Position: %f, Profit/Loss: %f\n",
				settlement.Timestamp, settlement.Type, settlement.InstrumentName, settlement.Position, settlement.ProfitLoss)
		}

		if lastSettlementsResponse.Result.Continuation != "" {
			fmt.Printf("Continuation token: %s\n", lastSettlementsResponse.Result.Continuation)
		}

		fmt.Println("")
		fmt.Println("")

		// @@ ------------ [10] GetLastTradesByCurrencyAndTime --------
		currency = "BTC"
		startTimestamp = time.Now().Add(-1*time.Hour).Unix() * 1000 // 1 hour ago
		endTimestamp = time.Now().Unix() * 1000                     // Current timestamp
		count = 5

		lastTradesByCurrencyAndTimeResponse, err := apiClient.Markets.GetLastTradesByCurrencyAndTime(currency, startTimestamp, endTimestamp, count)
		if err != nil {
			// Handle error
			log.Fatalf("failed [GetLastTradesByCurrencyAndTime]: %+v", err)
		}

		fmt.Println("Last Trades [GetLastTradesByCurrencyAndTime]:")
		for _, trade := range lastTradesByCurrencyAndTimeResponse.Result.Trades {
			entryTime := time.Unix(int64(trade.Timestamp/1000), 0).Format("2006-01-02 15:04:05")
			// fmt.Printf("Timestamp: %d, Instrument Name: %s, Direction: %s, Price: %f, Amount: %f\n",
			// 	trade.Timestamp, trade.InstrumentName, trade.Direction, trade.Price, trade.Amount)

			fmt.Printf("Timestamp: %s, Instrument Name: %s, Direction: %s, Price: %f, Amount: %f\n",
				entryTime, trade.InstrumentName, trade.Direction, trade.Price, trade.Amount)
		}

		if lastTradesByCurrencyAndTimeResponse.Result.HasMore {
			fmt.Println("More trades available")
		}

		fmt.Println("")
		fmt.Println("")

		// @@ ------------ [11] GetLastTradesByInstrument --------
		instrumentName = "BTC-PERPETUAL"
		startSeq := 0
		endSeq := 0
		startTimestamp = time.Now().Add(-1*time.Hour).Unix() * 1000 // 1 hour ago
		endTimestamp = time.Now().Unix() * 1000                     // Current timestamp
		count = 10
		sorting := "default"

		lastTradesByInstrumentResponse, err := apiClient.Markets.GetLastTradesByInstrument(instrumentName, startSeq, endSeq, startTimestamp, endTimestamp, count, sorting)
		if err != nil {
			// Handle error
			log.Fatalf("failed [GetLastTradesByInstrument]: %+v", err)
		}

		fmt.Println("Last Trades [GetLastTradesByInstrument]:")
		for _, trade := range lastTradesByInstrumentResponse.Result.Trades {
			// fmt.Printf("Timestamp: %d, Trade ID: %s, Instrument Name: %s, Direction: %s, Price: %f, Amount: %f\n",
			// 	trade.Timestamp, trade.TradeID, trade.InstrumentName, trade.Direction, trade.Price, trade.Amount)

			entryTime := time.Unix(int64(trade.Timestamp/1000), 0).Format("2006-01-02 15:04:05")
			// fmt.Printf("Timestamp: %d, Instrument Name: %s, Direction: %s, Price: %f, Amount: %f\n",
			// 	trade.Timestamp, trade.InstrumentName, trade.Direction, trade.Price, trade.Amount)

			fmt.Printf("Timestamp: %s, Instrument Name: %s, Direction: %s, Price: %f, Amount: %f\n",
				entryTime, trade.InstrumentName, trade.Direction, trade.Price, trade.Amount)
		}

		if lastTradesByInstrumentResponse.Result.HasMore {
			fmt.Println("More trades available")
		}

		fmt.Println("")
		fmt.Println("")

		// @@ ------------ [12] GetLastTradesByInstrumentAndTime --------
		instrumentName = "ETH-PERPETUAL"
		startTimestamp = time.Now().Add(-1*time.Hour).Unix() * 1000 // 1 hour ago
		endTimestamp = time.Now().Unix() * 1000                     // Current timestamp
		count = 10
		sorting = "default"

		lastTradesByInstrumentAndTimeResponse, err := apiClient.Markets.GetLastTradesByInstrumentAndTime(instrumentName, startTimestamp, endTimestamp, count, sorting)
		if err != nil {
			// Handle error
			log.Fatalf("failed [GetLastTradesByInstrumentAndTime]: %+v", err)
		}

		fmt.Println("Last Trades:")
		for _, trade := range lastTradesByInstrumentAndTimeResponse.Result.Trades {
			// fmt.Printf("Timestamp: %d, Trade ID: %s, Instrument Name: %s, Direction: %s, Price: %f, Amount: %f\n",
			// 	trade.Timestamp, trade.TradeID, trade.InstrumentName, trade.Direction, trade.Price, trade.Amount)

			entryTime := time.Unix(int64(trade.Timestamp/1000), 0).Format("2006-01-02 15:04:05")
			fmt.Printf("Timestamp: %s, Trade ID: %s, Instrument Name: %s, Direction: %s, Price: %f, Amount: %f\n",
				entryTime, trade.TradeID, trade.InstrumentName, trade.Direction, trade.Price, trade.Amount)

		}

		if lastTradesByInstrumentAndTimeResponse.Result.HasMore {
			fmt.Println("More trades available")
		}

		fmt.Println("")
		fmt.Println("")

		// @@ ------------ [13] GetMarkPriceHistory --------
		instrumentName = "BTC-PERPETUAL"
		startTimestamp = time.Now().Add(-1*time.Hour).Unix() * 1000 // 4 hour ago
		endTimestamp = time.Now().Unix() * 1000                     // Current timestamp
		// resolution := "5min"

		markPriceHistoryResponse, err := apiClient.Markets.GetMarkPriceHistory(instrumentName, startTimestamp, endTimestamp)
		if err != nil {
			// Handle error
			log.Fatalf("failed [GetMarkPriceHistory]: %+v", err)
		}

		fmt.Println("Mark Price History:")
		fmt.Printf("%+v", markPriceHistoryResponse.Result)
		for _, entry := range markPriceHistoryResponse.Result {
			timestamp := int64(entry[0])
			markPrice := entry[1]
			entryTime := time.Unix(int64(timestamp/1000), 0).Format("2006-01-02 15:04:05")
			// fmt.Printf("Timestamp: %d, Mark Price: %f\n", timestamp, markPrice)
			fmt.Printf("Timestamp: %s, Mark Price: %f\n", entryTime, markPrice)
		}

		fmt.Println("")
		fmt.Println("")

		// @@ ------------ [14] GetOrderBook --------
		instrumentName = "BTC_USDC"
		depth := 5

		orderBookResponse, err := apiClient.Markets.GetOrderBook(instrumentName, depth)
		if err != nil {
			// Handle error
			log.Fatalf("failed [GetOrderBook]: %+v", err)
		}

		fmt.Println("GetOrderBook:")
		fmt.Printf("Instrument Name: %s\n", orderBookResponse.Result.InstrumentName)
		fmt.Printf("Timestamp: %d\n", orderBookResponse.Result.Timestamp)
		fmt.Printf("Index Price: %f\n", orderBookResponse.Result.IndexPrice)
		fmt.Printf("Mark Price: %f\n", orderBookResponse.Result.MarkPrice)
		fmt.Printf("Last Price: %f\n", orderBookResponse.Result.LastPrice)
		fmt.Printf("Open Interest: %f\n", orderBookResponse.Result.OpenInterest)
		fmt.Printf("Current Funding: %f\n", orderBookResponse.Result.CurrentFunding)
		fmt.Printf("Funding 8h: %f\n", orderBookResponse.Result.Funding8h)

		fmt.Println("\nBids:")
		for _, bid := range orderBookResponse.Result.Bids {
			fmt.Printf("Price: %f, Amount: %f\n", bid[0], bid[1])
		}

		fmt.Println("\nAsks:")
		for _, ask := range orderBookResponse.Result.Asks {
			fmt.Printf("Price: %f, Amount: %f\n", ask[0], ask[1])
		}

		fmt.Println("")
		fmt.Println("")

		// @@ ------------ [15] GetOrderBookByInstrumentId --------
		instrumentID := 404150 // BTC_USDC
		if instrumentResponse.Result.InstrumentID != 0 {
			instrumentID = int(instrumentResponse.Result.InstrumentID)
		}
		depth = 10

		orderBookByInstrumentResponse, err := apiClient.Markets.GetOrderBookByInstrumentId(instrumentID, depth)
		if err != nil {
			// Handle error
			log.Fatalf("failed [GetOrderBookByInstrumentId]: %+v", err)
		}

		fmt.Println("GetOrderBookByInstrumentId: ")
		fmt.Printf("Instrument Name: %s\n", orderBookByInstrumentResponse.Result.InstrumentName)
		fmt.Printf("Timestamp: %d\n", orderBookByInstrumentResponse.Result.Timestamp)
		fmt.Printf("Index Price: %f\n", orderBookByInstrumentResponse.Result.IndexPrice)
		fmt.Printf("Mark Price: %f\n", orderBookByInstrumentResponse.Result.MarkPrice)
		fmt.Printf("Last Price: %f\n", orderBookByInstrumentResponse.Result.LastPrice)
		fmt.Printf("Open Interest: %f\n", orderBookByInstrumentResponse.Result.OpenInterest)
		fmt.Printf("Current Funding: %f\n", orderBookByInstrumentResponse.Result.CurrentFunding)
		fmt.Printf("Funding 8h: %f\n", orderBookByInstrumentResponse.Result.Funding8h)

		fmt.Println("\nBids:")
		for _, bid := range orderBookByInstrumentResponse.Result.Bids {
			fmt.Printf("Price: %f, Amount: %f\n", bid[0], bid[1])
		}

		fmt.Println("\nAsks:")
		for _, ask := range orderBookByInstrumentResponse.Result.Asks {
			fmt.Printf("Price: %f, Amount: %f\n", ask[0], ask[1])
		}

		fmt.Println("")
		fmt.Println("")

		// @@ ------------ [16] GetTradeVolumes --------
		extended := true

		tradeVolumesResponse, err := apiClient.Markets.GetTradeVolumes(extended)
		if err != nil {
			// Handle error
			log.Fatalf("failed [GetTradeVolumes]: %+v", err)
		}

		fmt.Println("Trade Volumes:")
		for _, volume := range tradeVolumesResponse.Result {
			fmt.Printf("Currency: %s\n", volume.Currency)
			fmt.Printf("Calls Volume: %f\n", volume.CallsVolume)
			fmt.Printf("Puts Volume: %f\n", volume.PutsVolume)
			fmt.Printf("Futures Volume: %f\n", volume.FuturesVolume)
			fmt.Printf("Spot Volume: %f\n", volume.SpotVolume)
			if extended {
				fmt.Printf("Calls Volume 7d: %f\n", volume.CallsVolume7d)
				fmt.Printf("Calls Volume 30d: %f\n", volume.CallsVolume30d)
				fmt.Printf("Puts Volume 7d: %f\n", volume.PutsVolume7d)
				fmt.Printf("Puts Volume 30d: %f\n", volume.PutsVolume30d)
				fmt.Printf("Futures Volume 7d: %f\n", volume.FuturesVolume7d)
				fmt.Printf("Futures Volume 30d: %f\n", volume.FuturesVolume30d)
				fmt.Printf("Spot Volume 7d: %f\n", volume.SpotVolume7d)
				fmt.Printf("Spot Volume 30d: %f\n", volume.SpotVolume30d)
			}
			fmt.Println()
		}

		fmt.Println("")
		fmt.Println("")

		// @@ ------------ [17] GetTradingViewChartData --------
		instrumentName = "BTC_USDC"
		startTimestamp = time.Now().Add(-4*time.Hour).Unix() * 1000 // 4 hour ago
		endTimestamp = time.Now().Unix() * 1000                     // Current timestamp
		resolution := "30"

		tradingViewChartDataResponse, err := apiClient.Markets.GetTradingViewChartData(instrumentName, startTimestamp, endTimestamp, resolution)
		if err != nil {
			// Handle error
			log.Fatalf("failed [GetTradingViewChartData]: %+v", err)
		}

		fmt.Println("GetTradingViewChartData: ")
		fmt.Printf("Instrument Name: %s\n", instrumentName)
		fmt.Printf("Status: %s\n", tradingViewChartDataResponse.Result.Status)

		fmt.Println("\nOpen:")
		for i, open := range tradingViewChartDataResponse.Result.Open {
			fmt.Printf("Timestamp: %d, Open: %f\n", tradingViewChartDataResponse.Result.Ticks[i], open)
		}

		fmt.Println("\nHigh:")
		for i, high := range tradingViewChartDataResponse.Result.High {
			fmt.Printf("Timestamp: %d, High: %f\n", tradingViewChartDataResponse.Result.Ticks[i], high)
		}

		fmt.Println("\nLow:")
		for i, low := range tradingViewChartDataResponse.Result.Low {
			fmt.Printf("Timestamp: %d, Low: %f\n", tradingViewChartDataResponse.Result.Ticks[i], low)
		}

		fmt.Println("\nClose:")
		for i, close := range tradingViewChartDataResponse.Result.Close {
			fmt.Printf("Timestamp: %d, Close: %f\n", tradingViewChartDataResponse.Result.Ticks[i], close)
		}

		fmt.Println("\nVolume:")
		for i, volume := range tradingViewChartDataResponse.Result.Volume {
			fmt.Printf("Timestamp: %d, Volume: %f\n", tradingViewChartDataResponse.Result.Ticks[i], volume)
		}

		fmt.Println("\nCost:")
		for i, cost := range tradingViewChartDataResponse.Result.Cost {
			fmt.Printf("Timestamp: %d, Cost: %f\n", tradingViewChartDataResponse.Result.Ticks[i], cost)
		}

		fmt.Println("")
		fmt.Println("")

		// @@ ------------ [18] GetVolatilityIndexData --------
		currency = "BTC"
		startTimestamp = time.Now().Add(-24*time.Hour).Unix() * 1000 // 24 hour ago
		endTimestamp = time.Now().Unix() * 1000                      // Current timestamp
		resolution = "1D"

		volatilityIndexDataResponse, err := apiClient.Markets.GetVolatilityIndexData(currency, startTimestamp, endTimestamp, resolution)
		if err != nil {
			// Handle error
			log.Fatalf("failed [GetVolatilityIndexData]: %+v", err)
		}

		fmt.Println("GetVolatilityIndexData: ")
		fmt.Printf("Currency: %s\n", currency)
		fmt.Printf("Continuation: %s\n", volatilityIndexDataResponse.Result.Continuation)

		fmt.Println("\nVolatility Index Data:")
		for _, data := range volatilityIndexDataResponse.Result.Data {
			timestamp := int64(data[0])
			open := data[1]
			high := data[2]
			low := data[3]
			close := data[4]
			fmt.Printf("Timestamp: %d, Open: %f, High: %f, Low: %f, Close: %f\n", timestamp, open, high, low, close)
		}

		fmt.Println("")
		fmt.Println("")

		// @@ ------------ [19] GetTicker --------
		instrumentName = "BTC_USDC"

		tickerResponse, err := apiClient.Markets.GetTicker(instrumentName)
		if err != nil {
			// Handle error
			log.Fatalf("failed [GetTicker]: %+v", err)
		}

		fmt.Println("GetTicker: ")
		// fmt.Printf("%+v\n", tickerResponse)
		// fmt.Printf("\n")

		fmt.Printf("Instrument Name: %s\n", tickerResponse.Result.InstrumentName)
		fmt.Printf("Timestamp: %d\n", tickerResponse.Result.Timestamp)
		fmt.Printf("Best Bid Price: %f\n", tickerResponse.Result.BestBidPrice)
		fmt.Printf("Best Bid Amount: %f\n", tickerResponse.Result.BestBidAmount)
		fmt.Printf("Best Ask Price: %f\n", tickerResponse.Result.BestAskPrice)
		fmt.Printf("Best Ask Amount: %f\n", tickerResponse.Result.BestAskAmount)
		fmt.Printf("Last Price: %f\n", tickerResponse.Result.LastPrice)
		fmt.Printf("Mark Price: %f\n", tickerResponse.Result.MarkPrice)
		fmt.Printf("Index Price: %f\n", tickerResponse.Result.IndexPrice)
		fmt.Printf("Settlement Price: %f\n", tickerResponse.Result.SettlementPrice)
		fmt.Printf("Open Interest: %f\n", tickerResponse.Result.OpenInterest)
		fmt.Printf("Current Funding: %f\n", tickerResponse.Result.CurrentFunding)
		fmt.Printf("Funding 8h: %f\n", tickerResponse.Result.Funding8h)
		fmt.Printf("Estimated Delivery Price: %f\n", tickerResponse.Result.EstimatedDeliveryPrice)
		fmt.Printf("State: %s\n", tickerResponse.Result.State)
		fmt.Printf("High: %f\n", tickerResponse.Result.Stats.High)
		fmt.Printf("Low: %f\n", tickerResponse.Result.Stats.Low)
		fmt.Printf("Price Change: %f\n", tickerResponse.Result.Stats.PriceChange)
		fmt.Printf("Volume: %f\n", tickerResponse.Result.Stats.Volume)
		fmt.Printf("Volume USD: %f\n", tickerResponse.Result.Stats.VolumeUSD)

		fmt.Println("")
		fmt.Println("")

	}

	// ## ------ [Public] Websocket Testing and usage --------------
	if isUsePublicWebSocket {

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
		// err = client.Subscribe(
		// 	"deribit_price_index.btc_usd",
		// 	"deribit_price_index.btc_usdc",
		// 	"deribit_price_index.btc_usdt",
		// )
		err = client.Subscribe(
			"deribit_price_index.btc_usdc",      // ## Index price
			"deribit_volatility_index.btc_usdc", // ## Volatility Index
			"deribit_price_statistics.btc_usdc", // ## Volatility Index
			"ticker.BTC_USDC.100ms",             // ## Ticker of btc_usdc pairs [raw/100ms] (raw not for public channel)
			"chart.trades.BTC_USDC.1",           // ## OHLCV of chart data [1/3/60 -> in minutes except 1D = 1 day]
			"quote.BTC_USDC",                    // ## Quote of btc_usdc pairs [raw/100ms]
			"perpetual.BTC-PERPETUAL.100ms",     // ## Perpetual of BTC-PERPETUAL pairs [raw/100ms]  (raw not for public channel)
			"markprice.options.btc_usdc",        // ## Options Market price by pairs of SPOT
		)
		if err != nil {
			client.Close()
			log.Fatalf("failed to subscribe : %v", err)
		}

		// ## UnSubscribe to multiple channels
		err = client.Unsubscribe(
			"deribit_price_index.btc_usdt",
		)
		if err != nil {
			client.Close()
			log.Fatalf("failed to Unsubscribe : %v", err)
		}

		// ## -------------- Main Loop (Concurrent GO) ---------------------
		// Start concurrent tasks for HandleReadMessage and HandleHeartBeatMessage
		// Start concurrent tasks for HandleReadMessage and HandleHeartBeatMessage
		go client.Run()

		// ## -------------- Termination ---------------------
		// ## Wait for termination signals
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
		<-sigChan

		// ## Gracefully shut down the client
		client.Close()
		log.Println("Shutting down...")
	}
}
