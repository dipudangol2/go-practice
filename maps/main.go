package main

import (
	"fmt"
	"sort"
)

func main() {
	// Create and populate a map
	capitals := map[string]string{
		"France":  "Paris",
		"Germany": "Berlin",
		"Japan":   "Tokyo",
	}
	fmt.Println("Capitals:", capitals)

	// Add and update
	capitals["India"] = "New Delhi"
	capitals["France"] = "Paris (updated)"

	// Lookup with existence check
	if capital, ok := capitals["Japan"]; ok {
		fmt.Println("Capital of Japan:", capital)
	}

	if _, ok := capitals["Australia"]; !ok {
		fmt.Println("Australia not found in map")
	}

	// Delete
	delete(capitals, "France")

	// Iterate in sorted key order
	keys := make([]string, 0, len(capitals))
	for k := range capitals {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	fmt.Println("Remaining entries (sorted):")
	for _, k := range keys {
		fmt.Printf("  %s -> %s\n", k, capitals[k])
	}

	// Map with make
	scores := make(map[string]int)
	scores["Alice"] = 95
	scores["Bob"] = 87
	scores["Charlie"] = 92
	fmt.Println("Scores:", scores)

	// Word frequency count
	words := []string{"go", "is", "fun", "go", "is", "great", "go"}
	freq := make(map[string]int)
	for _, w := range words {
		freq[w]++
	}
	fmt.Println("Word frequencies:", freq)
}
