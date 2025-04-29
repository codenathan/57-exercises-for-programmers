package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter input filename: ")
	inputFileName, _ := reader.ReadString('\n')
	inputFileName = strings.TrimSpace(inputFileName)

	fmt.Print("Enter output filename: ")
	outputFileName, _ := reader.ReadString('\n')
	outputFileName = strings.TrimSpace(outputFileName)

	inputFile, err := os.Open(inputFileName)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer func(inputFile *os.File) {
		err := inputFile.Close()
		if err != nil {
			fmt.Println("Error closing input file:", err)
		}
	}(inputFile)

	var outputLines []string
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		modified := strings.ReplaceAll(line, "utilize", "use")
		outputLines = append(outputLines, modified)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	outputFile, err := os.Create(outputFileName)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer func(outputFile *os.File) {
		err := outputFile.Close()
		if err != nil {
			fmt.Println("Error closing output file:", err)
		}
	}(outputFile)

	for _, line := range outputLines {
		_, _ = outputFile.WriteString(line + "\n")
	}

	fmt.Printf("Replaced occurrences written to %s\n", outputFileName)
}
