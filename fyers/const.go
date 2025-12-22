package fyers

const (
	ApplicationJson = "application/json"
)

const (
	APIURL                  = "https://api-t1.fyers.in"
	profileURL              = APIURL + "/api/v3/profile"
	validateAuthCodeURL     = APIURL + "/api/v3/validate-authcode"
	validateRefreshTokenURL = APIURL + "/api/v3/validate-refresh-token"
	holdingURL              = APIURL + "/api/v3/holdings"
	logoutURL               = APIURL + "/api/v3/logout"
	fundURL                 = APIURL + "/api/v3/funds"
	orderURL                = APIURL + "/api/v3/orders"
	positionURL             = APIURL + "/api/v3/positions"
	tradeURL                = APIURL + "/api/v3/trades"
	generateAuthCodeURL     = APIURL + "/api/v3/generate-authcode"
)

var (
	instrumentList = []string{"NSE_FO", "NSE_CM", "BSE_FO"}
)
