package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

//https://gobyexample.com/reading-files
//https://gobyexample.com/writing-files

func main() {
	inputFile, err := os.Open("names.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	var names []string
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			names = append(names, line)
		}
	}

	sort.Strings(names)
	fmt.Printf("Total of %d names\n", len(names))
	fmt.Println(strings.Repeat("-", 20))
	for _, name := range names {
		fmt.Println(name)
	}

	outputFile, err := os.Create("sorted_names.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)

	// Write output
	writer.WriteString(fmt.Sprintf("Total of %d names\n", len(names)))
	writer.WriteString("-----------------\n")
	for _, name := range names {
		writer.WriteString(name + "\n")
	}

	writer.Flush()

}
