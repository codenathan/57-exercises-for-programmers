package main

import (
	"fmt"
	numbers "github.com/codenathan/exercieses-for-programmers/00-numbers"
	"math"
)

func main() {
	euros := numbers.GetPositiveFloat("How many euros are you exchanging? ")
	exchangeRate := numbers.GetPositiveFloat("What is the exchange rate? ")

	dollar := euros * exchangeRate / 100
	dollar = math.Ceil(dollar*100) / 100
	fmt.Printf("%.0f euros at an exchange rate of %.2f is %.2f U.S. dollars.\n", euros, exchangeRate, dollar)

}
