package main

import (
	"testing"
	"time"

	"github.com/Giully314/snippetbox/internal/assert"
)

func TestHumanDate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		tm       time.Time
		expected string
	}{
		{
			name:     "UTC",
			tm:       time.Date(2024, 3, 17, 10, 15, 0, 0, time.UTC),
			expected: "17 Mar 2024 at 10:15",
		},
		{
			name:     "Empty",
			tm:       time.Time{},
			expected: "",
		},
		{
			name:     "CET",
			tm:       time.Date(2024, 3, 17, 10, 15, 0, 0, time.FixedZone("CET", 1*60*60)),
			expected: "17 Mar 2024 at 09:15",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.tm)
			assert.Equal(t, hd, tt.expected)
		})
	}

}
