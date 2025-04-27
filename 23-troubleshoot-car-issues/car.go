package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func askQuestion(question string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(question + " (y/n): ")
	answer, _ := reader.ReadString('\n')
	return strings.TrimSpace(strings.ToLower(answer))
}

func main() {
	q1 := askQuestion("Is the car silent when you turn the key?")

	if q1 == "y" {
		q2 := askQuestion("Are the battery terminals corroded?")

		if q2 == "y" {
			fmt.Println("Clean terminals and try starting again.")
		} else {
			fmt.Println("Replace cables and try again.")
		}
	} else {
		q3 := askQuestion("Does the car make a clicking noise?")

		if q3 == "y" {
			fmt.Println("Replace the battery")
		} else {
			q4 := askQuestion("Does the car crank up but fail to start?")

			if q4 == "y" {
				fmt.Println("Check the spark plug connections")
			} else {
				q5 := askQuestion("Does the engine start and then die?")

				if q5 == "y" {
					q6 := askQuestion("does you car have fuel injection")

					if q6 == "y" {
						fmt.Println("get it in for service")
					} else {
						fmt.Println("Check to ensure the choke is opening and closing.")
					}
				}
			}
		}
	}
}
