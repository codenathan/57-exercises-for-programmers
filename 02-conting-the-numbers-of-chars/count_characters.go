package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getUserInput() string {
	fmt.Print("What is the input string?")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func main() {
	input := getUserInput()

	if input == "" {
		fmt.Println("You must enter an input")
		return
	}

	fmt.Printf("%s has %d characters.\n", input, len(input))
}
