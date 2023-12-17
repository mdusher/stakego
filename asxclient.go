package stakego

import (
  "net/http"
  "net/url"
  "io"
)

// NewASXClient - create and initialise an ASXClient
func NewASXClient() ASXClient {
  c := ASXClient{}
  c.Init()
  return c
}

// ASXClient - Client for interacting with Stake ASX
type ASXClient struct {
  apiUrl string
  Credentials *Credentials
  User *User
  httpclient http.Client
}

// Init - initialise the ASX client with defaults
func (c *ASXClient) Init() {
  c.apiUrl = "https://global-prd-api.hellostake.com/api/"
  c.httpclient = NewHTTPClient()
}

// Login - create a user session
func (c *ASXClient) Login() (err error) {
  if c.Credentials.StakeSessionToken == "" {
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
      c.Credentials.StakeSessionToken = u.SessionKey
    }
  }

  if c.Credentials.StakeSessionToken == "" {
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
  if c.Credentials.StakeSessionToken == "" {
    return NewStakeError("logout", ErrSessionTokenMissing)
  }

  u, err := url.JoinPath(c.apiUrl, "userauth", c.Credentials.StakeSessionToken)
  if err != nil {
    return NewStakeError("logout", err)
  }

  req, _ := NewJSONRequest("DELETE", u, nil)
  resp, err := c.httpclient.Do(req)
  if err != nil {
    return NewStakeError("logout", err)
  }
  if resp.StatusCode == 200 {
    c.Credentials.StakeSessionToken = ""
    return nil
  }

  return NewStakeError("cash", ErrInvalidAPIResponse)
}

// GetMarket - Get the current market status
func (c *ASXClient) GetMarket() (*Market, error) {
  u := "https://early-bird-promo.hellostake.com/marketStatus"
  req, _ := NewJSONRequest("GET", u, nil)
  resp, err := c.httpclient.Do(req)
  if err != nil {
    return nil, NewStakeError("market", err)
  }

  if resp.StatusCode == 200 {
    defer resp.Body.Close()
    rbody, err := io.ReadAll(resp.Body)
    if err != nil {
      return nil, NewStakeError("market", err)
    }
    m := NewMarketFromJSON(rbody)
    return m, nil
  }

  return nil, NewStakeError("market", ErrInvalidAPIResponse)
}

// GetCash - get the current available cash
func (c *ASXClient) GetCash() (*Cash, error) {
  if c.Credentials.StakeSessionToken == "" {
    return nil, NewStakeError("cash", ErrSessionTokenMissing)
  }

  u, err := url.JoinPath(c.apiUrl, "asx/cash")
  if err != nil {
    return nil, NewStakeError("cash", err)
  }

  req, _ := NewJSONRequest("GET", u, nil)
  req.Header.Set("Stake-Session-Token", c.Credentials.StakeSessionToken)
  resp, err := c.httpclient.Do(req)
  if err != nil {
    return nil, NewStakeError("cash", err)
  }

  if resp.StatusCode == 200 {
    defer resp.Body.Close()
    rbody, _ := io.ReadAll(resp.Body)
    c := NewCashFromJSON(rbody)
    return c, nil
  }
  return nil, NewStakeError("cash", ErrInvalidAPIResponse)
}

// GetEquityPoisitions - get the current user's equity positions
func (c *ASXClient) GetEquityPositions() (*EquityPositions, error) {
  if c.Credentials.StakeSessionToken == "" {
    return nil, NewStakeError("equity positions", ErrSessionTokenMissing)
  }

  u, err := url.JoinPath(c.apiUrl, "asx/instrument/equityPositions")
  if err != nil {
    return nil, NewStakeError("equity positions", err)
  }

  req, _ := NewJSONRequest("GET", u, nil)
  req.Header.Set("Stake-Session-Token", c.Credentials.StakeSessionToken)
  resp, err := c.httpclient.Do(req)
  if err != nil {
    return nil, NewStakeError("equity positions", err)
  }

  if resp.StatusCode == 200 {
    defer resp.Body.Close()
    rbody, _ := io.ReadAll(resp.Body)
    e := NewEquityPositionsFromJSON(rbody)
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

  req, _ := NewJSONRequest("GET", u, nil)
  req.Header.Set("Stake-Session-Token", c.Credentials.StakeSessionToken)
  resp, err := c.httpclient.Do(req)
  if err != nil {
    return nil, NewStakeError("user", err)
  }

  if resp.StatusCode == 200 {
    defer resp.Body.Close()
    rbody, _ := io.ReadAll(resp.Body)
    user := NewUserFromJSON(rbody)
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

  req, _ := NewJSONRequest("GET", u, nil)
  req.Header.Set("Stake-Session-Token", c.Credentials.StakeSessionToken)
  resp, err := c.httpclient.Do(req)
  if err != nil {
    return nil, NewStakeError("orders", err)
  }

  if resp.StatusCode == 200 {
    defer resp.Body.Close()
    rbody, _ := io.ReadAll(resp.Body)
    orders := NewOrderListFromJSON(rbody)
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

  req, _ := NewJSONRequest("POST", u, order.AsJSON())
  req.Header.Set("Stake-Session-Token", c.Credentials.StakeSessionToken)
  resp, err := c.httpclient.Do(req)
  if err != nil {
    return nil, NewStakeError("orders/place", err)
  }

  if resp.StatusCode == 200 {
    defer resp.Body.Close()
    rbody, _ := io.ReadAll(resp.Body)
    orders := NewOrderResponseFromJSON(rbody)
    return orders, nil
  }

  return nil, NewStakeError("orders/place", ErrInvalidAPIResponse)
}

// CancelOrder - cancel an order
func (c *ASXClient) CancelOrder(uuid string) (error) {
  u, err := url.JoinPath(c.apiUrl, "asx/orders", uuid, "cancel")
  if err != nil {
    return NewStakeError("orders/cancel", err)
  }

  req, _ := NewJSONRequest("POST", u, nil)
  req.Header.Set("Stake-Session-Token", c.Credentials.StakeSessionToken)
  resp, err := c.httpclient.Do(req)
  if err != nil {
    return NewStakeError("orders/cancel", err)
  }

  if resp.StatusCode == 200 {
    return nil
  }

  return NewStakeError("orders/cancel", ErrInvalidAPIResponse)
}