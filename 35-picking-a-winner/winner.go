package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	var contestants []string

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter names of contestants (leave blank to finish):")

	for {
		fmt.Print("Enter a name: ")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		if name == "" {
			break
		}

		contestants = append(contestants, name)
	}

	if len(contestants) == 0 {
		fmt.Println("No contestants entered!")
		return
	}

	winner := contestants[rng.Intn(len(contestants))]
	fmt.Printf("The winner is... %s.\n", winner)
}
