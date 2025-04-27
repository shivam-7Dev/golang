package problems

import (
	"fmt"
	"os"
)

// taking command line argumes
func SimpleExample() {
	fmt.Println("this is basic exmaple")
	fmt.Println("the name of the program is", os.Args[0])
	// fmt.Println("First argument is:", os.Args[1])
	var sum float64
	var n int
	for {
		var val float64
		_, err := fmt.Fscanln(os.Stdin, &val)
		if err != nil {
			break
		}
		sum += sum
		n++
	}
	if n == 0 {
		fmt.Fprint(os.Stderr, "no values found")
		os.Exit(-1)
	}
	fmt.Println(sum)
	fmt.Println(n)

	fmt.Println("the average is", sum/float64(n))
}

/*
	In Go, you can use the os package to access command-line arguments.
	The os.Args slice contains the command-line arguments passed to the program.

	passing commannd line arguments:

		go run main.go argumentOne, argumentTwo ...

		go run main.go $(cat arguments.txt) //	When you want arguments in os.Args.

		go run main.go < arguments.txt //When you want to process input from stdin.
			but input redirect < does not poplate the os.Args

		ssing piper operator
			echo "Hello, World!" | go run main.go
			cat file.txt | go run main.go
			The pipe operator (|) redirects the standard output (stdout)
			of one command to the standard input (stdin) of another command.
			Since the pipe operator sends data to stdin, it does not affect

	Steps to Take Command-Line Arguments in Go:

	Access Command-Line Arguments:
		Use os.Args to get the arguments.
		os.Args[0] is the name of the program.
		os.Args[1:] contains the actual arguments passed to the program.
	Iterate Over Arguments:
		Use a for loop to iterate over os.Args.

	Pass a File as an Argument:
		You can pass the file path as an argument and access it using os.Args.

	Key Points
		os.Args is a slice of strings containing all command-line arguments.
		os.Args[0] is the program name.
		os.Args[1:] contains the actual arguments.
	Error Handling:
		Always check the length of os.Args to ensure arguments are provided.

	File Handling:
		Once you have the file path,
		you can use Go's os or io/ioutil packages to open and read the file.



*/
