package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	playAgain := true

	fmt.Println("Let's play Guess the Number.")

	for playAgain {
		fmt.Print("Pick a difficulty level (1, 2, or 3): ")
		difficultyStr, _ := reader.ReadString('\n')
		difficulty, err := strconv.Atoi(strings.TrimSpace(difficultyStr))

		if err != nil || difficulty < 1 || difficulty > 3 {
			fmt.Println("Invalid difficulty. Please choose 1, 2, or 3.")
			continue
		}

		maxNumber := 10
		if difficulty == 2 {
			maxNumber = 100
		} else if difficulty == 3 {
			maxNumber = 1000
		}

		target := rand.Intn(maxNumber) + 1
		guesses := 0

		fmt.Println("I have my number. What's your guess?")

		for {
			guessStr, _ := reader.ReadString('\n')
			guess, err := strconv.Atoi(strings.TrimSpace(guessStr))
			guesses++

			if err != nil {
				fmt.Print("Invalid input. Guess again: ")
				continue
			}

			if guess < target {
				fmt.Print("Too low. Guess again: ")
			} else if guess > target {
				fmt.Print("Too high. Guess again: ")
			} else {
				fmt.Printf("You got it in %d guesses!\n", guesses)

				var comment string
				switch {
				case guesses == 1:
					comment = "You're a mind reader!"
				case guesses <= 4:
					comment = "Most impressive."
				case guesses <= 6:
					comment = "You can do better than that."
				default:
					comment = "Better luck next time."
				}
				fmt.Println(comment)
				break
			}
		}

		fmt.Print("Play again? (y/n): ")
		againStr, _ := reader.ReadString('\n')
		playAgain = strings.ToLower(strings.TrimSpace(againStr)) == "y"
	}

	fmt.Println("Goodbye!")

}
