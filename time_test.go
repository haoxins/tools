package tools

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_ParseTime(t *testing.T) {
	r, local := ParseTime("Mon, 01/02/06, 03:04PM", "Thu, 05/19/11, 10:47PM", "Asia/Shanghai")
	assert.Equal(t, r.Format(time.RFC3339), "2011-05-19T22:47:00Z")
	assert.Equal(t, local.String(), "Asia/Shanghai")
}

func Test_StartOfDate(t *testing.T) {
	r := StartOfDate("Mon, 01/02/06, 03:04PM", "Thu, 05/19/11, 10:47PM", "Asia/Shanghai")
	assert.Equal(t, r.Format(time.RFC3339), "2011-05-19T00:00:00+08:00")
}

func Test_EndOfDate(t *testing.T) {
	r := EndOfDate("Mon, 01/02/06, 03:04PM", "Thu, 05/19/11, 10:47PM", "Asia/Shanghai")
	assert.Equal(t, r.Format(time.RFC3339), "2011-05-19T23:59:59+08:00")
}

func Test_Yesterday(t *testing.T) {
	t.Log("Yesterday is:", Yesterday("Asia/Shanghai", "2006-01-02"))
}

func Test_Today(t *testing.T) {
	t.Log("Today is:", Today("Asia/Shanghai", "2006-1-2"))
}
