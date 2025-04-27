package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

const (
	VeryWeak = iota
	Weak
	Strong
	VeryStrong
)

func passwordValidator(password string) int {
	var hasLetter, hasNumber, hasSpecial bool

	for _, char := range password {
		switch {
		case unicode.IsLetter(char):
			hasLetter = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if len(password) < 8 {
		if hasNumber && !hasLetter && !hasSpecial {
			return VeryWeak
		}
		if hasLetter && !hasNumber && !hasSpecial {
			return Weak
		}
	}

	if hasLetter && hasNumber && !hasSpecial && len(password) >= 8 {
		return Strong
	}
	if hasLetter && hasNumber && hasSpecial && len(password) >= 8 {
		return VeryStrong
	}

	return Weak
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	strength := passwordValidator(password)

	var strengthText string
	switch strength {
	case VeryWeak:
		strengthText = "very weak"
	case Weak:
		strengthText = "weak"
	case Strong:
		strengthText = "strong"
	case VeryStrong:
		strengthText = "very strong"
	}

	fmt.Printf("The password '%s' is a %s password.\n", password, strengthText)

}
