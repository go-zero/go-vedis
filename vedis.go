package vedis

// #include "vedis.h"
import "C"

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
// See http://vedis.symisc.net/cmd/msetnx.html
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

// Returns the values of all specified keys.
// For every key that does not hold a string value or does not exist, the special value null is returned.
// Because of this, the operation never fails.
//
// See http://vedis.symisc.net/cmd/mget.html
func (v *Vedis) MGet(keys ...string) ([]string, error) {
	command, args := massive("MGET", keys)
	return executeWithArrayResult(v, command, args...)
}

// Atomically sets key to value and returns the old value stored at key.
// Returns an error when key exists but does not hold a string value.
//
// See http://vedis.symisc.net/cmd/getset.html
func (v *Vedis) GetSet(key string, value string) (string, error) {
	return executeWithStringResult(v, "GETSET \"%s\" \"%s\"", key, value)
}

// Removes the specified keys.
// A key is ignored if it does not exist.
//
// See http://vedis.symisc.net/cmd/del.html
func (v *Vedis) Del(key string) (int, error) {
	return executeWithIntResult(v, "DEL \"%s\"", key)
}

// Increments the number stored at key by one.
// If the key does not exist, it is set to 0 before performing the operation.
// An error is returned if the key contains a value of the wrong type or contains a string that can not be represented as integer.
// This operation is limited to 64 bit signed integers.
//
// See http://vedis.symisc.net/cmd/incr.html
func (v *Vedis) Incr(key string) (int, error) {
	return executeWithIntResult(v, "INCR \"%s\"", key)
}

// Increments the number stored at key by increment.
// If the key does not exist, it is set to 0 before performing the operation.
// An error is returned if the key contains a value of the wrong type or contains a string that can not be represented as integer.
// This operation is limited to 64 bit signed integers.
//
// See http://vedis.symisc.net/cmd/incrby.html
func (v *Vedis) IncrBy(key string, increment int) (int, error) {
	return executeWithIntResult(v, "INCRBY \"%s\" %d", key, increment)
}

// Decrements the number stored at key by one.
// If the key does not exist, it is set to 0 before performing the operation.
// An error is returned if the key contains a value of the wrong type or contains a string that can not be represented as integer.
// This operation is limited to 64 bit signed integers.
//
// See http://vedis.symisc.net/cmd/decr.html
func (v *Vedis) Decr(key string) (int, error) {
	return executeWithIntResult(v, "DECR \"%s\"", key)
}

// Decrements the number stored at key by decrement.
// If the key does not exist, it is set to 0 before performing the operation.
// An error is returned if the key contains a value of the wrong type or contains a string that can not be represented as integer.
// This operation is limited to 64 bit signed integers.
//
// See http://vedis.symisc.net/cmd/decrby.html
func (v *Vedis) DecrBy(key string, decrement int) (int, error) {
	return executeWithIntResult(v, "DECRBY \"%s\" %d", key, decrement)
}

// Sets field in the hash stored at key to value.
// If key does not exist, a new key holding a hash is created.
// If field already exists in the hash, it is overwritten.
//
// See http://vedis.symisc.net/cmd/hset.html
func (v *Vedis) HSet(key string, field string, value string) (bool, error) {
	return executeWithBoolResult(v, "HSET \"%s\" \"%s\" \"%s\"", key, field, value)
}

// Returns the value associated with field in the hash stored at key
//
// See http://vedis.symisc.net/cmd/hget.html
func (v *Vedis) HGet(key string, field string) (string, error) {
	return executeWithStringResult(v, "HGET \"%s\" \"%s\"", key, field)
}

// Removes the specified fields from the hash stored at key.
// Specified fields that do not exist within this hash are ignored.
// If key does not exist, it is treated as an empty hash and this command returns 0.
//
// See http://vedis.symisc.net/cmd/hdel.html
func (v *Vedis) HDel(key string, fields ...string) (int, error) {
	command, args := massive("HDEL", append([]string{key}, fields...))
	return executeWithIntResult(v, command, args...)
}

// Returns the number of fields contained in the hash stored at key.
//
// See http://vedis.symisc.net/cmd/hlen.html
func (v *Vedis) HLen(key string) (int, error) {
	return executeWithIntResult(v, "HLEN \"%s\"", key)
}

// Returns if field is an existing field in the hash stored at key.
//
// See http://vedis.symisc.net/cmd/hexists.html
func (v *Vedis) HExists(key string, field string) (bool, error) {
	return executeWithBoolResult(v, "HEXISTS \"%s\" \"%s\"", key, field)
}

// Returns all field names in the hash stored at key.
//
// See http://vedis.symisc.net/cmd/hkeys.html
func (v *Vedis) HKeys(key string) ([]string, error) {
	return executeWithArrayResult(v, "HKEYS \"%s\"", key)
}

// Returns all field values in the hash stored at key.
//
// See http://vedis.symisc.net/cmd/hvals.html
func (v *Vedis) HVals(key string) ([]string, error) {
	return executeWithArrayResult(v, "HVALS \"%s\"", key)
}

// Sets the specified fields to their respective values in the hash stored at key.
// This command overwrites any existing fields in the hash.
// If key does not exist, a new key holding a hash is created.
//
// See http://vedis.symisc.net/cmd/hmset.html
func (v *Vedis) HMSet(key string, fv ...string) (int, error) {
	command, args := massive("HMSET", append([]string{key}, fv...))
	return executeWithIntResult(v, command, args...)
}

// Returns the values associated with the specified fields in the hash stored at key.
// For every field that does not exist in the hash, a nil value is returned.
// Because a non-existing keys are treated as empty hashes, running HMGET against a non-existing key will return a list of nil values.
//
// See http://vedis.symisc.net/cmd/hmget.html
func (v *Vedis) HMGet(key string, fields ...string) ([]string, error) {
	command, args := massive("HMGET", append([]string{key}, fields...))
	return executeWithArrayResult(v, command, args...)
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
