package rfc5322_test

import (
	"rfc5322-go"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewDate(t *testing.T) {
	testCases := []struct {
		name     string
		input    time.Time
		expected string
	}{
		{
			name:     "valid date",
			input:    time.Date(2023, 10, 1, 12, 0, 0, 0, time.UTC),
			expected: "Sun, 01 Oct 2023 12:00:00 +0000",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			d := rfc5322.NewDate(tc.input)
			assert.Equal(t, tc.expected, d.String())
		})
	}
}
