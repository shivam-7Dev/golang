package intermediate

import (
	"fmt"
)

func FMT() {
	// printingFunctions()
	// formattingFucntions()
	// scanningFunctions()
	errorFormating()

}

func printingFunctions() {

	fmt.Println("--------this is printing funcition---------")
	fmt.Print("Hello")
	fmt.Print("world!")
	fmt.Print("12")
	fmt.Println("----------")
	fmt.Println("Hello")
	fmt.Println("world!")
	fmt.Println("12")
	fmt.Println("----------")

	name := "shivam"
	age := 26

	fmt.Printf("name of the person is %s and age is %d\n", name, age)
	data, error := fmt.Printf("name of the person is %v and age is %v\n", name, age)

	if error != nil {
		newError := fmt.Errorf("there is an error: %v", error)
		fmt.Println(newError) // Optionally print the new error
	}

	fmt.Printf("the returned value of Print f is %v\n", data)

}

func formattingFucntions() {
	fmt.Println("-----formatting functions demo----------")

	name := "shivam"
	lastName := "sharma"
	greeting := "hello"
	fullGreet := greeting + name + lastName

	fmt.Println(fullGreet) //helloshivamhello

	//Sprint does not add add space in between the argumets
	//its like simple concatination
	//also the returne valued does not have the new line charter
	//this function is used extensively in grpc
	SprintReturn := fmt.Sprint(greeting, name, greeting)
	fmt.Print("The return value of sprint fun:", SprintReturn) //helloshivamhello

	//sprintln adds space in between the argumets
	//and return the concatinated strings
	//have new line character at the end of the retuned values
	SpirntLnReturn := fmt.Sprintln(greeting, name, lastName)
	fmt.Print("Returned value fo Sprintln:", SpirntLnReturn) //hello shivam sharma
	fmt.Print("Returned value fo Sprintln:", SpirntLnReturn) //hello shivam sharma

	//its line printf but it does not print the value
	// it does not add new line character the return value
	SprintfReturn := fmt.Sprintf("%s %s %s ", greeting, name, lastName)
	fmt.Print(SprintfReturn)
	fmt.Print(SprintfReturn)
	fmt.Print(SprintfReturn)
}

func scanningFunctions() {
	fmt.Println("----------this is scan functions-----------")
	var name string
	var age int
	fmt.Print("Enter your name and age:")
	//scan waits until you give all input
	// it dones not wait scanning until it has all the inputs
	fmt.Scan(&name, &age)

	fmt.Printf("name is :%s \n", name)
	fmt.Printf("age is %d \n", age)

	var address string
	var pin int
	fmt.Println("Enter address and Pin")
	//scanln is line scan but it stops scanning at new line
	// so if you press enter then it will stop scannig
	//scan is more porwerful at than scanln
	fmt.Scanln(&address, &pin)
	fmt.Printf("address is %q \n", address)
	fmt.Printf("pin is %v \n", pin)

	var dob int
	var gender string

	fmt.Println("Enter dob and gender:")
	//scanf need format specifiers
	fmt.Scanf("%d %s", &dob, &gender)

	fmt.Printf("value of dob is %v \n", dob)
	fmt.Printf("value of gender is %s\n", gender)

}

func errorFormating() {
	fmt.Println("---this is error formtin---------")

	isgreaterthan18, error := checkAge(17)

	if error != nil {
		//error exist
		fmt.Println("Error:", error)
	}
	fmt.Println(isgreaterthan18)
}

func checkAge(age int) (bool, error) {
	if age > 18 {
		return true, nil
	}
	return false, fmt.Errorf("age is not greaer than 18")
}
