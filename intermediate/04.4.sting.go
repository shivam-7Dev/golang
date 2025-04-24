package intermediate

import (
	"bytes"
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
)

func StringProblmes() {
	greetings := "hello 世界"
	// runeCounter(greetings)
	fmt.Println(safeSlice(greetings, 6, 8))
	stringConcatenationComparison()

}

/*
	1. Unicode Rune Counter and Printer

	Problem: Create a function that accurately counts the number of characters (not bytes)
	in a string containing various Unicode characters,
	and prints each character along with its Unicode code point.

	func runeCounter(input string) {
	// Count runes (characters) and print details for each one
	// Output: Total characters, and for each character:
	// - The character itself
	// - Its Unicode code point
	// - Its position in the string (both byte index and rune index)
    }
*/

func runeCounter(input string) {
	fmt.Println("the rune count is:", utf8.RuneCountInString(input))
	runes := []rune(input) // Convert the string to a slice of runes

	// Iterate over the runes to print details
	for _, r := range runes {
		fmt.Printf("The rune(%%d):%d, chracter(%%c):%c Unicode(%%U): %U\n", r, r, r)
	}

	// Print the length of the input in bytes and runes
	fmt.Println("Length of the input (bytes):", len(input)) // Number of bytes
	fmt.Println("Length of the input (runes):", len(runes)) // Number of characters
}

/*
	2. Correct String Slicer
	implement a function that slices a string by character indices (not byte indices)
	to extract substrings correctly, even with multi-byte Unicode characters.

	func safeSlice(s string, start, end int) string {
	// Slice the string by character positions, not byte positions
	// For example, safeSlice("hello 世界", 6, 8) should return "世界"
    }

*/

func safeSlice(s string, start, end int) string {

	runes := []rune(s)
	slice := runes[start:end]
	// var returnvalue string
	// for _, r := range slice {
	// 	returnvalue += fmt.Sprintf("%c", r)
	// 	/*
	// 		//Using returnvalue += fmt.Sprintf("%c", r) in a loop is inefficient because string
	// 		concatenation in Go creates a new string each time,
	// 		which can be slow for large strings.
	// 		Instead, use strings.Builder for efficient string concatenation.

	// 	*/
	// }
	// return returnvalue
	// Use strings.Builder for efficient concatenation
	var builder strings.Builder
	for _, r := range slice {
		builder.WriteRune(r) // Append each rune to the builder
	}

	return builder.String() // Return the concatenated string

}

/*
3. Efficient String Builder
	Compare the performance of different string concatenation methods
	 (+ operator, strings.Builder, bytes.Buffer)
	 by concatenating 1000 strings and measuring execution time.

	func stringConcatenationComparison() {
    // Implement and compare:
    // 1. Using + operator
    // 2. Using strings.Builder
    // 3. Using bytes.Buffer
    // Measure and print the time taken by each approach
    }

*/

func stringConcatenationComparison() {
	const iterations = 100000
	const baseString = "hello"

	// 1. Using + operator
	start := time.Now()
	result := ""
	for i := 0; i < iterations; i++ {
		result += baseString
	}
	fmt.Printf("Using + operator: %v\n", time.Since(start))

	// 2. Using strings.Builder
	start = time.Now()
	var builder strings.Builder
	for i := 0; i < iterations; i++ {
		builder.WriteString(baseString)
	}
	fmt.Printf("Using strings.Builder: %v\n", time.Since(start))

	// 3. Using bytes.Buffer
	start = time.Now()
	var buffer bytes.Buffer
	for i := 0; i < iterations; i++ {
		buffer.WriteString(baseString)
	}
	fmt.Printf("Using bytes.Buffer: %v\n", time.Since(start))

	/*
		Using + operator: 2.17190444s
		Using strings.Builder: 552.513µs
		Using bytes.Buffer: 476.496µs

	*/

	/*
		Result on mac
		Using + operator: 1.124576417s
		Using strings.Builder: 251.458µs
		Using bytes.Buffer: 234.916µs
	*/
}

/*
	4. Unicode-Aware String Reversal
	Problem: Create a function that correctly reverses a string containing Unicode characters
	(like emojis, combining marks, etc.) without breaking the characters.

	func reverseStringCorrectly(s string) string {
		// Implement string reversal that handles:
		// - Multi-byte characters
		// - Combining marks
		// - Surrogate pairs
		// Example: reverseStringCorrectly("Hello 世界😊") should properly reverse with intact characters
	}

*/

/*
5.
	 String Normalization and Comparison

	Problem: Create a function that normalizes two strings and correctly compares them,
	handling cases like different representations of the same character
	(e.g., "café" vs "cafe\u0301").

	func normalizedCompare(s1, s2 string) bool {
		// Normalize s1 and s2 before comparison
		// Return true if they're equivalent after normalization
		// Example: normalizedCompare("café", "cafe\u0301") should return true
	}
*/

/*
6. Word Tokenizer with Unicode Support
Problem: Implement a function that splits a string into words,
 properly handling Unicode whitespace and punctuation.

 func tokenizeWords(text string) []string {
    // Split text into words
    // Handle Unicode whitespace and punctuation
    // Return a slice of words
}


*/

/*
7. String Search and Replace with Context
Problem: Create a function that finds all occurrences of a substring and replaces them,
but only if they match specific context conditions.
func contextualReplace(text, find, replace string, contextFunc func(before, after string) bool) string {
    // Replace occurrences of 'find' with 'replace'
    // But only if contextFunc returns true for the surrounding context
    // Return the modified string
}

*/

/*
	8. UTF-8 Validation and Repair
	Problem: Write a function that checks if a string is valid UTF-8 and attempts to repair it if it's not.


	func validateAndRepairUTF8(input []byte) string {
		// Check if the input is valid UTF-8
		// If not, attempt to repair it (replace invalid sequences)
		// Return the valid UTF-8 string
	}

*/

/*
	9. String Memory Analysis
	Problem: Create a utility that analyzes the memory usage of a string, showing how bytes are allocated and stored.

	func analyzeStringMemory(s string) {
		// Print detailed information about the string's memory usage:
		// - Total bytes allocated
		// - Breakdown of bytes per character
		// - Memory address information
		// - String header structure
	}
*/

/*
	10. Custom String Formatter
	Problem: Implement a custom string formatter that supports various formatting directives,
	similar to fmt.Printf but with your own extensions.


		func customFormat(format string, args ...interface{}) string {
		// Implement a custom formatter that supports:
		// - Basic placeholders like %s, %d
		// - Custom placeholders (%r for reverse, %u for uppercase, etc.)
		// - Width and precision specifications
		// Return the formatted string
		}
*/
