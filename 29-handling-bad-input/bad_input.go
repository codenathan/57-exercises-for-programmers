package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var num float64
	var err error

	for {
		fmt.Print("What is the rate of return? ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		num, err = strconv.ParseFloat(input, 64)
		if err != nil || num <= 0 {
			// Ensure the input is a positive number (greater than 0)
			fmt.Println("Sorry. That's not a valid input. Please enter a number greater than 0.")
			continue
		}
		break
	}

	fmt.Printf("It will take %.0f years to double your initial investment.\n", checkInvestment(num))
}

func checkInvestment(rate float64) float64 {
	return 72 / rate
}
