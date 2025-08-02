package dhan

const (
	instrumentURL = "https://images.dhan.co/api-data/api-scrip-master.csv"

	APIURL           = "https://api.dhan.co/v2"
	ordersURL        = "/orders"
	fundURL          = "/fundlimit"
	postionsURL      = "/positions"
	tradesURL        = "/trades"
	profileURL       = "/profile"
	sliceOrderURL    = "/orders/slicing"
	orderExternalURL = "/orders/external"

	superOrdersURL   = "/super/orders"
	superOrderIDURL  = "/super/orders/%s"    // for PUT
	superOrderLegURL = "/super/orders/%s/%s" // for DELETE

	marketFeedLTPURL   = "/marketfeed/ltp"
	marketFeedOHLCURL  = "/marketfeed/ohlc"
	marketFeedQuoteURL = "/marketfeed/quote"

	chartsHistoricalURL = "/charts/historical"
	chartsIntradayURL   = "/charts/intraday"

	optionChainURL       = "/optionchain"
	optionChainExpiryURL = "/optionchain/expirylist"
)

const (
	contentType                = "content-type"
	accept                     = "accept"
	contentTypeJSON            = "application/json"
	contentTypeForm            = "application/x-www-form-urlencoded"
	authorization              = "access-token"
	grantTypeAuthorizationCode = "authorization_code"
	clientIDHeader             = "client-id"
)
