package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Print("Please enter the number of the month: ")
	var month string
	fmt.Scan(&month)

	monthNumber, err := strconv.Atoi(month)
	if err != nil {
		fmt.Println("Invalid input. Please enter a numeric value.")
		return
	}

	var stringMonth string
	switch monthNumber {
	case 1:
		stringMonth = "January"
	case 2:
		stringMonth = "February"
	case 3:
		stringMonth = "March"
	case 4:
		stringMonth = "April"
	case 5:
		stringMonth = "May"
	case 6:
		stringMonth = "June"
	case 7:
		stringMonth = "July"
	case 8:
		stringMonth = "August"
	case 9:
		stringMonth = "September"
	case 10:
		stringMonth = "October"
	case 11:
		stringMonth = "November"
	case 12:
		stringMonth = "December"
	default:
		fmt.Println("You must enter a number between 1 and 12.")
		return
	}

	fmt.Printf("The name of the month is %s.\n", stringMonth)
}
