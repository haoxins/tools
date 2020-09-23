package tools

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ExecCommand(t *testing.T) {
	out, err := ExecCommand("ls", "-a")
	t.Log(out)
	assert.Equal(t, err, nil)
	assert.Equal(t, strings.Contains(out, "exec.go"), true)
	assert.Equal(t, strings.Contains(out, "exec_test.go"), true)

	out, err = ExecCommand("foo")
	assert.NotEqual(t, err, nil)
	assert.NotEqual(t, out, "")
}

func Test_Split(t *testing.T) {

	assert.Equal(t, split("ls -a -l"), []string{"ls", "-a", "-l"})
	assert.Equal(t, split("ls -a -l "), []string{"ls", "-a", "-l"})
	assert.Equal(t, split(" ls -a -l"), []string{"ls", "-a", "-l"})
	assert.Equal(t, split(" ls -a -l "), []string{"ls", "-a", "-l"})
	assert.Equal(t, split(" ls   -a   -l "), []string{"ls", "-a", "-l"})

	assert.Equal(t, split("cmd \"a b c\""), []string{"cmd", "a b c"})
	assert.Equal(t, split("cmd -a \"a b c\""), []string{"cmd", "-a", "a b c"})
	assert.Equal(t, split(" cmd -a \"a b c\" "), []string{"cmd", "-a", "a b c"})
	assert.Equal(t, split(" cmd -a a/b \"a  b  c\" "), []string{"cmd", "-a", "a/b", "a  b  c"})
	assert.Equal(t, split(" cmd  -a  a/b  \"a  b  c\" "), []string{"cmd", "-a", "a/b", "a  b  c"})
}
