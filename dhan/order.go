package dhan

import (
	"errors"
	"fmt"
	"net/http"
)

// --- Structs for Request and Response ---

type OrderRequest struct {
	DhanClientID      string          `json:"dhanClientId"`
	CorrelationID     string          `json:"correlationId"`
	TransactionType   TransactionType `json:"transactionType"`
	ExchangeSegment   ExchangeSegment `json:"exchangeSegment"`
	ProductType       ProductType     `json:"productType"`
	OrderType         OrderType       `json:"orderType"`
	Validity          Validity        `json:"validity"`
	SecurityID        string          `json:"securityId"`
	Quantity          int64           `json:"quantity"`
	DisclosedQuantity int64           `json:"disclosedQuantity"`
	Price             float64         `json:"price"`
	TriggerPrice      float64         `json:"triggerPrice"`
	AfterMarketOrder  bool            `json:"afterMarketOrder"`
	AmoTime           AMOTime         `json:"amoTime"`
	BoProfitValue     float64         `json:"boProfitValue"`
	BoStopLossValue   float64         `json:"boStopLossValue"`
}

type ModifyOrderRequest struct {
	DhanClientID      string    `json:"dhanClientId"`
	OrderID           string    `json:"orderId"`
	OrderType         OrderType `json:"orderType"`
	LegName           LegName   `json:"legName"`
	Quantity          int64     `json:"quantity"`
	Price             float64   `json:"price"`
	DisclosedQuantity int64     `json:"disclosedQuantity"`
	TriggerPrice      float64   `json:"triggerPrice"`
	Validity          Validity  `json:"validity"`
}

type OrderResponse struct {
	OrderID     string      `json:"orderId"`
	OrderStatus OrderStatus `json:"orderStatus"`
}

type SliceOrderRequest struct {
	Symbol   string  `json:"symbol"`
	TotalQty int64   `json:"total_quantity"`
	Price    float64 `json:"price"`
	Legs     int64   `json:"legs"`
	Side     string  `json:"side"`
	Product  string  `json:"product"`
	Exchange string  `json:"exchange"`
}

// POST /orders
func (c *Client) PlaceOrder(accessToken string, order OrderRequest) (*OrderResponse, error) {
	var resp OrderResponse
	r, err := c.resty.R().
		SetHeader(accept, contentTypeJSON).
		SetHeader(authorization, accessToken).
		SetBody(order).
		SetResult(&resp).
		Post(ordersURL)
	if r.IsError() || r.StatusCode() != http.StatusOK {
		return nil, errors.New("failed to post order: " + string(r.Bytes()))
	}
	return &resp, err
}

// PUT /orders/{order-id}
func (c *Client) ModifyOrder(accessToken string, orderID string, updated ModifyOrderRequest) (*OrderResponse, error) {
	var resp OrderResponse
	r, err := c.resty.R().
		SetHeader(authorization, accessToken).
		SetBody(updated).
		SetResult(&resp).
		Put(fmt.Sprintf("%s/%s", ordersURL, orderID))
	if r.IsError() || r.StatusCode() != http.StatusOK {
		return nil, errors.New("failed to modify order: " + string(r.Bytes()))
	}
	return &resp, err
}

// DELETE /orders/{order-id}
func (c *Client) CancelOrder(accessToken, orderID string) (*OrderResponse, error) {
	var resp OrderResponse
	_, err := c.resty.R().
		SetHeader(authorization, accessToken).
		SetResult(&resp).
		Delete(fmt.Sprintf("%s/%s", ordersURL, orderID))
	return &resp, err
}

// POST /orders/slicing
func (c *Client) SliceOrder(request SliceOrderRequest) (*OrderResponse, error) {
	var resp OrderResponse
	_, err := c.resty.R().
		SetBody(request).
		SetResult(&resp).
		Post(sliceOrderURL)
	return &resp, err
}

// GET /orders
// Orders implements IDhan.
func (d *Client) Orders(accessToken string) (Orders, error) {
	res := Orders{}
	r, err := d.resty.R().
		SetHeader(accept, contentTypeJSON).
		SetHeader(authorization, accessToken).
		SetResult(&res).
		Get(ordersURL)
	if err != nil {
		return res, err
	}
	if r.IsError() || r.StatusCode() != http.StatusOK {
		return nil, errors.New("failed to fetch Orders: " + string(r.Bytes()))
	}
	return res, nil
}

// GET /orders/{order-id}
func (c *Client) GetOrderByID(accessToken, orderID string) (*OrderResponse, error) {
	var resp OrderResponse
	r, err := c.resty.R().
		SetHeader(authorization, accessToken).
		SetResult(&resp).
		Get(fmt.Sprintf("%s/%s", ordersURL, orderID))
	if r.IsError() || r.StatusCode() != http.StatusOK {
		return nil, errors.New("failed to fetch get order by id: " + string(r.Bytes()))
	}
	return &resp, err
}

// GET /orders/external/{correlation-id}
func (c *Client) GetOrderByCorrelationID(accessToken, correlationID string) (*OrderResponse, error) {
	var resp OrderResponse
	r, err := c.resty.R().
		SetHeader(authorization, accessToken).
		SetResult(&resp).
		Get(fmt.Sprintf("%s/%s", orderExternalURL, correlationID))
	if r.IsError() || r.StatusCode() != http.StatusOK {
		return nil, errors.New("failed to fetch order by correlation id: " + string(r.Bytes()))
	}
	return &resp, err
}

// GET /trades
func (c *Client) GetAllTrades(accessToken string) ([]Trade, error) {
	var trades []Trade
	r, err := c.resty.R().
		SetHeader(authorization, accessToken).
		SetResult(&trades).
		Get(tradesURL)
	if r.IsError() || r.StatusCode() != http.StatusOK {
		return nil, errors.New("failed to fetch trades: " + string(r.Bytes()))
	}
	return trades, err
}

// GET /trades/{order-id}
func (c *Client) GetTradeByOrderID(accessToken, orderID string) ([]Trade, error) {
	var trades []Trade
	r, err := c.resty.R().
		SetHeader(authorization, accessToken).
		SetResult(&trades).
		Get(fmt.Sprintf("%s/%s", tradesURL, orderID))
	if r.IsError() || r.StatusCode() != http.StatusOK {
		return nil, errors.New("failed to fetch trade by order id: " + string(r.Bytes()))
	}
	return trades, err
}
