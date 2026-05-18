package main

import "fmt"

// Any generic type can be passed in it
func printSlice[T any](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func main() {
	// fmt.Println("Generics and Enums in Go:")
	items := []int{1, 2, 3, 4, 5}
	printSlice(items)
	printSlice([]string{"Dipu", "Dangol"})
}
