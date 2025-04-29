package intermediate

import (
	"fmt"
	"os"
	"text/template"
)

func Temp() {
	tempOne()
}

func tempOne() {
	tmpl := template.Must(template.New("greet").Parse("hello {{.name}} how are you doing?\n"))

	// template data

	data := map[string]interface{}{
		"name": "shivam",
	}

	//execute the template

	err := tmpl.Execute(os.Stdout, data)
	if err != nil {
		fmt.Println("error while executing the data")
	}
}
