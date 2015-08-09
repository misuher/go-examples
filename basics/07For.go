/*
Defino las diferentes formas de expresar for
*/

package main 

import (
	"fmt"
)

func main(){
	//forma 1
	var i int = 1

	for i <= 10{
		fmt.Println(i)
		i = i +1
	}

	//forma 2, for sin condición es un while
	for{
		fmt.Println(i)
		i = i +1
		if i > 10 {
			break
		}
	}

	//forma 3, clásica inspirada en C
	for j := 1; j <= 10 ;j++{
		fm.Println(j)
	}	

	//forma 4, uso de range para iterar por todos los elementos de un contenedor 
	s := "vble string"
	for k,c := range s{  //range devuelve el índice de iteración y el valor del indice
		fmt.Println(k, c)
		fmt.Println(k, s[k]) //equivalente al anterior
	}

}