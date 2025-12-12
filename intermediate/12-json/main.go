package main

import (
	"encoding/json"
	"fmt"
)

// Address struct defines the structure for location data.
type Address struct {
	City  string `json:"city"`  // JSON tag ensures field maps to 'city' in JSON
	State string `json:"state"` // JSON tag ensures field maps to 'state' in JSON
}

// Person struct demonstrates basic JSON marshaling.
// Note: FirstName uses the default Go name "FirstName" in JSON unless tagged.
type Person struct {
	FirstName string  `json:"name"`          // Tagged field: Go's FirstName maps to JSON's "name"
	Age       int     `json:"age,omitempty"` // omitempty: Skips the field if Age is 0 or its zero value
	Email     string  `json:"email,omitempty"`
	Address   Address `json:"address"`
}

// Employee struct represents a standard data model for an API resource.
type Employee struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Active  bool    `json:"active"`
	Address Address `json:"address"`
}

func main() {
	// --- PART 1: Basic Marshaling Example (Go Struct -> JSON bytes) ---
	fmt.Println("--- PART 1: Basic Marshaling (Go -> JSON) ---")

	// Create a Person struct instance
	person := Person{
		FirstName: "John",
		Age:       19, // Age is non-zero, so it will be included
		Address:   Address{City: "Noida", State: "UP"},
	}

	fmt.Printf("Initial Go Struct: %+v\n", person)

	// json.Marshal converts the Go struct into a slice of JSON bytes.
	// It uses the struct tags (e.g., `json:"name"`) to determine the output keys.
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling Person struct:", err)
		return
	}

	// Outputting the raw JSON bytes (which are just numbers/integers)
	fmt.Println("Raw JSON Bytes:", jsonData)
	// Converting the bytes to a string for readable JSON output
	fmt.Println("Readable JSON:", string(jsonData))
	// Output will show 'name' instead of 'FirstName' due to the struct tag.

	// --- PART 2: Unmarshaling Example (JSON string -> Go Struct) ---
	fmt.Println("\n--- PART 2: Unmarshaling (JSON -> Go) ---")

	// Example JSON string representing an existing employee record
	jsonEmployee := `{"id":1,"name":"Jeanette","email":"jpenddreth0@census.gov","active":true,"address":{"city":"Old City","state":"Anytown"}}`

	var data Employee // Declare an empty struct to hold the parsed data

	// json.Unmarshal parses the JSON bytes and populates the fields of 'data'.
	// The SECOND argument MUST be a POINTER (&data) so that the function can
	// modify the original 'data' variable, not just a copy.
	err = json.Unmarshal([]byte(jsonEmployee), &data)

	if err != nil {
		fmt.Println("Error unmarshalling employee data:", err)
		return
	}

	fmt.Printf("Unmarshaled Employee Data: %+v\n", data) // descriptive output that shows the structure's data

	// --- PART 3: Updating and Re-Marshaling an Employee ---
	fmt.Println("\n--- PART 3: Update Employee Data and Re-Marshal ---")

	// 3a. Modify the data (Simulating an update logic)
	fmt.Println("--- Applying Updates ---")
	data.Name = "Changed Name"
	data.Address.City = "New York City"
	data.Address.State = "NY"

	fmt.Printf("Updated Go Struct: %+v\n", data)

	// 3b. Marshal the updated struct back into a JSON string
	// This is typically what an API would send back to the client or save to a DB.
	updatedJsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error re-marshalling updated data:", err)
		return
	}

	updatedJsonEmployee := string(updatedJsonData)

	fmt.Println("Re-Marshaled JSON (Updated Record):")
	fmt.Println(updatedJsonEmployee)

	// --- PART 4: Marshaling a List of Addresses (Demonstrates Slices) ---
	fmt.Println("\n--- PART 4: Marshaling a Slice of Structs ---")

	// Define a slice of Address structs
	locations := []Address{
		{City: "San Francisco", State: "CA"},
		{City: "Austin", State: "TX"},
		{City: "Miami", State: "FL"},
	}

	// Marshal the slice into a JSON array
	locationsData, err := json.Marshal(locations)
	if err != nil {
		fmt.Println("Error marshalling locations:", err)
		return
	}

	fmt.Println("JSON Array Output:")
	fmt.Println(string(locationsData))

	// --- PART 5 : Handling unknown json structs-----

	fmt.Println("\n--- PART 5: Handling unknown json structs ---")

	jsonData2 := `{"name":"person", "age":23, "address":{"city":"noida"}}`

	var unknownData map[string]interface{}

	json.Unmarshal([]byte(jsonData2), &unknownData)

	fmt.Println(unknownData)
	fmt.Println(unknownData["name"])
}
