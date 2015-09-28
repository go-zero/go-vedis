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
	if ok, err := suite.store.Set("name", "John"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.True(ok)
	}

	if value, err := suite.store.Get("name"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.Equal("John", value)
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

	if count, err := suite.store.Del("nothing"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.Equal(0, count)
	}
}

func (suite *VedisTestSuite) TestAppend() {
	if count, err := suite.store.Append("message", "hello"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.Equal(len("hello"), count)
	}

	if count, err := suite.store.Append("message", " world"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.Equal(len("hello world"), count)
	}

	if value, err := suite.store.Get("message"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.Equal("hello world", value)
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
	if ok, err := suite.store.Set("message", "hello"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.True(ok)
	}

	if ok, err := suite.store.Copy("message", "backup"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.True(ok)
	}

	if count, err := suite.store.Append("message", " world"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.Equal(len("hello world"), count)
	}

	if value, err := suite.store.Get("message"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.Equal("hello world", value)
	}

	if value, err := suite.store.Get("backup"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.Equal("hello", value)
	}
}

func (suite *VedisTestSuite) TestMove() {
	if ok, err := suite.store.Set("name", "TangZero"); err != nil {
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
		suite.Equal("TangZero", value)
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

func (suite *VedisTestSuite) TestIncr() {
	if ok, err := suite.store.Set("count", "12"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.True(ok)
	}

	if value, err := suite.store.Incr("count"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.Equal(13, value)
	}
}

func (suite *VedisTestSuite) TestIncrBy() {
	if ok, err := suite.store.Set("count", "12"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.True(ok)
	}

	if value, err := suite.store.IncrBy("count", 10); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.Equal(22, value)
	}
}

func (suite *VedisTestSuite) TestDecr() {
	if ok, err := suite.store.Set("count", "15"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.True(ok)
	}

	if value, err := suite.store.Decr("count"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.Equal(14, value)
	}
}

func (suite *VedisTestSuite) TestDecrBy() {
	if ok, err := suite.store.Set("count", "23"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.True(ok)
	}

	if value, err := suite.store.DecrBy("count", 3); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.Equal(20, value)
	}
}

func (suite *VedisTestSuite) TestHSetHGet() {
	if ok, err := suite.store.HSet("config", "url", "github.com"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.True(ok)
	}

	if value, err := suite.store.HGet("config", "url"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.Equal("github.com", value)
	}
}

func (suite *VedisTestSuite) TestHDel() {
	if ok, err := suite.store.HSet("config", "url", "github.com"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.True(ok)
	}

	if count, err := suite.store.HDel("config", "url"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.Equal(1, count)
	}

	if count, err := suite.store.HDel("config", "timeout"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.Equal(0, count)
	}
}

func (suite *VedisTestSuite) TestLen() {
	if ok, err := suite.store.HSet("config", "url", "github.com"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.True(ok)
	}

	if count, err := suite.store.HLen("config"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.Equal(1, count)
	}

	if count, err := suite.store.HDel("nothing"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.Equal(0, count)
	}
}

func (suite *VedisTestSuite) TestHExists() {
	if ok, err := suite.store.HSet("config", "url", "github.com"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.True(ok)
	}

	if exists, err := suite.store.HExists("config", "url"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.True(exists)
	}

	if exists, err := suite.store.HExists("config", "timeout"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.False(exists)
	}
}

func (suite *VedisTestSuite) TestHKeys() {
	if count, err := suite.store.HMSet("config", "url", "github.com", "timeout", "500"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.Equal(2, count)
	}

	if keys, err := suite.store.HKeys("config"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.Equal([]string{"url", "timeout"}, keys)
	}
}

func (suite *VedisTestSuite) TestHVals() {
	if count, err := suite.store.HMSet("config", "url", "github.com", "timeout", "500"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.Equal(2, count)
	}

	if values, err := suite.store.HVals("config"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.Equal([]string{"github.com", "500"}, values)
	}
}

func (suite *VedisTestSuite) TestMSetMGet() {
	if count, err := suite.store.HMSet("config", "url", "github.com", "timeout", "500", "retries", "3"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.Equal(3, count)
	}

	if values, err := suite.store.HMGet("config", "url", "retries"); err != nil {
		suite.Fail(err.Error())
	} else {
		suite.Equal([]string{"github.com", "3"}, values)
	}
}

func TestVedisTestSuite(t *testing.T) {
	suite.Run(t, new(VedisTestSuite))
}
