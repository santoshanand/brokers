package dhan

import (
	"strings"
	"time"
)

type Interval int

const (
	Interval1Min  Interval = 1
	Interval5Min  Interval = 5
	Interval15Min Interval = 15
	Interval30Min Interval = 25
	Interval1Hour Interval = 60
)

type PositionType string

const (
	PositionTypeLong   PositionType = "LONG"   // Cash & Carry (Equity Delivery)
	PositionTypeSHORT  PositionType = "SHORT"  // Intraday (Equity, F&O)
	PositionTypeCLOSED PositionType = "CLOSED" // Carry Forward (F&O)
)

type LegName string

const (
	LegNameENTRY         LegName = "ENTRY_LEG"
	LegNameTARGET_LEG    LegName = "TARGET_LEG"
	LegNameSTOP_LOSS_LEG LegName = "STOP_LOSS_LEG"
)

type Validity string

const (
	ValidityDAY Validity = "DAY" // Valid for the day
	ValidityIOC Validity = "IOC" // Immediate or Cancel
)

type TransactionType string

const (
	TransactionTypeBUY  TransactionType = "BUY"
	TransactionTypeSELL TransactionType = "SELL"
)

type OrderType string

const (
	OrderTypeMARKET         OrderType = "MARKET"           // Market Order
	OrderTypeLIMIT          OrderType = "LIMIT"            // Limit Order
	OrderTypeSTOPLOSS       OrderType = "STOP_LOSS"        // Stop Loss Order
	OrderTypeSTOPLOSSMARKET OrderType = "STOP_LOSS_MARKET" // Stop Limit Order
)

type ExchangeSegment string

const (
	ExchangeSegmentIDXI        ExchangeSegment = "IDX_I"
	ExchangeSegmentNSEEQ       ExchangeSegment = "NSE_EQ"
	ExchangeSegmentNSEFNO      ExchangeSegment = "NSE_FNO"
	ExchangeSegmentNSECURRENCY ExchangeSegment = "NSE_CURRENCY"
	ExchangeSegmentBSEEQ       ExchangeSegment = "BSE_EQ"
	ExchangeSegmentMCXCOMM     ExchangeSegment = "MCX_COMM"
	ExchangeSegmentBSECURRENCY ExchangeSegment = "BSE_CURRENCY"
	ExchangeSegmentBSEFNO      ExchangeSegment = "BSE_FNO"
)

type ProductType string

const (
	ProductTypeCNC      ProductType = "CNC"      // Cash & Carry (Equity Delivery)
	ProductTypeINTRADAY ProductType = "INTRADAY" // Intraday (Equity, F&O)
	ProductTypeMARGIN   ProductType = "MARGIN"   // Carry Forward (F&O)
	ProductTypeCO       ProductType = "CO"       // Cover Order (Intraday only)
	ProductTypeBO       ProductType = "BO"       // Bracket Order (Intraday only)
)

type OrderStatus string

const (
	OrderStatusTRANSIT     OrderStatus = "TRANSIT"     // Did not reach the exchange server
	OrderStatusPENDING     OrderStatus = "PENDING"     // Awaiting execution
	OrderStatusCLOSED      OrderStatus = "CLOSED"      // Super Order placed (entry & exit)
	OrderStatusTRIGGERED   OrderStatus = "TRIGGERED"   // Super Order leg triggered
	OrderStatusREJECTED    OrderStatus = "REJECTED"    // Rejected by broker/exchange
	OrderStatusCANCELLED   OrderStatus = "CANCELLED"   // Cancelled by user
	OrderStatusPART_TRADED OrderStatus = "PART_TRADED" // Partial Quantity traded
	OrderStatusTRADED      OrderStatus = "TRADED"      // Executed successfully
)

type AMOTime string

const (
	AMOTimePREOPEN AMOTime = "PRE_OPEN" // Pumped at pre-market session
	AMOTimeOPEN    AMOTime = "OPEN"     // Pumped at market open
	AMOTimeOPEN30  AMOTime = "OPEN_30"  // Pumped 30 minutes after open
	AMOTimeOPEN60  AMOTime = "OPEN_60"  // Pumped 60 minutes after open
)

type InstrumentType string

