package main

import "fmt"

func main() {
	fmt.Print("What is your name?:")
	var name string
	_, err := fmt.Scan(&name)
	if err != nil {
		return
	}
	fmt.Println(fmt.Sprintf("Hello %s, nice to meet you", name))
}
