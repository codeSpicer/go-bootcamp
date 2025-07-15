package main

import (
	"fmt"
	"math"
)

// Pair is a simple struct with two integer fields.
type Pair struct {
	first  int
	second int
}

// Method with value receiver:
// add is a method on the Pair type. It can be called using a Pair value (e.g., pair.add()).
// In Go, a method is a function with a special receiver argument. The receiver appears in its own argument list between the func keyword and the method name.
// Here, the receiver is (p Pair), so this method gets a copy of the Pair value.
func (p Pair) add() int {
	return p.first + p.second
}

// Regular function (not a method):
// add is a normal function that takes a Pair as an argument.
// It is not associated with the Pair type, so it must be called as add(pair).
func add(p Pair) int {
	return p.first + p.second
}

// MyFloat is a custom type based on float64.
// Methods can be declared on any type defined in the same package, not just structs.
type MyFloat float64

// Method on a non-struct type:
// Abs is a method with a value receiver of type MyFloat.
// It returns the absolute value of the MyFloat.
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Method with pointer receiver:
// Multiply is a method with a pointer receiver (*MyFloat).
// Methods with pointer receivers can modify the value that their receiver points to.
func (f *MyFloat) Multiply() {
	*f *= 2
}

func main() {
	// Creating a Pair instance
	pair := Pair{
		first:  10,
		second: 15,
	}

	// Calling the method and the function
	fmt.Println(pair.add()) // Method call: uses receiver syntax
	fmt.Println(add(pair))  // Function call: passes struct as argument

	// Creating a MyFloat instance
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs()) // Calls Abs method

	f.Multiply() // Calls Multiply method, modifies f in place
	fmt.Println(f)

	s := Shape{
		Rectangle: Rectangle{length: 4, breadth: 3},
	}
	fmt.Println("Area of reactange shape is: ", s.Area()) // embbeded struct magic
	fmt.Println("Area of reactange shape is: ", s.Rectangle.Area())
}

type Rectangle struct {
	length  int
	breadth int
}

func (r Rectangle) Area() int {
	return r.length * r.breadth
}

type Shape struct {
	Rectangle
}

/*
Summary of Go Methods:
- A method is a function with a special receiver argument.
- The receiver can be a value or a pointer.
- Methods can be declared on any type defined in the same package, not just structs.
- Use value receivers when the method does not need to modify the receiver.
- Use pointer receivers to modify the receiver or to avoid copying large structs.
- Methods provide a way to associate behavior with types, similar to classes in other languages.
*/
