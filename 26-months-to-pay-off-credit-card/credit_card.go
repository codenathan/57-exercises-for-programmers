package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("What is your balance? ")
	balanceInput, _ := reader.ReadString('\n')
	balanceInput = strings.TrimSpace(balanceInput)
	balance, _ := strconv.ParseFloat(balanceInput, 64)

	fmt.Print("What is the APR on the card (as a percent)? ")
	aprInput, _ := reader.ReadString('\n')
	aprInput = strings.TrimSpace(aprInput)
	apr, _ := strconv.ParseFloat(aprInput, 64)

	fmt.Print("What is the monthly payment you can make? ")
	paymentInput, _ := reader.ReadString('\n')
	paymentInput = strings.TrimSpace(paymentInput)
	monthlyPayment, _ := strconv.ParseFloat(paymentInput, 64)

	months := calculateMonthsUntilPaidOff(balance, apr, monthlyPayment)

	fmt.Printf("It will take you %d months to pay off this card.\n", months)

}

func calculateMonthsUntilPaidOff(balance float64, apr float64, payment float64) int {
	dailyRate := apr / 100 / 365
	i := dailyRate
	b := balance
	p := payment

	n := -1.0 / 30.0 * math.Log(1+b/p*(1-math.Pow(1+i, 30))) / math.Log(1+i)
	return int(math.Ceil(n))
}
