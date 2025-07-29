package fyers

// ProductType represents different product types for trading
type ProductType string

const (
	ProductCNC      ProductType = "CNC"      // For equity only
	ProductIntraday ProductType = "INTRADAY" // Applicable for all segments
	ProductMargin   ProductType = "MARGIN"   // Applicable only for derivatives
	ProductCO       ProductType = "CO"       // Cover Order
	ProductBO       ProductType = "BO"       // Bracket Order
	ProductMTF      ProductType = "MTF"      // Margin Trading Facility
)

// OrderType represents different types of orders in trading
type OrderType int

const (
	OrderLimit     OrderType = 1 // Limit order
	OrderMarket    OrderType = 2 // Market order
	OrderStop      OrderType = 3 // Stop order (SL-M)
	OrderStopLimit OrderType = 4 // Stoplimit order (SL-L)
)

// OrderStatus represents the current status of an order
type OrderStatus int

const (
	OrderCancelled OrderStatus = 1 // Cancelled
	OrderFilled    OrderStatus = 2 // Traded / Filled
	OrderReserved  OrderStatus = 3 // For future use
	OrderTransit   OrderStatus = 4 // Transit
	OrderRejected  OrderStatus = 5 // Rejected
	OrderPending   OrderStatus = 6 // Pending
)

// OrderSide represents whether the order is a buy or sell
type OrderSide int

const (
	OrderBuy  OrderSide = 1  // Buy
	OrderSell OrderSide = -1 // Sell
)

// PositionSide represents the direction of a position
type PositionSide int

const (
	PositionLong   PositionSide = 1  // Long
	PositionShort  PositionSide = -1 // Short
	PositionClosed PositionSide = 0  // Closed position
)

// HoldingType represents the type of holding based on share settlement status
type HoldingType string

const (
	HoldingT1  HoldingType = "T1"  // Shares purchased but not yet delivered to demat
	HoldingHLD HoldingType = "HLD" // Shares purchased and available in demat
)

// OrderSource represents the platform or origin of the order
type OrderSource string

const (
	OrderSourceMobile OrderSource = "M"   // Mobile
	OrderSourceWeb    OrderSource = "W"   // Web
	OrderSourceFyers  OrderSource = "R"   // Fyers One
	OrderSourceAdmin  OrderSource = "A"   // Admin
	OrderSourceAPI    OrderSource = "ITS" // API
)

// InstrumentType represents different types of financial instruments across segments
type InstrumentType int

const (
	// CM Segment
	InstrumentEQ         InstrumentType = 0  // EQ (EQUITY)
	InstrumentPrefShares InstrumentType = 1  // PREFSHARES
	InstrumentDebentures InstrumentType = 2  // DEBENTURES
	InstrumentWarrants   InstrumentType = 3  // WARRANTS
	InstrumentMisc       InstrumentType = 4  // MISC (NSE, BSE)
	InstrumentSGB        InstrumentType = 5  // SGB
	InstrumentGSecs      InstrumentType = 6  // G - Secs
	InstrumentTBills     InstrumentType = 7  // T - Bills
	InstrumentMF         InstrumentType = 8  // MF
	InstrumentETF        InstrumentType = 9  // ETF
	InstrumentIndex      InstrumentType = 10 // INDEX
	InstrumentMiscBSE    InstrumentType = 50 // MISC (BSE)

	// FO Segment
	InstrumentFutIdx InstrumentType = 11 // FUTIDX
	InstrumentFutIVX InstrumentType = 12 // FUTIVX
	InstrumentFutStk InstrumentType = 13 // FUTSTK
	InstrumentOptIdx InstrumentType = 14 // OPTIDX
	InstrumentOptStk InstrumentType = 15 // OPTSTK

	// CD Segment
	InstrumentFutCur  InstrumentType = 16 // FUTCUR
	InstrumentFutIRT  InstrumentType = 17 // FUTIRT
	InstrumentFutIRC  InstrumentType = 18 // FUTIRC
	InstrumentOptCur  InstrumentType = 19 // OPTCUR
	InstrumentUndCur  InstrumentType = 20 // UNDCUR
	InstrumentUndIRC  InstrumentType = 21 // UNDIRC
	InstrumentUndIRT  InstrumentType = 22 // UNDIRT
	InstrumentUndIRD  InstrumentType = 23 // UNDIRD
	InstrumentIndexCD InstrumentType = 24 // INDEX_CD
	InstrumentFutIRD  InstrumentType = 25 // FUTIRD

	// COM Segment
	InstrumentFutCom     InstrumentType = 30 // FUTCOM
	InstrumentOptFut     InstrumentType = 31 // OPTFUT
	InstrumentOptCom     InstrumentType = 32 // OPTCOM
	InstrumentFutBAS     InstrumentType = 33 // FUTBAS
	InstrumentFutBLN     InstrumentType = 34 // FUTBLN
	InstrumentFutENR     InstrumentType = 35 // FUTENR
	InstrumentOptBLN     InstrumentType = 36 // OPTBLN
	InstrumentOptFutNCOM InstrumentType = 37 // OPTFUT (NCOM)
)

// Segment represents different trading segments in the market
type Segment int

const (
	SegmentCapitalMarket        Segment = 10 // Capital Market
	SegmentEquityDerivatives    Segment = 11 // Equity Derivatives
	SegmentCurrencyDerivatives  Segment = 12 // Currency Derivatives
	SegmentCommodityDerivatives Segment = 20 // Commodity Derivatives
)

// Exchange represents different stock or commodity exchanges
type Exchange int

const (
	ExchangeNSE Exchange = 10 // NSE (National Stock Exchange)
	ExchangeMCX Exchange = 11 // MCX (Multi Commodity Exchange)
	ExchangeBSE Exchange = 12 // BSE (Bombay Stock Exchange)
)

type OrderValidity string

const (
	OrderValidityIntraday OrderValidity = "DAY" // Intraday order, valid for the day
	OrderValidityIOC      OrderValidity = "IOC" // Immediate or Cancel order
)
