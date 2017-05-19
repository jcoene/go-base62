package base62

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("0", Encode(0))
	assert.Equal("1b", Encode(99))
	assert.Equal("QNZ", Encode(101405))
	assert.Equal("15y79wVi9XQ", Encode(920110421043409228))
}

func TestEncodeUint64(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("0", EncodeUint64(0))
	assert.Equal("1b", EncodeUint64(99))
	assert.Equal("QNZ", EncodeUint64(101405))
	assert.Equal("15y79wVi9XQ", EncodeUint64(920110421043409228))
	assert.Equal("GU0iBVp7iAC", EncodeUint64(13835058055282163712))
}

func TestDecode(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(int64(0), Decode("0"))
	assert.Equal(int64(99), Decode("1b"))
	assert.Equal(int64(101405), Decode("QNZ"))
	assert.Equal(int64(920110421043409228), Decode("15y79wVi9XQ"))
}

func TestDecodeUint64(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(uint64(0), DecodeUint64("0"))
	assert.Equal(uint64(99), DecodeUint64("1b"))
	assert.Equal(uint64(101405), DecodeUint64("QNZ"))
	assert.Equal(uint64(920110421043409228), DecodeUint64("15y79wVi9XQ"))
	assert.Equal(uint64(13835058055282163712), DecodeUint64("GU0iBVp7iAC"))
}
