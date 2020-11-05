package tools

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Resolve(t *testing.T) {
	assert.Equal(t, true, strings.HasSuffix(Resolve("", "c"), "tools/c"))
	assert.Equal(t, "/a/b/c", Resolve("/a/b", "./c"))
	assert.Equal(t, "/a/c", Resolve("/a/b", "../c"))
	assert.Equal(t, "/a/b/c", Resolve("/a/b", "c"))
	assert.Equal(t, "/c", Resolve("/a/b", "/c"))
	assert.Equal(t, "/c", Resolve("", "/c"))
}
