package util

import (
	"fmt"
	"time"
)

type Time interface {
	Now(timeGMT *int) time.Time
	GenerateDateOfBirthFromString(s string) (time.Time, error)
}

type timesCustomImpl struct {
	gmt int
}

func ProvideNewTimesCustom() Time {
	// default GMT +7 (Asia/Jakarta)
	return &timesCustomImpl{gmt: 7}
}

func (t *timesCustomImpl) Now(timeGMT *int) time.Time {
	return time.Now().UTC()
}

func (t *timesCustomImpl) GenerateDateOfBirthFromString(s string) (time.Time, error) {
	// Parse the date string using Go's built-in time package
	// The format string used in this case assumes a date in the format "YYYY-MM-DD"
	date, err := time.Parse("2006-01-02", s)
	if err != nil {
		// If the date string couldn't be parsed, return a zero-value time.Time and the error
		return time.Time{}, fmt.Errorf("invalid date string: %v", err)
	}

	return date, nil
}
