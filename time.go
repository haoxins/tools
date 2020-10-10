package tools

import (
	"time"
)

//
// `zone` is always required in case you forget time zone in your business
// `Mon Jan 2 15:04:05 -0700 MST 2006`
//

// ParseTime - parse time and location
func ParseTime(layout, date, zone string) (time.Time, *time.Location) {
	t, e := time.Parse(layout, date)
	AssertError(e)
	local, e := time.LoadLocation(zone)
	AssertError(e)

	return t, local
}

// StartOfDate - return start time of date
func StartOfDate(layout, date, zone string) time.Time {
	t, local := ParseTime(layout, date, zone)
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, local)
}

// EndOfDate - return end time of date
func EndOfDate(layout, date, zone string) time.Time {
	t, local := ParseTime(layout, date, zone)
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, local)
}

// Yesterday - return yesterday date
func Yesterday(zone string, format string) string {
	if zone == "" {
		zone = "Asia/Shanghai"
	}
	if format == "" {
		format = "2006-01-02"
	}

	local, err := time.LoadLocation(zone)
	AssertError(err)

	return time.Now().In(local).AddDate(0, 0, -1).Format(format)
}

// Today - return today date
func Today(zone string, format string) string {
	if zone == "" {
		zone = "Asia/Shanghai"
	}
	if format == "" {
		format = "2006-01-02"
	}

	local, err := time.LoadLocation(zone)
	AssertError(err)

	return time.Now().In(local).Format(format)
}
