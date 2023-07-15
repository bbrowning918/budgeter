package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	fd, err := os.Open("book.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()

	reader := csv.NewReader(fd)
	l, err := build(reader)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("-- balance --")
	balance := l.balance()
	fmt.Println(toString(balance))

	fmt.Println("-- budget --")
	for _, category := range []Category{Needs, Wants, Savings} {
		total := l.totalFor(category)

		name, err := category.toString()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(name, ":", toString(total))
	}

	income := l.totalFor(Income)
	fmt.Println("-- targets --")
	for _, category := range []Category{Needs, Wants, Savings} {
		target, err := Target(income, category)
		if err != nil {
			log.Fatal(err)
		}

		categoryName, err := category.toString()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(categoryName, ":", toString(target))
	}
}

func build(r *csv.Reader) (Ledger, error) {
	l := Ledger{}
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return l, err
		}

		amount, err := strconv.Atoi(record[2])
		if err != nil {
			return l, err
		}
		category, err := toCategory(record[1])
		if err != nil {
			return l, err
		}
		l = append(l, Line{record[0], category, amount})
	}

	return l, nil
}

func toString(amount int) string {
	decimal := 2
	thousand := ","

	s := strconv.FormatInt(int64(amount), 10)
	for i := len(s) - decimal - 3; i > 0; i -= 3 {
		s = s[:i] + thousand + s[i:]
	}

	s = "$" + s[:len(s)-decimal] + "." + s[len(s)-decimal:]
	return s
}
