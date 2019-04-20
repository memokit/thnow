package thnow

import (
	"strings"
	"time"
)

// ISO8601Date uses ISO 8601 as a default for parsing and rendering
const ISO8601Date = "2006-01-02"

// DateNow struct
type DateNow struct {
	time.Time
}

// StringNow struct
type StringNow struct {
	string
}

// Date initialize Now with date time
func Date(t time.Time) *DateNow {
	return &DateNow{t}
}

// String initialize Now with date string
func String(str string) *StringNow {
	return &StringNow{strings.TrimSpace(str)}
}
