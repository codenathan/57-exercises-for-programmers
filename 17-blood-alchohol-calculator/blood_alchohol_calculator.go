package main

import (
	"fmt"
	"strings"
)

const (
	maleRatio   = 0.73
	femaleRatio = 0.66
)

type PersonData struct {
	weight     float64
	gender     string
	drinks     float64
	alcoholVol float64
	timeSince  float64
}

func main() {
	var data PersonData

	fmt.Print("Enter your weight in pounds: ")
	fmt.Scanln(&data.weight)

	fmt.Print("Enter your gender (M/F): ")
	fmt.Scanln(&data.gender)

	fmt.Print("Number of drinks: ")
	fmt.Scanln(&data.drinks)

	fmt.Print("Amount of alcohol by volume of the drinks consumed (in %): ")
	fmt.Scanln(&data.alcoholVol)

	fmt.Print("Hours since last drink: ")
	fmt.Scanln(&data.timeSince)

	bac := calculateBAC(data)

	fmt.Printf("\nYour BAC is %.3f\n", bac)

	if bac >= 0.08 {
		fmt.Println("It is not legal for you to drive.")
	} else {
		fmt.Println("It is legal for you to drive.")
	}
}

func calculateBAC(data PersonData) float64 {
	ratio := maleRatio
	if strings.ToUpper(data.gender) == "F" {
		ratio = femaleRatio
	}

	alcoholOunces := data.drinks * data.alcoholVol / 100.0
	bac := (alcoholOunces * 5.14 / data.weight * ratio) - (0.015 * data.timeSince)

	if bac < 0 {
		return 0
	}
	return bac

}
