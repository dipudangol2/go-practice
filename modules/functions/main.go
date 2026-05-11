package main

import (
	"fmt"
	"strconv"
)

// Declaration of function with multiple return values
func multiReturn(year int) (int, string) {
	return 2026 - year, "The age is"
}

// Functions can have functions as a parameter. Here is a demo function
func callFunc(num int, callable func(int) int) string {

	return "The double of " + strconv.Itoa(num) + "  is " + strconv.Itoa(callable(num))
}

func funcCall(callable func(int) int) int {
	return callable(15)
}

func returnDouble(number int) int {
	return number * 2
}

func main() {
	// usually all return types must be assigned to a variable. If a variable is not to be used use _ symbol to skip that value
	age, message := multiReturn(2002)
	fmt.Println(message, age)
	fmt.Println(callFunc(15, returnDouble))

	// The function can be passed inside the function call, they are known as anonymus functions
	fmt.Println(funcCall(func(x int) int {
		return x + 1
	})) // this should return 16

	// anonymus functions can also be stored in variables and passed to another functions.
	fun := func(x int) int {
		return x * 2
	}
	fmt.Println(funcCall(fun)) // This should print 30
}
