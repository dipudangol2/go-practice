package main

import (
	"fmt"
	"math"
)

// Declare an interface
// Nothing will be shape interface and it cannot be instantiated.
// Various other structs can implement this interface. It is like a map for a struct which derives this interface
type Shape interface {
	getPerimeter() float32
	getArea() float32
}

// Create a Triangle struct to later assign to the shape interface
type Triangle struct {
	a, b, c float32
}

// Create another Square struct
type Square struct {
	length float32
}

func (t *Triangle) setSides(a float32, b float32, c float32) {
	t.a = a
	t.b = b
	t.c = c
}
func (t Triangle) getPerimeter() float32 {
	return t.a + t.b + t.c
}

func (t Triangle) getArea() float32 {
	s := (t.a + t.b + t.c) / 2
	return float32(math.Sqrt(float64(s * (s - t.a) * (s - t.b) * (s - t.c))))
}

func (s Square) getPerimeter() float32 {
	return 4 * s.length
}
func (s Square) getArea() float32 {
	return float32(math.Pow(float64(s.length), 2))
}

func InterfaceDemo() {
	fmt.Println("Interfaces in Go")
	t := Triangle{}
	t.setSides(4, 3, 6)
	fmt.Printf("Sides a, b, c : %v, %v, %v \n", t.a, t.b, t.c)
	fmt.Printf("Perimeter: %v, Area: %v\n", t.getPerimeter(), t.getArea())
	// Declare the triangle object as shape interface
	var ti Shape = Triangle{2, 3, 4}
	// As the dimensions were not declared in the interface, they cannot be used with the shape object
	fmt.Printf("Perimeter: %v, Area: %v\n", ti.getPerimeter(), ti.getArea())
	// Create array for the interface with multiple structs, works as long as the passed structs implement the interface properly
	var shapes []Shape = []Shape{Triangle{5, 6, 3}, Square{5}}

	fmt.Println("Perimeter and area of triangle struct followed by square struct which both follow the shape interface so can be put in an array")
	for _, shape := range shapes {
		fmt.Printf("Perimeter: %v, Area: %v\n", shape.getPerimeter(), shape.getArea())
	}
}
