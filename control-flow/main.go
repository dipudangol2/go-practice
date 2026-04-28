package main

import "fmt"

func main() {
	// if / else
	x := 42
	if x > 0 {
		fmt.Println("x is positive")
	} else if x < 0 {
		fmt.Println("x is negative")
	} else {
		fmt.Println("x is zero")
	}

	// if with initialization statement
	if y := x * 2; y > 50 {
		fmt.Println("doubled x is greater than 50:", y)
	}

	// for loop (basic)
	fmt.Print("Counting: ")
	for i := 0; i < 5; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// for as while loop
	n := 1
	for n < 100 {
		n *= 2
	}
	fmt.Println("First power of 2 >= 100:", n)

	// range over slice
	fruits := []string{"apple", "banana", "cherry"}
	for i, fruit := range fruits {
		fmt.Printf("  fruits[%d] = %s\n", i, fruit)
	}

	// switch statement
	day := "Monday"
	switch day {
	case "Saturday", "Sunday":
		fmt.Println(day, "is a weekend")
	default:
		fmt.Println(day, "is a weekday")
	}

	// switch with no condition (acts like if/else chain)
	score := 85
	switch {
	case score >= 90:
		fmt.Println("Grade: A")
	case score >= 80:
		fmt.Println("Grade: B")
	case score >= 70:
		fmt.Println("Grade: C")
	default:
		fmt.Println("Grade: F")
	}
}
