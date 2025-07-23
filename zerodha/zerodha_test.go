package zerodha

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"resty.dev/v3"
)

type ZerodhaTestSuite struct {
	suite.Suite
	zInstance *Zerodha
}

// this function executes before the test suite begins execution
func (suite *ZerodhaTestSuite) SetupSuite() {
	suite.zInstance = NewZerodha(resty.New())
}

// before each test
func (suite *ZerodhaTestSuite) SetupTest() {
}

func (suite *ZerodhaTestSuite) TestLoadInstrument() {
	suite.Run("load instrument success", func() {
		instruments, err := suite.zInstance.LoadInstrument()
		suite.Require().Nil(err)
		suite.Require().NotNil(instruments)
		suite.Require().GreaterOrEqual(len(instruments), 1)
	})
}

// this function executes after all tests executed
func (suite *ZerodhaTestSuite) TearDownSuite() {
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ZerodhaTestSuite))
}
