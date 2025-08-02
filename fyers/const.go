package fyers

const (
	ApplicationJson = "application/json"
)

const (
	APIURL                  = "https://api-t1.fyers.in/api/v3"
	profileURL              = APIURL + "/profile"
	validateAuthCodeURL     = APIURL + "/validate-authcode"
	validateRefreshTokenURL = APIURL + "/validate-refresh-token"
	holdingURL              = APIURL + "/holdings"
	logoutURL               = APIURL + "/logout"
	fundURL                 = APIURL + "/funds"
	orderURL                = APIURL + "/orders"
	positionURL             = APIURL + "/positions"
	tradeURL                = APIURL + "/trades"
	generateAuthCodeURL     = APIURL + "/generate-authcode"
)

var (
	instrumentList = []string{"NSE_FO", "NSE_CM", "BSE_FO"}
)
