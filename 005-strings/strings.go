package main

import (
	"fmt"
)

func main() {
	book := "The color of magic"
	fmt.Println(book)

	fmt.Println(len(book))

	fmt.Printf("book[0] = %v (type %T)\n", book[0], book[0])
	// uint8 is a byte

	// Strings in Go are immutable

	// Slice (start, end - not included), 0 based, 1/2 empty range
	fmt.Println(book[4:12])

	// Slice (no end)
	fmt.Println(book[4:])

	// Slice (no start)
	fmt.Println(book[:4])

	// Use + to concatenate strings
	fmt.Println("t" + book[1:])

	// Go strings have Unicode enabled

	// Multi line raw strings
	poem := `
		The road goes ever on
		Down from the door where it began
		...
		`
	fmt.Println(poem)
}
