package fyers

import (
	"encoding/json"
	"errors"
	"fmt"
	"maps"
	"sync"

	"resty.dev/v3"
)

type Options struct {
	request *resty.Request
}

func NewFyers() *Options {
	return &Options{
		request: resty.New().R(),
	}
}

func (o *Options) ValidateAuthCode(appID, appSecret, code string) (*AuthResponse, error) {
	authReq := &AuthRequest{
		GrantType: "authorization_code",
		Code:      code,
	}
	authReq = authReq.WithAppIDHash(appID, appSecret)
	res := &AuthResponse{}
	resp, err := o.request.
		SetResult(res).
		SetHeader("Content-Type", ApplicationJson).
		SetForceResponseContentType(ApplicationJson).
		SetBody(authReq).
		Post(validateAuthCodeURL)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, errors.New("failed to authenticate: " + resp.String())
	}
	return res, nil
}
func (o *Options) RefreshToken(appID, appSecret, refreshToken, pin string) (*RefreshTokenResponse, error) {
	refreshReq := &RefreshTokenRequest{
		GrantType:    "refresh_token",
		RefreshToken: refreshToken,
		Pin:          pin,
	}
	refreshReq.AppIDHash = Sha256Hash(appID + appSecret)
	res := &RefreshTokenResponse{}
	resp, err := o.request.
		SetResult(res).
		SetHeader("Content-Type", ApplicationJson).
		SetForceResponseContentType(ApplicationJson).
		SetBody(refreshReq).
		Post(validateRefreshTokenURL)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, errors.New("failed to refresh token: " + resp.String())
	}
	return res, nil
}
func (o *Options) Profile(accessToken string) (*ProfileRes, error) {
	res := &ProfileRes{}
	resp, err := o.request.
		SetResult(res).
		SetForceResponseContentType(ApplicationJson).
		SetHeader("Authorization", accessToken).
		Get(profileURL)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, errors.New("failed to fetch profile: " + resp.String())
	}
	return res, nil
}

func (o *Options) Funds(accessToken string) (*FundResponse, error) {
	res := &FundResponse{}
	resp, err := o.request.
		SetResult(res).
		SetForceResponseContentType(ApplicationJson).
		SetHeader("Authorization", accessToken).
		Get(fundURL)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, errors.New("failed to fetch fund: " + resp.String())
	}
	return res, nil
}

func (o *Options) Holdings(accessToken string) (*HoldingResponse, error) {
	res := &HoldingResponse{}
	resp, err := o.request.
		SetResult(res).
		SetForceResponseContentType(ApplicationJson).
		SetHeader("Authorization", accessToken).
		Get(holdingURL)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, errors.New("failed to fetch holdings: " + resp.String())
	}
	return res, nil
}

func (o *Options) Logout(accessToken string) (*CommonResponse, error) {
	res := &CommonResponse{}
	resp, err := o.request.
		SetResult(res).
		SetForceResponseContentType(ApplicationJson).
		SetHeader("Authorization", accessToken).
		Post(logoutURL)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, errors.New("failed to post logout: " + resp.String())
	}
	return res, nil
}

func (o *Options) Orders(accessToken string) (*OrdersResponse, error) {
	res := &OrdersResponse{}
	resp, err := o.request.
		SetResult(res).
		SetForceResponseContentType(ApplicationJson).
		SetHeader("Authorization", accessToken).
		Get(orderURL)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, errors.New("failed to fetch orders: " + resp.String())
	}
	return res, nil
}

func (o *Options) OrdersByID(accessToken, orderID string) (*OrdersResponse, error) {
	res := &OrdersResponse{}
	resp, err := o.request.
		SetResult(res).
		SetForceResponseContentType(ApplicationJson).
		SetHeader("Authorization", accessToken).
		Get(orderURL + "?id=" + orderID)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, errors.New("failed to fetch order by id: " + resp.String())
	}
	return res, nil
}

func (o *Options) Positions(accessToken string) (*PositionsResponse, error) {
	res := &PositionsResponse{}
	resp, err := o.request.
		SetResult(res).
		SetForceResponseContentType(ApplicationJson).
		SetHeader("Authorization", accessToken).
		Get(positionURL)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, errors.New("failed to fetch positions: " + resp.String())
	}
	return res, nil
}

func (o *Options) Trades(accessToken string) (*TradesResponse, error) {
	res := &TradesResponse{}
	resp, err := o.request.
		SetResult(res).
		SetForceResponseContentType(ApplicationJson).
		SetHeader("Authorization", accessToken).
		Get(tradeURL)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, errors.New("failed to fetch trade: " + resp.String())
	}
	return res, nil
}

// GetInstruments implements IFyers.
func (o *Options) GetInstruments() (Instruments, error) {
	res := Instruments{}
	errs := make([]error, 0)
	var wg sync.WaitGroup
	var mutex sync.Mutex
	for _, instrument := range instrumentList {
		wg.Add(1)
		go func() {
			defer wg.Done()
			url := fmt.Sprintf("https://public.fyers.in/sym_details/%v_sym_master.json", instrument)
			tmpRes := Instruments{}
			rs, err := o.request.Get(url)
			if err != nil {
				errs = append(errs, fmt.Errorf("error loading symbol %s: %w", instrument, err))
			}
			if err := json.Unmarshal(rs.Bytes(), &tmpRes); err != nil {
				errs = append(errs, fmt.Errorf("error unmarshalling response for %s: %w", instrument, err))
			}
			mutex.Lock()
			maps.Copy(res, tmpRes)
			mutex.Unlock()
		}()
	}
	wg.Wait()
	return res, nil
}

func (o *Options) LoginLink(apiKey, redirectURI string) (string, error) {
	if apiKey == "" || redirectURI == "" {
		return "", errors.New("apiKey and redirectURI must not be empty")
	}
	loginLink := fmt.Sprintf("%v?client_id=%v&redirect_uri=%s&response_type=code&state=fyers", generateAuthCodeURL, apiKey, redirectURI)
	return loginLink, nil
}
