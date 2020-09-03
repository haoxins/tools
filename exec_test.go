package tools

import (
	"strings"
	"testing"

	. "github.com/pkg4go/assert"
)

	a := A{t}

func Test_ExecCommand(t *testing.T) {
	out, err := ExecCommand("ls", "-a")
	t.Log(out)
	a.Equal(err, nil)
	a.Equal(strings.Contains(out, "exec.go"), true)
	a.Equal(strings.Contains(out, "exec_test.go"), true)

	out, err = ExecCommand("foo")
	a.NotEqual(err, nil)
	a.NotEqual(out, "")
}

	a := A{t}
func Test_Split(t *testing.T) {

	a.Equal(split("ls -a -l"), []string{"ls", "-a", "-l"})
	a.Equal(split("ls -a -l "), []string{"ls", "-a", "-l"})
	a.Equal(split(" ls -a -l"), []string{"ls", "-a", "-l"})
	a.Equal(split(" ls -a -l "), []string{"ls", "-a", "-l"})
	a.Equal(split(" ls   -a   -l "), []string{"ls", "-a", "-l"})

	a.Equal(split("cmd \"a b c\""), []string{"cmd", "a b c"})
	a.Equal(split("cmd -a \"a b c\""), []string{"cmd", "-a", "a b c"})
	a.Equal(split(" cmd -a \"a b c\" "), []string{"cmd", "-a", "a b c"})
	a.Equal(split(" cmd -a a/b \"a  b  c\" "), []string{"cmd", "-a", "a/b", "a  b  c"})
	a.Equal(split(" cmd  -a  a/b  \"a  b  c\" "), []string{"cmd", "-a", "a/b", "a  b  c"})
}
