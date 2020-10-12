package tools

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_ParseTime(t *testing.T) {
	r, local := ParseTime("Mon, 01/02/06, 03:04PM", "Thu, 05/19/11, 10:47PM", TimeZoneShanghai)
	assert.Equal(t, "2011-05-19T22:47:00Z", r.Format(time.RFC3339))
	assert.Equal(t, TimeZoneShanghai, local.String())

	r, local = ParseTime(DateFmt1, "2020-10-10", TimeZoneShanghai)
	assert.Equal(t, "2020-10-10T00:00:00Z", r.Format(time.RFC3339))
	assert.Equal(t, TimeZoneShanghai, local.String())
}

func Test_StartOfDate(t *testing.T) {
	r := StartOfDate("Mon, 01/02/06, 03:04PM", "Thu, 05/19/11, 10:47PM", TimeZoneShanghai)
	assert.Equal(t, "2011-05-19T00:00:00+08:00", r.Format(time.RFC3339))

	r = StartOfDate(DateFmt1, "2020-10-10", TimeZoneShanghai)
	assert.Equal(t, "2020-10-10T00:00:00+08:00", r.Format(time.RFC3339))
}

func Test_EndOfDate(t *testing.T) {
	r := EndOfDate("Mon, 01/02/06, 03:04PM", "Thu, 05/19/11, 10:47PM", TimeZoneShanghai)
	assert.Equal(t, "2011-05-19T23:59:59+08:00", r.Format(time.RFC3339))

	r = EndOfDate(DateFmt1, "2020-10-10", TimeZoneShanghai)
	assert.Equal(t, "2020-10-10T23:59:59+08:00", r.Format(time.RFC3339))
}

func Test_Yesterday(t *testing.T) {
	t.Log("Yesterday is:", Yesterday(TimeZoneShanghai, DateFmt1))
}

func Test_Today(t *testing.T) {
	t.Log("Today is:", Today(TimeZoneShanghai, DateFmt2))
}
