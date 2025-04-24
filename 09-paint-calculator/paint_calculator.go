package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const conversionFactor = 350

func getPositiveFloat(prompt string) float64 {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		num, err := strconv.ParseFloat(input, 64)
		if err != nil || num < 0 {
			fmt.Println("Please enter a valid non-negative number.")
			continue
		}
		return num
	}
}

func main() {
	length := getPositiveFloat("Enter the length of the ceiling in feet: ")
	width := getPositiveFloat("Enter the width of the ceiling in feet: ")
	area := length * width

	// math.ceil rounds up any decimal number to the next whole number.
	gallonsRequired := math.Ceil(area / conversionFactor)
	fmt.Printf("You will need to purchase %.0f gallons of paint to cover %.0f square feet.\n", gallonsRequired, area)
}
