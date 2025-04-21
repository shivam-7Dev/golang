package basic

import "fmt"

func Operators() {
	a := 10
	b := 3

	fmt.Println("Add:", a+b)
	fmt.Println("Subtract:", a-b)
	fmt.Println("Multiply:", a*b)
	fmt.Println("Divide:", a/b)
	fmt.Println("Modulus:", a%b)
	// With float
	f := float64(a) / float64(b)
	fmt.Println("Float divide:", f)
	// Assignment ops
	a += 5
	fmt.Println("a += 5:", a)

	// Increment
	a++
	fmt.Println("a++:", a)
}
