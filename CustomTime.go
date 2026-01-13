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
// Supports multiple formats:
// - "2006-01-02T15:04:05" (original format)
// - "2006-01-02T15:04:05Z" (RFC3339 with Z)
// - "2006-01-02T15:04:05+00:00" (RFC3339 with offset)
// - time.RFC3339 (full RFC3339 format)
func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return nil
	}

	// Try multiple formats
	formats := []string{
		time.RFC3339,                // "2006-01-02T15:04:05Z07:00"
		"2006-01-02T15:04:05Z",      // "2006-01-02T15:04:05Z"
		"2006-01-02T15:04:05+00:00", // "2006-01-02T15:04:05+00:00"
		ctLayout,                    // "2006-01-02T15:04:05" (original)
	}

	var err error
	for _, layout := range formats {
		parsedTime, parseErr := time.Parse(layout, s)
		if parseErr == nil {
			ct.Time = parsedTime
			return nil
		}
		err = parseErr
	}

	// If all formats failed, return the last error
	return fmt.Errorf("failed to parse time %q: %v", s, err)
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
