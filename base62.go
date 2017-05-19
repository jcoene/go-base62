// Package base62 implements conversion to and from base62. Useful for url shorteners.
package base62

import (
	"bytes"
	"math"
)

// characters used for conversion
const alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// Encode encodes an int64 to a base62 encoded string.
func Encode(number int64) string {
	return EncodeUint64(uint64(number))
}

// Decode decodes a base62 encoded string to an int64.
func Decode(token string) int64 {
	return int64(DecodeUint64(token))
}

// EncodeUint64 encodes a uint64 to a base62 encoded string.
func EncodeUint64(number uint64) string {
	if number == 0 {
		return string(alphabet[0])
	}

	chars := make([]byte, 0)

	length := uint64(len(alphabet))

	for number > 0 {
		result := number / length
		remainder := number % length
		chars = append(chars, alphabet[remainder])
		number = result
	}

	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}

	return string(chars)
}

// DecodeUint64 decodes a base62 encoded string to an uint64.
func DecodeUint64(token string) uint64 {
	number := uint64(0)
	idx := float64(0.0)
	chars := []byte(alphabet)

	charsLength := float64(len(chars))
	tokenLength := float64(len(token))

	for _, c := range []byte(token) {
		power := tokenLength - (idx + 1)
		index := uint64(bytes.IndexByte(chars, c))
		number += index * uint64(math.Pow(charsLength, power))
		idx++
	}

	return number
}
