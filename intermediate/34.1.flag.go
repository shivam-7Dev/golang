package intermediate

import (
	"flag"
	"fmt"
)

func FlagIntro() {

	//defining flags
	var name string
	var age int
	var male bool

	flag.StringVar(&name, "name", "shivam", "name of the user")
	flag.IntVar(&age, "age", 29, "age of the user")
	flag.BoolVar(&male, "male", true, "is user male")

	//parse the cli arguments

	flag.Parse()
	//print the flags
	fmt.Println("name", name)
	fmt.Println("name", age)
	fmt.Println("name", male)

}
