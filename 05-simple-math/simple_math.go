package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getNumber(prompt string) int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		text, _ := reader.ReadString('\n')
		input := strings.TrimSpace(text)

		number, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Invalid number")
			continue
		}

		if number < 0 {
			fmt.Println("Please enter a positive number")
			continue
		}

		return number
	}

}

func main() {
	firstNumber := getNumber("What is the first number? ")
	secondNumber := getNumber("What is the second number? ")

	sum := firstNumber + secondNumber
	difference := firstNumber - secondNumber
	product := firstNumber * secondNumber
	var quotient string

	if secondNumber == 0 {
		quotient = "undefined (cannot divide by zero)"
	} else {
		quotient = fmt.Sprintf("%d", firstNumber/secondNumber)
	}

	fmt.Printf(
		"%d + %d = %d\n%d - %d = %d\n%d * %d = %d\n%d / %d = %s\n",
		firstNumber, secondNumber, sum,
		firstNumber, secondNumber, difference,
		firstNumber, secondNumber, product,
		firstNumber, secondNumber, quotient,
	)
}
