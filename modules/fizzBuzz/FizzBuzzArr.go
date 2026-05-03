// FizzBuzz problem: Print 100 numbers from 1 to 100 but for every multiple of 3 print Fizz, every multiple of 5 print Buzz and every multiple of both print FizzBuzz
package fizzbuzz

import (
	"fmt"
	"strconv"
)

func FizzBuzzArr() {
	fmt.Println("Fizz-Buzz Problem (Array-only Implemetation):")
	var size int
	fmt.Print("Enter the number you want to print up to:")
	fmt.Scanf("%d", &size)

	arr := make([]string, size)

	for idx, _ := range arr {
		if (idx+1)%3 == 0 && (idx+1)%5 == 0 {
			arr[idx] = "FizzBuzz"
		} else if (idx+1)%3 == 0 {
			arr[idx] = "Fizz"
		} else if (idx+1)%5 == 0 {
			arr[idx] = "Buzz"
		} else {
			arr[idx] = strconv.Itoa(idx + 1)
		}
	}

	fmt.Print("[ ")
	for _, element := range arr {
		fmt.Printf("%s, ", element)
	}

	fmt.Print("]\n")
}
