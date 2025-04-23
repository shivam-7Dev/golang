package intermediate

import (
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
	"unsafe"
)

func StringAndRunes2() {
	fmt.Println("\n=== Strings and Runes in Go ===\n")

	stringDeclarationAndProperties()
	stringMemoryLayout()
	runeBasics2()
	stringIteration2()
	stringSlicingPitfalls()
	stringManipulation()
	unicodeHandling()
	performanceConsiderations()
	commonPitfalls()
	bestPractices()
}

// 1. String declaration and basic properties
func stringDeclarationAndProperties() {
	fmt.Println("\n--- String Declaration and Properties ---")

	// Basic string declaration
	asciiStr := "hello"
	unicodeStr := "hello 世界"
	rawStr := `raw\nstring`
	multilineStr := `This is
a multi-line
string`

	fmt.Println("ASCII string:", asciiStr)
	fmt.Println("Unicode string:", unicodeStr)
	fmt.Println("Raw string:", rawStr)
	fmt.Println("Multiline string:", multilineStr)

	// String properties
	fmt.Println("\nProperties:")
	fmt.Println("Length of 'hello':", len(asciiStr), "bytes")      // 5
	fmt.Println("Length of 'hello 世界':", len(unicodeStr), "bytes") // 11 (5 + 3*2)
	fmt.Println("Immutable check:")
	fmt.Printf("Address of 'hello': %p\n", &asciiStr)
	asciiStr += " world"
	fmt.Printf("After modification: %p (new string created)\n", &asciiStr)
}

