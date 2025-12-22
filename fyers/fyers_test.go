package fyers

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type FyersSuite struct {
	suite.Suite
	fyers *FyersClient
}

func (suite *FyersSuite) SetupTest() {
	suite.fyers = NewFyers()
}

func (suite *FyersSuite) TestGetInstruments() {
	suite.Run("GetInstruments", func() {
		instruments, err := suite.fyers.GetInstruments()
		suite.NoError(err, "Expected no error when fetching instruments")
		suite.NotEmpty(instruments, "Expected instruments to be non-empty")

		instruments = instruments.GetMapInstrumentByUnderSymbol()
		suite.NotEmpty(instruments, "Expected instruments map to be non-empty")
		suite.GreaterOrEqual(len(instruments), 1, "Expected instruments map to have one entry")
	})
}

func (suite *FyersSuite) TestLoginLink() {
	suite.T().Run("LoginLink", func(t *testing.T) {
		link, err := suite.fyers.LoginLink("xp16278", "http://localhost:8080/callback")
		suite.NoError(err, "Expected no error when generating login link")
		suite.NotEmpty(link, "Expected login link to be non-empty")
	})
}
func TestFyersSuite(t *testing.T) {
	suite.Run(t, new(FyersSuite))
}
