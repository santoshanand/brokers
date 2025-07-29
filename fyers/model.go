package fyers

import (
	"crypto/sha256"
	"encoding/hex"
)

type CommonResponse struct {
	S       string `json:"s"`
	Code    int64  `json:"code"`
	Message string `json:"message"`
}
type ProfileRes struct {
	CommonResponse
	Data Profile `json:"data"`
}

type Profile struct {
	Name          string `json:"name"`
	Image         string `json:"image"`
	DisplayName   string `json:"display_name"`
	EmailID       string `json:"email_id"`
	Pan           string `json:"PAN"`
	FyID          string `json:"fy_id"`
	PinChangeDate string `json:"pin_change_date"`
	MobileNumber  string `json:"mobile_number"`
	Totp          bool   `json:"totp"`
	PwdChangeDate string `json:"pwd_change_date"`
	PwdToExpire   int64  `json:"pwd_to_expire"`
	DdpiEnabled   bool   `json:"ddpi_enabled"`
	MtfEnabled    bool   `json:"mtf_enabled"`
}

type FundResponse struct {
	CommonResponse
	Funds []Fund `json:"fund_limit"`
}

type Fund struct {
	ID              int64   `json:"id"`
	Title           string  `json:"title"`
	EquityAmount    float64 `json:"equityAmount"`
	CommodityAmount int64   `json:"commodityAmount"`
}

type HoldingResponse struct {
	CommonResponse
	Holdings []Holding      `json:"holdings"`
	Overall  OverallHolding `json:"overall"`
}

type Holding struct {
	HoldingType             string  `json:"holdingType"`
	Quantity                int64   `json:"quantity"`
	CostPrice               float64 `json:"costPrice"`
	MarketVal               float64 `json:"marketVal"`
	RemainingQuantity       int64   `json:"remainingQuantity"`
	Pl                      float64 `json:"pl"`
	Ltp                     float64 `json:"ltp"`
	ID                      int64   `json:"id"`
	FyToken                 int64   `json:"fyToken"`
	Exchange                int64   `json:"exchange"`
	Symbol                  string  `json:"symbol"`
	Segment                 int64   `json:"segment"`
	Isin                    string  `json:"isin"`
	QtyT1                   int64   `json:"qty_t1"`
	RemainingPledgeQuantity int64   `json:"remainingPledgeQuantity"`
	CollateralQuantity      int64   `json:"collateralQuantity"`
}

type OverallHolding struct {
	CountTotal        int64   `json:"count_total"`
	TotalInvestment   float64 `json:"total_investment"`
	TotalCurrentValue float64 `json:"total_current_value"`
	TotalPl           float64 `json:"total_pl"`
	PnlPerc           float64 `json:"pnl_perc"`
}

type OrdersResponse struct {
	CommonResponse
	OrderBook []OrderBook `json:"orderBook"`
}

type OrderBook struct {
	ClientID          string        `json:"clientId"`
	ID                string        `json:"id"`
	ExchOrdID         string        `json:"exchOrdId"`
	Qty               int64         `json:"qty"`
	RemainingQuantity int64         `json:"remainingQuantity"`
	FilledQty         int64         `json:"filledQty"`
	DiscloseQty       int64         `json:"discloseQty"`
	LimitPrice        float64       `json:"limitPrice"`
	StopPrice         int64         `json:"stopPrice"`
	TradedPrice       float64       `json:"tradedPrice"`
	Type              OrderType     `json:"type"`
	FyToken           string        `json:"fyToken"`
	Exchange          int64         `json:"exchange"`
	Segment           Segment       `json:"segment"`
	Symbol            string        `json:"symbol"`
	Instrument        int64         `json:"instrument"`
	Message           string        `json:"message"`
	OfflineOrder      bool          `json:"offlineOrder"`
	OrderDateTime     string        `json:"orderDateTime"`
	OrderValidity     OrderValidity `json:"orderValidity"`
	Pan               string        `json:"pan"`
	ProductType       string        `json:"productType"`
	Side              OrderSide     `json:"side"`
	Status            OrderStatus   `json:"status"`
	Source            string        `json:"source"`
	ExSym             string        `json:"ex_sym"`
	Description       string        `json:"description"`
	Ch                float64       `json:"ch"`
	Chp               float64       `json:"chp"`
	Lp                float64       `json:"lp"`
	SlNo              int64         `json:"slNo"`
	DqQtyRem          int64         `json:"dqQtyRem"`
	OrderNumStatus    string        `json:"orderNumStatus"`
	DisclosedQty      int64         `json:"disclosedQty"`
	OrderTag          string        `json:"orderTag"`
}

