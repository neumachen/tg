package tg

import (
	"math/rand"
	"time"
)

// letterBytes defines the characters used for generating random strings.
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

// letterIdxBits represents the number of bits required to represent a letter index.
const letterIdxBits = 6

// letterIdxMask is a bitmask used to extract the letter index from the random number.
const letterIdxMask = 1<<letterIdxBits - 1

// letterIdxMax is the maximum number of letter indices that can fit in 63 bits.
const letterIdxMax = 63 / letterIdxBits

var src = rand.NewSource(time.Now().UnixNano())

// RandStringBytes generates a random string of the specified length.
func RandStringBytes(n int) string {
	b := make([]byte, n)

	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
