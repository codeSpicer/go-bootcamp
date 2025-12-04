package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	msg := "Hello\tworld"
	raw_msg1 := `Hello \nworld` // string literal (raw string)
	msg2 := "Hello \rworld!"
	msg3 := "12345\r678"

	fmt.Println(msg)
	fmt.Println(raw_msg1)
	fmt.Println(msg2)
	fmt.Println(msg3)

	// runes are int32 values that represent Unicode code points
	// strings are sequences of bytes (UTF-8 encoded in Go)

	fmt.Println("len(raw_msg1):", len(raw_msg1))                             // length in bytes
	fmt.Println("utf8.RuneCountInString(msg):", utf8.RuneCountInString(msg)) // number of runes (characters)

	var ch rune = 'a'
	fmt.Println("rune 'a' as int32:", ch)

	// Create a string with multiple Unicode characters (Chinese)
	anotherLanguage := ""
	for i := 26080; i < 26100; i++ {
		anotherLanguage += string(rune(i)) // go supports multiple languages
	}
	fmt.Println("anotherLanguage:", anotherLanguage)

	// Show what happens when you index a string (byte, not rune)
	fmt.Printf("anotherLanguage[0] as value: %v, as char: %c, type: %T\n", anotherLanguage[0], anotherLanguage[0], anotherLanguage[0])

	// Correct way: get the first rune (character) using utf8.DecodeRuneInString
	firstRune, size := utf8.DecodeRuneInString(anotherLanguage)
	fmt.Printf("First rune: %U '%c', size in bytes: %d\n", firstRune, firstRune, size)

	// Iterate over runes in the string
	fmt.Println("All runes in anotherLanguage:")
	for idx, r := range anotherLanguage {
		fmt.Printf("Index: %d, Rune: %U '%c', Type: %T, Value: %v\n", idx, r, r, r, r)
	}
}
