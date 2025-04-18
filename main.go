package main

import (
	"fmt"

	"github.com/shivam/personal/golang/basic" // Import the basic package
)

func main() {
	Test() //Directly call the Test funtion from index.go
	fmt.Println("this is main file")
	basic.PrintHello() // Call the exported PrintHello function
	basic.Variable()
}
