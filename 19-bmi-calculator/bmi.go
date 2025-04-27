package main

import (
	"fmt"
	numbers "github.com/codenathan/exercises-for-programmers/00-numbers"
)

func main() {
	weight := numbers.GetPositiveFloat("Enter your weight in pounds: ")
	height := numbers.GetPositiveFloat("Enter your height in inches: ")

	bmi := calculateBMI(weight, height)

	fmt.Printf("Your BMI is %.2f\n", bmi)

	if bmi < 18.5 {
		fmt.Println("You are underweight.")
	} else if bmi <= 25 {
		fmt.Println("You are within the ideal weight range.")
	} else {
		fmt.Println("You are overweight.")
	}

}

func calculateBMI(weight, height float64) float64 {
	return (weight / (height * height)) * 703
}
