package main

import "fmt"

// Create a person struct
type Person struct {
	name string
	age  uint8
}

// Getters and Setters for the Person struct(Kind of like class methods in object oriented programming)

// Create a method for the struct person. It is called reciever argument. The variable for the struct can access this function just like the attributes
func (p Person) getName() string {
	return p.name
}

func (p Person) getAge() uint8 {
	return p.age
}

// Add pointers for if you want to modify the original content of the struct
func (p *Person) setName(name string) {
	p.name = name
}

func (p *Person) setAge(age uint8) {
	p.age = age
}

func StructDemo() {
	// Create an object for the struct type
	p1 := Person{age: 23}
	p1.setName("Dipu Dangol")
	fmt.Println(p1.name, p1.age)
	name := p1.getName()
	fmt.Println(name)
}
