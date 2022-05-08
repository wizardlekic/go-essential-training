package main

import (
	"fmt"
)

func main() {
	var x int
	var y int

	x = 1
	y = 2

	fmt.Printf("x=%v, type of %T\n", x, x)
	fmt.Printf("y=%v, type of %T\n", y, y)

	mean := (x + y) / 2.0
	fmt.Printf("result: %v, type of %T\n", mean, mean)

	fmt.Println()

	var fx float64
	var fy float64

	fx = 1.0
	fy = 2.0

	fmt.Printf("fx=%v, type of %T\n", fx, fx)
	fmt.Printf("fy=%v, type of %T\n", fy, fy)

	fmean := (fx + fy) / 2.0

	fmt.Printf("result: %v, type of %T\n", fmean, fmean)

	fmt.Println()

	/* Type inferrence */
	xshort := 1.0
	yshort := 2.0
	// xshort, yshort := 1.0, 2.0 <- also works

	fmt.Printf("xshort=%v, type of %T\n", xshort, xshort)
	fmt.Printf("yshort=%v, type of %T\n", yshort, yshort)
}
