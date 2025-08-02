package dhan

type HoldingResponse []HoldingResponseElement

type HoldingResponseElement struct {
	Exchange      string  `json:"exchange"`
	TradingSymbol string  `json:"tradingSymbol"`
	SecurityID    string  `json:"securityId"`
	Isin          string  `json:"isin"`
	TotalQty      int64   `json:"totalQty"`
	DPQty         int64   `json:"dpQty"`
	T1Qty         int64   `json:"t1Qty"`
	AvailableQty  int64   `json:"availableQty"`
	CollateralQty int64   `json:"collateralQty"`
	AvgCostPrice  float64 `json:"avgCostPrice"`
}

type PositionResponse []PositionResponseElement

type PositionResponseElement struct {
	DhanClientID          string          `json:"dhanClientId"`
	TradingSymbol         string          `json:"tradingSymbol"`
	SecurityID            string          `json:"securityId"`
	PositionType          PositionType    `json:"positionType"`
	ExchangeSegment       ExchangeSegment `json:"exchangeSegment"`
	ProductType           ProductType     `json:"productType"`
	BuyAvg                float64         `json:"buyAvg"`
	BuyQty                int64           `json:"buyQty"`
	CostPrice             int64           `json:"costPrice"`
	SellAvg               float64         `json:"sellAvg"`
	SellQty               int64           `json:"sellQty"`
	NetQty                int64           `json:"netQty"`
	RealizedProfit        float64         `json:"realizedProfit"`
	UnrealizedProfit      float64         `json:"unrealizedProfit"`
	RbiReferenceRate      float64         `json:"rbiReferenceRate"`
	Multiplier            int64           `json:"multiplier"`
	CarryForwardBuyQty    int64           `json:"carryForwardBuyQty"`
	CarryForwardSellQty   int64           `json:"carryForwardSellQty"`
	CarryForwardBuyValue  float64         `json:"carryForwardBuyValue"`
	CarryForwardSellValue float64         `json:"carryForwardSellValue"`
	DayBuyQty             int64           `json:"dayBuyQty"`
	DaySellQty            int64           `json:"daySellQty"`
	DayBuyValue           float64         `json:"dayBuyValue"`
	DaySellValue          float64         `json:"daySellValue"`
	DrvExpiryDate         string          `json:"drvExpiryDate"`
	DrvOptionType         string          `json:"drvOptionType"`
	DrvStrikePrice        float64         `json:"drvStrikePrice"`
	CrossCurrency         bool            `json:"crossCurrency"`
}
