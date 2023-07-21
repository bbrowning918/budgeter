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

func (c category) toString() string {
	switch c {
	case income:
		return "income"
	case needs:
		return "needs"
	case wants:
		return "wants"
	case savings:
		return "savings"
	default:
		return "unknown"
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
		return 0, fmt.Errorf("no category match for: %s", s)
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
			return nil, err
		}

		if len(record) != 3 {
			return nil, fmt.Errorf("line does not have 3 values: %s", record)
		}

		amount, err := strconv.Atoi(record[2])
		if err != nil {
			return nil, fmt.Errorf("could not parse as int: %s", record[2])
		}

		c, err := toCategory(record[1])
		if err != nil {
			return nil, err
		}

		l = append(l, line{record[0], c, amount})
	}
	return l, nil
}
