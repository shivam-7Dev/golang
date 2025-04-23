package intermediate

import "fmt"

func Formating() {
	// generalFormattingVerbs()
	// generalFormattingVerbs2()
	// integerFormattingVerb()
	// stringFormattingverb()
	// boolenFormattingVerb()
	floatingFormattingVerb()
}

func generalFormattingVerbs() {
	fmt.Println("this is general formatting verb----------")
	// %v, %#v %T %%

	i := 15.5
	fmt.Println("this is print ln for int:", i)
	fmt.Printf("%v\n", i)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%v%%\n", i)
	fmt.Printf("%T\n", i)
	fmt.Println("---------------------")

	name := "shivam kumar sharma"
	fmt.Println("this is print ln for string:", name)
	fmt.Printf("%v\n", name)
	fmt.Printf("%#v\n", name)
	fmt.Printf("%T\n", name)
	fmt.Println("---------------------")

	arr := []int{1, 2, 3}
	fmt.Println("this is print ln for arr:", arr)
	fmt.Printf("%v\n", arr)
	fmt.Printf("%#v\n", arr)
	fmt.Printf("%T\n", arr)
	fmt.Println("---------------------")

	type Person struct {
		Name string
		Age  int
	}
	p := Person{Name: "Alice", Age: 30}
	fmt.Println("this is print ln for struct:", p)
	fmt.Printf("%v\n", p)
	fmt.Printf("%#v\n", p)
	fmt.Printf("%T\n", p)
	fmt.Println("---------------------")

}
func generalFormattingVerbs2() {
	fmt.Println("this is general formatting verb----------")

	// Example values
	num := 123
	str := "hello"
	arr := []int{1, 2, 3}
	type Person struct {
		Name string
		Age  int
	}
	p := Person{Name: "Alice", Age: 30}

	// %v: Default format
	fmt.Printf("Default format (%%v): %v\n", num)
	fmt.Printf("Default format (%%v): %v\n", str)
	fmt.Printf("Default format (%%v): %v\n", arr)

	// %+v: Struct with field names
	fmt.Printf("Struct with field names (%%+v): %+v\n", p)

	// %#v: Go-syntax representation
	fmt.Printf("Go-syntax representation (%%#v): %#v\n", p)

	// %T: Type of the value
	fmt.Printf("Type of num (%%T): %T\n", num)
	fmt.Printf("Type of str (%%T): %T\n", str)
	fmt.Printf("Type of arr (%%T): %T\n", arr)

	// %%: Literal percent sign
	fmt.Printf("Literal percent sign (%%%%): 100%% complete\n")
}

func integerFormattingVerb() {
	fmt.Println("this is interger fomatting verb------")
	number := 69

	fmt.Println("printing number using Println:", number)
	fmt.Printf("decimal (%%d) represention of the number %v is %d \n", number, number)
	fmt.Printf("binary(%%b) represention of number %v is %b \n", number, number)
	fmt.Printf("octal(%%o) representaion of the nubmer %v is %o \n", number, number)
	fmt.Printf("hexadecial(lowercase)(%%x) representaion of the number %v is %x \n", number, number)
	fmt.Printf("Hedadecimal(upper case)(%%X) representaiton of the number %v is %X \n", number, number)
	fmt.Printf("Unicode(%%U) represetation of the number %v is %U \n", number, number)
	fmt.Printf("character(%%c) representation of the number %v is %c \n", number, number)
	fmt.Printf("quoted(%%q) chracter of the number %v is %q \n", number, number)

}

func stringFormattingverb() {
	fmt.Println("------------ string formatting verb-------------")
	greeting := "hello Shivam"

	fmt.Println("Println output for string:", greeting)
	fmt.Printf("defult format outpur using (%%v): %v \n", greeting)
	fmt.Printf("Type (%%T) for the string: %v is : %T \n", greeting, greeting)
	fmt.Printf("Simple string format (%%s):%s \n", greeting)
	fmt.Printf("Simple string(width 8, (right justifed) format (%%8s):%8s \n", greeting)
	fmt.Printf("Simple string(width 8, (left justifed) format (%%-8s):%-8s \n", greeting)
	fmt.Printf("Quoted string format (%%q):%q \n", greeting)
	fmt.Printf("hexadecimal(lowercase) format (%%x):%x \n", greeting)
	fmt.Printf("Hexadeical(upper case) format (%%X):%X \n", greeting)
	fmt.Printf("GO syntax format (%%#v):%#v \n", greeting)
	fmt.Printf("plain string format (%%s):%s \n", greeting)
	fmt.Printf("plain string format (%%s):%s \n", greeting)
	fmt.Printf("plain string format (%%s):%s \n", greeting)
}

func boolenFormattingVerb() {
	fmt.Println("------------ boolean formatting verb------------")

	// Example boolean values
	isTrue := true
	isFalse := false

	// %t: Boolean representation
	fmt.Printf("Boolean value (%%t): %t\n", isTrue)
	fmt.Printf("Boolean value (%%t): %t\n", isFalse)
	// Using boolean values in a sentence
	fmt.Printf("The statement 'Go is awesome' is %t.\n", isTrue)
	fmt.Printf("The statement 'Go is difficult' is %t.\n", isFalse)
}

func floatingFormattingVerb() {
	fmt.Println("------------ floating-point formatting verb-------------")

	number := 123.456

	// %f: Decimal point notation
	//Default precision is 6 digits after the decimal point.

	fmt.Printf("Decimal point notation (%%f): %f\n", number)

	// %e: Scientific notation (lowercase)
	//Prints the number in scientific notation.
	fmt.Printf("Scientific notation (%%e): %e\n", number)

	// %E: Scientific notation (uppercase)
	//Prints the number in scientific notation.
	fmt.Printf("Scientific notation (%%E): %E\n", number)

	// %g: Compact representation (lowercase)
	//Automatically chooses between %f and %e (or %E) based on the value and precision.
	fmt.Printf("Compact representation (%%g): %g\n", number)

	// %G: Compact representation (uppercase)
	//Automatically chooses between %f and %e (or %E) based on the value and precision.
	fmt.Printf("Compact representation (%%G): %G\n", number)

	// %.nf: Precision control
	fmt.Printf("Precision control (%%.2f): %.2f\n", number)

	// %nf: Width control
	//Ensures the output is at least n characters wide, padding with spaces if necessary.
	fmt.Printf("Width control (%%10f): %10f\n", number)

	// %n.mf: Width and precision control
	//Combines width and precision formatting for precise control over the output.
	fmt.Printf("Width and precision control (%%10.2f): %10.2f\n", number)
}
