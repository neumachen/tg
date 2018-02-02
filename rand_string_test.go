package tg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandStringBytes(t *testing.T) {
	str := RandStringBytes(12)
	assert.Equal(t, len(str), 12)
}
