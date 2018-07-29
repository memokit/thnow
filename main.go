package thnow

import "time"

// ISO8601Date uses ISO 8601 as a default for parsing and rendering
const ISO8601Date = "2006-01-02"

type THNow struct {
	time.Time
}
