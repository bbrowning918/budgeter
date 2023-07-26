package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
)

func main() {
	f := flag.String("file", "budget.csv.example", "csv file to load")
	flag.Parse()

	file, err := os.Open(*f)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	l, err := parse(file)
	if err != nil {
		log.Fatal(err)
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, l.balance())

	fmt.Fprintln(w, "----- budget -----")
	fmt.Fprintln(w, "category\tamount\t")
	for _, category := range []category{income, needs, wants, savings} {
		fmt.Fprintln(w, l.totalFor(category))
	}

	fmt.Fprintln(w, "----- targets -----")
	fmt.Fprintln(w, "category\tamount\t")
	for _, category := range []category{needs, wants, savings} {
		target, err := target(l, category)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintln(w, target)
	}

	w.Flush()
}
