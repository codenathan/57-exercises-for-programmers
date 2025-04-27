package main

import (
	"bufio"
	"fmt"
	numbers "github.com/codenathan/exercises-for-programmers/00-numbers"
	"os"
	"strings"
)

func main() {
	fmt.Println("Press C to convert from Fahrenheit to Celsius.")
	fmt.Println("Press F to convert from Celsius to Fahrenheit.")
	fmt.Print("Your choice: ")
	reader := bufio.NewReader(os.Stdin)

	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)
	choice = strings.ToLower(choice)

	var temperature float64

	if choice == "c" {
		temperature = numbers.GetPositiveFloat("Enter the temperature in Fahrenheit: ")
	} else {
		temperature = numbers.GetPositiveFloat("Enter the temperature in Celsius: ")
	}

	result := calculateTemp(choice, temperature)

	if choice == "C" {
		fmt.Printf("The temperature in Celsius is %.2f.\n", result)
	} else {
		fmt.Printf("The temperature in Fahrenheit is %.2f.\n", result)
	}

}

func calculateTemp(choice string, temp float64) float64 {
	if choice == "c" {
		return (temp - 32) * 5 / 9
	}
	return (temp * 9 / 5) + 32
}
