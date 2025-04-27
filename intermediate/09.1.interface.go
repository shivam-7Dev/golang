package intermediate

import (
	"fmt"
	"math"
)

// here geometry in interace with two methods area() and perim()
// both methods return float64 values

//how to implement an interface?
//A type implicitly satisfies an interface if it implemets all the methods defined by interface

type geometry interface {
	area() float64
	perim() float64
	//This interface specifies that any type implementing the area() and perim() methods
	// satisfies the geometry interface.
}

type rect struct {
	length float64
	widht  float64
}

func (r *rect) area() float64 {
	return r.length * r.widht
}
func (r *rect) perim() float64 {
	return 2 * (r.length + r.widht)
}

type circle struct {
	radius float64
}

func (r *circle) area() float64 {
	return 2 * math.Pi * r.radius * r.radius
}
func (r *circle) perim() float64 {
	return 2 * math.Pi * r.radius
}

func (r *circle) dia() float64 {
	return 2 * r.radius
}

// type square struct {
// 	side float64
// }

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func Interface() {
	fmt.Println("this is interace lecture")
	interfacebasic()
}

func interfacebasic() {

	r := rect{
		length: 2,
		widht:  2,
	}

	fmt.Println(r)

	c := circle{
		radius: 2,
	}

	fmt.Println(c)

	// fmt.Println("ara of rectangle is:", r.area())
	// fmt.Println("perim of rectangle is:", r.perim())

	// fmt.Println("ara of rectangle is:", c.area())
	// fmt.Println("perim of circle is:", c.perim())
	// fmt.Println("DIA of circle is:", c.dia())
	measure(&r)
	measure(&c)

}
