package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestCategoryToString(t *testing.T) {
	cases := []struct {
		name     string
		category category
		expected string
	}{
		{"zero is not a valid category", 0, "unknown"},
		{"income", income, "income"},
		{"needs", needs, "needs"},
		{"wants", wants, "wants"},
		{"savings", savings, "savings"},
		{"no match", 5, "unknown"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := c.category.toString()
			if c.expected != result {
				t.Errorf("expected: '%s', received: '%s'", c.expected, result)
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
		{"no match", "wrong", 0, "no category match for: 'wrong'"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result, err := toCategory(c.input)

			if c.expected != result {
				t.Errorf("expected: '%d', received: '%d'", c.expected, result)
			}

			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if c.errMsg != errMsg {
				t.Errorf("expected error: '%s', received: '%s'", c.errMsg, errMsg)
			}
		})
	}
}

func TestLedgerTotalFor(t *testing.T) {
	ledger := ledger{
		{"income", income, 1000.00},
		{"need1", needs, 45000},
		{"need2", needs, 15000},
	}

	if result := ledger.totalFor(savings); result != "savings\t$0.00\t" {
		t.Errorf("expected: 'savings\t$0.00\t', received: '%s'", result)
	}

	if result := ledger.totalFor(needs); result != "needs\t$600.00\t" {
		t.Errorf("expected: 'needs\t$600.00\t', received: '%s'", result)
	}
}

func TestLedgerBalance(t *testing.T) {
	empty := ledger{}
	if result := empty.balance(); result != "balance\t$0.00\t" {
		t.Errorf("expected: 'balance\t$0.00\t', received: '%s'", result)
	}

	ledger := ledger{
		{"income", income, 100000},
		{"need1", needs, 45000},
		{"need2", needs, 15000},
	}

	if result := ledger.balance(); result != "balance\t$400.00\t" {
		t.Errorf("expected: 'balance\t$400.00\t', received: '%s'", result)
	}
}

func TestParse(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected ledger
		errMsg   string
	}{
		{"line too short", "job,income", nil, "line does not have 3 values: '[job income]'"},
		{"line too long", "job,income,500,100", nil, "line does not have 3 values: '[job income 500 100]'"},
		{"amount is not int", "job,income,loads", nil, "could not parse as int: 'loads'"},
		{"unknown category", "job,money,1000", nil, "could not parse as category: 'money', no category match for: 'money'"},
		{"valid line", "job,income,10000", ledger{{"job", income, 10000}}, ""},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result, err := parse(strings.NewReader(c.input))
			if !reflect.DeepEqual(c.expected, result) {
				t.Errorf("expected: '%+v', received: '%+v'", c.expected, result)
			}

			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if c.errMsg != errMsg {
				t.Errorf("expected error: '%s', received: '%s'", c.errMsg, errMsg)
			}
		})
	}
}
