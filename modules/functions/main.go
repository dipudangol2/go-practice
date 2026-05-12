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

// Return functions from a function
func getFunc(str string) func(string) string {
	// This function returns a function which takes another string as input and returns the concatenation of the parameter of the parent function and the parameter of itself.
	return func(str2 string) string {
		return str + str2
	}
}

// Variadic functions: These can take any number of parameters of the specified type in the function definition
func sum(nums ...int) int { // The parameters get taken in as a slice of numbers
	total := 0
	// Loop through the parameter values
	for _, num := range nums {
		total += num
	}
	return total
}

// Named return values
func greet(str string) (greeting string, name string) {
	greeting = "Hello, " // Already declared at the top
	name = str + "!"
	return // Doesnot need variable names to be included in the return statement.
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

	f1 := getFunc("Hello, ") // Gets a function returned which takes another parameter
	value := f1("World!!")   // Invokes the child function which concatenates the parameter of the previous function and the current parameter
	fmt.Println(value)

	// Call the sum functions with multiple inputs
	sum1 := sum(1, 2, 3)
	sum2 := sum(2, 4, 6, 8)
	sum3 := sum(1, 3, 5, 7, 9)
	fmt.Println(sum1, sum2, sum3)
	// You can also spread a slice similar to javascript spread operator
	sum4 := sum([]int{1, 2, 3, 4, 5}...) // adding ... at the end passes the slice as individual elements inside the function. Only works in a variadic parameter declaration
	fmt.Println(sum4)

	// Named return values
	greeting, name := greet("Dipu")

	fmt.Print(greeting, name, "\n")
}
