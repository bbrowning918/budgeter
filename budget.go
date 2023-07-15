package main

import (
	"fmt"
)

const (
	needsTarget   = 50
	wantsTarget   = 30
	savingsTarget = 20
)

func Target(totalIncome int, c Category) (int, error) {
	if totalIncome < 0 {
		return 0, fmt.Errorf("income cannot be less than 0")
	}
	switch c {
	case Income:
		return 0, fmt.Errorf("cannot find target for income")
	case Needs:
		return (needsTarget * totalIncome) / 100, nil
	case Wants:
		return (wantsTarget * totalIncome) / 100, nil
	case Savings:
		return (savingsTarget * totalIncome) / 100, nil
	default:
		return 0, fmt.Errorf("no target match for category %d", c)
	}
}
