package dhan

import (
	"bytes"
	"encoding/csv"
	"errors"
	"net/http"

	"github.com/gocarina/gocsv"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"resty.dev/v3"
)

type Client struct {
	resty *resty.Client
}

// --- Client Constructor ---

func NewClient(baseURL string) *Client {
	client := resty.New().SetBaseURL(baseURL)
	return &Client{resty: client}
}

// Profile implements IDhan.
func (d *Client) Profile(accessToken string) (Profile, error) {
	res := Profile{}
	r, err := d.resty.R().
		SetHeader(accept, contentTypeJSON).
		SetHeader(authorization, accessToken).
		SetResult(&res).
		Get(ordersURL)
	if err != nil {
		return res, err
	}
	if r.IsError() || r.StatusCode() != http.StatusOK {
		return res, errors.New("failed to fetch profile: " + string(r.Bytes()))
	}
	return res, nil
}

// Trades implements IDhan.
func (d *Client) Trades(accessToken string) (Trades, error) {
	res := Trades{}
	r, err := d.resty.R().
		SetHeader(accept, contentTypeJSON).
		SetHeader(authorization, accessToken).
		SetResult(&res).
		Get(tradesURL)
	if err != nil {
		return res, err
	}
	if r.IsError() || r.StatusCode() != http.StatusOK {
		return res, errors.New("failed to fetch trades: " + string(r.Bytes()))
	}
	return res, nil
}

// Positions implements IDhan.
func (d *Client) Positions(accessToken string) (Postions, error) {
	res := Postions{}
	r, err := d.resty.R().
		SetHeader(accept, contentTypeJSON).
		SetHeader(authorization, accessToken).
		SetResult(&res).
		Get(postionsURL)
	if err != nil {
		return res, err
	}
	if r.IsError() || r.StatusCode() != http.StatusOK {
		return res, errors.New("failed to fetch positions: " + string(r.Bytes()))
	}
	return res, err
}

// GetInstruments implements IDhan.
func (d *Client) GetInstruments() (Instruments, error) {
	res := Instruments{}
	rs, err := d.resty.R().Get(instrumentURL)
	if err != nil {
		return res, err
	}
	in := csv.NewReader(transform.NewReader(bytes.NewReader(rs.Bytes()), unicode.UTF8BOM.NewDecoder()))
	if err := gocsv.UnmarshalCSV(in, &res); err != nil {
		return nil, err
	}
	return res, nil
}

// Buy implements IDhan.
func (d *Client) Buy(token string, req BuyRequest) (BuyResponse, error) {
	res := BuyResponse{}
	r, err := d.resty.R().
		SetHeader(accept, contentTypeJSON).
		SetHeader(contentType, contentTypeJSON).
		SetHeader(authorization, token).
		SetBody(&req).
		SetResult(&res).
		Post(ordersURL)
	if err != nil {
		return res, err
	}
	if r.IsError() || r.StatusCode() != http.StatusOK {
		return res, errors.New("failed to fetch positions: " + string(r.Bytes()))
	}
	return res, err
}
