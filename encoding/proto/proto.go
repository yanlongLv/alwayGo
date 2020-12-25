package proto

import (
	"github.com/alwayGo/encoding"
	"google.golang.org/protobuf/proto"
)

//Name is the name registered
const Name = "proto"

type codec struct{}

func init() {
	encoding.RegisterCodec(codec{})
}

func (codec) Marshal(v interface{}) ([]byte, error) {
	return proto.Marshal(v.(proto.Message))
}

func (codec) Unmarshal(data []byte, v interface{}) error {
	return proto.Unmarshal(data, v.(proto.Message))
}

func (codec) Name() string {
	return Name
}
