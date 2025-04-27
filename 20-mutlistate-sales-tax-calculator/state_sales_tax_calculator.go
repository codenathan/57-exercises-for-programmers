package main

import (
	"bufio"
	"fmt"
	numbers "github.com/codenathan/exercises-for-programmers/00-numbers"
	"math"
	"os"
	"strings"
)

func main() {
	var taxRate float64
	orderAmount := numbers.GetPositiveFloat("What is the order amount: ")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is the state? ")
	stateInput, _ := reader.ReadString('\n')
	state := strings.ToLower(strings.TrimSpace(stateInput))

	if state == "wi" || state == "wisconsin" {
		taxRate = 0.05

		fmt.Print("What county do you live in? ")
		countyInput, _ := reader.ReadString('\n')
		county := strings.ToLower(strings.TrimSpace(countyInput))

		if county == "eau claire" {
			taxRate += 0.005
		} else if county == "dunn" {
			taxRate += 0.004
		}
	} else if state == "il" || state == "illinois" {
		taxRate = 0.08
	}

	tax := roundUp(orderAmount * taxRate)
	total := orderAmount + tax

	if taxRate > 0 {
		fmt.Printf("The tax is $%.2f.\nThe total is $%.2f.\n", tax, total)
	} else {
		fmt.Printf("The total is $%.2f.\n", total)
	}
}

func roundUp(amount float64) float64 {
	return math.Ceil(amount*100) / 100
}
