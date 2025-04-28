package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Employee struct {
	LastName  string
	FirstName string
	Salary    string
}

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var employees []Employee
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) == 3 {
			employees = append(employees, Employee{
				LastName:  parts[0],
				FirstName: parts[1],
				Salary:    parts[2],
			})
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	maxLastName := len("Last")
	maxFirstName := len("First")
	maxSalary := len("Salary")

	for _, emp := range employees {
		if len(emp.LastName) > maxLastName {
			maxLastName = len(emp.LastName)
		}
		if len(emp.FirstName) > maxFirstName {
			maxFirstName = len(emp.FirstName)
		}
		if len(emp.Salary) > maxSalary {
			maxSalary = len(emp.Salary)
		}
	}

	maxLastName++
	maxFirstName++
	maxSalary++

	headerFormat := fmt.Sprintf("%%-%ds%%-%ds%%-%ds\n", maxLastName, maxFirstName, maxSalary)
	fmt.Printf(headerFormat, "Last", "First", "Salary")

	totalWidth := maxLastName + maxFirstName + maxSalary
	fmt.Println(strings.Repeat("-", totalWidth))

	for _, emp := range employees {
		fmt.Printf(headerFormat, emp.LastName, emp.FirstName, emp.Salary)
	}
}
