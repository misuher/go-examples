//the empty interface can accept a parameter of any type
package main2

import (
	"fmt"
	"strconv"
)

type example struct {
	i int
	s string
}

func assign_types_to_interfaces() {
	var a interface{}

	//define some random types
	var i int = 2
	var s string = "Hello world"

	//we can assign this types to the interface
	a = i
	a = s
	fmt.Println(a)
}

//how could be this usefull??? This is implemented for example in fmt.Println where it accept as a
//parameter a variadic empty interfaces, so we can pass any type argument like:
//			fmt.Println(3)
//			fmt.Println(1.5)
//			fmt.Println("Hello")
//
//Go make an internal type asserttion to get the type and act consequently, but what happens when is owr own type?
//go define the interface fmt.Stringer that is called when the type is unknown
//		type Stringer interface{
//			String() string
//		}
//so we can redefine it for owr types to generate
func (e example) String() string {
	return "(num:" + strconv.Itoa(e.i) + ", " + "string:" + e.s + ")"
}

func main() {
	e := example{5, "hello"}
	fmt.Println("My example type have the stringer representation: ", e)
	//output: My example type have the stringer representation:  (num:5, string:hello)

}
