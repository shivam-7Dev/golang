package intermediate

import "fmt"

type Animal struct {
	Name string
}

func (a Animal) Speak() {
	fmt.Println("I am", a.Name)
}

type Dog struct {
	Animal // <--- embedding Animal
	Breed  string
}

func StructEmbeddig() {
	d := Dog{
		Animal: Animal{Name: "Tommy"},
		Breed:  "Labrador",
	}

	d.Speak()            // Can directly call Speak() from Animal
	fmt.Println(d.Name)  // Can access Name directly
	fmt.Println(d.Breed) // Dog's own field
}
