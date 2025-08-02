package fyers

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"resty.dev/v3"
)

type FyersSuite struct {
	suite.Suite
	fyers *Options
}

func (suite *FyersSuite) SetupTest() {
	suite.fyers = NewFyers(resty.New())
}

func (suite *FyersSuite) TestGetInstruments() {
	suite.T().Run("GetInstruments", func(t *testing.T) {
		instruments, err := suite.fyers.GetInstruments()
		suite.NoError(err, "Expected no error when fetching instruments")
		suite.NotEmpty(instruments, "Expected instruments to be non-empty")
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
