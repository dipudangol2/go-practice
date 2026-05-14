package main

import "fmt"

// Create a Job struct
type Job struct {
	name string
	role string
}

// Create a person struct
type Person struct {
	name string
	age  uint8
	// You can also declare functions as fields inside a struct
	introduce func(Person)
	// Embedded structs (Structs inside struct)
	jobs Job
	job  []Job
	// Can also have a slice like []Job

}

// Getters and Setters for the Person struct(Kind of like class methods in object oriented programming)

// Create a method for the struct person. It is called reciever argument. The variable for the struct can access this function just like the attributes
func (p Person) getName() string {
	return p.name
}

func (p Person) getAge() uint8 {
	return p.age
}

func (p Person) Introduction() {
	fmt.Printf("Hello, My name is %s and I am %v years old.\n", p.name, p.age)
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
	// Function needs to be initialized first to be used.
	p1.introduce = func(p Person) {
		fmt.Printf("Hello, My name is %s and I am %v years old.\n", p.name, p.age)
	}
	fmt.Println(p1.name, p1.age)
	name := p1.getName()
	fmt.Println(name)
	// Function invocation
	p1.introduce(p1)
	// Reciever arguments look better in these types of cases. unless the function needs to be specific for each object type reciever arguments work for general methods.
	p1.Introduction()

	// Add the struct in the Person object
	p1.jobs = Job{name: "Frontend Developer", role: "Junior"}
	// Slice of Job struct
	p1.job = make([]Job, 3)
	fmt.Println(p1.jobs, p1.job)
	jobNames := []string{"Backend Developer", "Software Engineer (DevOps)", "Project Manager"}
	roles := []string{"Junior", "Mid", "Senior"}
	// Populating like a normal slice
	for idx := 0; idx < len(p1.job); idx++ {
		p1.job[idx] = Job{name: jobNames[idx], role: roles[idx]}
	}
	fmt.Println(p1.job)

}
