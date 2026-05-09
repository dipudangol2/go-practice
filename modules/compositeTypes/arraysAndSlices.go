package main

import "fmt"

func ArrayDemo() {
	fmt.Println("Array Demo:")

	// 1D Array declaration and initialization. Must have the type specified during initialization
	var arr [2]bool = [2]bool{true, false}
	// Using composite literal
	array := [2]int{1, 2}

	fmt.Printf("Boolean Array: %v\n", arr)
	fmt.Printf("Integer Array: %v\n", array)

	// 2D Array declaration and initialization. [...][2] allows dynamic numbers of array for arrays of size 2
	var twoDimensionalArr [2][2]int = [2][2]int{{1, 2}, {1, 3}}

	// Using composite literal to declare. Has some flexibility for number of rows as they can be dynamic
	twoDimensionalArr2 := [...][2]int{{1, 2}, {2, 3}, {3, 4}}

	fmt.Printf("Two dimensional Array of size [2][2]: %v\n", twoDimensionalArr)
	fmt.Printf("Two dimensional Array of size [%d][%d]: %v\n", len(twoDimensionalArr2), len(twoDimensionalArr2[0]), twoDimensionalArr2)

	// Iterating on an array
	//1D array: Iterating on arr
	// Usual method with indexes

	// 1D array
	fmt.Print("Boolean Array iterated over with idx variable: [")
	for idx := 0; idx < len(arr); idx++ {
		fmt.Printf(" %v", arr[idx])
	}
	fmt.Print(" ]\n")

	fmt.Printf("Two dimensional Array of size [2][2] iterated over with idx variable:[")
	for idx := 0; idx < len(twoDimensionalArr); idx++ {
		fmt.Print("[")
		for idxCol := 0; idxCol < len(twoDimensionalArr[0]); idxCol++ {
			fmt.Printf("%v", twoDimensionalArr[idx][idxCol])
			if idxCol != len(twoDimensionalArr[0])-1 {
				fmt.Print(",")
			}
		}
		fmt.Print("]")
		if idx != len(twoDimensionalArr)-1 {
			fmt.Print(" ")
		}

	}
	fmt.Print("]\n")

	fmt.Print("Integer Array iterated over with range function implementation: [")
	// If any value is not intended to be used then just use _ for the value.
	//1D Array
	for idx, elem := range array {
		fmt.Printf("%v", elem)
		if idx != len(array)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Print("]\n")
	// 2D Array
	fmt.Printf("Two dimensional Array of size [%d][%d] iterated over range funciton: [", len(twoDimensionalArr2), len(twoDimensionalArr2[0]))
	for idx, nested := range twoDimensionalArr2 {
		fmt.Print("[")
		for idx2, elem := range nested {
			fmt.Printf("%v", elem)
			if idx2 != len(nested)-1 {
				fmt.Print(",")
			}
		}
		fmt.Print("]")
		if idx != len(twoDimensionalArr2)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Print("]\n")
}

func SliceDemo() {

	arr := []int{1, 2, 3, 4, 5, 6}

	//Creating a slice from the arr array from index 1 to 5. Elements of index 1,2,3,4 are available, 5 is not included in the slice [] denotes slice type in go
	slice := arr[1:5]

	fmt.Printf("The original array is: %v\n", arr)
	fmt.Printf("The slice from the array [1:5] is: %v. It does not include the elements of the index 5.\n", slice)

	// Mutating the slice affects the array as well
	slice[2] = 99
	fmt.Printf("slice[2] was changed to 99. The array should also be affected by the slice mutation.\nArray: %v\nSlice:%v\n", arr, slice)

	// Slices have three properties:
	// - pointer: Which points to the first element of the array defined in the slice if derived from the array
	// - length: The total number of elements in the slice. function to access length is len().
	// - capacity: Different from length. Refers to how many elements the slice can handle. Slices can be redefined unlike arrays. function to access capacity is cap().
	fmt.Printf("Slice: %v, length:%v, capacity:%v\n", slice, len(slice), cap(slice))

	//Reslicing the slice. It can increase the capacity of the slice now if the original capacity was larger than the length of the slice.
	slice = slice[:5]
	fmt.Printf("Slice: %v, length:%v, capacity:%v\n", slice, len(slice), cap(slice))

	// Demonstrating slice's auto expanding feature.
	// Create a slice of integers with two numbers
	sl := []int{1, 2}

	fmt.Println("Demonstration for slice auto expand nature: ")
	fmt.Printf("Intial Array: %v\n", sl)

	// Add elements to the slice but print the slice, length and capacity of the slice on each step when adding ten elements to the slice
	// Whenever appending to the slice, if the capacity is full for the slice, the capacity for the slice gets doubled.
	for idx := 2; idx < 12; idx++ {
		sl = append(sl, idx+1)
		fmt.Printf("Slice: %v, length:%v, capacity:%v\n", sl, len(sl), cap(sl))
	}
	// Slices are mutable when passed to a function unlike arrays which create a copy of the array passed in the function
	// Create a test slice for checking mutation
	slce := []int{1, 3, 4, 5, 6}
	fmt.Printf("Before passing the slice to the function: %v\n", slce)
	testSliceMutation(slce)
	fmt.Printf("After passing the slice to the function: %v\n", slce)

}

func testSliceMutation(sl []int) {
	sl[0] = 2
}
