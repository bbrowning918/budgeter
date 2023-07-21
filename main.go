package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f := flag.String("f", "budget.csv.example", "csv file to load")
	flag.Parse()

	file, err := os.Open(*f)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	l, err := parse(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("-- balance --")
	balance := l.balance()
	fmt.Println(toString(balance))

	fmt.Println("-- budget --")
	for _, category := range []category{needs, wants, savings} {
		total := l.totalFor(category)

		fmt.Println(category.toString(), ":", toString(total))
	}

	income := l.totalFor(income)
	fmt.Println("-- targets --")
	for _, category := range []category{needs, wants, savings} {
		target, err := target(income, category)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(category.toString(), ":", toString(target))
	}
}

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
