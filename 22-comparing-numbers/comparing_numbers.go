package main

import (
	"fmt"
	numbers "github.com/codenathan/exercises-for-programmers/00-numbers"
)

func main() {
	num1 := numbers.GetPositiveInt("Enter the first number: ")
	num2 := numbers.GetPositiveInt("Enter the second number: ")
	num3 := numbers.GetPositiveInt("Enter the third number: ")

	if num1 == num2 || num2 == num3 || num3 == num1 {
		fmt.Println("Numbers must be different")
		return
	}

	largest := num1
	if num2 > largest {
		largest = num2
	}
	if num3 > largest {
		largest = num3
	}

	fmt.Printf("The largest number is %d.\n", largest)

}
