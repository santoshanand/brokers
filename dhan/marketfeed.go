package dhan

import (
	"fmt"
	"net/http"
)

type LTPRequest map[ExchangeSegment][]int64

type LTPResponse struct {
	Data   LTPData `json:"data"`
	Status string  `json:"status"`
}

type LTPData map[ExchangeSegment]map[int64]LTPOhlc

type LTP struct {
	LastPrice float64 `json:"last_price"`
}

type OHLCResponse struct {
	Data   OHLCData `json:"data"`
	Status string   `json:"status"`
}

type LTPOhlc struct {
	LastPrice float64 `json:"last_price"`
	OHLC      OHLC    `json:"ohlc"`
}

type OHLC struct {
	Open   float64 `json:"open"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Close  float64 `json:"close"`
	Volume int64   `json:"volume"`
}
type OHLCData map[ExchangeSegment]map[int64]LTP

type OHLCRequest LTPRequest

func (c *Client) GetLTP(accessToken, clientID string, req LTPRequest) (*LTPResponse, error) {
	var resp LTPResponse
	r, err := c.resty.R().
		SetHeader(accept, contentTypeJSON).
		SetHeader(authorization, accessToken).
		SetHeader(clientIDHeader, clientID).
		SetHeader("content-type", contentTypeJSON).
		SetForceResponseContentType(contentTypeJSON).
		SetBody(req).
		SetResult(&resp).
		Post(marketFeedLTPURL)
	if r.IsError() || r.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch LTP: %s", string(r.Bytes()))
	}
	return &resp, err
}

func (c *Client) GetOHLC(accessToken, clientID string, req OHLCRequest) (*OHLCResponse, error) {
	var resp OHLCResponse
	r, err := c.resty.R().
		SetHeader(accept, contentTypeJSON).
		SetHeader(authorization, accessToken).
		SetHeader(clientIDHeader, clientID).
		SetBody(req).
		SetResult(&resp).
		Post(marketFeedOHLCURL)
	if r.IsError() || r.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch OHLC: %s", string(r.Bytes()))
	}
	return &resp, err
}
