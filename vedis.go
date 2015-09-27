package vedis

// #include "vedis.h"
import "C"
import "encoding/json"

// Vedis datastore.
type Vedis struct {
	ptr *C.vedis
}

// Get a new Vedis datastore.
func New() *Vedis {
	return new(Vedis)
}

// Open the datastore.
func (v *Vedis) Open() (bool, error) {
	if status := C.vedis_open(&v.ptr, C.CString(":mem:")); status != C.VEDIS_OK {
		return false, newError(status, v.ptr)
	}
	return true, nil
}

// Close the datastore.
func (v *Vedis) Close() (bool, error) {
	if status := C.vedis_close(v.ptr); status != C.VEDIS_OK {
		return false, newError(status, v.ptr)
	}
	return true, nil
}

// Set key to hold the string value.
// If key already holds a value, it is overwritten, regardless of its type.
// Any previous time to live associated with the key is discarded on successful SET operation.
//
// See http://vedis.symisc.net/cmd/set.html
func (v *Vedis) Set(key string, value string) (bool, error) {
	return executeWithBoolResult(v, "SET \"%s\" \"%s\"", key, value)
}

// Set key to hold string value if key does not exist.
// In that case, it is equal to SET.
// When key already holds a value, no operation is performed.
// SETNX is short for "SET if N ot e X ists".
//
// See http://vedis.symisc.net/cmd/setnx.html
func (v *Vedis) SetNX(key string, value string) (bool, error) {
	return executeWithBoolResult(v, "SETNX \"%s\" \"%s\"", key, value)
}

// Sets the given keys to their respective values.
// MSET replaces existing values with new values, just as regular SET.
// See MSETNX if you don't want to overwrite existing values.
//
// See http://vedis.symisc.net/cmd/mset.html
func (v *Vedis) MSet(kv ...string) (bool, error) {
	command, args := massive("MSET", kv)
	return executeWithBoolResult(v, command, args...)
}

// Sets the given keys to their respective values.
// MSETNX replaces existing values with new values only if the key does not exits, just as regular SETNX.
//
// See http://vedis.symisc.net/cmd/mset.html
func (v *Vedis) MSetNX(kv ...string) (bool, error) {
	command, args := massive("MSETNX", kv)
	return executeWithBoolResult(v, command, args...)
}

// Check if a key already exists in the datastore.
//
// See http://vedis.symisc.net/cmd/exists.html
func (v *Vedis) Exists(key string) (bool, error) {
	return executeWithBoolResult(v, "EXISTS \"%s\"", key)
}

// Copy key values.
//
// See http://vedis.symisc.net/cmd/copy.html
func (v *Vedis) Copy(oldkey string, newkey string) (bool, error) {
	return executeWithBoolResult(v, "COPY \"%s\" \"%s\"", oldkey, newkey)
}

// Move key values (remove old key).
//
// See http://vedis.symisc.net/cmd/move.html
func (v *Vedis) Move(oldkey string, newkey string) (bool, error) {
	return executeWithBoolResult(v, "MOVE \"%s\" \"%s\"", oldkey, newkey)
}

// Get the value of key.
// If the key does not exist the special value null is returned.
//
// See http://vedis.symisc.net/cmd/get.html
func (v *Vedis) Get(key string) (string, error) {
	return executeWithStringResult(v, "GET \"%s\"", key)
}

// Atomically sets key to value and returns the old value stored at key.
// Returns an error when key exists but does not hold a string value.
//
// See http://vedis.symisc.net/cmd/getset.html
func (v *Vedis) GetSet(key string, value string) (string, error) {
	return executeWithStringResult(v, "GETSET \"%s\" \"%s\"", key, value)
}

// Returns the values of all specified keys.
// For every key that does not hold a string value or does not exist, the special value null is returned.
// Because of this, the operation never fails.
//
// See http://vedis.symisc.net/cmd/mget.html
func (v *Vedis) MGet(keys ...string) ([]string, error) {
	command, args := massive("MGET", keys)

	if err := execute(v, command, args...); err != nil {
		return nil, err
	}
	if result, err := result(v); err != nil {
		return nil, err
	} else {
		var values []string
		if err := json.Unmarshal([]byte(toString(result)), &values); err != nil {
			return nil, err
		}
		return values, nil
	}
}

// Removes the specified keys.
// A key is ignored if it does not exist.
//
// See http://vedis.symisc.net/cmd/del.html
func (v *Vedis) Del(key string) (int, error) {
	if err := execute(v, "DEL \"%s\"", key); err != nil {
		return 0, err
	}
	if result, err := result(v); err != nil {
		return 0, err
	} else {
		return toInt(result), nil
	}
}

// If key already exists and is a string, this command appends the value at the end of the string.
// If key does not exist it is created and set as an empty string, so APPEND will be similar to SET in this special case.
//
// See http://vedis.symisc.net/cmd/append.html
func (v *Vedis) Append(key string, value string) (int, error) {
	if err := execute(v, "APPEND \"%s\" \"%s\"", key, value); err != nil {
		return 0, err
	}
	if result, err := result(v); err != nil {
		return 0, err
	} else {
		if int(C.vedis_value_is_int(result)) == 1 {
			return toInt(result), nil
		} else {
			return len(value), nil
		}
	}
}
