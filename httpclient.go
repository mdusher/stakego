package stakego

import (
  "io"
  "bytes"
  "net"
  "net/http"
  "net/url"
  "time"
)

// NewHTTPClient - create a http.Client with a connect timeout
func NewHTTPClient() http.Client {
  transport := &http.Transport{
    DialContext: (&net.Dialer{
      Timeout:   10 * time.Second,
      KeepAlive: 30 * time.Second,
      DualStack: true,
    }).DialContext,
    MaxIdleConns:          100,
    MaxIdleConnsPerHost:   100,
    IdleConnTimeout:       90 * time.Second,
    TLSHandshakeTimeout:   10 * time.Second,
    ResponseHeaderTimeout: 5 * time.Second,
  }
  client := http.Client{
    Transport: transport,
    CheckRedirect: func(req *http.Request, via []*http.Request) error {
      // Disable following redirects
      return http.ErrUseLastResponse
    },
  }
  return client
}

// NewRequest - create a new request using details defined in WebDAV
func NewRequest(method string, fullurl string, body io.Reader) (*http.Request, error) {
  purl, err := url.Parse(fullurl)
  if err != nil {
    return nil, err
  }

  r, err := http.NewRequest(method, purl.String(), body)

  if err != nil {
    return nil, err
  }
  r.Close = true // Close the request after reading response

  return r, nil
}

// NewJSONRequest - creates a new request, expecting a json byte slice to be passed as the body.
func NewJSONRequest(method string, fullurl string, jsonBody []byte) (*http.Request, error) {
  var bodyReader io.Reader
  if string(jsonBody) != "" {
    bodyReader = bytes.NewBuffer(jsonBody)
  }
  r, err := NewRequest(method, fullurl, bodyReader)
  if err != nil {
    return nil, err
  }
  r.Header.Set("Content-Type", "application/json;charset=UTF-8")
  r.Header.Set("Accept", "application/json")
  return r, nil
}