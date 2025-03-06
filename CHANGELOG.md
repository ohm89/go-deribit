# 1.2.2 

- [NEW] api/client.go add Debug message on do request error


# 1.2.1 

- [NEW] ws/auth.go add Refresh auth for refresh token every 1 hour

# 1.1.1 

- [BUG] ws/client.go add heartbeat management also in Receive function

# 1.1.0 

- [NEW-FEATURE] ws/client.go add Received function to read data and send out 


# 1.0.0 

- [NEW-FEATURE] api/position.go to get Positions value of future position


# 0.3.1 

- [NEW-FEATURE] ws/client.go add PrivateSubscribe, PrivateUnsubscribe, PrivateUnsubscribeAll, Unsubscribe, UnsubscribeAll function to use

- [TEST] main.go add example usage of different channel from api playground

# 0.3.0 

- [NEW] api/market.go

# 0.2.0 

- [NEW] api/account.go
- [NEW] api/auth.go
- [NEW] api/client.go
- [NEW] api/order.go
- [NEW] api/util.go

- [TBD] api/position.go
- [TBD] api/market.go

# 0.1.0 

- [NEW] ws/account.go
- [NEW] ws/position.go
- [NEW] ws/ausbaccount.go

# 0.0.4

- [NEW] cancelAll api
- [NEW] cancelOne api
- [NEW] createBuy api
- [NEW] createSell api

# 0.0.3

- [CHANGE] client.go to be go routine to used

# 0.0.2

- [CHANGE] add integer for SetHeartBeat function in ws/client

# 0.0.1

- Websocket connection and heartbeat setup