package main

import (
	"fmt"
)

func main() {
	var char byte = 'c'
	//printing byte as a string or a character
	fmt.Println(string(char))
	fmt.Printf("%c\n", char)
	// %v formats the variable as a int8
	fmt.Printf("%v\n", char)
	fmt.Println("Hello world in Go!")

}
