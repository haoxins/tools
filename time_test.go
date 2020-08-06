package tools

import (
	"testing"
)

func TestYesterday(t *testing.T) {
	t.Log("Yesterday is:", Yesterday("Asia/Shanghai", "2006-01-02"))
}
