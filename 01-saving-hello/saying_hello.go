package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getNameInput() string {
	fmt.Print("What is your name? ")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	return strings.TrimSpace(name)
}

func createGreeting(name string) string {
	return fmt.Sprintf("Hello, %s, nice to meet you", name)
}

func printGreeting(greet string) {
	fmt.Println(greet)
}

func main() {
	name := getNameInput()
	greet := createGreeting(name)
	printGreeting(greet)
}
