package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type square struct {
	width, height float64
}

//every types that have the methods defined in this interface can use the functions
//that receive a geometry as a parameter, decoupling the implementation of this functions
//and the interface functions
type geometry interface {
	area() (float64, error)
	perim() (float64, error)
}

func (c circle) area() (float64, error) {
	return math.Pi * c.radius * c.radius, nil
}

func (c circle) perim() (float64, error) {
	return 2 * math.Pi * c.radius, nil
}

func (s square) area() (float64, error) {
	return s.width * s.height, nil
}

func (s square) perim() (float64, error) {
	return 2*s.width + 2*s.height, nil
}

//this function receive an interface as parameter using its methods but dont care about its implementations
func calcAll(g geometry) {
	fmt.Print("Interface values:")
	fmt.Println(g)
	fmt.Print("area, error values:")
	fmt.Println(g.area())
	fmt.Print("perim, error values:")
	fmt.Println(g.perim())
}

func main() {
	c := circle{5.0}
	s := square{3.0, 3.0}

	//As circle and square have the area() and perim() methods implemented they automatically meets the
	//requirements to be a geometry interface:
	calcAll(c)
	calcAll(s)
}
