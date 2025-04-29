package intermediate

import (
	"fmt"
	"os"
	"text/template"
)

type PersonStruct struct {
	Name string
	Age  int
}

func TemplateBasic() {
	// template1()
	variableInsertion()
	vairiabelTemplate()

}

func template1() {
	// Define a simple template
	templ := "hello {{.}}"
	data := "world \n"

	// parse the template
	t, err := template.New("greetings").Parse(templ)
	if err != nil {
		fmt.Println("error while parsing")
	}
	//execute the template

	err = t.Execute(os.Stdout, data)

	if err != nil {
		fmt.Println("error while executing")
	}

}

func variableInsertion() {
	shivam := PersonStruct{
		Name: "shivam",
		Age:  29,
	}

	greetings := "hello {{.Name}} your age is {{.Age}} \n"

	// parse template\
	t := template.Must(template.New("greet").Parse(greetings))

	// execute the template

	err := t.Execute(os.Stdout, shivam)

	if err != nil {
		fmt.Println("error while executing")
	}

}

func vairiabelTemplate() {
	t := template.Must(template.New("vars").Parse(`
		{{$name := "Bob"}}
		{{$age := 25}}
		Name: {{$name}}, Age: {{$age}}
		`))
	t.Execute(os.Stdout, nil)
}
