package main

import (
	"fmt"
	numbers "github.com/codenathan/exercises-for-programmers/00-numbers"
	"strconv"
)

type BasketItem struct {
	price    float64
	quantity int
}

const tax = 0.055

func main() {
	//prompt for prices of three items
	var basket []BasketItem

	for i := 0; i < 3; i++ {
		item := numbers.GetPositiveFloat("Enter the price of item " + strconv.Itoa(i+1) + ": ")
		quantity := numbers.GetPositiveInt("Enter the quantity of item " + strconv.Itoa(i+1) + ": ")
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
