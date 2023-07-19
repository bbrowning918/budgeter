package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
)

type category int

const (
	income category = iota
	needs
	wants
	savings
)

func (c category) toString() (string, error) {
	switch c {
	case income:
		return "income", nil
	case needs:
		return "needs", nil
	case wants:
		return "wants", nil
	case savings:
		return "savings", nil
	default:
		return "", fmt.Errorf("no toString() match for category %d", c)
	}
}

func toCategory(s string) (category, error) {
	switch s {
	case "income":
		return income, nil
	case "needs":
		return needs, nil
	case "wants":
		return wants, nil
	case "savings":
		return savings, nil
	default:
		return 0, fmt.Errorf("no category match for string %s", s)
	}
}

type line struct {
	name     string
	category category
	amount   int
}

type ledger []line

func (l ledger) totalFor(c category) int {
	total := 0
	for _, item := range l {
		if item.category == c {
			total += item.amount
		}
	}
	return total
}

func (l ledger) balance() int {
	balance := 0
	for _, item := range l {
		if item.category == income {
			balance += item.amount
		} else {
			balance -= item.amount
		}
	}
	return balance
}

func parse(i io.Reader) (ledger, error) {
	r := csv.NewReader(i)
	var l ledger
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return l, nil
		}

		if len(record) != 3 {
			return l, fmt.Errorf("line does not have 3 values, got %s", record)
		}

		amount, err := strconv.Atoi(record[2])
		if err != nil {
			return l, err
		}

		c, err := toCategory(record[1])
		if err != nil {
			return l, err
		}

		l = append(l, line{record[0], c, amount})
	}
	return l, nil
}
