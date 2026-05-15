package helper

import "time"

func MustConvertStringToTime(timeString string) time.Time {
	loc, _ := time.LoadLocation("Asia/Bangkok")

	timeNew, _ := time.ParseInLocation(
		"2006-01-02 15:04:05",
		timeString,
		loc,
	)

	return timeNew
}
