package basic

import "fmt"

func ArithmeticOperator() {
	a := 10
	b := 3
	fmt.Println(a + b) // 13
	fmt.Println(a - b) // 7
	fmt.Println(a * b) // 30
	fmt.Println(a / b) // 3 (integer division)
	fmt.Println(a % b) // 1
	a++                // a becomes 11
	b--                // b becomes 2
}

func RelationShipOperator() {
	x := 5
	y := 7
	fmt.Println(x == y) // false
	fmt.Println(x != y) // true
	fmt.Println(x > y)  // false
	fmt.Println(x < y)  // true
	fmt.Println(x >= 5) // true
	fmt.Println(y <= 6) // false
}

func LogicaOperators() {
	a := true
	b := false
	fmt.Println(a && b) // false
	fmt.Println(a || b) // true
	fmt.Println(!a)     // false

	// Short-circuit example
	funcA := func() bool {
		fmt.Println("funcA called")
		return false
	}
	funcB := func() bool {
		fmt.Println("funcB called")
		return true
	}

	fmt.Println(funcA() && funcB()) // Only "funcA called" prints
	fmt.Println(funcB() || funcA()) // Only "funcB called" prints
}

func BitwiseOperators() {
	a := 10 // 1010 in binary
	b := 3  // 0011 in binary

	fmt.Println(a & b)  // 2  (0010)
	fmt.Println(a | b)  // 11 (1011)
	fmt.Println(a ^ b)  // 9  (1001)
	fmt.Println(a &^ b) // 8  (1000) - clears bits in a where b has 1
	fmt.Println(a << 2) // 40 (101000)
	fmt.Println(a >> 1) // 5  (0101)
}

func AssignmentOperator() {
	var a int

	a = 10
	a += 5  // a = 15
	a -= 3  // a = 12
	a *= 2  // a = 24
	a /= 4  // a = 6
	a %= 4  // a = 2
	a <<= 2 // a = 8
	a >>= 1 // a = 4
	a |= 1  // a = 5
	a &^= 1 // a = 4
}

func AddressAndDereferenceOperator() {
	x := 10
	ptr := &x // ptr holds the memory address of x =>address-of

	fmt.Println(ptr)  // Memory address (e.g., 0xc0000180a8)
	fmt.Println(*ptr) // 10 (dereferencing the pointer)

	*ptr = 20      // Change the value at the address
	fmt.Println(x) // 20
}

func ChannelOperator() {
	ch := make(chan int)

	// Send to channel (in a goroutine)
	go func() {
		ch <- 42
	}()

	// Receive from channel
	value := <-ch
	fmt.Println(value) // 42
}
