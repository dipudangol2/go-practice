package main

import "fmt"

// Any generic type can be passed in it
func printSlice[T any](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// Create a custom type that takes specific data types
type number interface {
	int | float32 | float64
}

// Function which only operates on numbers
// Can use union with types without declaring like [ T int | float64 | float32 ] this works normally
func sum[T number](items ...T) T {
	sum := T(0)
	for _, num := range items {
		sum += num
	}
	return sum
}

// using generics on structs
type stack[T int | string] struct {
	elements []T
}

func main() {
	// fmt.Println("Generics and Enums in Go:")
	items := []int{1, 2, 3, 4, 5}
	printSlice(items)                      // This runs without any hiccup
	printSlice([]string{"Dipu", "Dangol"}) // This also runs with seperate type passed

	// invoke the sum function with a slice of integers as well as float
	intSlice := []int{1, 2, 3, 4, 5, 6}
	floatSlice := []float64{1.35, 4.5533, 5.4, 6, 77.89}

	fmt.Println("Sum for the slice(", intSlice, "):", sum(intSlice...))
	fmt.Println("Sum for the slice (", floatSlice, ") with the same function using generic type:", sum(floatSlice...))

	// Create an object of stack for int and string. They need to be instantiated to be used.
	intStack := stack[int]{[]int{1, 2, 3, 4, 5}}
	stringStack := stack[string]{[]string{"Dipu", "Dangol"}}
	fmt.Println(intStack)
	fmt.Println(stringStack)
}
