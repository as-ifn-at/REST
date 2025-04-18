package common

import (
	"fmt"
	"time"
)

// Has to be DD-MM-YYYY
func IsValidFormat(date string) bool {
	_, err := time.Parse(Layout, date)
	return err == nil
}

func ConvertStrToDate(dateStr string) (time.Time, error) {
	date, err := time.Parse(Layout, dateStr)
	if err != nil {
		return time.Time{}, err
	}

	return date, nil
}

func CheckValidStartEndDate(start, end string) error {
	startDate, err := time.Parse(Layout, start)
	if err != nil {
		return err
	}

	if !startDate.After(time.Now()) {
		return fmt.Errorf("invalid date: start date cannot be before current date")
	}

	endDate, err := time.Parse(Layout, end)
	if err != nil {
		return err
	}

	if startDate.Before(endDate) {
		return nil
	}

	return fmt.Errorf("end Date cannot be greater than start date")
}

func CheckClassAvailability(start, end, date string) error {

	var err error
	startDate, err := time.Parse(Layout, start)
	if err != nil {
		return err
	}
	endDate, err := time.Parse(Layout, end)
	if err != nil {
		return err
	}
	currentDate, err := time.Parse(Layout, date)
	if err != nil {
		return err
	}
	if !currentDate.After(time.Now()) {
		return fmt.Errorf("invalid date: date of booking cannot be before current date")
	}
	if currentDate.After(startDate) && currentDate.Before(endDate) {
		return nil
	}

	return fmt.Errorf("date not available for booking")
}
