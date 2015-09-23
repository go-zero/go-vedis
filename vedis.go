package vedis

// #include "vedis.h"
import "C"

import (
	"fmt"
)

type Vedis interface {
	Open() (bool, error)
	Close() (bool, error)

	// commands
	Set(string, string) (bool, error)
	Get(string) (string, error)
	Del(string) (int, error)
	Append(string, string) (int, error)
	Exists(string) (bool, error)
}

// internal Vedis implementation
type store struct {
	ptr *C.vedis
}

// constructor
func New() Vedis {
	return new(store)
}

func (s *store) Open() (bool, error) {
	if status := C.vedis_open(&s.ptr, C.CString(":mem:")); status != C.VEDIS_OK {
		return false, NewError(status, s.ptr)
	}
	return true, nil
}

func (s *store) Close() (bool, error) {
	if status := C.vedis_close(s.ptr); status != C.VEDIS_OK {
		return false, NewError(status, s.ptr)
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

// private functions

func execute(s *store, format string, values ...interface{}) error {
	command := fmt.Sprintf(format, values...)
	if status := C.vedis_exec(s.ptr, C.CString(command), -1); status != C.VEDIS_OK {
		return NewError(status, s.ptr)
	}
	return nil
}

func result(s *store) (*C.vedis_value, error) {
	value := new(C.vedis_value)
	if status := C.vedis_exec_result(s.ptr, &value); status != C.VEDIS_OK {
		return nil, NewError(status, s.ptr)
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
