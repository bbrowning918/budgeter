package main

import (
	"testing"
)

func TestCategoryToString(t *testing.T) {
	cases := []struct {
		name     string
		category category
		expected string
		errMsg   string
	}{
		{"income", income, "income", ""},
		{"needs", needs, "needs", ""},
		{"wants", wants, "wants", ""},
		{"savings", savings, "savings", ""},
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

func TestStringToCategory(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected category
		errMsg   string
	}{
		{"income case sensitive", "income", income, ""},
		{"needs case sensitive", "needs", needs, ""},
		{"wants case sensitive", "wants", wants, ""},
		{"savings case sensitive", "savings", savings, ""},
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

func TestLedgerTotalFor(t *testing.T) {
	ledger := ledger{
		{"income", income, 100000},
		{"need1", needs, 45000},
		{"need2", needs, 15000},
	}

	if result := ledger.totalFor(savings); result != 0 {
		t.Errorf("expected default of 0, got %d", result)
	}

	if result := ledger.totalFor(needs); result != 60000 {
		t.Errorf("expected needs total of 60000, got %d", result)
	}
}

func TestLedgerBalance(t *testing.T) {
	empty := ledger{}
	if result := empty.balance(); result != 0 {
		t.Errorf("expected default of 0, got %d", result)
	}

	ledger := ledger{
		{"income", income, 100000},
		{"need1", needs, 45000},
		{"need2", needs, 15000},
	}

	if result := ledger.balance(); result != 40000 {
		t.Errorf("expected balance of 40000, got %d", result)
	}
}
