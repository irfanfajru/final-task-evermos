package utils

import "time"

// @TODO : make function parsing date
func DateTimeToDate(dateTime string) time.Time {
	date, _ := time.Parse("YYYY-MM-DD", dateTime)
	return date
}
