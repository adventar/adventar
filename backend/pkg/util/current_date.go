package util

import (
	"os"
	"time"

	"github.com/m-mizutani/goerr"
)

type Date struct {
	Year  int
	Month int
	Day   int
}

func CurrentDate() (*Date, error) {
	currentDate := os.Getenv("CURRENT_DATE")
	var t time.Time
	var err error
	if currentDate != "" {
		t, err = time.Parse("2006-01-02 15:04:05", currentDate)
		if err != nil {
			return nil, goerr.Wrap(err, "Failed to parse CURRENT_DATE").With("currentDate", currentDate)
		}
	} else {
		t = time.Now()
	}
	return &Date{Year: t.Year(), Month: int(t.Month()), Day: t.Day()}, nil
}
