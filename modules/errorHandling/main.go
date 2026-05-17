package main

import "fmt"

// Error handling in go is a bit different from other languages
// Instead of try catch and finally there is panic, recover and defer

// Create a divide function for demonstration of divide by zero
func divide(a int, b int) int {
	return a / b
}

// Function which is deferred
func deferredFunc() {
	r := recover() // Store the error messages and details of the error

	if r == nil {
		fmt.Println("There were no errors in the program")
	} else {
		fmt.Println(r) // The message passed inside the panic
	}

	fmt.Println("This is a deferred function which executes after the main block execution is finished")
}

func main() {
	// Defer usually executes after the main block finishes execution but if there is a panic defer runs first and panic is invoked
	defer deferredFunc()

	// Create an error
	// panic("This caused a crash") // any code after panic will not work

	fmt.Println("Run") // This is unreachable if there is a panic above it

	// Invoke a panic from the divide function
	divide(1, 0)

}

