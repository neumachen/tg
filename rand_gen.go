// Package tg is a token generator.
package tg

import (
	"errors"
	"math/rand"
	"strings"
	"time"
)

// None represents no character set.
const None = 0

// Lower represents the lowercase character set.
const Lower = 1 << 0

// Upper represents the uppercase character set.
const Upper = 1 << 1

// Digit represents the numeric character set.
const Digit = 1 << 2

// Punct represents the punctuation character set.
const Punct = 1 << 3

// LowerUpper represents the lowercase and uppercase character sets.
const LowerUpper = Lower | Upper

// LowerDigit represents the lowercase and numeric character sets.
const LowerDigit = Lower | Digit

// UpperDigit represents the uppercase and numeric character sets.
const UpperDigit = Upper | Digit

// LowerUpperDigit represents the lowercase, uppercase, and numeric character sets.
const LowerUpperDigit = LowerUpper | Digit

// All represents all character sets, including lowercase, uppercase, numeric, and punctuation.
const All = LowerUpperDigit | Punct

// lower represents the lowercase characters.
const lower = "abcdefghijklmnopqrstuvwxyz"

// upper represents the uppercase characters.
const upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// digit represents the numeric characters.
const digit = "0123456789"

// punct represents the punctuation characters.
const punct = "~!@#$%^&*()_+-="

// RandGen generates a random string of specified size, using the specified character set and exclusion criteria.
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
		return "", errors.New("too much to exclude")
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
