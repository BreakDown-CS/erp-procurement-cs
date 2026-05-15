package helper

import "time"

func MustConvertStringToTime(timeString string) time.Time {
	loc, _ := time.LoadLocation("Asia/Bangkok")

	timeConvert, _ := time.ParseInLocation("2006-01-02", timeString, loc)

	return timeConvert
}
