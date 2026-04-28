package main

import (
	"fmt"
	"math"
)

// Shape interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64      { return math.Pi * c.Radius * c.Radius }
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.Radius }

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64      { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64 { return 2 * (r.Width + r.Height) }

type Triangle struct {
	A, B, C float64 // side lengths
}

func (t Triangle) Area() float64 {
	s := (t.A + t.B + t.C) / 2
	return math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
}

func (t Triangle) Perimeter() float64 { return t.A + t.B + t.C }

// printShapeInfo accepts any Shape
func printShapeInfo(s Shape) {
	fmt.Printf("%T: area=%.2f, perimeter=%.2f\n", s, s.Area(), s.Perimeter())
}

// Stringer interface (from fmt package)
type Celsius float64
type Fahrenheit float64

func (c Celsius) String() string    { return fmt.Sprintf("%.1f°C", float64(c)) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%.1f°F", float64(f)) }

// Empty interface and type assertions
func describe(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("int: %d\n", v)
	case string:
		fmt.Printf("string: %q\n", v)
	case bool:
		fmt.Printf("bool: %t\n", v)
	case Shape:
		fmt.Printf("Shape with area %.2f\n", v.Area())
	default:
		fmt.Printf("unknown type: %T\n", v)
	}
}

func main() {
	shapes := []Shape{
		Circle{Radius: 3},
		Rectangle{Width: 4, Height: 5},
		Triangle{A: 3, B: 4, C: 5},
	}

	for _, s := range shapes {
		printShapeInfo(s)
	}

	fmt.Println()

	// Stringer
	temp := Celsius(100)
	fmt.Println("Boiling point:", temp)

	// Type assertions
	describe(42)
	describe("hello")
	describe(true)
	describe(Circle{Radius: 7})
	describe(3.14)
}
