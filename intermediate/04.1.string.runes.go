package intermediate

import (
	"fmt"
	"unicode/utf8"
)

func StringAndRunes() {
	fmt.Println("this is strings package")
	stringBasics()
	stringIteration()
	stringImmutability()
	stringIteration2()
	stringSlicing()
	runeBasics()
	runeIteration()
	stringToRuneConversion()
	runeToStringConversion()
	stringLengthVsRuneLength()
	utf8Validation()
}

// 1. Basics of Strings
func stringBasics() {
	fmt.Println("\n--- String Basics ---")
	s1 := "hello"
	s2 := `raw\nstring` // Raw string literal
	fmt.Println("String 1:", s1)
	fmt.Println("String 2:", s2)
	fmt.Println("Length of String 1:", len(s1)) // Length in bytes
}

// 2. String Immutability
func stringImmutability() {
	fmt.Println("\n--- String Immutability ---")
	s := "hello"
	// s[0] = 'H' // Uncommenting this will cause a compile error
	fmt.Println("Strings are immutable in Go. You cannot modify them directly.", s)
}

// 3. Iterating Over Strings
func stringIteration() {
	fmt.Println("\n--- String Iteration ---")
	s := "hello 世界"
	fmt.Println("Iterating by bytes:")
	for i := 0; i < len(s); i++ {
		fmt.Printf("Byte %d: %x\n", i, s[i])
	}

	fmt.Println("Iterating by runes:")
	for i, r := range s {
		fmt.Printf("Rune %d: %c (Unicode: %U)\n", i, r, r)
	}
}

// 4. String Slicing
func stringSlicing() {
	fmt.Println("\n--- String Slicing ---")
	s := "hello 世界"
	fmt.Println("Original String:", s)
	fmt.Println("Slice (bytes):", s[0:5]) // Slicing by bytes
	runes := []rune(s)
	fmt.Println("Slice (runes):", string(runes[6:8])) // Slicing by runes
}

// 5. Basics of Runes
func runeBasics() {
	fmt.Println("\n--- Rune Basics ---")
	r := '世' // A rune literal
	fmt.Printf("Rune: %c, Unicode: %U, Type: %T\n", r, r, r)
}

// 6. Iterating Over Runes
func runeIteration() {
	fmt.Println("\n--- Rune Iteration ---")
	s := "hello 世界"
	runes := []rune(s)
	for i, r := range runes {
		fmt.Printf("Rune %d: %c (Unicode: %U)\n", i, r, r)
	}
}

// 7. Converting String to Runes
func stringToRuneConversion() {
	fmt.Println("\n--- String to Rune Conversion ---")
	s := "hello 世界"
	runes := []rune(s)
	fmt.Println("String:", s)
	fmt.Println("Runes:", runes)
}

// 8. Converting Runes to String
func runeToStringConversion() {
	fmt.Println("\n--- Rune to String Conversion ---")
	runes := []rune{'h', 'e', 'l', 'l', 'o', ' ', '世', '界'}
	s := string(runes)
	fmt.Println("Runes:", runes)
	fmt.Println("String:", s)
}

// 9. String Length vs Rune Length
func stringLengthVsRuneLength() {
	fmt.Println("\n--- String Length vs Rune Length ---")
	s := "hello 世界"
	fmt.Println("Length in bytes:", len(s))         // Length in bytes
	fmt.Println("Length in runes:", len([]rune(s))) // Length in runes
}

// 10. UTF-8 Validation
func utf8Validation() {
	fmt.Println("\n--- UTF-8 Validation ---")
	valid := "hello 世界"
	invalid := string([]byte{0xff, 0xfe, 0xfd}) // Invalid UTF-8 sequence

	fmt.Println("Valid UTF-8 String:", valid, "Is Valid?", utf8.ValidString(valid))
	fmt.Println("Invalid UTF-8 String:", invalid, "Is Valid?", utf8.ValidString(invalid))
}
