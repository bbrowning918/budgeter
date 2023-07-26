package main

import (
	"fmt"
	"strconv"
)

func toString(amount int) string {
	if amount < 0 {
		return "$0.00"
	}
	if amount < 100 {
		return "$0." + fmt.Sprintf("%02d", amount)
	}

	decimal := 2
	thousand := ","

	s := strconv.FormatInt(int64(amount), 10)
	for i := len(s) - decimal - 3; i > 0; i -= 3 {
		s = s[:i] + thousand + s[i:]
	}

	s = "$" + s[:len(s)-decimal] + "." + s[len(s)-decimal:]
	return s
}
