package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	employees := []map[string]string{
		{"firstName": "John", "lastName": "Johnson", "position": "Manager", "separationDate": "2016-12-31"},
		{"firstName": "Tou", "lastName": "Xiong", "position": "Software Engineer", "separationDate": "2016-10-05"},
		{"firstName": "Michaela", "lastName": "Michaelson", "position": "District Manager", "separationDate": "2015-12-19"},
		{"firstName": "Jake", "lastName": "Jacobson", "position": "Programmer", "separationDate": ""},
		{"firstName": "Jacquelyn", "lastName": "Jackson", "position": "DBA", "separationDate": ""},
		{"firstName": "Sally", "lastName": "Weber", "position": "Web Developer", "separationDate": "2015-12-18"},
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a search string: ")
	searchString, _ := reader.ReadString('\n')
	searchString = strings.TrimSpace(searchString)
	searchString = strings.ToLower(searchString)

	fmt.Println("\nResults:")
	fmt.Printf("%-20s| %-20s| %-15s\n", "Name", "Position", "Separation Date")
	fmt.Println("--------------------|---------------------|-----------------")

	for _, employee := range employees {
		firstName := employee["firstName"]
		lastName := employee["lastName"]
		position := employee["position"]
		separationDate := employee["separationDate"]

		if strings.Contains(strings.ToLower(firstName), searchString) || strings.Contains(strings.ToLower(lastName), searchString) {
			fullName := firstName + " " + lastName
			fmt.Printf("%-20s| %-20s| %-15s\n", fullName, position, separationDate)
		}
	}

}
