//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/packet_decoder.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/packet_decoder.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/packet_decoder.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/packet_decoder.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/packet_decoder.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/packet_decoder.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/packet_decoder.go:1
)

// PacketDecoder is the interface providing helpers for reading with Kafka's encoding rules.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/packet_decoder.go:3
// Types implementing Decoder only need to worry about calling methods like GetString,
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/packet_decoder.go:3
// not about how a string is represented in Kafka.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/packet_decoder.go:6
type packetDecoder interface {
	// Primitives
	getInt8() (int8, error)
	getInt16() (int16, error)
	getInt32() (int32, error)
	getInt64() (int64, error)
	getVarint() (int64, error)
	getUVarint() (uint64, error)
	getFloat64() (float64, error)
	getArrayLength() (int, error)
	getCompactArrayLength() (int, error)
	getBool() (bool, error)
	getEmptyTaggedFieldArray() (int, error)

	// Collections
	getBytes() ([]byte, error)
	getVarintBytes() ([]byte, error)
	getCompactBytes() ([]byte, error)
	getRawBytes(length int) ([]byte, error)
	getString() (string, error)
	getNullableString() (*string, error)
	getCompactString() (string, error)
	getCompactNullableString() (*string, error)
	getCompactInt32Array() ([]int32, error)
	getInt32Array() ([]int32, error)
	getInt64Array() ([]int64, error)
	getStringArray() ([]string, error)

	// Subsets
	remaining() int
	getSubset(length int) (packetDecoder, error)
	peek(offset, length int) (packetDecoder, error)	// similar to getSubset, but it doesn't advance the offset
	peekInt8(offset int) (int8, error)		// similar to peek, but just one byte

	// Stacks, see PushDecoder
	push(in pushDecoder) error
	pop() error
}

// PushDecoder is the interface for decoding fields like CRCs and lengths where the validity
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/packet_decoder.go:45
// of the field depends on what is after it in the packet. Start them with PacketDecoder.Push() where
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/packet_decoder.go:45
// the actual value is located in the packet, then PacketDecoder.Pop() them when all the bytes they
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/packet_decoder.go:45
// depend upon have been decoded.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/packet_decoder.go:49
type pushDecoder interface {
	// Saves the offset into the input buffer as the location to actually read the calculated value when able.
	saveOffset(in int)

	// Returns the length of data to reserve for the input of this encoder (eg 4 bytes for a CRC32).
	reserveLength() int

	// Indicates that all required data is now available to calculate and check the field.
	// SaveOffset is guaranteed to have been called first. The implementation should read ReserveLength() bytes
	// of data from the saved offset, and verify it based on the data between the saved offset and curOffset.
	check(curOffset int, buf []byte) error
}

// dynamicPushDecoder extends the interface of pushDecoder for uses cases where the length of the
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/packet_decoder.go:62
// fields itself is unknown until its value was decoded (for instance varint encoded length
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/packet_decoder.go:62
// fields).
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/packet_decoder.go:62
// During push, dynamicPushDecoder.decode() method will be called instead of reserveLength()
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/packet_decoder.go:66
type dynamicPushDecoder interface {
	pushDecoder
	decoder
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/packet_decoder.go:69
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/packet_decoder.go:69
var _ = _go_fuzz_dep_.CoverTab
