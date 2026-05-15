package helper

import (
	"fmt"
	"time"
)

func ConvertStringToTime(timeString string) (time.Time, error) {
	loc, _ := time.LoadLocation("Asia/Bangkok")

	layouts := []string{
		"2006-01-02 15:04:05",
		time.RFC3339,
	}

	for _, layout := range layouts {
		if t, err := time.ParseInLocation(layout, timeString, loc); err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("invalid time format: %s", timeString)
}
