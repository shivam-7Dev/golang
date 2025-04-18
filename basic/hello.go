package basic //package declaration

import "fmt" //Imported package

func printhello() { //main funtion (entry point)
	fmt.Println("hello from go ")
}

func PrintHello() {
	printhello()
}

/*
*Key points
*Every file must have a package declartion
*the main package should have a main funtion

* to run a file => go run file.go
* to build a file => go build file.go
*In Go, if a function is defined but not used within the same package,
 the compiler will throw an error or warning.
*Functions that start with an uppercase letter (e.g., PrintHello)
 are exported and can be accessed from other packages.
*Functions that start with a lowercase letter (e.g., printhello)
are unexported and can only be used within the same package.
*/
