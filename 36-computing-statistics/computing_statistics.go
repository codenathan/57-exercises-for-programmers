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
	var numbers []float64
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter a number: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "done" {
			break
		}

		num, err := strconv.ParseFloat(input, 64)
		if err != nil {
			continue
		}
		numbers = append(numbers, num)
	}

	if len(numbers) == 0 {
		fmt.Println("No numbers entered!")
		return
	}

	fmt.Print("Numbers: ")
	for i, num := range numbers {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Printf("%.0f", num)
	}
	fmt.Println()

	average := calculateAverage(numbers)
	minimum := findMin(numbers)
	maximum := findMax(numbers)
	stdDev := calculateStandardDeviation(numbers, average)

	fmt.Printf("The average is %.0f.\n", average)
	fmt.Printf("The minimum is %.0f.\n", minimum)
	fmt.Printf("The maximum is %.0f.\n", maximum)
	fmt.Printf("The standard deviation is %.2f.\n", stdDev)

}

func calculateAverage(nums []float64) float64 {
	sum := 0.0
	for _, num := range nums {
		sum += num
	}
	return sum / float64(len(nums))
}

func findMin(nums []float64) float64 {
	minimum := nums[0]
	for _, num := range nums {
		if num < minimum {
			minimum = num
		}
	}
	return minimum
}

func findMax(nums []float64) float64 {
	maximum := nums[0]
	for _, num := range nums {
		if num > maximum {
			maximum = num
		}
	}
	return maximum
}

func calculateStandardDeviation(nums []float64, mean float64) float64 {
	var sumSquares float64
	for _, num := range nums {
		diff := num - mean
		sumSquares += diff * diff
	}
	variance := sumSquares / float64(len(nums))
	return math.Sqrt(variance)
}
