package main

import (
	"fmt"
)

func main() {
	//var and const keyword can be used when declaring variable. Similar to JavaScript
	// syntax for declaring would be; var/const var_name type "= value"
	// This is an explicit assignment, you specify type of the variable and can assign values to it later.
	var num int8 = 120
	fmt.Println(num)
	num = 121
	fmt.Println(num)
	// num =129 not allowed  due to overflow int8: signed values representable in 8 bits (-128 to 127) use greater values with int 16,32,64

	// byte is alias for uint8 but can also be used to hold a single character if converted properly
	// rune is alias for int32 can be used for general number operations unless sign and size is not really important for the variable
	var char byte = 'c'
	//printing byte as a string or a character
	fmt.Println(string(char))
	fmt.Printf("%c\n", char)
	fmt.Printf("%T\n", char)

	// %v formats the variable as a int8
	fmt.Printf("%v\n", char)
	/*
	 * This is an implicit assignment when the type of the variable is determined by the compiler. This is what is 	 going to be used most of the time.
	 * Implicit assigns character value as default int32 unless specified as a byte. or perform typecast. In the below example it has int32 first then it is type casted as a byte
	 */
	character := 'c'
	secondCharacter := byte('d')
	fmt.Printf("%c\n", character)
	// %T prints the type of the variable that is passed in the function
	fmt.Printf("%T\n", character)
	fmt.Printf("%c: %T\n", secondCharacter, secondCharacter)
	// NOTE: Conversion from integer to string if intended to just get the same value but as a string, use fmt.Sprint
	// char(value) returns ASCII representation of the value
	a := 65
	fmt.Printf("%s\n", string(a)) // Gives A instead of 65 as a string
	fmt.Printf("%s\n", fmt.Sprint(a))

}
