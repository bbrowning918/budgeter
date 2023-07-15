package main

import (
	"fmt"
)

type Category int

const (
	Income Category = iota
	Needs
	Wants
	Savings
)

func (c Category) toString() (string, error) {
	switch c {
	case Income:
		return "income", nil
	case Needs:
		return "needs", nil
	case Wants:
		return "wants", nil
	case Savings:
		return "savings", nil
	default:
		return "", fmt.Errorf("no toString() match for category %d", c)
	}
}

func toCategory(s string) (Category, error) {
	switch s {
	case "income":
		return Income, nil
	case "needs":
		return Needs, nil
	case "wants":
		return Wants, nil
	case "savings":
		return Savings, nil
	default:
		return 0, fmt.Errorf("no category match for string %s", s)
	}
}

type Line struct {
	name     string
	category Category
	amount   int
}

type Ledger []Line

func (l Ledger) totalFor(category Category) int {
	total := 0
	for _, item := range l {
		if item.category == category {
			total += item.amount
		}
	}
	return total
}

func (l Ledger) balance() int {
	balance := 0
	for _, item := range l {
		if item.category == Income {
			balance += item.amount
		} else {
			balance -= item.amount
		}
	}
	return balance
}
