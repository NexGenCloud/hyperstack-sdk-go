package PACKAGE_REPLACE

import (
	"fmt"
	"strings"
	"time"
)

type CustomTime struct {
	time.Time
}

const ctLayout = "2006-01-02T15:04:05" // Specify your time format here

func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	ct.Time, err = time.Parse(ctLayout, s)
	return
}

func (ct *CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format(ctLayout))), nil
}
