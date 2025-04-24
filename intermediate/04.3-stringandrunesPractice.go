package intermediate

import (
	"fmt"
	"unsafe"
)

func StringsPractice() {
	// one()
	// stringIteration3()
	// properIteration()
	// runeExample()
	// stringMemoryLayout2()
	runeBasics3()
}

func one() {
	var name string = "shivam"
	middleName := "kumar"
	lastName := `Sharma`

	fmt.Println(name)
	fmt.Println(middleName)
	fmt.Println(lastName)

	fullName := name + middleName + lastName

	fmt.Println(fullName)
}

func stringIteration3() {
	fmt.Println("\n--- String Iteration ---")
	s := "hello 世界"

	fmt.Println(s[0]) // Output: 104 (byte value of 'h')
	fmt.Println(s[6]) // Output: 228 (first byte of '世')

	fmt.Println("The lenght of of string s:", len(s)) //12

	/*
		len(s) returns the number of bytes in the string, not the number of characters.
		The first 5 characters (hello) are ASCII characters,
		each represented by a single byte.
		The 6th character (世) is a Unicode character, represented by 3 bytes (228 184 150).
		Accessing s[6] gives the first byte of 世 (228), not the entire character.
	*/
	fmt.Println("Iterating by bytes:")
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
	/*
		Iterating by bytes:
		104
		101
		108
		108
		111
		32
		228
		184
		150
		231
		149
		140
	*/

	fmt.Println("Iterating by runes:")
	for i, r := range s {
		fmt.Println(i, r)
	}

	/*
		Iterating by runes:
		//all are decimal value and they represent codepoint
		//"世".charCodeAt() =>19990
		0 104
		1 101
		2 108
		3 108
		4 111
		5 32
		6 19990
		9 30028
		console.log("世".charCodeAt(0)); // Output: 19990
	*/
}

func properIteration() {
	s := "hello 世界"
	for i := 0; i < len(s); i++ {
		//%x is hexadecimal representaion
		fmt.Printf("Byte %d: %x\n", i, s[i])
	}
	/*
		Byte 0: 68
		Byte 1: 65
		Byte 2: 6c
		Byte 3: 6c
		Byte 4: 6f
		Byte 5: 20
		Byte 6: e4
		Byte 7: b8
		Byte 8: 96
		Byte 9: e7
		Byte 10: 95
		Byte 11: 8c
	*/

	fmt.Println("Iterating by runes:")
	for i, r := range s {
		/*
			%d is decimal representaion
			%c is is chracter, you can use it for rune formatting
			%U is unicode
		*/
		fmt.Printf("Rune %d: %c (Unicode: %U, Raw Value: %d, Hex value:%x)\n", i, r, r, r, r)
	}
	/*
		Rune 0: h (Unicode: U+0068, Raw Value: 104, Hex value:68)
		Rune 1: e (Unicode: U+0065, Raw Value: 101, Hex value:65)
		Rune 2: l (Unicode: U+006C, Raw Value: 108, Hex value:6c)
		Rune 3: l (Unicode: U+006C, Raw Value: 108, Hex value:6c)
		Rune 4: o (Unicode: U+006F, Raw Value: 111, Hex value:6f)
		Rune 5:   (Unicode: U+0020, Raw Value: 32, Hex value:20)
		Rune 6: 世 (Unicode: U+4E16, Raw Value: 19990, Hex value:4e16)
		Rune 9: 界 (Unicode: U+754C, Raw Value: 30028, Hex value:754c)

		"世".charCodeAt()=>19990 //deimal value
		"界".charCodeAt()=>30028 // decimal value
	*/

}

func runeExample() {
	s := "界"
	r := []rune(s) // Convert string to runes
	fmt.Printf("Character: %c, Unicode: %U, Decimal: %d\n", r[0], r[0], r[0])
	//Character: 界, Unicode: U+754C, Decimal: 30028
}

// 2. Memory layout of strings
func stringMemoryLayout2() {
	fmt.Println("\n--- String Memory Layout ---")

	str := "hello"
	fmt.Println("string pointer", &str)

	// Using unsafe to examine string internals
	hdr := (*struct {
		data uintptr
		len  int
	})(unsafe.Pointer(&str))

	fmt.Printf("String: %q\n", str)
	fmt.Printf("Data pointer: %#x\n", hdr.data)
	fmt.Printf("Length: %d bytes\n", hdr.len)

	// Demonstrating string interning
	str1 := "hello"
	str2 := "hello"
	hdr1 := (*struct{ data uintptr })(unsafe.Pointer(&str1))
	hdr2 := (*struct{ data uintptr })(unsafe.Pointer(&str2))
	fmt.Printf("\nString interning (same contents point to same data):\n")
	fmt.Printf("str1 data: %#x\nstr2 data: %#x\n", hdr1.data, hdr2.data)
}

// 3. Rune basics and conversion
func runeBasics3() {
	fmt.Println("\n--- Rune Basics ---")

	// Rune declarations
	r1 := 'A'
	r2 := '世'
	r3 := '\u4e16'     // 世
	r4 := '\U0001f600' // 😀

	fmt.Printf("Runes: %c, %c, %c, %c\n", r1, r2, r3, r4)
	fmt.Printf("Rune types: %T, %T, %T, %T\n", r1, r2, r3, r4)
	fmt.Printf("Rune values: %d, %d, %d, %d\n", r1, r2, r3, r4)

	// String to rune conversion
	s := "hello 世界"
	runes := []rune(s)
	bytes := []byte(s)

	fmt.Println("\nString:", s)
	fmt.Println("As runes:", runes)
	fmt.Println("As bytes:", bytes)
	fmt.Printf("Byte length: %d, Rune count: %d\n", len(bytes), len(runes))

	// Rune to string conversion
	fmt.Println("\nRune to string:")
	fmt.Println(string(r1), string(r2), string(r3), string(r4))
}
