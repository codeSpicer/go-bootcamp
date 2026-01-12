package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// User represents a data model with Struct Tags.
// Tags like `json:"name"` tell Go exactly how to name the keys in the resulting JSON.
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// --- PART 1: Marshal & Unmarshal (In-Memory) ---
	// Use these when you already have the data as a variable or a byte slice.

	user := User{Name: "akshat", Email: "test@test.com"}

	// Marshal: Go Object -> []byte (JSON)
	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Marshaled JSON:", string(jsonData))

	// Unmarshal: []byte (JSON) -> Go Object
	var user1 User
	err = json.Unmarshal(jsonData, &user1)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("Unmarshaled Struct: %+v\n", user1)

	// --- PART 2: NewDecoder & NewEncoder (Streaming) ---
	// Use these when reading from/writing to a network connection or a file (io.Reader/Writer).

	jsonData1 := `{"name":"username","email":"testuseremail@email"}`

	// NewReader simulates a network stream (io.Reader)
	reader := strings.NewReader(jsonData1)
	
	// Create a Decoder that "listens" to the stream
	decoder := json.NewDecoder(reader)

	var user2 User
	// Decode reads the stream and populates the struct directly
	err = decoder.Decode(&user2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Decoded from Stream: %+v\n", user2)
}