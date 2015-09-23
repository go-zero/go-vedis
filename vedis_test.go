package vedis

import (
	assert "github.com/stretchr/testify/require"
	"testing"
)

func NewServer(t *testing.T) *Vedis {
	server := New()
	if ok, err := server.Open(); !ok {
		assert.FailNow(t, err.Error())
	}
	return server
}

func TestSetAndGet(t *testing.T) {
	server := NewServer(t)
	defer server.Close()

	name := "John"

	if ok, err := server.Set("name", name); !ok {
		assert.FailNow(t, err.Error())
	}

	if value, err := server.Get("name"); err != nil {
		assert.FailNow(t, err.Error())
	} else {
		assert.Equal(t, name, value)
	}
}

func TestDel(t *testing.T) {
	server := NewServer(t)
	defer server.Close()

	if ok, err := server.Set("foo", "bar"); !ok {
		assert.FailNow(t, err.Error())
	}

	if count, err := server.Del("foo"); err != nil {
		assert.FailNow(t, err.Error())
	} else {
		assert.Equal(t, 1, count)
	}
}

func TestAppend(t *testing.T) {
	server := NewServer(t)
	defer server.Close()

	hello := "hello"
	world := " world"

	if count, err := server.Append("message", hello); err != nil {
		assert.FailNow(t, err.Error())
	} else {
		assert.Equal(t, len(hello), count)
	}

	if count, err := server.Append("message", world); err != nil {
		assert.FailNow(t, err.Error())
	} else {
		assert.Equal(t, len(hello+world), count)
	}

	if value, err := server.Get("message"); err != nil {
		assert.FailNow(t, err.Error())
	} else {
		assert.Equal(t, hello+world, value)
	}
}

func TestExists(t *testing.T) {
	server := NewServer(t)
	defer server.Close()

	if ok, err := server.Set("foo", "bar"); !ok {
		assert.FailNow(t, err.Error())
	}

	if exists, err := server.Exists("foo"); err != nil {
		assert.FailNow(t, err.Error())
	} else {
		assert.True(t, exists)
	}

	if exists, err := server.Exists("nothing"); err != nil {
		assert.FailNow(t, err.Error())
	} else {
		assert.False(t, exists)
	}
}
