package main

import ("go-practice/modules/fizzBuzz"
"fmt"
)
func main(){
	fmt.Println("Choose FizzBuzz module you would like to run:")
	fmt.Println("1. FizzBuzz with Array.")
	fmt.Println("2. FizzBuzz with Map.")
	var a int ;
	fmt.Print("Enter your choice: ")
	fmt.Scanf("%d",&a)
	fmt.Printf("%v\n",a)

	switch a {
		case 1:
		fizzbuzz.FizzBuzzArr()

		case 2: 
		fizzbuzz.FizzBuzzMap()

		default:
		fmt.Println("Invalid option!")
	}

}

