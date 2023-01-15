// Count how many times each words appears in a text.
// To split text to words, use the "Fields" function from the "strings" package.
// Also use "ToLower" from the same package to conver the words to lowercase.
// Finally, print the map

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	input := string(file)

	words := strings.Fields(strings.ToLower(input))

	word_map := map[string]int{}

	for _, w := range words {
		_, ok := word_map[w]

		if !ok {
			word_map[w] = 1
			continue
		}

		word_map[w]++
	}

	for key := range word_map {
		fmt.Println("Word: " + key + " appeared " + strconv.Itoa(word_map[key]) + " times.")
	}
}
