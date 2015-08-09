/*
Son arrays
Son agrupaciones en secuencia de vbles del mismo tipo
*/

package main 

import(
	"fmt"
)

func main(){
	var a [5]int
	var b [3]bool

	a[0] = -1
	a[4] = 7

	fmt.Println(a) //-1,0,0,0,7
	fmt.Println(len(a)) // 5
	fmt.Println(b) //false, false, false

	//recorrer tablas
	var c[100]int
	for i := 0; i < 100; i++{
		c[i] = -1
	}
	//literales de tabla, iniciar la tabla con un valor de sus casillas
	d := [5]int{1,2,3,4,5} //tabla de 5 elementos con esos valores
	//puedo poner menos valores del máximo
	//tabla de tamaño definido por la inicializacion:
	e := [...]string{"hola", "que", "tal"} 
}