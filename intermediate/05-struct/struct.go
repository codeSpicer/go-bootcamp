package main

import "fmt"

// Address represents a physical address with country and city fields.
type Address struct {
	country string
	city    string
}

// PhoneHomeCell demonstrates an embedded (anonymous) struct for phone numbers.
type PhoneHomeCell struct {
	home string
	cell string
}

// Person demonstrates a struct with nested and anonymous fields.
type Person struct {
	firstName     string
	lastName      string
	age           int
	address       Address
	PhoneHomeCell // Embedded struct: fields 'home' and 'cell' are promoted to Person
}

func main() {
	// Initializing a struct using composite literal (all fields specified)
	p := Person{
		firstName: "Akshat",
		lastName:  "Kumar",
		age:       23,
		address: Address{
			country: "india",
			city:    "delhi",
		},
		PhoneHomeCell: PhoneHomeCell{
			cell: "93345632626",
			home: "xm-780",
		},
	}

	// Initializing a struct and setting fields later (dot notation)
	p1 := Person{
		firstName: "person",
		age:       35,
	}
	p1.address.city = "haridwar"
	p1.address.country = "india"
	p1.cell = "2345246345" // Accessing promoted field from embedded struct

	fmt.Println("Person p:", p)
	fmt.Println("Person p1:", p1)

	// Anonymous struct: useful for one-off data structures
	user := struct {
		userName string
		age      int
	}{
		userName: "user",
		age:      21,
	}
	fmt.Println("Anonymous user struct:", user)

	// Struct comparison: possible if all fields are comparable
	p2 := Person{
		firstName: "Akshat",
		lastName:  "Kumar",
		age:       23,
		address: Address{
			country: "india",
			city:    "delhi",
		},
		PhoneHomeCell: PhoneHomeCell{
			cell: "93345632626",
			home: "xm-780",
		},
	}
	fmt.Println("p == p2:", p == p2) // true if all fields are equal
}
