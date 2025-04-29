package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {

	inputFile, err := os.Open("words.txt")
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

	wordCounts := make(map[string]int)
	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		line := strings.ToLower(scanner.Text()) // Make it case-insensitive
		words := strings.Fields(line)
		for _, word := range words {
			word = strings.Trim(word, ".,!?\"';:-()[]{}") // Clean up punctuation
			wordCounts[word]++
		}
	}

	type wordFreq struct {
		word  string
		count int
	}
	var freqs []wordFreq
	for word, count := range wordCounts {
		freqs = append(freqs, wordFreq{word, count})
	}

	sort.Slice(freqs, func(i, j int) bool {
		return freqs[i].count > freqs[j].count
	})

	for _, wf := range freqs {
		fmt.Printf("%-10s: %s\n", wf.word, strings.Repeat("*", wf.count))
	}

}
