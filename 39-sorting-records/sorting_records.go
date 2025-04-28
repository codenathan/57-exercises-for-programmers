package main

import (
	"fmt"
	"sort"
)

// Employee struct for easier management
type Employee struct {
	FirstName      string
	LastName       string
	Position       string
	SeparationDate string
}

func main() {
	employees := []map[string]string{
		{"firstName": "John", "lastName": "Johnson", "position": "Manager", "separationDate": "2016-12-31"},
		{"firstName": "Tou", "lastName": "Xiong", "position": "Software Engineer", "separationDate": "2016-10-05"},
		{"firstName": "Michaela", "lastName": "Michaelson", "position": "District Manager", "separationDate": "2015-12-19"},
		{"firstName": "Jake", "lastName": "Jacobson", "position": "Programmer", "separationDate": ""},
		{"firstName": "Jacquelyn", "lastName": "Jackson", "position": "DBA", "separationDate": ""},
		{"firstName": "Sally", "lastName": "Weber", "position": "Web Developer", "separationDate": "2015-12-18"},
	}

	var employeeList []Employee
	for _, e := range employees {
		employeeList = append(employeeList, Employee{
			FirstName:      e["firstName"],
			LastName:       e["lastName"],
			Position:       e["position"],
			SeparationDate: e["separationDate"],
		})
	}

	sort.Slice(employeeList, func(i, j int) bool {
		return employeeList[i].LastName < employeeList[j].LastName
	})

	fmt.Printf("%-20s| %-18s| %-15s\n", "Name", "Position", "Separation Date")
	fmt.Println("--------------------|-------------------|-----------------")

	// Print each employee
	for _, emp := range employeeList {
		fullName := emp.FirstName + " " + emp.LastName
		fmt.Printf("%-20s| %-18s| %-15s\n", fullName, emp.Position, emp.SeparationDate)
	}
}
