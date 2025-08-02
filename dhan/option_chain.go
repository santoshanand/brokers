package dhan

import (
	"fmt"
	"net/http"
)

type OptionChainRequest struct {
	UnderlyingScrip int64           `json:"UnderlyingScrip"`
	UnderlyingSeg   ExchangeSegment `json:"UnderlyingSeg"`
	Expiry          string          `json:"Expiry"`
}

type OptionExpiryRequest struct {
	UnderlyingScrip int64           `json:"UnderlyingScrip"`
	UnderlyingSeg   ExchangeSegment `json:"UnderlyingSeg"`
}

type OptionExpiryResponse struct {
	Data   []string `json:"data"`
	Status string   `json:"status"`
}

type OptionChainResponse struct {
	Data OptionChainData `json:"data"`
}

type OptionChainData struct {
	LastPrice float64                      `json:"last_price"`
	OC        map[string]OptionChainStrike `json:"oc"` // strike-wise array
}

type OptionChainStrike struct {
	CE *OptionData `json:"ce,omitempty"`
	PE *OptionData `json:"pe,omitempty"`
}

type OptionData struct {
	Greeks             Greeks  `json:"greeks"`
	ImpliedVolatility  float64 `json:"implied_volatility"`
	LastPrice          float64 `json:"last_price"`
	OI                 int     `json:"oi"`
	PreviousClosePrice float64 `json:"previous_close_price"`
	PreviousOI         int     `json:"previous_oi"`
	PreviousVolume     int     `json:"previous_volume"`
	TopAskPrice        float64 `json:"top_ask_price"`
	TopAskQuantity     int     `json:"top_ask_quantity"`
	TopBidPrice        float64 `json:"top_bid_price"`
	TopBidQuantity     int     `json:"top_bid_quantity"`
	Volume             int     `json:"volume"`
}

type Greeks struct {
	Delta float64 `json:"delta"`
	Theta float64 `json:"theta"`
	Gamma float64 `json:"gamma"`
	Vega  float64 `json:"vega"`
}

func (c *Client) GetOptionChain(accessToken, clientID string, req OptionChainRequest) (*OptionChainResponse, error) {
	var resp OptionChainResponse
	r, err := c.resty.R().
		SetHeader(accept, contentTypeJSON).
		SetHeader(authorization, accessToken).
		SetHeader(clientIDHeader, clientID).
		SetBody(req).
		SetResult(&resp).
		Post(optionChainURL)
	if r.IsError() || r.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch option chain: %s", string(r.Bytes()))
	}
	return &resp, err
}

func (c *Client) GetOptionExpiryList(accessToken, clientID string, req OptionExpiryRequest) (*OptionExpiryResponse, error) {
	var resp OptionExpiryResponse
	r, err := c.resty.R().
		SetHeader(accept, contentTypeJSON).
		SetHeader(authorization, accessToken).
		SetHeader(clientIDHeader, clientID).
		SetBody(req).
		SetResult(&resp).
		Post(optionChainExpiryURL)
	if r.IsError() || r.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch option expiry list: %s", string(r.Bytes()))
	}
	return &resp, err
}
