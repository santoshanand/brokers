package dhan

import (
	"fmt"
	"net/http"
	"time"
)

type SuperOrderRequest struct {
	DhanClientID    string          `json:"dhanClientId"`
	CorrelationID   string          `json:"correlationId"`
	TransactionType TransactionType `json:"transactionType"`
	ExchangeSegment ExchangeSegment `json:"exchangeSegment"`
	ProductType     ProductType     `json:"productType"`
	OrderType       OrderType       `json:"orderType"`
	SecurityID      string          `json:"securityId"`
	Quantity        int64           `json:"quantity"`
	Price           float64         `json:"price"`
	TargetPrice     float64         `json:"targetPrice"`
	StopLossPrice   float64         `json:"stopLossPrice"`
	TrailingJump    float64         `json:"trailingJump"`
}
type SuperOrderStatus OrderStatus
type SuperOrderResponse struct {
	OrderID     string           `json:"orderId"`
	OrderStatus SuperOrderStatus `json:"orderStatus"`
}

type SuperOrdersResponse []SuperOrdersResponseElement

type SuperOrdersResponseElement struct {
	DhanClientID        string          `json:"dhanClientId"`
	OrderID             string          `json:"orderId"`
	CorrelationID       string          `json:"correlationId"`
	OrderStatus         OrderStatus     `json:"orderStatus"`
	TransactionType     TransactionType `json:"transactionType"`
	ExchangeSegment     ExchangeSegment `json:"exchangeSegment"`
	ProductType         ProductType     `json:"productType"`
	OrderType           OrderType       `json:"orderType"`
	Validity            Validity        `json:"validity"`
	TradingSymbol       string          `json:"tradingSymbol"`
	SecurityID          string          `json:"securityId"`
	Quantity            int64           `json:"quantity"`
	RemainingQuantity   int64           `json:"remainingQuantity"`
	Ltp                 float64         `json:"ltp"`
	Price               int64           `json:"price"`
	AfterMarketOrder    bool            `json:"afterMarketOrder"`
	LegName             LegName         `json:"legName"`
	ExchangeOrderID     string          `json:"exchangeOrderId"`
	CreateTime          time.Time       `json:"createTime"`
	UpdateTime          time.Time       `json:"updateTime"`
	ExchangeTime        time.Time       `json:"exchangeTime"`
	OmsErrorDescription string          `json:"omsErrorDescription"`
	AverageTradedPrice  float64         `json:"averageTradedPrice"`
	FilledQty           int64           `json:"filledQty"`
	LegDetails          []LegDetail     `json:"legDetails"`
}

type LegDetail struct {
	OrderID           string  `json:"orderId"`
	LegName           LegName `json:"legName"`
	TransactionType   string  `json:"transactionType"`
	TotalQuatity      *int64  `json:"totalQuatity,omitempty"`
	RemainingQuantity int64   `json:"remainingQuantity"`
	TriggeredQuantity int64   `json:"triggeredQuantity"`
	Price             float64 `json:"price"`
	OrderStatus       string  `json:"orderStatus"`
	TrailingJump      float64 `json:"trailingJump"`
}

func (c *Client) CreateSuperOrder(accessToken string, order SuperOrderRequest) (*SuperOrderResponse, error) {
	var resp SuperOrderResponse
	r, err := c.resty.R().
		SetHeader(accept, contentTypeJSON).
		SetHeader(authorization, accessToken).
		SetBody(order).
		SetResult(&resp).
		Post(superOrdersURL)
	if r.IsError() || r.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("failed to create super order: %s", string(r.Bytes()))
	}
	return &resp, err
}

func (c *Client) ModifySuperOrder(accessToken, orderID string, order SuperOrderRequest) (*SuperOrderResponse, error) {
	var resp SuperOrderResponse
	r, err := c.resty.R().
		SetHeader(accept, contentTypeJSON).
		SetHeader(authorization, accessToken).
		SetBody(order).
		SetResult(&resp).
		Put(fmt.Sprintf(superOrderIDURL, orderID))
	if r.IsError() || r.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("failed to modify super order: %s", string(r.Bytes()))
	}
	return &resp, err
}

func (c *Client) CancelSuperOrderLeg(accessToken, orderID, orderLeg string) error {
	r, err := c.resty.R().
		SetHeader(accept, contentTypeJSON).
		SetHeader(authorization, accessToken).
		Delete(fmt.Sprintf(superOrderLegURL, orderID, orderLeg))
	if r.IsError() || r.StatusCode() != http.StatusOK {
		return fmt.Errorf("failed to cancel super order leg: %s", string(r.Bytes()))
	}
	return err
}

func (c *Client) GetSuperOrders(accessToken string) (SuperOrdersResponse, error) {
	var resp SuperOrdersResponse
	r, err := c.resty.R().
		SetHeader(accept, contentTypeJSON).
		SetHeader(authorization, accessToken).
		SetResult(&resp).
		Get(superOrdersURL)
	if r.IsError() || r.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch super orders: %s", string(r.Bytes()))
	}
	return resp, err
}
