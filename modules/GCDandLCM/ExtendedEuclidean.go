package GCDandLCM

import "fmt"

// passing pointers for a and b with *
func gcdExtended(x, y int, a, b *int) int {
	//base case after x becomes 0
	if x == 0 {
		*a = 0
		*b = 1
		return y
	}
	var a1, b1 int
	// During call pointers need to be passed addresses
	gcd := gcdExtended(y%x, x, &a1, &b1)

	// Update a and b with results of recursive call
	*a = b1 - (y/x)*a1
	*b = a1

	return gcd
}

func calculateGCD(x, y int) int {
	a, b := 1, 1
	return gcdExtended(x, y, &a, &b)
}

func lcmExtended(x int, y int) int {
	if x == 0 && y == 0 {
		return 0
	}
	return Abs(x*y) / calculateGCD(x, y)

}
func ExtendedEuclidean() {
	fmt.Println("GCD and LCM (Extended Euclidean approach)")
	fmt.Print("Enter the two numbers: ")
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)

	gcdResult := calculateGCD(a, b)
	lcmResult := lcmExtended(a, b)
	fmt.Printf("GCD(%d,%d) = %d\nLCM(%d,%d) = %d\n", a, b, gcdResult, a, b, lcmResult)
}
