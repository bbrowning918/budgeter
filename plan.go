package main

import (
	"fmt"
)

const (
	needsTarget   = 50
	wantsTarget   = 30
	savingsTarget = 20
)

func target(totalIncome int, c category) (int, error) {
	if totalIncome < 0 {
		return 0, fmt.Errorf("income cannot be less than 0")
	}
	switch c {
	case income:
		return 0, fmt.Errorf("cannot find target for income")
	case needs:
		return (needsTarget * totalIncome) / 100, nil
	case wants:
		return (wantsTarget * totalIncome) / 100, nil
	case savings:
		return (savingsTarget * totalIncome) / 100, nil
	default:
		return 0, fmt.Errorf("no target match for: %d", c)
	}
}
