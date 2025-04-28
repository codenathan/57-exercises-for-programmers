package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	employeeList := []string{"John Smith", "Jackie Jackson", "Chris Jones", "Amanda Cullen", "Jeremy Goodwin"}

	fmt.Println("There are", len(employeeList), "employees:")

	for _, employee := range employeeList {
		fmt.Println(employee)
	}

	fmt.Print("Enter an employee name to remove: ")
	reader := bufio.NewReader(os.Stdin)
	nameToRemove, _ := reader.ReadString('\n')
	nameToRemove = strings.TrimSpace(nameToRemove)

	for i, employee := range employeeList {
		if employee == nameToRemove {
			employeeList = append(employeeList[:i], employeeList[i+1:]...)
			break
		}
	}

	fmt.Printf("\nThere are %d employees:\n", len(employeeList))
	for _, employee := range employeeList {
		fmt.Println(employee)
	}
}
