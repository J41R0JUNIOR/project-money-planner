package domain

import "time"

// Occurrences == nil && EndDate == nil → no end date (infinite).
// Occurrences != nil && EndDate == nil → finish after N occurrences.
// Occurrences == nil && EndDate != nil → terminate on date.
// Occurrences != nil && EndDate != nil → invalid.
type Recurrence struct {
	Frequency   RecurrenceFrequency
	Interval    int
	Occurrences *int
	EndDate     *time.Time
}

type RecurrenceFrequency string

const (
	Daily   RecurrenceFrequency = "DAILY"
	Weekly  RecurrenceFrequency = "WEEKLY"
	Monthly RecurrenceFrequency = "MONTHLY"
	Yearly  RecurrenceFrequency = "YEARLY"
)