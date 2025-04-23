package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func convertStringToInt(prompt string) int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		num, err := strconv.Atoi(input)
		if err != nil || num < 0 {
			fmt.Println("Please enter a valid non-negative number.")
			continue
		}
		return num
	}

}

func main() {
	age := convertStringToInt("What is your current age? ")
	retireAge := convertStringToInt("At what age would you like to retire? ")

	yearsLeft := retireAge - age
	currentYear := time.Now().Year()
	retireYear := currentYear + yearsLeft

	if yearsLeft <= 0 {
		fmt.Printf("You can already retire. It's %d, so you could have retired in %d.\n", currentYear, retireYear)
	} else {
		fmt.Printf("You have %d years left until you can retire\n", yearsLeft)
		fmt.Printf("It's %d, so you can retire in %d.\n", currentYear, retireYear)
	}
}
