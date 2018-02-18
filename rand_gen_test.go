package tg

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandGen_Digit(t *testing.T) {
	digits, err := regexp.Compile("[0-9]+")
	assert.NoError(t, err)
	token, err := RandGen(10, Digit, "", "")
	assert.NoError(t, err)
	b := digits.MatchString(token)
	assert.True(t, b)
}
