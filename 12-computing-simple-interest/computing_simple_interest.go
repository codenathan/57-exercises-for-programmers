package main

import (
	"fmt"
	numbers "github.com/codenathan/exercises-for-programmers/00-numbers"
)

func main() {
	principal := numbers.GetPositiveFloat("Enter the principal: ")
	interest := numbers.GetPositiveFloat("Enter the rate of interest: ")
	years := numbers.GetPositiveInt("Enter the number of years: ")

	investment := principal * (1 + (interest/100)*float64(years))

	fmt.Printf("After %d years, the investment will be worth $%.2f\n", years, investment)
}
