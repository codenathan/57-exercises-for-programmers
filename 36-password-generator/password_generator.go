package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	letterBytes  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberBytes  = "0123456789"
	specialBytes = "!@#$%^&*()_+-=[]{}|;:,.<>?"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	fmt.Print("What's the minimum length? ")
	minLengthStr, _ := reader.ReadString('\n')
	minLength, _ := strconv.Atoi(strings.TrimSpace(minLengthStr))

	fmt.Print("How many special characters? ")
	specialCountStr, _ := reader.ReadString('\n')
	specialCount, _ := strconv.Atoi(strings.TrimSpace(specialCountStr))

	fmt.Print("How many numbers? ")
	numberCountStr, _ := reader.ReadString('\n')
	numberCount, _ := strconv.Atoi(strings.TrimSpace(numberCountStr))

	// Generate password
	password := make([]byte, 0, minLength)

	// Add required special characters
	for i := 0; i < specialCount; i++ {
		password = append(password, specialBytes[rng.Intn(len(specialBytes))])
	}

	// Add required numbers
	for i := 0; i < numberCount; i++ {
		password = append(password, numberBytes[rng.Intn(len(numberBytes))])
	}

	// Fill the rest with letters
	remainingLength := minLength - specialCount - numberCount
	for i := 0; i < remainingLength; i++ {
		password = append(password, letterBytes[rng.Intn(len(letterBytes))])
	}

	// Shuffle the password
	rng.Shuffle(len(password), func(i, j int) {
		password[i], password[j] = password[j], password[i]
	})

	fmt.Println("Your password is")
	fmt.Println(string(password))
}
