package main

import (
	"testing"
)

func TestTarget(t *testing.T) {
	cases := []struct {
		name        string
		totalIncome int
		category    category
		expected    int
		errMsg      string
	}{
		{"cannot find target for income", 10000, income, 0, "cannot find target for income"},
		{"no negative income", -1, needs, 0, "income cannot be less than 0"},
		{"needs calculates to 50%", 10000, needs, 5000, ""},
		{"wants calculates to 30%", 10000, wants, 3000, ""},
		{"savings calculates to 20%", 10000, savings, 2000, ""},
		{"no category match", 10000, 5, 0, "no target match for: 5"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result, err := target(c.totalIncome, c.category)

			if c.expected != result {
				t.Errorf("expected '%d', got '%d'", c.expected, result)
			}

			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if c.errMsg != errMsg {
				t.Errorf("expectedf error '%s', got '%s'", c.errMsg, errMsg)
			}
		})
	}
}
