package json

import (
	"encoding/json"

	"github.com/alwayGo/encoding"
	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/encoding/protojson"
)

//Name is the name registered
const Name = "json"

type codec struct{}

func init() {
	encoding.RegisterCodec(codec{})
}

func (codec) Marshal(v interface{}) ([]byte, error) {
	m, ok := v.(proto.Message)
	if ok {
		return protojson.Marshal(m)
	}
	return json.Marshal(v)
}

func (codec) Unmarshal(data []byte, v interface{}) error {
	m, ok := v.(proto.Message)
	if ok {
		return protojson.Unmarshal(data, m)
	}
	return json.Unmarshal(data, v)
}

func (codec) Name() string {
	return Name
}
