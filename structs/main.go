package main

import (
	"fmt"
	"math"
)

// Basic struct
type Person struct {
	Name string
	Age  int
}

// Method on struct (value receiver)
func (p Person) String() string {
	return fmt.Sprintf("%s (age %d)", p.Name, p.Age)
}

// Method with pointer receiver (can modify the struct)
func (p *Person) Birthday() {
	p.Age++
}

// Embedding (composition)
type Employee struct {
	Person
	Company string
	Salary  float64
}

func (e Employee) String() string {
	return fmt.Sprintf("%s at %s, salary: $%.2f", e.Person.String(), e.Company, e.Salary)
}

// Struct with methods for a geometric shape
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func main() {
	// Struct literal
	p := Person{Name: "Alice", Age: 30}
	fmt.Println(p)

	// Pointer to struct
	pp := &Person{Name: "Bob", Age: 25}
	pp.Birthday()
	fmt.Println(pp)

	// Embedding
	emp := Employee{
		Person:  Person{Name: "Charlie", Age: 28},
		Company: "Gopher Inc.",
		Salary:  75000,
	}
	fmt.Println(emp)
	fmt.Println("Employee name:", emp.Name) // promoted field

	// Shapes
	c := Circle{Radius: 5}
	fmt.Printf("Circle  – area=%.2f perimeter=%.2f\n", c.Area(), c.Perimeter())

	r := Rectangle{Width: 4, Height: 6}
	fmt.Printf("Rectangle – area=%.2f perimeter=%.2f\n", r.Area(), r.Perimeter())
}
