package tools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RightPad(t *testing.T) {
	assert.Equal(t, len(RightPad("abc","8")), 8)
}