const (
	InstrumentTypeINDEX  InstrumentType = "INDEX"  // Index
	InstrumentTypeFUTIDX InstrumentType = "FUTIDX" // Futures of Index
	InstrumentTypeOPTIDX InstrumentType = "OPTIDX" // Options of Index
	InstrumentTypeEQUITY InstrumentType = "EQUITY" // Equity
	InstrumentTypeFUTSTK InstrumentType = "FUTSTK" // Futures of Stock
	InstrumentTypeOPTSTK InstrumentType = "OPTSTK" // Options of Stock
	InstrumentTypeFUTCOM InstrumentType = "FUTCOM" // Futures of Commodity
	InstrumentTypeOPTFUT InstrumentType = "OPTFUT" // Options of Commodity Futures
	InstrumentTypeFUTCUR InstrumentType = "FUTCUR" // Futures of Currency
	InstrumentTypeOPTCUR InstrumentType = "OPTCUR" // Options of Currency
)

type Profile struct {
	DhanClientID  string    `json:"dhanClientId"`
	TokenValidity string    `json:"tokenValidity"`
	ActiveSegment string    `json:"activeSegment"`
	Ddpi          string    `json:"ddpi"`
	Mtf           string    `json:"mtf"`
	DataPlan      string    `json:"dataPlan"`
	DataValidity  time.Time `json:"dataValidity"`
}

type BuyRequest struct {
	DhanClientID      string  `json:"dhanClientId"`
	CorrelationID     string  `json:"correlationId"`
	TransactionType   string  `json:"transactionType"`
	ExchangeSegment   string  `json:"exchangeSegment"`
	ProductType       string  `json:"productType"`
	OrderType         string  `json:"orderType"`
	Validity          string  `json:"validity"`
	TradingSymbol     string  `json:"tradingSymbol"`
	SecurityID        string  `json:"securityId"`
	Quantity          int64   `json:"quantity"`
	DisclosedQuantity int64   `json:"disclosedQuantity"`
	Price             float64 `json:"price"`
	TriggerPrice      float64 `json:"triggerPrice"`
	AfterMarketOrder  bool    `json:"afterMarketOrder"`
	AmoTime           string  `json:"amoTime"`
	BoProfitValue     float64 `json:"boProfitValue"`
	BoStopLossValue   float64 `json:"boStopLossValue"`
	DrvExpiryDate     string  `json:"drvExpiryDate"`
	DrvOptionType     string  `json:"drvOptionType"`
	DrvStrikePrice    float64 `json:"drvStrikePrice"`
}

type BuyResponse struct {
	OrderID     string `json:"orderId"`
	OrderStatus string `json:"orderStatus"`
}

type Instruments []Instrument

func (values Instruments) ToMapInstrument() map[string]Instrument {
	mpInstrument := make(map[string]Instrument)
	for _, v := range values {
		mpInstrument[strings.ToUpper(v.CustomSymbol)] = v
	}
	return mpInstrument
}

type Instrument struct {
	ExmExchID          string  `csv:"SEM_EXM_EXCH_ID"`
	Segment            string  `csv:"SEM_SEGMENT"`
	SecurityID         int64   `csv:"SEM_SMST_SECURITY_ID"`
	InstrumentName     string  `csv:"SEM_INSTRUMENT_NAME"`
	ExpiryCode         int64   `csv:"SEM_EXPIRY_CODE"`
	TradingSymbol      string  `csv:"SEM_TRADING_SYMBOL"`
	LotUnits           int64   `csv:"SEM_LOT_UNITS"`
	CustomSymbol       string  `csv:"SEM_CUSTOM_SYMBOL"`
	ExpiryDate         string  `csv:"SEM_EXPIRY_DATE"`
	StrikePrice        float64 `csv:"SEM_STRIKE_PRICE"`
	OptionType         string  `csv:"SEM_OPTION_TYPE"`
	TickSize           float64 `csv:"SEM_TICK_SIZE"`
	ExpiryFlag         string  `csv:"SEM_EXPIRY_FLAG"`
	ExchInstrumentType string  `csv:"SEM_EXCH_INSTRUMENT_TYPE"`
	Series             string  `csv:"SEM_SERIES"`
	SymbolName         string  `csv:"SM_SYMBOL_NAME"`
}

type Postions []Postion

type ProfitAndLoss struct {
	Total     float64
	Manual    float64
	Auto      float64
	Comission float64
}

