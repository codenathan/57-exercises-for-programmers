package main

import (
	"fmt"
	numbers "github.com/codenathan/exercises-for-programmers/00-numbers"
)

func main() {

	numbersArray := [5]int{}

	for i := 0; i < 5; i++ {
		numbersArray[i] = numbers.GetPositiveInt("Enter a number: ")
	}

	var sum int
	for _, number := range numbersArray {
		sum += number
	}

	fmt.Printf("The total is is %d\n", sum)

}
