package tools

import (
	"time"
)

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
