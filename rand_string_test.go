package tg_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/neumachen/tg"
)

func TestRandStringBytes(t *testing.T) {
	// Test case 1: Generate a random string of length 10
	result := tg.RandStringBytes(10)
	require.Len(t, result, 10, "Generated string length should match specified length")

	// Test case 2: Generate a random string of length 0
	result = tg.RandStringBytes(0)
	require.Len(t, result, 0, "Generated string length should be 0")

	// Test case 3: Generate a random string of length 100
	result = tg.RandStringBytes(100)
	require.Len(t, result, 100, "Generated string length should match specified length")
}
