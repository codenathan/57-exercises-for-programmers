package main

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"
	"log"
	"os"
	"strings"
)

var users = map[string]string{
	"foo": hashPassword("abc$123"),
	"bar": hashPassword("pa55w0rd!"),
}

func hashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hash)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your username: ")
	usernameInput, _ := reader.ReadString('\n')
	username := strings.TrimSpace(usernameInput)

	fmt.Print("Enter your password: ")
	passwordBytes, _ := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()
	password := string(passwordBytes)

	storedHash, exists := users[username]
	if !exists {
		fmt.Println("I don't know you.")
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password))
	if err != nil {
		fmt.Println("I don't know you.")
	} else {
		fmt.Println("Welcome!")
	}
}
