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
	if err := execute(v, "SET %s \"%s\"", key, value); err != nil {
		return false, err
	}
	if result, err := result(v); err != nil {
		return false, err
	} else {
		return toString(result) == "true", nil
	}
}

// Get the value of key.
// If the key does not exist the special value null is returned.
//
// See http://vedis.symisc.net/cmd/get.html
func (v *Vedis) Get(key string) (string, error) {
	if err := execute(v, "GET %s", key); err != nil {
		return "", err
	}
	if result, err := result(v); err != nil {
		return "", err
	} else {
		return toString(result), nil
	}
}

// Removes the specified keys.
// A key is ignored if it does not exist.
//
// See http://vedis.symisc.net/cmd/del.html
func (v *Vedis) Del(key string) (int, error) {
	if err := execute(v, "DEL %s", key); err != nil {
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
	if err := execute(v, "APPEND %s \"%s\"", key, value); err != nil {
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

// Check if a key already exists in the datastore.
//
// See http://vedis.symisc.net/cmd/exists.html
func (v *Vedis) Exists(key string) (bool, error) {
	if err := execute(v, "EXISTS %s", key); err != nil {
		return false, err
	}
	if result, err := result(v); err != nil {
		return false, err
	} else {
		return toInt(result) == 1, nil
	}
}
