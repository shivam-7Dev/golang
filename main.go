package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type school struct {
	name   string
	medium string
}

func main() {
	// fmt.Println("this is main file")
	// basic.PrintHello() // Call the exported PrintHello function
	// basic.Variable()
	// basic.Operators()
	// star.PrintStar()
	// basic.SliceDemo()
	/*
	*Intermediate
	 */
	// intermediate.Pointer()
	// intermediate.StringAndRunes()
	// intermediate.StringAndRunes2()
	// intermediate.StringsPractice()
	// intermediate.Formating()
	// intermediate.FMT()
	// intermediate.StringProblmes()
	// intermediate.StructExamples()
	// intermediate.MethodsInGo()
	// intermediate.Interface()
	// intermediate.Errors()

	// intermediate.TemplateBasic()
	// intermediate.Temp()
	/*Formating
	*advance
	 */

	/*
	 problmes examples package
	*/
	// problems.SimpleExample()

	// inputOne()
	inputTwo()
}

func inputOne() {
	var s school
	fmt.Println("enter school details")
	fmt.Println("enter school name:")
	fmt.Scan(&s.name)

	fmt.Println("entere school medum:")
	fmt.Scan(&s.medium)
	fmt.Printf("Final school details: %+v\n", s)

}
func inputTwo() {

	for {
		reader := bufio.NewReader(os.Stdin)

		fmt.Println("press 1 to exit or any key oterh key to continue")

		chooise, _ := reader.ReadString('\n')

		chooise = strings.TrimSpace(chooise)
		if chooise == "1" {
			break
		}

		var s school

		fmt.Print("Enter school name: ")
		nameInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		s.name = strings.TrimSpace(nameInput)

		fmt.Print("Enter school medium: ")
		mediumInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		s.medium = strings.TrimSpace(mediumInput)

		fmt.Printf("School details: %+v\n", s)
	}
}
