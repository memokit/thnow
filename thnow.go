package thnow

import (
	"time"
)

// BeginningOfMinute beginning of minute
func (thnow *THNow) BeginningOfMinute() time.Time {
	return thnow.Truncate(time.Minute)
}
