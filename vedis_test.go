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
	if ok, err := suite.store.Open(); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.True(ok)
	}
}

func (suite *VedisTestSuite) TearDownTest() {
	suite.store.Close()
}

func (suite *VedisTestSuite) TestSetAndGet() {
	name := "John"

	if ok, err := suite.store.Set("name", name); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.True(ok)
	}

	if value, err := suite.store.Get("name"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.Equal(name, value)
	}
}

func (suite *VedisTestSuite) TestDel() {
	if ok, err := suite.store.Set("foo", "bar"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.True(ok)
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
	if ok, err := suite.store.Set("foo", "bar"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.True(ok)
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

	if ok, err := suite.store.Set("message", hello); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.True(ok)
	}

	if ok, err := suite.store.Copy("message", "backup"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.True(ok)
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

func (suite *VedisTestSuite) TestMove() {
	name := "TangZero"

	if ok, err := suite.store.Set("name", name); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.True(ok)
	}

	if ok, err := suite.store.Move("name", "nickname"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.True(ok)
	}

	if exists, err := suite.store.Exists("name"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.False(exists)
	}

	if value, err := suite.store.Get("nickname"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.Equal(name, value)
	}
}

func (suite *VedisTestSuite) TestMassiveSetAndMassiveGet() {
	if ok, err := suite.store.MSet("name", "John", "age", "29"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.True(ok)
	}

	if values, err := suite.store.MGet("name", "age", "email"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.Equal([]string{"John", "29", ""}, values)
	}
}

func (suite *VedisTestSuite) TestSetNX() {
	if ok, err := suite.store.SetNX("name", "John"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.True(ok)
	}

	if ok, err := suite.store.SetNX("name", "Smith"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.False(ok)
	}

	if ok, err := suite.store.SetNX("age", "25"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.True(ok)
	}

	if values, err := suite.store.MGet("name", "age"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.Equal([]string{"John", "25"}, values)
	}
}

func (suite *VedisTestSuite) TestMSetNX() {
	if ok, err := suite.store.MSetNX("name", "John", "age", "29"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.True(ok)
	}

	if ok, err := suite.store.MSetNX("name", "Smith", "email", "smith@gmail.com"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.True(ok)
	}

	if values, err := suite.store.MGet("name", "age", "email"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.Equal([]string{"John", "29", "smith@gmail.com"}, values)
	}
}

func (suite *VedisTestSuite) TestGetSet() {
	if ok, err := suite.store.Set("message", "Foo"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.True(ok)
	}

	if value, err := suite.store.GetSet("message", "Bar"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.Equal("Foo", value)
	}

	if value, err := suite.store.Get("message"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.Equal("Bar", value)
	}
}

func TestVedisTestSuite(t *testing.T) {
	suite.Run(t, new(VedisTestSuite))
}
