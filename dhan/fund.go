package dhan

import (
	"errors"
	"net/http"
)

type FundResponse struct {
	DhanClientID        string  `json:"dhanClientId"`
	AvailabelBalance    float64 `json:"availabelBalance"`
	SodLimit            float64 `json:"sodLimit"`
	CollateralAmount    float64 `json:"collateralAmount"`
	ReceiveableAmount   float64 `json:"receiveableAmount"`
	UtilizedAmount      float64 `json:"utilizedAmount"`
	BlockedPayoutAmount float64 `json:"blockedPayoutAmount"`
	WithdrawableBalance float64 `json:"withdrawableBalance"`
}

func (fs FundResponse) EquityAmount() float64 {
	return fs.AvailabelBalance
}

// Fund implements IDhan.
func (d *Client) Fund(accessToken string) (FundResponse, error) {
	res := FundResponse{}
	r, err := d.resty.R().
		SetHeader(accept, contentTypeJSON).
		SetHeader(authorization, accessToken).
		SetResult(&res).
		Get(fundURL)
	if err != nil {
		return res, err
	}
	if r.IsError() || r.StatusCode() != http.StatusOK {
		return res, errors.New("failed to fetch fund: " + string(r.Bytes()))
	}
	return res, err
}
