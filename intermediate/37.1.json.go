package intermediate

import (
	"encoding/json"
	"fmt"
)

/*
only exported fields (starting with a capital letter) are marshalled to JSON.
*/
type person struct {
	Firstname string `json:"firstName"`
	Age       int    `json:"age"`
	Email     string `json:"email,omitempty"`
}

type pressonAddress struct {
	FirstName    string `json:"firstname,omitempty"`
	EmailAddress string `json:"email"`
	Age          int    `json:"age"`
	Add          `json:"address"`
}

type Add struct {
	City  string `json:"city"`
	State string `json:"state"`
}

func (p *pressonAddress) printInfo() {
	fmt.Printf("info about the user is %v", *p)
}

func (p *person) printInfo() {
	fmt.Printf("info about the user is %v", *p)
}

func JSONINTRO() {
	fmt.Println("this is json package")
	jsonone()
	fmt.Println()
	jsontwo()
	fmt.Println()
	unMarshell()
	fmt.Println()
	unknowJsondecoding()
}

func jsonone() {
	shiv := person{
		Firstname: "shivam kumar sharma",
		Age:       29,
	}
	fmt.Printf(" shivam value is: %+v \n", shiv)
	shiv.printInfo()

	//marshelling

	jsonData, err := json.Marshal(shiv)
	if err != nil {
		fmt.Println("error while converting to json")
		return
	}
	fmt.Println("data=====", string(jsonData))
}

func jsontwo() {
	shiv := pressonAddress{
		FirstName:    "shivam sharma",
		EmailAddress: "nlmshmcr@gmail.com",
		Age:          29,
		Add: Add{
			City:  "aligarh",
			State: "up",
		},
	}

	jsonData, err := json.Marshal(shiv)
	if err != nil {
		fmt.Println("erro while marshelling")
		return
	}

	fmt.Println("josn data is ----------", string(jsonData))
	shiv.printInfo()
}

func unMarshell() {
	jsondata := `{"firstname":"Atul kumar sharma", "age":25, "address":{"city":"delhi","state":"delhi"}}`

	var atul pressonAddress
	json.Unmarshal([]byte(jsondata), &atul)

	fmt.Println("atul--------", atul)

}

func unknowJsondecoding() {
	jsondata := `{"name":"shivamkumar sharm","favaoutitemove":"interstellar","company":{"firstcompnay":"fiftyfivetech","secondCompany":"unknow"}}`

	var data map[string]interface{}

	json.Unmarshal([]byte(jsondata), &data)

	fmt.Println("unhandlled data", data)
}

/*
	type Printer interface {
		Print()
	}

	type Person struct{}

	func (p *Person) Print() {} // pointer receiver

	// Only *Person implements Printer, not Person



	Pointer receiver: Only pointer type implements the interface.
	Value receiver: Both value and pointer types implement the interface.
*/
