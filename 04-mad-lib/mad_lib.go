package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getUserInput(reader *bufio.Reader, message string) string {

	fmt.Print(message)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	noun := getUserInput(reader, "Enter a noun: ")
	verb := getUserInput(reader, "Enter a verb: ")
	adjective := getUserInput(reader, "Enter an adjective: ")
	adverb := getUserInput(reader, "Enter an adverb: ")

	fmt.Printf("Do you %s your %s %s %s? That's hilarious!\n", verb, adjective, noun, adverb)
}
