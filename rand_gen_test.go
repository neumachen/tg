package tg_test

import (
	"strings"
	"testing"
	"unicode"

	"github.com/stretchr/testify/require"

	"github.com/neumachen/tg"
)

func TestRandGen(t *testing.T) {
	checkOutput := func(s string, checks ...func(c rune) bool) bool {
		checkResult := false
		for _, c := range s {
			for i := range checks {
				if checkResult = checks[i](c); checkResult {
					break
				}
			}
			if !checkResult {
				break
			}
		}
		return checkResult
	}

	// Test case 1: Generate a random string of length 8 using only lowercase characters
	result, err := tg.RandGen(8, tg.Lower, "", "")
	require.NoError(t, err)
	require.Equal(t, 8, len(result))
	require.True(t, checkOutput(result, unicode.IsLower))

	result, err = tg.RandGen(8, tg.Upper, "", "")
	require.NoError(t, err)
	require.Equal(t, 8, len(result))
	require.True(t, checkOutput(result, unicode.IsUpper))

	// Test case 2: Generate a random string of length 12 using both uppercase and numeric characters
	result, err = tg.RandGen(12, tg.Upper|tg.Digit, "", "")
	require.NoError(t, err)
	require.Equal(t, 12, len(result))
	// require.True(t, checkUppercase(result))
	// require.True(t, checkNumeric(result))
	require.True(t, checkOutput(result, unicode.IsUpper, unicode.IsNumber))

	// Test case 3: Generate a random string of length 10 using all character sets except punctuation, excluding certain characters
	result, err = tg.RandGen(10, tg.All^tg.Punct, "", "AEIOU")
	require.NoError(t, err)
	require.Equal(t, 10, len(result))
	require.True(t, checkOutput(result, func(c rune) bool {
		return unicode.IsLower(c) || unicode.IsUpper(c)
	},
		unicode.IsLower,
		unicode.IsUpper,
		unicode.IsNumber,
		unicode.IsPunct,
		func(r rune) bool {
			return strings.ContainsRune("AEIOU", r)
		},
	))
}
