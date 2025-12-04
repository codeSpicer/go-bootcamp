package main

import "fmt"

type person struct {
	name string
	age  uint8
}

type employee struct {
	person // anonymous struct with field promotion
	// employeeInfo person // named field
	employeeId string
	salary     int
}

func (p person) introduce() {
	fmt.Printf("Hi , I am  %s and I'm %d years old.\n", p.name, p.age)
}

func (e employee) introduce() { // method overriding
	fmt.Printf("Hi, I am %s and I'm %d years old. I have %d monthly salary.\n", e.name, e.age, e.salary)
}

type Speaker interface {
	speak()
}
type Animal struct{}

func (a Animal) speak() {
	fmt.Println("Animal is speaking")
}

func main() {

	emp := employee{employeeId: "xyz", salary: 40000, person: person{name: "emp 1", age: 12}}
	fmt.Println(emp)

	fmt.Println(emp.name)

	emp.introduce()

	// var s Speaker
	s := Speaker(Animal{})

	d := Animal{}

	fmt.Println(s == d)

}
