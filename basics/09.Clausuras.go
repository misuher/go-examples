/*
Las clausuras son literales de funcion que usan vbles que estan
fuera de su rango de visión
*/

package main 

import "fmt"

func progresion() func() int{
	/*se le pasa como parametro una funcion y devuelve un entero*/
	s := 0
	return func() int{
		s++
		return s
	}	
}

func main(){
	n := 1

	inc := func(k int) int{
		n += k
		return n
	}

	fmt.Println(inc(2)) //sale 3
	fmt.Println(n) 		//sale 3

	//llamada a la funcion con clausura
	p1 := progresion()
	p2 := progresion()
	fmt.Println(p1()) //sale 1
	fmt.Println(p1()) //sale 2
	fmt.Println(p1()) //sale 3
	fmt.Println(p2()) //sale 1
	fmt.Println(p2()) //sale 2
	fmt.Println(p1()) //sale 4
}