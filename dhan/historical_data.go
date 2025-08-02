package dhan

import (
	"fmt"
	"net/http"
)

type HistoricalChartRequest struct {
	SecurityID      string          `json:"securityId"`
	ExchangeSegment ExchangeSegment `json:"exchangeSegment"`
	Instrument      InstrumentType  `json:"instrument"`
	ExpiryCode      int             `json:"expiryCode"`
	Oi              bool            `json:"oi"`
	FromDate        string          `json:"fromDate"`
	ToDate          string          `json:"toDate"`
}

type HistoricalChartResponse struct {
	Open         []float64 `json:"open"`
	High         []float64 `json:"high"`
	Low          []float64 `json:"low"`
	Close        []float64 `json:"close"`
	Volume       []float64 `json:"volume"`
	Timestamp    []float64 `json:"timestamp"`
	OpenInterest []float64 `json:"open_interest"`
}

type IntradayChartRequest struct {
	SecurityID      string          `json:"securityId"`
	ExchangeSegment ExchangeSegment `json:"exchangeSegment"`
	Instrument      InstrumentType  `json:"instrument"`
	Interval        Interval        `json:"interval"`
	Oi              bool            `json:"oi"`
	FromDate        string          `json:"fromDate"` //"2024-09-11 09:30:00",
	ToDate          string          `json:"toDate"`   //"2024-09-15 13:00:00"
}

func (c *Client) GetHistoricalOHLC(accessToken string, req HistoricalChartRequest) (*HistoricalChartResponse, error) {
	var resp HistoricalChartResponse
	r, err := c.resty.R().
		SetHeader(accept, contentTypeJSON).
		SetHeader(authorization, accessToken).
		SetBody(req).
		SetResult(&resp).
		Post(chartsHistoricalURL)
	if r.IsError() || r.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch historical OHLC: %s", string(r.Bytes()))
	}
	return &resp, err
}

func (c *Client) GetIntradayOHLC(accessToken string, req IntradayChartRequest) (*HistoricalChartResponse, error) {
	var resp HistoricalChartResponse
	r, err := c.resty.R().
		SetHeader(accept, contentTypeJSON).
		SetHeader(authorization, accessToken).
		SetBody(req).
		SetResult(&resp).
		Post(chartsIntradayURL)
	if r.IsError() || r.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch intraday OHLC: %s", string(r.Bytes()))
	}
	return &resp, err
}
