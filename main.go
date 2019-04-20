package thnow

import (
	"strings"
	"time"
)

// ISO8601Date uses ISO 8601 as a default for parsing and rendering
const ISO8601Date = "2006-01-02"

type DateNow struct {
	time.Time
}

type StringNow struct {
	string
}

// New initialize Now with time
func Date(t time.Time) *DateNow {
	return &DateNow{t}
}

func String(str string) *StringNow {
	return &StringNow{strings.TrimSpace(str)}
}

// BeginningOfMinute beginning of minute.
// func BeginningOfMinute() time.Time {
// 	return New(time.Now()).BeginningOfMinute()
// }
