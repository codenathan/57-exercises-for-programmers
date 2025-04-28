package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Product struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type ProductData struct {
	Products []Product `json:"products"`
}

func main() {
	file, err := os.Open("data.json")
	if err != nil {
		fmt.Println("Error opening products.json:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}(file)

	var data ProductData
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("What is the product name? ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		found := false
		for _, product := range data.Products {
			if strings.EqualFold(product.Name, input) {
				fmt.Printf("Name: %s\n", product.Name)
				fmt.Printf("Price: $%.2f\n", product.Price)
				fmt.Printf("Quantity on hand: %d\n", product.Quantity)
				found = true
				break
			}
		}

		if found {
			break
		} else {
			fmt.Println("Sorry, that product was not found in our inventory.")
		}
	}
}
