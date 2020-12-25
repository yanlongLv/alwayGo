package errors

import "github.com/golang/protobuf/proto"

//StatusError ..
type StatusError struct {
	Code    int32           `json:"code"`
	Message string          `json:"message"`
	Details []proto.Message `json:"details"`
}
