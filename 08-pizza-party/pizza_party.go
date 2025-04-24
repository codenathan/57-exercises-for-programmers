package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getPositiveInt(prompt string) int {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		num, err := strconv.Atoi(input)
		if err != nil || num < 0 {
			fmt.Println("Please enter a valid non-negative number.")
			continue
		}
		return num
	}
}

func pluralize(word string, count int) string {
	if count == 1 {
		return word
	}
	return word + "s"
}

func main() {
	numberOfPeople := getPositiveInt("How many people? ")
	numberOfPizza := getPositiveInt("How many pizzas do you have? ")
	slicesPerPizza := getPositiveInt("How many slices per pizza? ")

	totalSlices := slicesPerPizza * numberOfPizza
	slicesEach := totalSlices / numberOfPeople
	leftovers := totalSlices % numberOfPeople

	fmt.Printf("%d people with %d pizzas\n", numberOfPeople, numberOfPizza)

	fmt.Printf("Each person gets %d %s of pizza.\n",
		slicesEach, pluralize("piece", slicesEach),
	)

	fmt.Printf("There are %d leftover pieces\n", leftovers)

}
