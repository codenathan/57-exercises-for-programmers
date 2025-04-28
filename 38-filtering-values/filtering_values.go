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
	fmt.Print("Enter a list of numbers, separated by spaces: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	numbers := convertInputToArray(input)
	evenNumbers := filterEvenNumbers(numbers)

	fmt.Print("The even numbers are ")
	for _, num := range evenNumbers {
		fmt.Printf("%d ", num)
	}
	fmt.Println()

}

func convertInputToArray(input string) []int {
	parts := strings.Split(input, " ")
	var numbers []int

	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err == nil {
			numbers = append(numbers, num)
		}
	}

	return numbers
}

func filterEvenNumbers(numbers []int) []int {
	var even []int
	for _, num := range numbers {
		if num%2 == 0 {
			even = append(even, num)
		}
	}
	return even
}
