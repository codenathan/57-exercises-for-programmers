package main

import (
	"fmt"
	numbers "github.com/codenathan/exercises-for-programmers/00-numbers"
)

func main() {
	age := numbers.GetPositiveInt("Enter your age")
	restingHR := numbers.GetPositiveInt("Enter your resting heart rate")
	fmt.Printf("\nResting Pulse: %d\tAge: %d\n", restingHR, age)
	fmt.Println("Intensity | Rate")
	fmt.Println("----------|--------")

	for intensity := 55; intensity <= 95; intensity += 5 {
		targetHR := calculateTargetHeartRate(age, restingHR, float64(intensity)/100)
		fmt.Printf("%d%%       | %d bpm\n", intensity, targetHR)
	}
}

func calculateTargetHeartRate(age int, restingHR int, intensity float64) int {
	heartRateReserve := float64(220 - age - restingHR)
	targetHeartRate := (heartRateReserve * intensity) + float64(restingHR)
	return int(targetHeartRate + 0.5)
}
