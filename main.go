package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/shivam/personal/golang/intermediate"
)

type school struct {
	name   string
	medium string
}

func main() {
	intermediate.LogIntro()
	// USER := os.Getenv("USER")
	// HOME := os.Getenv("HOME")
	// fmt.Println("USER env is:", USER)
	// fmt.Println("USER env is:", HOME)

	// os.Setenv("FRUIT", "MANGo")
	// fmt.Println("USER env is:", os.Getenv("FRUIT"))

	// for _, e := range os.Environ() {
	// 	fmt.Println("e--->", e)
	// }
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
	// intermediate.TimeOperations()
	// intermediate.RandNumberInro()
	// intermediate.URLParsingIntro()
	// intermediate.BuffioInto()
	// intermediate.Base64Intro()
	// intermediate.HasingInto()

	// intermediate.FlagIntro()
	// intermediate.SubcommandIntro()
	/*Formating
	*advance
	 */

	/*
	 problmes examples package
	*/
	// problems.SimpleExample()

	// inputOne()
	// inputTwo()
	// inputThree()

	// useScan()
	// useScanln()
	// useScanf()

	// for ind, value := range os.Args {
	// 	fmt.Println("arg----", ind, value)
	// }
}

func inputOne() {
	var s school
	fmt.Print("enter your schoo name:")
	fmt.Print("Enter your school name and medium (e.g., DPS English): ")

	// Proper Scanf with format
	fmt.Scanf("%s %s", &s.name, &s.medium)
	fmt.Print("\nschool name is %v", s)

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

func inputThree() {

	// step one => declare template string
	templateMap := map[string]string{
		"welcome":      "welcome {{.}}, we are glad you join",
		"notification": "{{.}}, you have one notification",
		"error":        "Opps! an error occured {{.}}",
	}

	//step 2 => parse and store template
	parsedTemplate := make(map[string]*template.Template)

	//loop over the template map and get the key and value
	for key, value := range templateMap {
		parsedTemplate[key] = template.Must(template.New(key).Parse(value))
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("------------------------------")
		fmt.Println("1. welcome Template")
		fmt.Println("2. Notification template")
		fmt.Println("3. Error  template")
		fmt.Println("4. exit program")
		fmt.Print("\nEnter your choise:")
		fmt.Println("------------------------------")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			fmt.Print("\nEnter you name:")
			name, _ := reader.ReadString('\n')
			parsedTemplate["welcome"].Execute(os.Stdout, name)

		case "2":
			fmt.Print("\nEnter notification message:")
			noti, _ := reader.ReadString('\n')
			parsedTemplate["notification"].Execute(os.Stdout, noti)

		case "3":
			fmt.Print("\nEnter error message:")
			err, _ := reader.ReadString('\n')
			parsedTemplate["error"].Execute(os.Stdout, err)

		case "4":
			return
		default:
			fmt.Println("wrong option:try again")
			fmt.Println("-------------------------")
		}

	}
}
func useScan() {
	var s school
	fmt.Print("Enter school name and medium (separated by space or newlines): ")
	fmt.Scan(&s.name, &s.medium)
	fmt.Printf("Scan -> Name: %s, Medium: %s\n", s.name, s.medium)
}
func useScanln() {
	var s school
	fmt.Print("Enter school name and medium (on the same line): ")
	fmt.Scanln(&s.name, &s.medium)
	fmt.Printf("Scanln -> Name: %s, Medium: %s\n", s.name, s.medium)
}
func useScanf() {
	var s school
	fmt.Print("Enter school name and medium (format: %s %s): ")
	fmt.Scanf("%s %s", &s.name, &s.medium)
	fmt.Printf("Scanf -> Name: %s, Medium: %s\n", s.name, s.medium)
}
