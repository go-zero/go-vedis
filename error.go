package vedis

// #include "vedis.h"
//
// void GetErrorMessage(vedis *store, const char **message) {
//     vedis_config(store, VEDIS_CONFIG_ERR_LOG, message, 0);
// }
//
import "C"
import "fmt"

type Error struct {
	Code    int
	Message string
}

func (e Error) Error() string {
	return fmt.Sprintf("(%d) %s", e.Code, e.Message)
}

func newError(code C.int, ptr *C.vedis) Error {
	var message *C.char
	C.GetErrorMessage(ptr, &message)
	return Error{int(code), C.GoString(message)}
}
