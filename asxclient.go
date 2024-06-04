package stakego

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
)

// NewASXClient - create and initialise an ASXClient
func NewASXClient() *ASXClient {
	c := ASXClient{}
	c.Init()
	return &c
}

// ASXClient - Client for interacting with Stake ASX
type ASXClient struct {
	apiUrl      string
	Credentials *Credentials
	User        *User
	httpclient  http.Client
	tokenMutex  sync.Mutex
	authMutex   sync.Mutex
}

// ResponseData - holds http response
type ResponseData struct {
	StatusCode int
	Body       []byte
}

// Init - initialise the ASX client with defaults
func (c *ASXClient) Init() {
	c.apiUrl = "https://global-prd-api.hellostake.com/api/"
	c.httpclient = NewHTTPClient()
}

// Login - create a user session
func (c *ASXClient) Login() (err error) {
	c.authMutex.Lock()
	defer c.authMutex.Unlock()

	if c.Credentials.GetSessionToken() == "" {
		u, err := url.JoinPath(c.apiUrl, "sessions/v2/createSession")
		if err != nil {
			return NewStakeError("login", err)
		}

		req, _ := NewJSONRequest("POST", u, c.Credentials.AsJSON())
		resp, err := c.httpclient.Do(req)
		if err != nil {
			return NewStakeError("login", err)
		}

		if resp.StatusCode == 200 {
			defer resp.Body.Close()
			rbody, err := io.ReadAll(resp.Body)
			if err != nil {
				return NewStakeError("login", err)
			}
			u := NewUserSessionFromJSON(rbody)
			c.Credentials.SetSessionToken(u.SessionKey)
		}
	}

	if c.Credentials.GetSessionToken() == "" {
		return NewStakeError("login", ErrSessionTokenMissing)
	}

	c.User, err = c.GetUser()
	if err != nil {
		return NewStakeError("login", err)
	}

	return nil
}

// Logout - end a user session
func (c *ASXClient) Logout() (err error) {
	c.authMutex.Lock()
	defer c.authMutex.Unlock()

	if c.Credentials.GetSessionToken() == "" {
		return NewStakeError("logout", ErrSessionTokenMissing)
	}

	u, err := url.JoinPath(c.apiUrl, "userauth", c.Credentials.GetSessionToken())
	if err != nil {
		return NewStakeError("logout", err)
	}

	req, _ := NewJSONRequest("DELETE", u, nil)
	resp, err := c.httpclient.Do(req)
	if err != nil {
		return NewStakeError("logout", err)
	}

	if resp.StatusCode == 200 {
		c.Credentials.SetSessionToken("")
		return nil
	}

	return NewStakeError("cash", ErrInvalidAPIResponse)
}

// GetMarket - Get the current market status
func (c *ASXClient) GetMarket() (*Market, error) {
	// early-bird-promo.hellostake.com has been taken down,
	// so pretend we're always open for now
	m := Market{}
	m.Status.Current = "OPEN"
	return &m, nil
	// u := "https://early-bird-promo.hellostake.com/marketStatus"
	// req, _ := NewJSONRequest("GET", u, nil)
	// resp, err := c.httpclient.Do(req)
	// if err != nil {
	// 	return nil, NewStakeError("market", err)
	// }

	// if resp.StatusCode == 200 {
	// 	defer resp.Body.Close()
	// 	rbody, err := io.ReadAll(resp.Body)
	// 	if err != nil {
	// 		return nil, NewStakeError("market", err)
	// 	}
	// 	m := NewMarketFromJSON(rbody)
	// 	return m, nil
	// }

	// return nil, NewStakeError("market", ErrInvalidAPIResponse)
}

// GetCash - get the current available cash
func (c *ASXClient) GetCash() (*Cash, error) {
	u, err := url.JoinPath(c.apiUrl, "asx/cash")
	if err != nil {
		return nil, NewStakeError("cash", err)
	}

	rd, err := c.AuthedRequest("GET", u, nil)
	if err != nil {
		return nil, NewStakeError("cash", err)
	}

	if rd.StatusCode == 200 {
		c := NewCashFromJSON(rd.Body)
		return c, nil
	}
	return nil, NewStakeError("cash", ErrInvalidAPIResponse)
}

// GetEquityPositions - get the current user's equity positions
func (c *ASXClient) GetEquityPositions() (*EquityPositions, error) {
	u, err := url.JoinPath(c.apiUrl, "asx/instrument/equityPositions")
	if err != nil {
		return nil, NewStakeError("equity positions", err)
	}

	rd, err := c.AuthedRequest("GET", u, nil)
	if err != nil {
		return nil, NewStakeError("equity positions", err)
	}

	if rd.StatusCode == 200 {
		e := NewEquityPositionsFromJSON(rd.Body)
		return e, nil
	}
	return nil, NewStakeError("equity positions", ErrInvalidAPIResponse)
}

