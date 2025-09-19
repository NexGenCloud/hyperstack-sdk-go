package hyperstack

import (
	"fmt"
	"strings"
	"time"
)

type CustomTime struct {
	time.Time
}

const ctLayout = "2006-01-02T15:04:05"

// UnmarshalJSON converts JSON to CustomTime
func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return nil
	}
	parsedTime, err := time.Parse(ctLayout, s)
	if err != nil {
		return err
	}
	ct.Time = parsedTime
	return nil
}

// MarshalJSON converts CustomTime to JSON
func (ct CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format(ctLayout))), nil
}

// NewCustomTime creates a CustomTime from time.Time
func NewCustomTime(t time.Time) CustomTime {
	return CustomTime{Time: t}
}

// Now returns the current time as CustomTime
func Now() CustomTime {
	return CustomTime{Time: time.Now()}
}

// ParseCustomTime parses a string into CustomTime
func ParseCustomTime(layout, value string) (CustomTime, error) {
	t, err := time.Parse(layout, value)
	return CustomTime{Time: t}, err
}