// 2. Memory layout of strings
func stringMemoryLayout() {
	fmt.Println("\n--- String Memory Layout ---")

	str := "hello"

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
func runeBasics2() {
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

// 4. String iteration methods
func stringIteration2() {
	fmt.Println("\n--- String Iteration ---")

	s := "hello 世界"

	fmt.Println("\nByte-by-byte iteration (incorrect for Unicode):")
	for i := 0; i < len(s); i++ {
		fmt.Printf("Byte %d: %x %q\n", i, s[i], s[i])
	}

	fmt.Println("\nRune-by-rune iteration (correct):")
	for i, r := range s {
		fmt.Printf("Rune %d: %U %c (%d bytes)\n", i, r, r, utf8.RuneLen(r))
	}

	fmt.Println("\nUsing utf8.DecodeRuneInString:")
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("Rune at %d: %c (%d bytes)\n", i, r, size)
		i += size
	}
}

// 5. String slicing pitfalls
func stringSlicingPitfalls() {
	fmt.Println("\n--- String Slicing Pitfalls ---")

	s := "hello 世界"

	fmt.Println("Original string:", s)
	fmt.Println("Byte length:", len(s))
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// Problematic slicing
	broken := s[6:9] // Cuts a multi-byte rune in half
	fmt.Println("\nProblematic slice s[6:9]:", broken)

	// Correct slicing
	runes := []rune(s)
	correct := string(runes[6:8])
	fmt.Println("Correct slice runes[6:8]:", correct)
}

// 6. String manipulation techniques
func stringManipulation() {
	fmt.Println("\n--- String Manipulation ---")

	// Concatenation
	s1 := "hello"
	s2 := "世界"

	fmt.Println("Concatenation (+):", s1+" "+s2)
	fmt.Println("Concatenation (fmt.Sprintf):", fmt.Sprintf("%s %s", s1, s2))

	// Using strings.Builder
	var sb strings.Builder
	sb.WriteString(s1)
	sb.WriteString(" ")
	sb.WriteString(s2)
	fmt.Println("Concatenation (strings.Builder):", sb.String())

	// String reversal
	fmt.Println("\nString reversal:")
	fmt.Println("Original:", s1+s2)
	fmt.Println("Reversed (byte-wise):", reverseStringByte(s1+s2))
	fmt.Println("Reversed (rune-wise):", reverseStringRune(s1+s2))

	// Case conversion
	mixed := "Hello World 你好"
	fmt.Println("\nCase conversion:")
	fmt.Println("Upper:", strings.ToUpper(mixed))
	fmt.Println("Lower:", strings.ToLower(mixed))
	fmt.Println("Title:", strings.ToTitle(mixed))
}

func reverseStringByte(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func reverseStringRune(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// 7. Unicode handling
func unicodeHandling() {
	fmt.Println("\n--- Unicode Handling ---")

	emoji := "Hello 😀 世界"

	fmt.Println("String:", emoji)
	fmt.Println("Byte length:", len(emoji))
	fmt.Println("Rune count:", utf8.RuneCountInString(emoji))

	fmt.Println("\nRune breakdown:")
	for i, r := range emoji {
		fmt.Printf("%d: %U %c %X (%d bytes)\n",
			i, r, r, []byte(string(r)), utf8.RuneLen(r))
	}

	// Normalization
	fmt.Println("\nNormalization:")
	s := "café"
	fmt.Printf("Original: % x\n", []byte(s))
	fmt.Printf("NFD: % x\n", []byte(normalizeNFD(s)))
	fmt.Printf("NFC: % x\n", []byte(normalizeNFC(s)))
}

// Placeholder for normalization functions
func normalizeNFD(s string) string { return s }
func normalizeNFC(s string) string { return s }

// 8. Performance considerations
func performanceConsiderations() {
	fmt.Println("\n--- Performance Considerations ---")

	const count = 10000

	// Inefficient concatenation
	start := time.Now()
	var s string
	for i := 0; i < count; i++ {
		s += "a"
	}
	duration := time.Since(start)
	fmt.Printf("+= concatenation: %v\n", duration)

	// Efficient concatenation
	start = time.Now()
	var sb strings.Builder
	for i := 0; i < count; i++ {
		sb.WriteString("a")
	}
	_ = sb.String()
	duration = time.Since(start)
	fmt.Printf("strings.Builder: %v\n", duration)

	// String vs []byte conversion
	str := "hello"
	start = time.Now()
	for i := 0; i < count; i++ {
		_ = []byte(str)
	}
	duration = time.Since(start)
	fmt.Printf("string->[]byte: %v\n", duration)

	// Rune count methods
	str = strings.Repeat("世界", 1000)
	start = time.Now()
	_ = len([]rune(str))
	duration = time.Since(start)
	fmt.Printf("len([]rune(s)): %v\n", duration)

	start = time.Now()
	_ = utf8.RuneCountInString(str)
	duration = time.Since(start)
	fmt.Printf("utf8.RuneCountInString: %v\n", duration)
}

// 9. Common pitfalls
func commonPitfalls() {
	fmt.Println("\n--- Common Pitfalls ---")

	// Pitfall 1: Length misunderstanding
	s := "世界"
	fmt.Println("String:", s)
	fmt.Println("len(s):", len(s), "(bytes, not characters!)")
	fmt.Println("utf8.RuneCountInString(s):", utf8.RuneCountInString(s))

	// Pitfall 2: Slicing multi-byte characters
	broken := s[:1]
	fmt.Println("\nBroken slice s[:1]:", broken, "(invalid UTF-8)")

	// Pitfall 3: String modification attempt
	// s[0] = 'H' // Compile error: cannot assign to s[0]

	// Pitfall 4: Range with index
	fmt.Println("\nRange loop index is byte position:")
	for i, r := range s {
		fmt.Printf("Index %d: %c\n", i, r)
	}

	// Pitfall 5: Comparing strings with different normalizations
	s1 := "café"       // Precomposed é
	s2 := "cafe\u0301" // e + combining acute accent
	fmt.Println("\nEquality check without normalization:")
	fmt.Printf("%s == %s: %v\n", s1, s2, s1 == s2)
}

// 10. Best practices
func bestPractices() {
	fmt.Println("\n--- Best Practices ---")

	fmt.Println("1. Use `range` for string iteration (handles runes correctly)")
	fmt.Println("2. Use []rune for character manipulation")
	fmt.Println("3. Use strings.Builder for efficient concatenation")
	fmt.Println("4. Be mindful of string encoding (UTF-8)")
	fmt.Println("5. Use utf8 package for Unicode operations")
	fmt.Println("6. Consider normalization for string comparison")
	fmt.Println("7. Avoid unnecessary string conversions")
	fmt.Println("8. Prefer rune operations for character counting")
	fmt.Println("9. Use proper string slicing with runes")
	fmt.Println("10. Remember strings are immutable - design accordingly")
}

var _ = time.Now // Avoid unused import error
