package main

import "fmt"

func main() {
	// Declare variables with var keyword
	var name string = "Go"
	var version int = 1

	// Short variable declaration
	language := "Go Programming"
	isAwesome := true
	pi := 3.14159

	fmt.Println("Name:", name)
	fmt.Println("Version:", version)
	fmt.Println("Language:", language)
	fmt.Println("Is Awesome:", isAwesome)
	fmt.Println("Pi:", pi)

	// Multiple variable declaration
	var (
		a int     = 10
		b float64 = 20.5
		c string  = "hello"
	)
	fmt.Println("a:", a, "b:", b, "c:", c)

	// Constants
	const maxSize = 100
	fmt.Println("Max size:", maxSize)

	// Zero values
	var zeroInt int
	var zeroFloat float64
	var zeroBool bool
	var zeroString string
	fmt.Printf("Zero values: int=%d, float=%f, bool=%t, string=%q\n",
		zeroInt, zeroFloat, zeroBool, zeroString)
}
