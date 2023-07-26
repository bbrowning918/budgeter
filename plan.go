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
	// TODO this is awkward, should pass in income totaled already
	totalIncome := 0
	for _, item := range l {
		if item.category == income {
			totalIncome += item.amount
		}
	}
	if totalIncome <= 0 {
		return "", fmt.Errorf("income cannot be less than 0")
	}

	switch c {
	case income:
		return "", fmt.Errorf("cannot find target for income")
	case needs:
		target := (needsTarget * totalIncome) / 100
		return fmt.Sprintf("%s\t%s\t", c.toString(), toString(target)), nil
	case wants:
		target := (wantsTarget * totalIncome) / 100
		return fmt.Sprintf("%s\t%s\t", c.toString(), toString(target)), nil
	case savings:
		target := (savingsTarget * totalIncome) / 100
		return fmt.Sprintf("%s\t%s\t", c.toString(), toString(target)), nil
	default:
		return "", fmt.Errorf("no target match for: %d", c)
	}
}
