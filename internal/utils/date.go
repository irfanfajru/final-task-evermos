package utils

import "time"

// @TODO : make function parsing date
func ParseDate(date string) time.Time {
	dateParsed, _ := time.Parse("02/01/2006", date)
	return dateParsed
}

func DateResponse(time time.Time) string {
	return time.Format("01/02/2006")
}
