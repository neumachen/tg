// Package tg is a token generator
package tg

import (
	"errors"
	"math/rand"
	"strings"
	"time"
)

// None ...
const None = 0

// Lower ...
const Lower = 1 << 0

// Upper ...
const Upper = 1 << 1

// Digit ...
const Digit = 1 << 2

// Punct ...
const Punct = 1 << 3

// LowerUpper ...
const LowerUpper = Lower | Upper

// LowerDigit ...
const LowerDigit = Lower | Digit

// UpperDigit ...
const UpperDigit = Upper | Digit

// LowerUpperDigit ...
const LowerUpperDigit = LowerUpper | Digit

// All ...
const All = LowerUpperDigit | Punct

const lower = "abcdefghijklmnopqrstuvwxyz"
const upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const digit = "0123456789"
const punct = "~!@#$%^&*()_+-="

// RandGen ...
func RandGen(size int, set int, include string, exclude string) (string, error) {
	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)
	all := include
	if set&Lower > 0 {
		all += lower
	}
	if set&Upper > 0 {
		all += upper
	}
	if set&Digit > 0 {
		all += digit
	}
	if set&Punct > 0 {
		all += punct
	}

	lenAll := len(all)
	if len(exclude) >= lenAll {
		return "", errors.New("Too much to exclude.")
	}
	buf := make([]byte, size)
	for i := 0; i < size; i++ {
		b := all[rnd.Intn(lenAll)]
		if strings.Contains(exclude, string(b)) {
			i--
			continue
		}
		buf[i] = b
	}
	return string(buf), nil
}
