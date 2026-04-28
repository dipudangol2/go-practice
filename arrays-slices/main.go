package main

import "fmt"

func main() {
	// Arrays – fixed size
	var arr [5]int
	arr[0] = 10
	arr[4] = 50
	fmt.Println("Array:", arr)

	colors := [3]string{"red", "green", "blue"}
	fmt.Println("Colors:", colors)

	// Slices – dynamic, backed by an array
	s := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice:", s)
	fmt.Printf("len=%d cap=%d\n", len(s), cap(s))

	// Slicing
	fmt.Println("s[1:3]:", s[1:3])
	fmt.Println("s[:2]:", s[:2])
	fmt.Println("s[3:]:", s[3:])

	// append
	s = append(s, 6, 7)
	fmt.Println("After append:", s)

	// make
	made := make([]int, 3, 5)
	fmt.Printf("made: %v len=%d cap=%d\n", made, len(made), cap(made))

	// copy
	dst := make([]int, len(s))
	n := copy(dst, s)
	fmt.Printf("Copied %d elements: %v\n", n, dst)

	// 2D slice
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("Matrix:")
	for _, row := range matrix {
		fmt.Println(" ", row)
	}
}
