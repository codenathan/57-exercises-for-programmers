package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func main() {
	answers := []string{
		"Yes",
		"No",
		"Maybe",
		"Ask again later",
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What's your question? Will I be rich and famous?")
	question, _ := reader.ReadString('\n')
	question = strings.TrimSpace(question)

	randomIndex := rand.Intn(len(answers))
	fmt.Println(answers[randomIndex])
}
