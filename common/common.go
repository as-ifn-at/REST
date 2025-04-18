package common

import (
	"fmt"
	"time"
)

func ConvertStrToDate(date string) (int, int, int, error) {
	parsedTime, err := time.Parse(Layout, date)
	if err != nil {
		return -1, -1, -1, fmt.Errorf("invalid date cannot be parsed in the format")
	}
	day := parsedTime.Day()
	year := parsedTime.Year()
	month := parsedTime.Month()

	return day, int(month), year, nil
}
