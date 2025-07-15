package main

import (
	"fmt"
	"math"
)

// --- INTERFACES IN GO ---
// An interface type specifies a set of method signatures (behavior).
// Any type that implements those methods satisfies the interface, implicitly (no explicit declaration needed).

// geometry interface defines two methods: area() and perim()
type geometry interface {
	area() float64
	perim() float64
}

// measure takes any geometry type and calls its area and perim methods
// This demonstrates polymorphism: measure works with any type that implements geometry
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// rectangle struct implements geometry by defining area() and perim()
type rectangle struct {
	width, height float64
}

// area method for rectangle (implements geometry.area)
func (r rectangle) area() float64 {
	return r.height * r.width
}

// perim method for rectangle (implements geometry.perim)
func (r rectangle) perim() float64 {
	return 2 * (r.height + r.width)
}

// circle struct also implements geometry by defining area() and perim()
type circle struct {
	radius float64
}

// area method for circle (implements geometry.area)
func (c circle) area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}

// perim method for circle (implements geometry.perim)
func (c circle) perim() float64 {
	return 2 * c.radius * math.Pi
}

// circle has an extra method not required by geometry
func (c circle) diameter() float64 {
	return 2 * c.radius
}

func main() {
	// Both rectangle and circle implement geometry, so they can be used with measure
	r := rectangle{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
