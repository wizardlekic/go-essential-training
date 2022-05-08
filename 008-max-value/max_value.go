/*
 Extract maximal value from a slice
*/

package main

import (
	"fmt"
)

func main() {
	nums := []int{16, 8, 42, 4, 23, 15}
	fmt.Println(nums)

	maxValue := nums[0]

	for _, current := range nums[1:] {
		if current > maxValue {
			maxValue = current
		}
	}

	fmt.Println(maxValue)
}
