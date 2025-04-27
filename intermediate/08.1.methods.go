package intermediate

import "fmt"

type Shape struct {
	Rectange
}

type Rectange struct {
	length int
	width  int
}

func (r *Rectange) increaseLength(length int) {
	r.length += length

	fmt.Println(r)

}

func (r *Rectange) scale(factor int) {
	r.length *= factor
	r.width *= factor
	fmt.Println(r)

}

// value reciever
func (r Rectange) printArea() {
	fmt.Printf("the are of the rectange is %d \n", r.length*r.width)

	r.length = 3423 //// This change affects only the copy, not the original
}

type myInt int

//methods on user defined type

func (my *myInt) printData() {
	fmt.Println("the instave value is ", *my)
}

func MethodsInGo() {

	r := Rectange{
		length: 2,
		width:  2,
	}
	r.increaseLength(5)
	r.printArea()
	r.scale(2)
	r.printArea()

	var age myInt = myInt(5)

	age.printData()

	fmt.Println("===================")
	structEmbedding()

}

func structEmbedding() {

	r := Rectange{
		length: 2,
		width:  2,
	}
	s := Shape{
		Rectange: r,
	}
	fmt.Println(s)

	s2 := Shape{
		Rectange: Rectange{
			length: 2,
			width:  3,
		},
	}

	fmt.Println(s2)
	fmt.Println(s == s2)

	s2.Rectange.printArea()
	s2.printArea()
}
