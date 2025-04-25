package main

import (
	"bufio"
	"fmt"
	numbers "github.com/codenathan/exercises-for-programmers/00-numbers"
	"os"
	"strings"
)

func main() {
	const taxRate = 0.055
	orderAmount := numbers.GetPositiveFloat("What is the order amount: ")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is the state? ")
	stateInput, _ := reader.ReadString('\n')
	state := strings.ToLower(strings.TrimSpace(stateInput))

	total := orderAmount

	if state == "wi" || state == "wisconsin" {
		tax := orderAmount * taxRate
		total = orderAmount + tax
		fmt.Printf("The Subtotal is %.2f\n", orderAmount)
		fmt.Printf("The Tax is %.2f\n", tax)
	}

	fmt.Printf("The Total is %.2f\n", total)

}
