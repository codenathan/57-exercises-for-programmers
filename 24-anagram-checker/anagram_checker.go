package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func isAnagram(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	word1Slice := strings.Split(str1, "")
	word2Slice := strings.Split(str2, "")
	sort.Strings(word1Slice)
	sort.Strings(word2Slice)

	return strings.Join(word1Slice, "") == strings.Join(word2Slice, "")
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter two strings and I'll tell you if they are anagrams:")

	fmt.Print("Enter the first string: ")
	str1, _ := reader.ReadString('\n')
	str1 = strings.TrimSpace(str1)

	fmt.Print("Enter the second string: ")
	str2, _ := reader.ReadString('\n')
	str2 = strings.TrimSpace(str2)

	if isAnagram(str1, str2) {
		fmt.Printf("\"%s\" and \"%s\" are anagrams.\n", str1, str2)
	} else {
		fmt.Printf("\"%s\" and \"%s\" are not anagrams.\n", str1, str2)
	}
}
