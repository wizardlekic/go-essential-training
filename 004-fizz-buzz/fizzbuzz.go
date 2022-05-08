/*
Print the numbers between 1 and 20, one per line.
- If the number is divisible by 3 print "fizz"
- If the number is divisible by 5 print "buzz"
- If the number is divisible by both 3 and 5 print "fizz buzz"
- Otherwise print the number
*/

package main

import "fmt"

func main() {

	for i := 1; i < 21; i++ {
		fizz := isFizz(i)
		buzz := isBuzz(i)

		fizzBuzz(fizz, buzz, i)
	}
}

func fizzBuzz(fizz bool, buzz bool, value int) {
	switch {
	case fizz && buzz:
		fmt.Println("fizz buzz")
	case fizz:
		fmt.Println("fizz")
	case buzz:
		fmt.Println("buzz")
	default:
		fmt.Println(value)
	}
}

func isFizz(value int) bool {
	return (value % 3) == 0
}

func isBuzz(value int) bool {
	return (value % 5) == 0
}
