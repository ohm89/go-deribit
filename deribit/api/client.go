package api

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/valyala/fasthttp"
)

const (
	maxRetries     = 5
	defaultBaseURL = "https://www.deribit.com"
	defaultAPIURL  = "/api/v2"

	userAgent = "go-deribit"

	HeaderKey        = "DERIBIT-KEY"
	HeaderSign       = "DERIBIT-SIGN"
	HeaderTS         = "DERIBIT-TS"
	HeaderSubaccount = "DERIBIT-SUBACCOUNT"
)

type service struct {
	client *Client
}

type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Id      uint64      `json:"id"`
	Jsonrpc string      `json:"jsonrpc"`
	Result  interface{} `json:"result,omitempty"`
	Error   *AuthError  `json:"error,omitempty"`
}

type Client struct {
	baseURL string
	client  *fasthttp.Client

	clientID     string
	clientSecret string
	accessToken  string
	refreshToken string
	// subaccountId uint64

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	Accounts *AccountService
	Markets  *MarketService
	// Wallets  *WalletService
	Orders *OrderService
	// Fills    *FillsService
	Positions *PositionService
}

func New(baseUrl string, clientID string, clientSecret string) *Client {
	httpClient := &fasthttp.Client{
		Name:         userAgent,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	c := &Client{
		baseURL:      baseUrl,
		client:       httpClient,
		clientID:     clientID,
		clientSecret: clientSecret,
	}
	c.common.client = c
	c.Accounts = (*AccountService)(&c.common)
	c.Markets = (*MarketService)(&c.common)
	// c.Wallets = (*WalletService)(&c.common)
	c.Orders = (*OrderService)(&c.common)
	c.Positions = (*PositionService)(&c.common)
	// c.Fills = (*FillsService)(&c.common)

	return c
}

// ## Basic http driver to request
func (c *Client) do(uri string, method string, in, out interface{}, isPrivate bool) error {
	req, resp := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}()

	req.SetRequestURI(uri)
	req.Header.SetMethod(method)

	if in != nil {
		req.Header.SetContentType("application/json")
		if err := json.NewEncoder(req.BodyWriter()).Encode(in); err != nil {
			return err
		}
	}

	if isPrivate {
		if c.accessToken == "" {
			if _, err := Authenticate(c); err != nil {
				return err
			}
		}

		req.Header.Set("Authorization", "Bearer "+c.accessToken)
	}

	var numRetries int
	var data Response
	for {
		if err := c.client.Do(req, resp); err != nil {
			return err
		}

		// Check the response status code
		if resp.StatusCode() == fasthttp.StatusOK {
			// Check if the response body is empty
			if resp.Body() == nil || len(resp.Body()) == 0 {
				// Return an error, as the response should not be empty
				return fmt.Errorf("unexpected empty response body with status code %d", resp.StatusCode())
			}

			if err := json.Unmarshal(resp.Body(), &data); err != nil {
				return fmt.Errorf("unmarshal: [%v] body: %v, error: %v", resp.StatusCode(), string(resp.Body()), err)
			}

			if data.Error != nil {
				// Check if the error code is 13009 (unauthorized)
				if data.Error.Code == 13009 && numRetries < maxRetries {
					// Retry the request after authenticating
					if _, err := Authenticate(c); err != nil {
						return err
					}
					numRetries++
					continue
				}
				return fmt.Errorf("request failed: code: %d, message: %s", data.Error.Code, data.Error.Message)
			}

			break
		} else {
			// ## [DEBUG]
			fmt.Printf("-- do Request Error !! -- \n")
			fmt.Printf("Error Request URL: %#v \n\n", string(req.RequestURI()))
			fmt.Printf("Error Request Body: %#v \n\n", string(req.Body()))
			fmt.Printf("Error Response body: %#v \n\n", string(resp.Body()))

			// Handle the error response
			return fmt.Errorf(
				"request URI: %s \n request Body: %s \n\n Request failed with status code: %d \n Response body: %s",
				string(req.RequestURI()),
				string(req.Body()),
				resp.StatusCode(),
				string(resp.Body()),
			)
		}
	}

	if out != nil {
		// Assign the data.Result to the out parameter
		if err := json.Unmarshal(resp.Body(), out); err != nil {
			return fmt.Errorf("unmarshal out: [%v] body: %v, error: %v", resp.StatusCode(), string(resp.Body()), err)
		}
	}

	// ## [DEBUG]
	// fmt.Printf("-- do Request Done !! -- \n")
	// fmt.Printf("Received a message: %#v \n\n", out)

	return nil
}

func (c *Client) DoPublic(uri string, method string, in, out interface{}) error {
	return c.do(uri, method, in, out, false)
}

func (c *Client) DoPrivate(uri string, method string, in, out interface{}) error {
	return c.do(uri, method, in, out, true)
}
