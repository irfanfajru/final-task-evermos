package utils

import "time"

// @TODO : make function parsing date
func DateTimeToDate(dateTime string) time.Time {
	date, _ := time.Parse("2006-01-02", dateTime)
	return date
}

func DateResponse(time time.Time) string {
	return time.Format("01/02/2006")
}
