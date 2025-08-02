package dhan

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type DhanSuite struct {
	suite.Suite
	dhan        *Client
	accessToken string
	clientID    string
}

func (suite *DhanSuite) SetupTest() {
	suite.accessToken = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiJ9.eyJpc3MiOiJkaGFuIiwicGFydG5lcklkIjoiIiwiZXhwIjoxNzU2NzA4MzA3LCJ0b2tlbkNvbnN1bWVyVHlwZSI6IlNFTEYiLCJ3ZWJob29rVXJsIjoiIiwiZGhhbkNsaWVudElkIjoiMTEwMDY0MjM2NiJ9.pT0wPbeWKtiGQQTbE2b5BClmmE28H876MmlIXSbK-u_aa02UVyBOz57js4Ew7Z4Gf7fDVFstO1d5SKtxoUso8Q"
	suite.clientID = "1100642366"
	suite.dhan = NewClient(APIURL)
}

func (suite *DhanSuite) TestGetInstruments() {
	suite.T().Run("GetInstruments", func(t *testing.T) {
		instruments, err := suite.dhan.GetInstruments()
		suite.NoError(err, "Expected no error when fetching instruments")
		suite.NotEmpty(instruments, "Expected instruments to be non-empty")
	})
}

func (suite *DhanSuite) TestGetLTP() {
	suite.T().Run("GetLTP", func(t *testing.T) {
		req := LTPRequest{}
		req["NSE_EQ"] = []int64{1048706} // Example security ID
		res, err := suite.dhan.GetLTP(suite.accessToken, suite.clientID, req)
		suite.NoError(err, "Expected no error when fetching LTP")
		suite.NotNil(res, "Expected LTP response to be non-nil")
	})
}

func (suite *DhanSuite) TestGetIntradayOHLC() {
	suite.T().Run("GetIntradayOHLC", func(t *testing.T) {
		req := IntradayChartRequest{
			SecurityID:      "1333",
			ExchangeSegment: ExchangeSegmentNSEEQ,
			Instrument:      InstrumentTypeEQUITY,
			Interval:        Interval1Min,
			Oi:              true,
			FromDate:        "2025-08-01 09:30:00",
			ToDate:          "2025-08-10 09:30:00",
		}
		// req["NSE_EQ"] = []int64{1048706} // Example security ID
		res, err := suite.dhan.GetIntradayOHLC(suite.accessToken, req)
		suite.NoError(err)
		suite.NotNil(res)
	})
}

func (suite *DhanSuite) TestGetHistoricalOHLC() {
	suite.T().Run("GetHistoricalOHLC", func(t *testing.T) {
		req := HistoricalChartRequest{
			SecurityID:      "1333",
			ExchangeSegment: ExchangeSegmentNSEEQ,
			Oi:              true,
			FromDate:        "2025-08-01",
			ToDate:          "2025-08-10",
			Instrument:      InstrumentTypeEQUITY,
			ExpiryCode:      0,
		}
		res, err := suite.dhan.GetHistoricalOHLC(suite.accessToken, req)
		suite.NoError(err)
		suite.NotNil(res)
	})
}

func (suite *DhanSuite) TestGetOptionChain() {
	suite.T().Run("GetOptionChain", func(t *testing.T) {

		res, err := suite.dhan.GetOptionChain(suite.accessToken, suite.clientID, OptionChainRequest{
			UnderlyingScrip: 13,
			UnderlyingSeg:   ExchangeSegmentNSEEQ,
			Expiry:          "2025-08-28",
		})
		suite.NoError(err)
		suite.NotNil(res)
	})
}

func (suite *DhanSuite) TestGetOptionExpiryList() {
	suite.T().Run("GetOptionExpiryList", func(t *testing.T) {
		res, err := suite.dhan.GetOptionExpiryList(suite.accessToken, suite.clientID, OptionExpiryRequest{
			UnderlyingScrip: 13,
			UnderlyingSeg:   ExchangeSegmentNSEEQ,
		})
		suite.NoError(err)
		suite.NotNil(res)
	})
}
func TestFyersSuite(t *testing.T) {
	suite.Run(t, new(DhanSuite))
}
