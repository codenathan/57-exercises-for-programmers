package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type InputData struct {
	firstName  string
	lastName   string
	employeeID string
	zipCode    string
}

func validateFirstName(name string) (bool, string) {
	if len(name) == 0 {
		return false, "The first name must be filled in."
	}
	if len(name) < 2 {
		return false, "\"" + name + "\" is not a valid first name. It is too short."
	}
	return true, ""
}

func validateLastName(name string) (bool, string) {
	if len(name) == 0 {
		return false, "The last name must be filled in."
	}
	if len(name) < 2 {
		return false, "\"" + name + "\" is not a valid last name. It is too short."
	}
	return true, ""
}

func validateEmployeeID(id string) (bool, string) {
	pattern := `^[A-Za-z]{2}-\d{4}$`
	match, _ := regexp.MatchString(pattern, id)
	if !match {
		return false, id + " is not a valid ID."
	}
	return true, ""
}

func validateZipCode(zip string) (bool, string) {
	_, err := strconv.Atoi(zip)
	if err != nil {
		return false, "The ZIP code must be numeric."
	}
	return true, ""
}

func validateInput(input InputData) []string {
	var errors []string

	if valid, msg := validateFirstName(input.firstName); !valid {
		errors = append(errors, msg)
	}
	if valid, msg := validateLastName(input.lastName); !valid {
		errors = append(errors, msg)
	}
	if valid, msg := validateZipCode(input.zipCode); !valid {
		errors = append(errors, msg)
	}
	if valid, msg := validateEmployeeID(input.employeeID); !valid {
		errors = append(errors, msg)
	}

	return errors
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the first name: ")
	firstName, _ := reader.ReadString('\n')
	firstName = strings.TrimSpace(firstName)

	fmt.Print("Enter the last name: ")
	lastName, _ := reader.ReadString('\n')
	lastName = strings.TrimSpace(lastName)

	fmt.Print("Enter the ZIP code: ")
	zipCode, _ := reader.ReadString('\n')
	zipCode = strings.TrimSpace(zipCode)

	fmt.Print("Enter an employee ID: ")
	employeeID, _ := reader.ReadString('\n')
	employeeID = strings.TrimSpace(employeeID)

	input := InputData{
		firstName:  firstName,
		lastName:   lastName,
		employeeID: employeeID,
		zipCode:    zipCode,
	}

	errors := validateInput(input)

	if len(errors) == 0 {
		fmt.Println("There were no errors found.")
	} else {
		for _, err := range errors {
			fmt.Println(err)
		}
	}
}
