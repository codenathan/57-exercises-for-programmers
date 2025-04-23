package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("What is the quote? ")
	quote, _ := reader.ReadString('\n')
	quote = strings.TrimSpace(quote)

	fmt.Print("Who said it? ")
	author, _ := reader.ReadString('\n')
	author = strings.TrimSpace(author)

	fmt.Println(author + " says, \"" + quote + "\"")
}
