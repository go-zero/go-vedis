package vedis

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type VedisTestSuite struct {
	suite.Suite
	store *Vedis
}

func (suite *VedisTestSuite) SetupTest() {
	suite.store = New()
	if ok, err := suite.store.Open(); !ok {
		suite.Fail(err.Error())
	}
}

func (suite *VedisTestSuite) TearDownTest() {
	suite.store.Close()
}

func (suite *VedisTestSuite) TestSetAndGet() {
	name := "John"

	if ok, err := suite.store.Set("name", name); !ok {
		suite.Fail(err.Error())
	}

	if value, err := suite.store.Get("name"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.Equal(name, value)
	}
}

func (suite *VedisTestSuite) TestDel() {
	if ok, err := suite.store.Set("foo", "bar"); !ok {
		suite.Fail(err.Error())
	}

	if count, err := suite.store.Del("foo"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.Equal(1, count)
	}
}

func (suite *VedisTestSuite) TestAppend() {
	hello := "hello"
	world := " world"

	if count, err := suite.store.Append("message", hello); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.Equal(len(hello), count)
	}

	if count, err := suite.store.Append("message", world); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.Equal(len(hello+world), count)
	}

	if value, err := suite.store.Get("message"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.Equal(hello+world, value)
	}
}

func (suite *VedisTestSuite) TestExists() {
	if ok, err := suite.store.Set("foo", "bar"); !ok {
		suite.Fail(err.Error())
	}

	if exists, err := suite.store.Exists("foo"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.True(exists)
	}

	if exists, err := suite.store.Exists("nothing"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.False(exists)
	}
}

func (suite *VedisTestSuite) TestCopy() {
	hello := "hello"
	world := " world"

	if ok, err := suite.store.Set("message", hello); !ok {
		suite.Fail(err.Error())
	}

	if ok, err := suite.store.Copy("message", "backup"); !ok {
		suite.Fail(err.Error())
	}

	if count, err := suite.store.Append("message", world); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.Equal(len(hello+world), count)
	}

	if value, err := suite.store.Get("message"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.Equal(hello+world, value)
	}

	if value, err := suite.store.Get("backup"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.Equal(hello, value)
	}
}

func TestVedisTestSuite(t *testing.T) {
	suite.Run(t, new(VedisTestSuite))
}
