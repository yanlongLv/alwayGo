package errors

import (
	"errors"

	"github.com/golang/protobuf/proto"
)

//Cancelled ..
func Cancelled(reason, format string, a ...interface{}) errors {
	return &StatusError{
		Code:    1,
		Message: "Canceled",
		Details: []proto.Message{
			//&ErrorItem{}
		},
	}
}

//IsCancelled ..
func IsCancelled(err error) bool {
	if se := new(StatusError); errors.As(err, se) {
		return se.Code == 1
	}
	return false
}
