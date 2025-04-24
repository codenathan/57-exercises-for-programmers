package main

import (
	"fmt"
	numbers "github.com/codenathan/exercises-for-programmers/00-numbers"
	"math"
)

func main() {
	principal := numbers.GetPositiveFloat("Enter the principal amount: ")
	rate := numbers.GetPositiveFloat("Enter the rate of interest: ")
	years := numbers.GetPositiveFloat("Enter the number of years: ")
	periods := numbers.GetPositiveFloat("Enter the number of times the interest is compounded per year: ")

	amount := principal * math.Pow(1+((rate/100)/periods), periods*years)

	// Output
	fmt.Printf("After %.0f years, the investment will be worth $%.2f\n", years, amount)
}
