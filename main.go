package thnow

import "time"

// ISO8601Date uses ISO 8601 as a default for parsing and rendering
const ISO8601Date = "2006-01-02"

type THNow struct {
	time.Time
}

// New initialize Now with time
func New(t time.Time) *THNow {
	return &THNow{t}
}

// BeginningOfMinute beginning of minute
func BeginningOfMinute() time.Time {
	return New(time.Now()).BeginningOfMinute()
}
