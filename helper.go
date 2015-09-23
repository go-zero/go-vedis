package vedis

// #include "vedis.h"
import "C"
import "fmt"

func execute(v *Vedis, format string, values ...interface{}) error {
	command := fmt.Sprintf(format, values...)
	if status := C.vedis_exec(v.ptr, C.CString(command), -1); status != C.VEDIS_OK {
		return newError(status, v.ptr)
	}
	return nil
}

func result(v *Vedis) (*C.vedis_value, error) {
	value := new(C.vedis_value)
	if status := C.vedis_exec_result(v.ptr, &value); status != C.VEDIS_OK {
		return nil, newError(status, v.ptr)
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
