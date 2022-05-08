// slices are like lists or arrays in other languages
package main

import (
	"fmt"
)

func main() {
	// same type
	loons := []string{"bugs", "daffy", "taz"}
	fmt.Printf("loons = %v (type %T)\n", loons, loons)

	// length
	fmt.Println(len(loons)) // 3

	// indexing
	fmt.Println(loons[1]) // daffy

	// slicing slices
	fmt.Println(loons[1:]) // [daffy taz]

	// looping
	for i := 0; i < len(loons); i++ {
		fmt.Println(loons[i])
	}

	// range - better looping option - index with one variable
	for i := range loons {
		fmt.Println(i) // index
	}

	// range with two variables
	for i, name := range loons {
		fmt.Printf("%s at %d\n", name, i)
	}

	// ignore index by using _
	for _, name := range loons {
		fmt.Println(name)
	}

	// append - add to end of slice
	loons = append(loons, "elmer")
	fmt.Println(loons) // [bugs daffy taz elmer]
}
