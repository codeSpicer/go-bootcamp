package main

import "fmt"

func main() {

	var a int = 32
	b := int32(a)
	c := float64(b)

	e := 45.23
	f := int(e)

	fmt.Println(a, b, c, e, f)

	// Type( value )

	g := "Hello"

	var h []byte
	h = []byte(g)

	fmt.Println(h)

	// The Chinese character for "World"
	chineseChar := "界"

	fmt.Println("--- 1. String to Byte Slice ---")

	// Convert the string to a raw byte slice
	d := []byte(chineseChar)

	fmt.Printf("String: %s\n", chineseChar)
	fmt.Printf("Length (bytes): %d\n", len(d)) // Output: 3 (3 bytes)

	// Print the decimal value of each byte in the slice
	fmt.Print("Byte values (decimal): [")
	for _, val := range d {
		// All these values are uint8, max 255.
		fmt.Printf("%d ", val)
	}
	fmt.Println("]") // Output: [231 150 140 ] (The individual byte pieces)

	// --- 2. String to Rune Slice ---

	// Convert the string to a rune slice (slice of Unicode code points/characters)
	r := []rune(chineseChar)

	fmt.Println("\n--- 2. String to Rune Slice ---")
	fmt.Printf("Length (runes): %d\n", len(r)) // Output: 1 (1 character)

	// Print the decimal value of the rune (the Unicode code point)
	fmt.Printf("Rune value (decimal): %d\n", r[0]) // Output: 29900
	fmt.Printf("Rune value (hex): 0x%X\n", r[0])   // Output: 0x754C (U+754C)

	// --- 3. Byte Slice to String ---

	// Converting the byte slice back to a string correctly reconstructs the character
	s := string(b)
	fmt.Println("\n--- 3. Byte Slice back to String ---")
	fmt.Printf("Reconstructed String: %s\n", s) // Output: 界

}
