package main

import (
	"fmt"
)

const (
	needsTarget   = 50
	wantsTarget   = 30
	savingsTarget = 20
)

func target(l ledger, c category) (string, error) {
	totalIncome := 0
	for _, item := range l {
		if item.category == income {
			totalIncome += item.amount
		}
	}
	if totalIncome <= 0 {
		return "", fmt.Errorf("income cannot be 0 or negative")
	}

	var target int
	switch c {
	case income:
		return "", fmt.Errorf("cannot find target for income")
	case needs:
		target = needsTarget
	case wants:
		target = wantsTarget
	case savings:
		target = savingsTarget
	default:
		return "", fmt.Errorf("no target match for: %d", c)
	}

	amount := (target * totalIncome) / 100
	return fmt.Sprintf("%s\t%s\t", c.toString(), toString(amount)), nil
}
