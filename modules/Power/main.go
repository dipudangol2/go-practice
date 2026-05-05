package main

import "fmt"

// Declaring an interface for ambiguous function parameter for power calculation
type Number interface {
	int | float32 | float64
}

func Power[T Number](base T, exponent int) T {
	if exponent == 0 {
		return T(1)
	}
	if exponent < 0 {
		return T(1) / Power(base, -exponent)
	}
	return base * Power(base, exponent-1)

}
func main() {
	var (
		base     float64
		exponent int
	)
	fmt.Println("Power function implementation (No fractional exponents:)")
	fmt.Print("Enter the Base: ")
	fmt.Scan(&base)
	fmt.Print("Enter the Exponent: ")
	fmt.Scan(&exponent)
	result := Power(base, exponent)
	fmt.Printf("pow(%v,%d)=%v\n", base, exponent, result)

}
