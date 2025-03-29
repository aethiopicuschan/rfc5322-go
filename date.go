package rfc5322

import "time"

// Date represents a date as per RFC 5322.
type Date struct {
	value string
}

// NewDate creates a new Date instance with the given time.Time value.
func NewDate(t time.Time) (d *Date) {
	d = &Date{
		value: t.Format(time.RFC1123Z),
	}
	return
}

// String returns the string representation of the Date.
func (d *Date) String() string {
	return d.value
}
