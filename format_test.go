package main

import (
	"testing"
)

func TestAmountToString(t *testing.T) {
	cases := []struct {
		name     string
		amount   int
		expected string
	}{
		{"negative", -1, "-$0.01"},
		{"zero", 0, "$0.00"},
		{"couple cents", 2, "$0.02"},
		{"a dollar", 100, "$1.00"},
		{"almost sixteen bucks", 1599, "$15.99"},
		{"couple hundred", 45130, "$451.30"},
		{"two grand", 200000, "$2,000.00"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := toString(c.amount)
			if c.expected != result {
				t.Errorf("expected '%s', got '%s", c.expected, result)
			}
		})
	}
}