// GetUser - get information about the current user
func (c *ASXClient) GetUser() (*User, error) {
	u, err := url.JoinPath(c.apiUrl, "user")
	if err != nil {
		return nil, NewStakeError("user", err)
	}

	rd, err := c.AuthedRequest("GET", u, nil)
	if err != nil {
		return nil, NewStakeError("user", err)
	}

	if rd.StatusCode == 200 {
		user := NewUserFromJSON(rd.Body)
		return user, nil
	}

	return nil, NewStakeError("user", ErrInvalidAPIResponse)
}

// GetOrders - get pending orders
func (c *ASXClient) GetOrders() (*[]OrderDetails, error) {
	u, err := url.JoinPath(c.apiUrl, "asx/orders")
	if err != nil {
		return nil, NewStakeError("orders", err)
	}

	rd, err := c.AuthedRequest("GET", u, nil)
	if err != nil {
		return nil, NewStakeError("orders", err)
	}

	if rd.StatusCode == 200 {
		orders := NewOrderListFromJSON(rd.Body)
		return orders, nil
	}
	return nil, NewStakeError("orders", ErrInvalidAPIResponse)
}

// PlaceOrder - place an order
func (c *ASXClient) PlaceOrder(order Order) (*OrderResponse, error) {
	u, err := url.JoinPath(c.apiUrl, "asx/orders")
	if err != nil {
		return nil, NewStakeError("orders/place", err)
	}

	rd, err := c.AuthedRequest("POST", u, order.AsJSON())
	if err != nil {
		return nil, NewStakeError("orders/place", err)
	}

	if rd.StatusCode == 200 {
		orders := NewOrderResponseFromJSON(rd.Body)
		return orders, nil
	}

	return nil, NewStakeError("orders/place", ErrInvalidAPIResponse)
}

// CancelOrder - cancel an order
func (c *ASXClient) CancelOrder(uuid string) error {
	u, err := url.JoinPath(c.apiUrl, "asx/orders", uuid, "cancel")
	if err != nil {
		return NewStakeError("orders/cancel", err)
	}

	rd, err := c.AuthedRequest("POST", u, nil)
	if err != nil {
		return NewStakeError("orders/cancel", err)
	}

	if rd.StatusCode == 200 {
		return nil
	}

	return NewStakeError("orders/cancel", ErrInvalidAPIResponse)
}

// GetBrokerage - get the current available cash
func (c *ASXClient) GetBrokerage(price float64) (*Brokerage, error) {
	u, err := url.JoinPath(c.apiUrl, "asx/orders/brokerage")
	if err != nil {
		return nil, NewStakeError("brokerage", err)
	}

	u = fmt.Sprintf("%s?orderAmount=%.2f", u, price)

	rd, err := c.AuthedRequest("GET", u, nil)
	if rd.StatusCode == 200 {
		b := NewBrokerageFromJSON(rd.Body)
		return b, nil
	}

	return nil, NewStakeError("brokerage", ErrInvalidAPIResponse)
}

// LookupInstrument - get an instrument by symbol
func (c *ASXClient) LookupInstrument(keyword string) (*Instrument, error) {
	u, err := url.JoinPath(c.apiUrl, "asx/instrument/search")
	if err != nil {
		return nil, NewStakeError("instrument", err)
	}

	u = fmt.Sprintf("%s?searchKey=%s&max=1", u, keyword)

	rd, err := c.AuthedRequest("GET", u, nil)
	if rd.StatusCode == 200 {
		ir := NewInstrumentResponseFromJSON(rd.Body)

		if len(ir.Instruments) > 0 {
			return &ir.Instruments[0], nil
		}
		return nil, NewStakeError("instrument", fmt.Errorf("No instrument for '%s' found", keyword))
	}

	return nil, NewStakeError("instrument", ErrInvalidAPIResponse)
}

// AuthedRequest - perform a http request and send auth token
func (c *ASXClient) AuthedRequest(method string, fullurl string, jsonBody []byte) (*ResponseData, error) {
	if c.Credentials.GetSessionToken() == "" {
		return nil, ErrSessionTokenMissing
	}

	req, _ := NewJSONRequest(method, fullurl, jsonBody)
	req.Header.Set("Stake-Session-Token", c.Credentials.GetSessionToken())
	resp, err := c.httpclient.Do(req)
	if err != nil {
		return nil, err
	}

	var rd ResponseData
	rd.StatusCode = resp.StatusCode

	defer resp.Body.Close()
	rbody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	rd.Body = rbody

	return &rd, nil
}
