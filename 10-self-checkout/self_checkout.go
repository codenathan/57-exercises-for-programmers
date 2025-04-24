package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BasketItem struct {
	price    float64
	quantity int
}

func getPositiveFloat(prompt string) float64 {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		num, err := strconv.ParseFloat(input, 64)
		if err != nil || num < 0 {
			fmt.Println("Please enter a valid non-negative number.")
			continue
		}
		return num
	}
}

func getPositiveInt(prompt string) int {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		num, err := strconv.Atoi(input)
		if err != nil || num < 0 {
			fmt.Println("Please enter a valid non-negative number.")
			continue
		}
		return num
	}
}

const tax = 0.055

func main() {
	//prompt for prices of three items
	var basket []BasketItem

	for i := 0; i < 3; i++ {
		item := getPositiveFloat("Enter the price of item " + strconv.Itoa(i+1) + ": ")
		quantity := getPositiveInt("Enter the quantity of item " + strconv.Itoa(i+1) + ": ")
		basket = append(basket, BasketItem{item, quantity})
	}

	var subtotal = 0.0
	for _, item := range basket {
		subtotal += item.price * float64(item.quantity)
	}

	taxAmount := subtotal * tax
	total := subtotal + taxAmount

	fmt.Printf("Subtotal: %.2f\nTax: %.2f\nTotal: %.2f\n", subtotal, taxAmount, total)

}
