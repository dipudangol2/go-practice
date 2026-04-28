package main

import "fmt"

// Basic function
func greet(name string) string {
	return "Hello, " + name + "!"
}

// Multiple return values
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

// Named return values
func minMax(nums []int) (min, max int) {
	min, max = nums[0], nums[0]
	for _, v := range nums[1:] {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return
}

// Variadic function
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// Higher-order function: accepts a function as argument
func apply(nums []int, f func(int) int) []int {
	result := make([]int, len(nums))
	for i, v := range nums {
		result[i] = f(v)
	}
	return result
}

// Closure
func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	fmt.Println(greet("Gopher"))

	result, err := divide(10, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 3 = %.4f\n", result)
	}

	_, err = divide(5, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}

	nums := []int{3, 1, 4, 1, 5, 9, 2, 6}
	min, max := minMax(nums)
	fmt.Printf("min=%d, max=%d\n", min, max)

	fmt.Println("sum(1,2,3,4,5) =", sum(1, 2, 3, 4, 5))

	doubled := apply([]int{1, 2, 3, 4}, func(n int) int { return n * 2 })
	fmt.Println("doubled:", doubled)

	counter := makeCounter()
	fmt.Println(counter(), counter(), counter())
}
