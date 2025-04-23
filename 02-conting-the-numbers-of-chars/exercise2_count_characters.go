package main

import "fmt"

func main() {
	fmt.Print("What is the input string?")
	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("%s has %d characters", input, len(input)))
}
