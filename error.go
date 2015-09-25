package vedis

// #include "vedis_extra.h"
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
	C.vedis_error_message(ptr, &message)
	return Error{int(code), C.GoString(message)}
}
