package main

import (
	"testing"
)

func TestTarget(t *testing.T) {
	empty := ledger{}
	l := ledger{
		{"income", income, 10000},
	}

	cases := []struct {
		name     string
		ledger   ledger
		category category
		expected string
		errMsg   string
	}{
		{"cannot find target for income", l, income, "", "cannot find target for income"},
		{"no negative income", empty, needs, "", "income cannot be 0 or negative"},
		{"needs calculates to 50%", l, needs, "needs\t$50.00\t", ""},
		{"wants calculates to 30%", l, wants, "wants\t$30.00\t", ""},
		{"savings calculates to 20%", l, savings, "savings\t$20.00\t", ""},
		{"no category match", l, 5, "", "no target match for: 5"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result, err := target(c.ledger, c.category)

			if c.expected != result {
				t.Errorf("expected '%s', got '%s'", c.expected, result)
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
