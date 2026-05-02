package main

import "fmt"

func main() {
	condition := 7 < 9
	fmt.Printf("%T\n", condition)
	if condition {
		fmt.Println("7 is less than 9")
	} else {
		fmt.Println("7 is not less than 9")
	}
	if 12 == 13 {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
	fmt.Println("Hello")

	a := 1
	//No need for break statements after cases, it automatically breaks after a case is executed
	switch a {
	case 1:
		fmt.Println("one")

	case 2:
		fmt.Println("two")

	case 3:
		fmt.Println("three")

	default:
		fmt.Println("Default case")
	}

	x := 1
	//You can also create switch statements without passing variables, Sort of like a if-else statement
	// fallthrough allows us to fall with the current case into the lower one
	switch {
	case x < 1:
		fmt.Println("The value is less than one")

	case x > 1:
		fmt.Println("The value is more than one")

	case x == 1:
		fmt.Println("The value is equal to one")
		fallthrough

	default:
		fmt.Println("None of the options were true!")

	}

}
