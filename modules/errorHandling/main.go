package main

import (
	"errors"
	"fmt"
)

// Error handling in go is a bit different from other languages
// Instead of try catch and finally there is panic, recover and defer

// Create a divide function for demonstration of divide by zero
func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not possible")
	}
	return a / b, nil
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
	// Method 1: Using panic, defer and recover

	// Defer usually executes after the main block finishes execution but if there is a panic defer runs first and panic is invoked
	// defer deferredFunc()

	// Create an error
	// panic("This caused a crash") // any code after panic will not work

	// fmt.Println("Run") // This is unreachable if there is a panic above it

	// Invoke a panic from the divide function
	// divide(1, 0)

	// Method 2: Use the errors package for error handling
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
	// this is a better way to handle errors as it does not crash the program and allows us to handle the error gracefully
}
