# A monthly budget in Go
#### _I wanted an excuse to play with Go_

---
Using `book.csv.example` where each line is:
 
- a description for that line item
- a category of either
  - income
  - needs
  - wants
  - savings
- an amount as an integer (representing cents, ex: $1 would be 100)

---

 Will calculate and display:
1. the overall balance of income minus needs, wants, and savings
2. totals for each category
3. baseline targets for each category
   1. needs as 50% of total income
   2. wants as 30% of total income
   3. savings as 20% of total income