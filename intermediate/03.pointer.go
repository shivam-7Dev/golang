package intermediate

import "fmt"

func Pointer() {
	basicExample()
	basic2()
	pointerToString()
	pointerToSlice()
	pointerToStruct()
	functionAcceptingPointer()
	functionReturningPointer()
	pointerToPointer()
}

// Basic example of pointers
func basicExample() {
	name := "shivam"
	nameAddress := &name // Get the memory address of the variable
	fmt.Println("Address of name:", nameAddress)
	fmt.Println("Value at address:", *nameAddress) // Dereference the pointer
}

// Pointer to an integer
func basic2() {
	var x int = 10
	var xpointer *int = &x // xpointer stores the memory address of x
	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", xpointer)
	fmt.Println("Value at xpointer:", *xpointer) // Dereference the pointer
}

// Pointer to a string
func pointerToString() {
	str := "Hello, Go!"
	ptr := &str // Pointer to the string
	fmt.Println("Pointer Address:", ptr)
	fmt.Println("Value via Pointer:", *ptr) // Dereference to get the value
}

// Pointer to a slice
func pointerToSlice() {
	nums := []int{1, 2, 3}
	ptr := &nums // Pointer to the slice
	fmt.Println("Original Slice:", *ptr)
	(*ptr)[0] = 100 // Modify the slice via the pointer
	fmt.Println("Modified Slice:", *ptr)
}

// Pointer to a struct
type Person struct {
	Name string
	Age  int
}

func pointerToStruct() {
	person := Person{Name: "Alice", Age: 25}
	ptr := &person // Pointer to the struct
	fmt.Println("Original Struct:", *ptr)
	ptr.Age = 30 // Modify the struct field via the pointer
	fmt.Println("Modified Struct:", *ptr)
}

// Function accepting a pointer
func functionAcceptingPointer() {
	x := 10
	increment(&x) // Pass the pointer to the function
	fmt.Println("Value after increment:", x)
}

func increment(ptr *int) {
	*ptr++ // Increment the value at the memory address
}

// Function returning a pointer
func functionReturningPointer() {
	ptr := newInt()                           // Get a pointer from the function
	fmt.Println("Value from function:", *ptr) // Dereference to get the value
}

func newInt() *int {
	x := 42
	return &x // Return the address of x
}

// Pointer to a pointer
func pointerToPointer() {
	x := 10
	ptr := &x
	ptrToPtr := &ptr // Pointer to a pointer
	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", ptr)
	fmt.Println("Pointer to Pointer:", ptrToPtr)
	fmt.Println("Value via Pointer to Pointer:", **ptrToPtr) // Dereference twice
}
