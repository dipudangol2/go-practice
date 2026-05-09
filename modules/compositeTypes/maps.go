package main

import "fmt"

func MapDemo() {
	// Declare a map [key]value
	var mp map[int]string = map[int]string{
		0: "Dipu",
		1: "Dangol",
	}
	fmt.Println(mp)
	// Implicit assignment. for empty map just leave braces
	mp2 := map[string]int{"a": 1}
	fmt.Println(mp2)

	// Using make function, the below example initializes empty map of string keys and slices of int
	mp3 := make(map[string][]int)
	mp3["integers"] = []int{-2, -1, 0, 1, 2, 3}
	mp3["positive"] = []int{1, 2, 3, 4, 5}
	mp3["keyToDelete"] = []int{4, 5, 6, 67, 7, 33}
	fmt.Println(mp3)

	// To delete a pair use delete function. doesnot need assignment like append function for slices
	delete(mp3, "keyToDelete")
	fmt.Println(mp3)

	// Check if the key is present inside the MapDemo
	value, ok := mp3["integers"]
	fmt.Println(value, ok)
	value, ok = mp3["keyToDelete"]
	fmt.Println(value, ok)

}
