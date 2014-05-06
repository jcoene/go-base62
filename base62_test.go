package base62

import (
	"testing"

	assert "github.com/pilu/miniassert"
)

func TestEncode(t *testing.T) {
	assert.Equal(t, "0", Encode(0))
	assert.Equal(t, "1b", Encode(99))
	assert.Equal(t, "QNZ", Encode(101405))
	assert.Equal(t, "15y79wVi9XQ", Encode(920110421043409228))

}

func TestDecode(t *testing.T) {
	assert.Equal(t, int64(0), Decode("0"))
	assert.Equal(t, int64(99), Decode("1b"))
	assert.Equal(t, int64(101405), Decode("QNZ"))
	assert.Equal(t, int64(920110421043409228), Decode("15y79wVi9XQ"))
}
