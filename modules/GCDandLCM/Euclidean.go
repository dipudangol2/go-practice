package GCDandLCM

import "fmt"

// helper abs function as go doesnot have one for int
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Calculate gcd with euclidean algorithm
func gcd(x int, y int) int {
	if x == 0 {
		return y
	}
	return gcd(y%x, x)

}

func lcm(x int, y int) int {
	if x == 0 && y == 0 {
		return 0
	}
	return abs(x*y) / gcd(x, y)

}

func Euclidean() {
	fmt.Println("GCD and LCM of two numbers using Euclidean Algorithm")
	fmt.Print("Enter the two numbers: ")
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)

	gcdResult := gcd(a, b)
	lcmResult := lcm(a, b)
	fmt.Printf("GCD(%d,%d) = %d\nLCM(%d,%d) = %d\n", a, b, gcdResult, a, b, lcmResult)
}