type PositionsResponse struct {
	S            string          `json:"s"`
	Code         int64           `json:"code"`
	Message      string          `json:"message"`
	NetPositions []Position      `json:"netPositions"`
	Overall      OverallPosition `json:"overall"`
}

type Position struct {
	NetQty           int64        `json:"netQty"`
	Qty              int64        `json:"qty"`
	AvgPrice         int64        `json:"avgPrice"`
	NetAvg           int64        `json:"netAvg"`
	Side             PositionSide `json:"side"`
	ProductType      ProductType  `json:"productType"`
	RealizedProfit   int64        `json:"realized_profit"`
	UnrealizedProfit int64        `json:"unrealized_profit"`
	Pl               int64        `json:"pl"`
	Ltp              int64        `json:"ltp"`
	BuyQty           int64        `json:"buyQty"`
	BuyAvg           int64        `json:"buyAvg"`
	BuyVal           int64        `json:"buyVal"`
	SellQty          int64        `json:"sellQty"`
	SellAvg          int64        `json:"sellAvg"`
	SellVal          int64        `json:"sellVal"`
	SlNo             int64        `json:"slNo"`
	FyToken          string       `json:"fyToken"`
	CrossCurrency    string       `json:"crossCurrency"`
	RbiRefRate       int64        `json:"rbiRefRate"`
	QtyMultiCOM      int64        `json:"qtyMulti_com"`
	Segment          Segment      `json:"segment"`
	Symbol           string       `json:"symbol"`
	ID               string       `json:"id"`
	CFBuyQty         int64        `json:"cfBuyQty"`
	CFSellQty        int64        `json:"cfSellQty"`
	DayBuyQty        int64        `json:"dayBuyQty"`
	DaySellQty       int64        `json:"daySellQty"`
	Exchange         Exchange     `json:"exchange"`
}

type OverallPosition struct {
	CountTotal   int64 `json:"count_total"`
	CountOpen    int64 `json:"count_open"`
	PlTotal      int64 `json:"pl_total"`
	PlRealized   int64 `json:"pl_realized"`
	PlUnrealized int64 `json:"pl_unrealized"`
}

type TradesResponse struct {
	CommonResponse
	TradeBook []TradeBook `json:"tradeBook"`
}

type TradeBook struct {
	ClientID        string      `json:"clientId"`
	OrderDateTime   string      `json:"orderDateTime"`
	OrderNumber     string      `json:"orderNumber"`
	ExchangeOrderNo string      `json:"exchangeOrderNo"`
	Exchange        Exchange    `json:"exchange"`
	Side            OrderSide   `json:"side"`
	Segment         Segment     `json:"segment"`
	OrderType       int64       `json:"orderType"`
	FyToken         string      `json:"fyToken"`
	ProductType     ProductType `json:"productType"`
	TradedQty       int64       `json:"tradedQty"`
	TradePrice      float64     `json:"tradePrice"`
	TradeValue      float64     `json:"tradeValue"`
	TradeNumber     string      `json:"tradeNumber"`
	Row             int64       `json:"row"`
	Symbol          string      `json:"symbol"`
	OrderTag        string      `json:"orderTag"`
}

type Order struct {
	Symbol       string    `json:"symbol"`
	Qty          int64     `json:"qty"`
	Type         OrderType `json:"type"`
	Side         OrderSide `json:"side"`
	ProductType  string    `json:"productType"`
	LimitPrice   float64   `json:"limitPrice"`
	StopPrice    int64     `json:"stopPrice"`
	DisclosedQty int64     `json:"disclosedQty"`
	Validity     string    `json:"validity"`
	OfflineOrder bool      `json:"offlineOrder"`
	StopLoss     int64     `json:"stopLoss"`
	TakeProfit   int64     `json:"takeProfit"`
}
type AuthRequest struct {
	GrantType string `json:"grant_type"`
	AppIDHash string `json:"appIdHash"`
	Code      string `json:"code"`
}

func (a *AuthRequest) WithAppIDHash(appID, appSecret string) *AuthRequest {
	a.AppIDHash = Sha256Hash(appID + appSecret)
	return a

}
func Sha256Hash(input string) string {
	hash := sha256.Sum256([]byte(input))
	return hex.EncodeToString(hash[:])
}

type AuthResponse struct {
	CommonResponse
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenRequest struct {
	GrantType    string `json:"grant_type"`
	AppIDHash    string `json:"appIdHash"`
	RefreshToken string `json:"refresh_token"`
	Pin          string `json:"pin"`
}

type RefreshTokenResponse struct {
	CommonResponse
	AccessToken string `json:"access_token"`
}
