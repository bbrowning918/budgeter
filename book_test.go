package main

import (
	"testing"
)

func TestCategory_toString(t *testing.T) {
	cases := []struct {
		name     string
		category Category
		expected string
		errMsg   string
	}{
		{"income", Income, "income", ""},
		{"needs", Needs, "needs", ""},
		{"wants", Wants, "wants", ""},
		{"savings", Savings, "savings", ""},
		{"no match", -1, "", "no toString() match for category -1"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result, err := c.category.toString()
			if result != c.expected {
				t.Errorf("expected '%s', got '%s'", c.expected, result)
			}

			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if errMsg != c.errMsg {
				t.Errorf("expected error '%s', got error '%s'", c.errMsg, errMsg)
			}

		})
	}
}

func Test_toCategory(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected Category
		errMsg   string
	}{
		{"income case sensitive", "income", Income, ""},
		{"needs case sensitive", "needs", Needs, ""},
		{"wants case sensitive", "wants", Wants, ""},
		{"savings case sensitive", "savings", Savings, ""},
		{"no match", "wrong", 0, "no category match for string wrong"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result, err := toCategory(c.input)

			if result != c.expected {
				t.Errorf("expected %c, got %c", c.expected, result)
			}

			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if errMsg != c.errMsg {
				t.Errorf("expected error '%s', got '%s'", c.errMsg, errMsg)
			}
		})
	}
}

func TestLedger_totalFor(t *testing.T) {
	ledger := Ledger{
		{"income", Income, 100000},
		{"need1", Needs, 45000},
		{"need2", Needs, 15000},
	}

	if result := ledger.totalFor(Savings); result != 0 {
		t.Errorf("expected default of 0, got %d", result)
	}

	if result := ledger.totalFor(Needs); result != 60000 {
		t.Errorf("expected needs total of 60000, got %d", result)
	}
}

func TestLedger_balance(t *testing.T) {
	empty := Ledger{}
	if result := empty.balance(); result != 0 {
		t.Errorf("expected default of 0, got %d", result)
	}

	ledger := Ledger{
		{"income", Income, 100000},
		{"need1", Needs, 45000},
		{"need2", Needs, 15000},
	}

	if result := ledger.balance(); result != 40000 {
		t.Errorf("expected balance of 40000, got %d", result)
	}
}
