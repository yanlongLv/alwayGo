package encoding

import "google.golang.org/grpc/encoding"

//Codec ..
type Codec encoding.Codec

//Compressor ..
type Compressor encoding.Compressor

//GetCodec ..
func GetCodec(contentSubtype string) Codec {
	return encoding.GetCodec(contentSubtype)
}

//RegisterCodec ..
func RegisterCodec(codec Codec) {
	encoding.RegisterCodec(codec)
}

//RegisterCompressor ..
func RegisterCompressor(c Compressor) {
	encoding.RegisterCompressor(c)
}

//GetCompressor ..
func GetCompressor(name string) Compressor {
	return encoding.GetCompressor(name)
}
