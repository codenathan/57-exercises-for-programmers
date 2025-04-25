package main

import (
	"fmt"
	numbers "github.com/codenathan/exercises-for-programmers/00-numbers"
)

const legalDrivingAge = 16

func main() {

	age := numbers.GetPositiveInt("What is your age? ")

	message := "You are not old enough to legally drive."

	if age >= legalDrivingAge {
		message = "You are old enough to legally drive."
	}

	fmt.Println(message)
}
