/*
Las asignaciones en go pueden ser en paralelo, permitiendo evitar tener que
crear una tercera vble para intercambiar los valores de dos vbles
*/

package main

import(
	"fmt"
)

func main(){
	var a int 
	var b bool
	var s string

	a, b , s = 5, true, "hola" //inicialización múltiple
	fmt.Println(a, b, s)

	var a2 int = 7
	a, a2 = a2, a //intercambio
	fmt.Println(a, a2)

	//Para ignorar alguno de los valores se usa _

	a, _, s = 1, false, "cambio"
	fmt.Println(a, b, s)
}