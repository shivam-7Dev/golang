package intermediate

import "fmt"

type Emp struct {
	name     string
	lastname string
	age      int
	adress   Address
	//adding anonymous field, fields in anonymous struct are promoted to parent struct
	company
}

type Address struct {
	city    string
	country string
}

type company struct {
	name string
	exp  int
}

// methods are fuctions associated with a specific type
// methods are defined with a reciever =>struct type on which the method operates
func (e *Emp) printData() {
	fmt.Println(e)  //&{shivam sharma 29}
	fmt.Println(*e) //{shivam sharma 29}
	fmt.Printf("type of *Emp is:%T, value is:%[1]v \n", e)
	//type of *Emp is:*intermediate.Emp, value is:&{shivam sharma 29}

	fmt.Println("name is :", e.name)
	fmt.Println("last naem is:", e.lastname)
	fmt.Println("age is", e.lastname)
	fmt.Printf("\nType of *Emp is: %T, value is: %+v\n", e, e)

}

//methods to  modify the values of struct

func (e *Emp) increaseAge(amount int) {
	fmt.Println("increasing age by---------", amount)
	e.age += amount
	fmt.Println("printing the value of the struct", e)

}

//updat the fields in the instace of Emp

func (e *Emp) modifycompnay(name string, exp int) {
	e.exp = exp
	e.company.name = name

	fmt.Print("company information modified", *e)
}

func StructExamples() {
	basic()
	anonymousStruct()
}

func basic() {
	//instace of Emp stuct
	var shivam = Emp{
		name:     "shivam",
		lastname: "sharma",
		age:      29,
	}
	shivam.adress.city = "aligarn"
	shivam.adress.country = "India"

	shivam.modifycompnay("fiftyfive technology", 4)

	shivam.printData()
	shivam.increaseAge(2)

	fmt.Printf("shivam is %+v-------------/n", shivam)
	//{name:fiftyfive technology lastname:sharma age:31 adress:{city:aligarn country:India} company:{name:fiftyfive technology exp:4}}

}

func anonymousStruct() {
	nodeData := struct {
		data     int
		nextVaue string
	}{data: 12, nextVaue: "some address"}
	fmt.Println(nodeData)
}

/*
	By design go requries

	methods and stuct to go declared at package level
	i.e they should not be enclosed in any function

	Methods associated with struct are declared outside the stuct body
	and not inside the body of the struct. This is done so that data and behaviour
	are defined differetly


	To modify the value inside the struct we use a pointer reciever not a vlue receiver
**/
