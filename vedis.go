package vedis

// #include "vedis.h"
import "C"

import (
	"fmt"
)

type Vedis interface {

	// Open the datastore.
	Open() (bool, error)

	// Close the datastore.
	Close() (bool, error)

	// Set key to hold the string value.
	// If key already holds a value, it is overwritten, regardless of its type.
	// Any previous time to live associated with the key is discarded on successful SET operation.
	//
	// See http://vedis.symisc.net/cmd/set.html
	Set(key string, value string) (bool, error)

	// Get the value of key.
	// If the key does not exist the special value null is returned.
	//
	// See http://vedis.symisc.net/cmd/get.html
	Get(key string) (string, error)

	// Removes the specified keys.
	// A key is ignored if it does not exist.
	//
	// See http://vedis.symisc.net/cmd/del.html
	Del(key string) (int, error)

	// If key already exists and is a string, this command appends the value at the end of the string.
	// If key does not exist it is created and set as an empty string, so APPEND will be similar to SET in this special case.
	//
	// See http://vedis.symisc.net/cmd/append.html
	Append(key string, value string) (int, error)

	// Check if a key already exists in the datastore.
	//
	// See http://vedis.symisc.net/cmd/exists.html
	Exists(key string) (bool, error)
}

// Get a new Vedis datastore.
func New() Vedis {
	return new(store)
}

func (s *store) Open() (bool, error) {
	if status := C.vedis_open(&s.ptr, C.CString(":mem:")); status != C.VEDIS_OK {
		return false, newError(status, s.ptr)
	}
	return true, nil
}

func (s *store) Close() (bool, error) {
	if status := C.vedis_close(s.ptr); status != C.VEDIS_OK {
		return false, newError(status, s.ptr)
	}
	return true, nil
}

func (s *store) Set(key string, value string) (bool, error) {
	if err := execute(s, "SET %s \"%s\"", key, value); err != nil {
		return false, err
	}
	if result, err := result(s); err != nil {
		return false, err
	} else {
		return toString(result) == "true", nil
	}
}

func (s *store) Get(key string) (string, error) {
	if err := execute(s, "GET %s", key); err != nil {
		return "", err
	}
	if result, err := result(s); err != nil {
		return "", err
	} else {
		return toString(result), nil
	}
}

func (s *store) Del(key string) (int, error) {
	if err := execute(s, "DEL %s", key); err != nil {
		return 0, err
	}
	if result, err := result(s); err != nil {
		return 0, err
	} else {
		return toInt(result), nil
	}
}

func (s *store) Append(key string, value string) (int, error) {
	if err := execute(s, "APPEND %s \"%s\"", key, value); err != nil {
		return 0, err
	}
	if result, err := result(s); err != nil {
		return 0, err
	} else {
		if int(C.vedis_value_is_int(result)) == 1 {
			return toInt(result), nil
		} else {
			return len(value), nil
		}
	}
}

func (s *store) Exists(key string) (bool, error) {
	if err := execute(s, "EXISTS %s", key); err != nil {
		return false, err
	}
	if result, err := result(s); err != nil {
		return false, err
	} else {
		return toInt(result) == 1, nil
	}
}

// internal Vedis implementation
type store struct {
	ptr *C.vedis
}

// private functions

func execute(s *store, format string, values ...interface{}) error {
	command := fmt.Sprintf(format, values...)
	if status := C.vedis_exec(s.ptr, C.CString(command), -1); status != C.VEDIS_OK {
		return newError(status, s.ptr)
	}
	return nil
}

func result(s *store) (*C.vedis_value, error) {
	value := new(C.vedis_value)
	if status := C.vedis_exec_result(s.ptr, &value); status != C.VEDIS_OK {
		return nil, newError(status, s.ptr)
	}
	return value, nil
}

func toString(value *C.vedis_value) string {
	length := new(C.int)
	return C.GoString(C.vedis_value_to_string(value, length))
}

func toInt(value *C.vedis_value) int {
	return int(C.vedis_value_to_int(value))
}
