package tools

import (
	"time"
)

const (
	// TimeZoneShanghai - The time zone Shanghai
	TimeZoneShanghai = "Asia/Shanghai"
	// DateFmt1 - 2006-01-02
	DateFmt1 = "2006-01-02"
	// DateFmt2 - 2006-1-2
	DateFmt2 = "2006-1-2"
)

//
// `zone` is always required in case you forget time zone in your business
// `Mon Jan 2 15:04:05 -0700 MST 2006`
//

// ParseTime - parse time and location
func ParseTime(zone, format, date string) (time.Time, *time.Location) {
	t, e := time.Parse(format, date)
	PanicError(e)
	local, e := time.LoadLocation(zone)
	PanicError(e)

	return t, local
}

// StartOfDate - return start time of date
func StartOfDate(zone, format, date string) time.Time {
	t, local := ParseTime(zone, format, date)
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, local)
}

// EndOfDate - return end time of date
func EndOfDate(zone, format, date string) time.Time {
	t, local := ParseTime(zone, format, date)
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, local)
}

// Yesterday - return yesterday date
func Yesterday(zone string, format string) string {
	if zone == "" {
		zone = TimeZoneShanghai
	}
	if format == "" {
		format = DateFmt1
	}

	local, err := time.LoadLocation(zone)
	PanicError(err)

	return time.Now().In(local).AddDate(0, 0, -1).Format(format)
}

// Today - return today date
func Today(zone string, format string) string {
	if zone == "" {
		zone = TimeZoneShanghai
	}
	if format == "" {
		format = DateFmt1
	}

	local, err := time.LoadLocation(zone)
	PanicError(err)

	return time.Now().In(local).Format(format)
}
