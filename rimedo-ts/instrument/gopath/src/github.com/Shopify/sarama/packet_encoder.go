//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/packet_encoder.go:1
package sarama

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/packet_encoder.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/packet_encoder.go:1
	_go_fuzz_dep_ "go-fuzz-dep"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/packet_encoder.go:1
)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/packet_encoder.go:1
import (
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/packet_encoder.go:1
	_atomic_ "sync/atomic"
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/packet_encoder.go:1
)

import "github.com/rcrowley/go-metrics"

// PacketEncoder is the interface providing helpers for writing with Kafka's encoding rules.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/packet_encoder.go:5
// Types implementing Encoder only need to worry about calling methods like PutString,
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/packet_encoder.go:5
// not about how a string is represented in Kafka.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/packet_encoder.go:8
type packetEncoder interface {
	// Primitives
	putInt8(in int8)
	putInt16(in int16)
	putInt32(in int32)
	putInt64(in int64)
	putVarint(in int64)
	putUVarint(in uint64)
	putFloat64(in float64)
	putCompactArrayLength(in int)
	putArrayLength(in int) error
	putBool(in bool)

	// Collections
	putBytes(in []byte) error
	putVarintBytes(in []byte) error
	putCompactBytes(in []byte) error
	putRawBytes(in []byte) error
	putCompactString(in string) error
	putNullableCompactString(in *string) error
	putString(in string) error
	putNullableString(in *string) error
	putStringArray(in []string) error
	putCompactInt32Array(in []int32) error
	putNullableCompactInt32Array(in []int32) error
	putInt32Array(in []int32) error
	putInt64Array(in []int64) error
	putEmptyTaggedFieldArray()

	// Provide the current offset to record the batch size metric
	offset() int

	// Stacks, see PushEncoder
	push(in pushEncoder)
	pop() error

	// To record metrics when provided
	metricRegistry() metrics.Registry
}

// PushEncoder is the interface for encoding fields like CRCs and lengths where the value
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/packet_encoder.go:48
// of the field depends on what is encoded after it in the packet. Start them with PacketEncoder.Push() where
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/packet_encoder.go:48
// the actual value is located in the packet, then PacketEncoder.Pop() them when all the bytes they
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/packet_encoder.go:48
// depend upon have been written.
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/packet_encoder.go:52
type pushEncoder interface {
	// Saves the offset into the input buffer as the location to actually write the calculated value when able.
	saveOffset(in int)

	// Returns the length of data to reserve for the output of this encoder (eg 4 bytes for a CRC32).
	reserveLength() int

	// Indicates that all required data is now available to calculate and write the field.
	// SaveOffset is guaranteed to have been called first. The implementation should write ReserveLength() bytes
	// of data to the saved offset, based on the data between the saved offset and curOffset.
	run(curOffset int, buf []byte) error
}

// dynamicPushEncoder extends the interface of pushEncoder for uses cases where the length of the
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/packet_encoder.go:65
// fields itself is unknown until its value was computed (for instance varint encoded length
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/packet_encoder.go:65
// fields).
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/packet_encoder.go:68
type dynamicPushEncoder interface {
	pushEncoder

	// Called during pop() to adjust the length of the field.
	// It should return the difference in bytes between the last computed length and current length.
	adjustLength(currOffset int) int
}

//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/packet_encoder.go:74
var _ = _atomic_.LoadUint32(&_go_fuzz_dep_.NoUse)
//line /home/tianchang/go/pkg/mod/github.com/!shopify/sarama@v1.31.1/packet_encoder.go:74
var _ = _go_fuzz_dep_.CoverTab
