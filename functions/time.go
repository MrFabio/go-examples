package functions

import (
	"time"
)

// TimeExamples demonstrates time operations in Go
// including current time, formatting, and time arithmetic
func TimeExamples() {
	// Get current time
	now := time.Now()
	pl("The year is", now.Year())
	pl("The month is", now.Month())
	pl("The day is", now.Day())
	nowUTC := now.UTC()
	pl("The Local time is", now)
	pl("The UTC time is", nowUTC)
	loc, _ := time.LoadLocation("Europe/Paris")
	nowParis := now.In(loc)
	pl("The time in Paris is", nowParis)
	// print in ISO 8601 format
	pl("The time in ISO 8601 format is", now.Format(time.RFC3339))
}
