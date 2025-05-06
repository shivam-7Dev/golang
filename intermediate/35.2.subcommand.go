/*
go run main.go

	go => main command
	run =>sub command
	main.go => first argument

git init

	git => primary command/ main command
	init=> subcommand

sub command has their own  flags
*/
/*
	creating subcommand using flag.NewFlagset("commandName",flag.ExitOnError)
*/
package intermediate

import (
	"flag"
	"fmt"
	"os"
)

func SubcommandIntro() {

	subcommand1 := flag.NewFlagSet("firstSub", flag.ExitOnError)

	subcommand2 := flag.NewFlagSet("secondSub", flag.ExitOnError)

	/*
	 Each pointer to flagset has methods associated wiht them
	*/

	fistFlag := subcommand1.Bool("processing", false, "this is about the bool val")
	secondFlag := subcommand1.Int("bytes", 1024, "Byte length of the result")

	flagsc2 := subcommand2.String("language", "GO", "Enter your language")

	/*
		Throw the eror is length of arugment is less than two
	*/

	if len(os.Args) < 2 {
		fmt.Println("this program requires additional commands")
		os.Exit(1)

	}
	//switch the os argumetns and check for both the cases

	switch os.Args[1] {
	case "firstSub":
		subcommand1.Parse(os.Args[2:])
		fmt.Println("processing :first flag of first subcommand:", *fistFlag)
		fmt.Println("bytes:second flag of first subcommand:", *secondFlag)

	case "secondSub":
		subcommand2.Parse(os.Args[2:])
		fmt.Println("language:first flag of second subcommand:", *flagsc2)

	default:
		fmt.Println("No subcommand entred!")
		os.Exit(1)

	}
}
