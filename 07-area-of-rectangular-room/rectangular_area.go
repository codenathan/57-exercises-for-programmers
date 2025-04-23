package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const conversionFactor = 0.09290304

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
	length := getPositiveFloat("What is the length of the room in feet? ")
	width := getPositiveFloat("What is the width of the room in feet? ")

	areaFeet := length * width
	areaMeters := areaFeet * conversionFactor

	fmt.Printf("You entered dimensions of %.0f feet by %.0f feet.\n", length, width)
	fmt.Println("The area is")
	fmt.Printf("%.0f square feet\n", areaFeet)
	fmt.Printf("%.3f square meters\n", areaMeters)
}
