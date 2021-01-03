package time

import "time"

// Clock represents struct that handle current time.
type Clock struct{}

// NewClock generates Clock.
func NewClock() *Clock {
	return &Clock{}
}

// Now returns current time.
func (c *Clock) Now() time.Time {
	return time.Now()
}
