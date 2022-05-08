/*
An "even ended number" is a number whose first and last digit are
the same.

Count how many "even ended numbers" are there that are a
multiplication of two 4 digit numbers.
*/

package main

import (
	"fmt"
)

func main() {
	count := 0

	for i := 1000; i <= 9999; i++ {
		for j := i; j <= 9999; j++ {
			result := i * j

			resultString := convertIntToString(result)

			if isEvenEnded(resultString) {
				count++
			}
		}
	}

	fmt.Printf("There are %d even ended numbers", count)
}

func convertIntToString(number int) string {
	return fmt.Sprintf("%d", number)
}

func isEvenEnded(numberString string) bool {
	return numberString[0] == numberString[len(numberString)-1]
}
