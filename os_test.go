package tools

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Getenv(t *testing.T) {
	k := "not_exists_xxx"
	assert.Equal(t, "欧耶", Getenv(k, "欧耶"))
	os.Setenv(k, "咻")
	assert.Equal(t, "咻", Getenv(k, "欧耶"))
}
