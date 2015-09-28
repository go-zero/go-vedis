package vedis

// #include "vedis.h"
import "C"
import (
	"encoding/json"
	"fmt"
)

func execute(v *Vedis, format string, values ...interface{}) error {
	command := fmt.Sprintf(format, values...)
	if status := C.vedis_exec(v.ptr, C.CString(command), -1); status != C.VEDIS_OK {
		return newError(status, v.ptr)
	}
	return nil
}

func result(v *Vedis) (*C.vedis_value, error) {
	var value *C.vedis_value
	if status := C.vedis_exec_result(v.ptr, &value); status != C.VEDIS_OK {
		return nil, newError(status, v.ptr)
	}
	return value, nil
}

func massive(command string, values []string) (string, []interface{}) {
	args := []interface{}{}
	for _, value := range values {
		command += " \"%s\""
		args = append(args, value)
	}
	return command, args
}

func executeWithIntResult(v *Vedis, cmd string, values ...interface{}) (int, error) {
	if err := execute(v, cmd, values...); err != nil {
		return 0, err
	}
	if result, err := result(v); err != nil {
		return 0, err
	} else {
		return toInt(result), nil
	}
}

func executeWithStringResult(v *Vedis, cmd string, values ...interface{}) (string, error) {
	if err := execute(v, cmd, values...); err != nil {
		return "", err
	}
	if result, err := result(v); err != nil {
		return "", err
	} else {
		return toString(result), nil
	}
}

func executeWithBoolResult(v *Vedis, cmd string, values ...interface{}) (bool, error) {
	if result, err := executeWithStringResult(v, cmd, values...); err != nil {
		return false, err
	} else {
		return result == "true", nil
	}
}

func executeWithArrayResult(v *Vedis, cmd string, values ...interface{}) ([]string, error) {
	if result, err := executeWithStringResult(v, cmd, values...); err != nil {
		return nil, err
	} else {
		var values []string
		if err := json.Unmarshal([]byte(result), &values); err != nil {
			return nil, err
		}
		return values, nil
	}
}

func toString(value *C.vedis_value) string {
	var length *C.int
	return C.GoString(C.vedis_value_to_string(value, length))
}

func toInt(value *C.vedis_value) int {
	return int(C.vedis_value_to_int(value))
}
