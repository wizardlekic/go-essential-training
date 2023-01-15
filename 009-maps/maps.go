package main

import (
	"fmt"
)

func main() {
	stocks := map[string]float64{
		"AMZN": 2087.98,
		"GOOG": 2540.85,
		"MSFT": 287.70,
	}

	// Number of items
	fmt.Println(len(stocks))

	// Get a value
	fmt.Println(stocks["MSFT"])

	// Get zero value if not found
	fmt.Println(stocks["TSLA"]) // 0

	// Use two values to see if found
	value, ok := stocks["TSLA"]

	if !ok {
		fmt.Println("TSLA not found")
	} else {
		fmt.Println(value)
	}

	// Set a value
	stocks["TSLA"] = 822.12
	fmt.Println(stocks)

	// Delete a value
	delete(stocks, "AMZN")
	fmt.Println(stocks)

	// Iterate for keys
	for key := range stocks {
		fmt.Println(key)
	}

	// Iterate for keys and values
	for key, value := range stocks {
		fmt.Printf("%s -> %.2f\n", key, value)
	}
}
