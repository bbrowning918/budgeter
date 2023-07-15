package main

import (
	"testing"
)

func TestTarget(t *testing.T) {
	cases := []struct {
		name        string
		totalIncome int
		category    Category
		expected    int
		errMsg      string
	}{
		{"cannot find target for income", 10000, Income, 0, "cannot find target for income"},
		{"no negative income", -1, Needs, 0, "income cannot be less than 0"},
		{"needs calculates to 50%", 10000, Needs, 5000, ""},
		{"wants calculates to 30%", 10000, Wants, 3000, ""},
		{"savings calculates to 20%", 10000, Savings, 2000, ""},
		{"no category match", 10000, 5, 0, "no target match for category 5"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result, err := Target(c.totalIncome, c.category)

			if result != c.expected {
				t.Errorf("expected %c, got %c", c.expected, result)
			}

			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if errMsg != c.errMsg {
				t.Errorf("expectedf error '%s', got '%s'", c.errMsg, errMsg)
			}
		})
	}
}