type Postion struct {
	DhanClientID          string      `json:"dhanClientId"`
	TradingSymbol         string      `json:"tradingSymbol"`
	SecurityID            string      `json:"securityId"`
	PositionType          string      `json:"positionType"`
	ExchangeSegment       string      `json:"exchangeSegment"`
	ProductType           string      `json:"productType"`
	BuyAvg                float64     `json:"buyAvg"`
	BuyQty                int64       `json:"buyQty"`
	CostPrice             float64     `json:"costPrice"`
	SellAvg               float64     `json:"sellAvg"`
	SellQty               int64       `json:"sellQty"`
	NetQty                int64       `json:"netQty"`
	RealizedProfit        float64     `json:"realizedProfit"`
	UnrealizedProfit      float64     `json:"unrealizedProfit"`
	RbiReferenceRate      float64     `json:"rbiReferenceRate"`
	Multiplier            float64     `json:"multiplier"`
	CarryForwardBuyQty    int64       `json:"carryForwardBuyQty"`
	CarryForwardSellQty   int64       `json:"carryForwardSellQty"`
	CarryForwardBuyValue  float64     `json:"carryForwardBuyValue"`
	CarryForwardSellValue float64     `json:"carryForwardSellValue"`
	DayBuyQty             float64     `json:"dayBuyQty"`
	DaySellQty            float64     `json:"daySellQty"`
	DayBuyValue           float64     `json:"dayBuyValue"`
	DaySellValue          float64     `json:"daySellValue"`
	DrvExpiryDate         string      `json:"drvExpiryDate"`
	DrvOptionType         interface{} `json:"drvOptionType"`
	DrvStrikePrice        float64     `json:"drvStrikePrice"`
	CrossCurrency         bool        `json:"crossCurrency"`
}

type Trades []Trade

type Trade struct {
	DhanClientID    string      `json:"dhanClientId"`
	OrderID         string      `json:"orderId"`
	ExchangeOrderID string      `json:"exchangeOrderId"`
	ExchangeTradeID string      `json:"exchangeTradeId"`
	TransactionType string      `json:"transactionType"`
	ExchangeSegment string      `json:"exchangeSegment"`
	ProductType     string      `json:"productType"`
	OrderType       string      `json:"orderType"`
	TradingSymbol   string      `json:"tradingSymbol"`
	SecurityID      string      `json:"securityId"`
	TradedQuantity  int64       `json:"tradedQuantity"`
	TradedPrice     float64     `json:"tradedPrice"`
	CreateTime      string      `json:"createTime"`
	UpdateTime      string      `json:"updateTime"`
	ExchangeTime    string      `json:"exchangeTime"`
	DrvExpiryDate   interface{} `json:"drvExpiryDate"`
	DrvOptionType   interface{} `json:"drvOptionType"`
	DrvStrikePrice  float64     `json:"drvStrikePrice"`
}

type Orders []Order

func (or Orders) Total() int {
	total := 0
	for _, v := range or {
		if strings.ToLower(v.OrderStatus) == "traded" {
			total += 1
		}
	}
	return total
}

type Order struct {
	DhanClientID        string      `json:"dhanClientId"`
	OrderID             string      `json:"orderId"`
	CorrelationID       string      `json:"correlationId"`
	OrderStatus         string      `json:"orderStatus"`
	TransactionType     string      `json:"transactionType"`
	ExchangeSegment     string      `json:"exchangeSegment"`
	ProductType         string      `json:"productType"`
	OrderType           string      `json:"orderType"`
	Validity            string      `json:"validity"`
	TradingSymbol       string      `json:"tradingSymbol"`
	SecurityID          string      `json:"securityId"`
	Quantity            int64       `json:"quantity"`
	DisclosedQuantity   int64       `json:"disclosedQuantity"`
	Price               float64     `json:"price"`
	TriggerPrice        float64     `json:"triggerPrice"`
	AfterMarketOrder    bool        `json:"afterMarketOrder"`
	BoProfitValue       float64     `json:"boProfitValue"`
	BoStopLossValue     float64     `json:"boStopLossValue"`
	LegName             string      `json:"legName"`
	CreateTime          string      `json:"createTime"`
	UpdateTime          string      `json:"updateTime"`
	ExchangeTime        string      `json:"exchangeTime"`
	DrvExpiryDate       interface{} `json:"drvExpiryDate"`
	DrvOptionType       interface{} `json:"drvOptionType"`
	DrvStrikePrice      float64     `json:"drvStrikePrice"`
	OmsErrorCode        interface{} `json:"omsErrorCode"`
	OmsErrorDescription interface{} `json:"omsErrorDescription"`
	FilledQty           int64       `json:"filled_qty"`
	AlgoID              string      `json:"algoId"`
}
