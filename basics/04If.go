package main 

import (
	"fmt"
	"os"
	)

func main(){
	/*Programa que hace una calculadora*/
	var a, b float64
	var op byte
	//forma de declararlo 1
	_, err := fmt.Scanf("%f %c %f", &a, &op, &b)
	//metemos el if para maejar un error de formato
	if err != nil{
		fmt.Println("Formato incorrecto")
		return
	}
	

	//forma de declararlo 2, errr definido solo dentro del if
	if _,errr := fmt.Scanf("%f %c %f", &a, &op, &b); errr != nil{
		fmt.Fprintln(os.Stderr, "Formato incorrecto")  //da el mensaje por el canal de errores
		os.Exit(1)
	}

	if op == '+'{
		fmt.Println(a + b)
	}else if op == '-'{
		fmt.Println(a - b)
	}else if op == '*'{
		fmt.Println(a * b)
	}else if op == '/'{
		fmt.Println(a / b)
	}else{
		fmt.Fprintf(os.Stderr, "operación incorrecta: '%c'\n", op)
		os.Exit(1)
	}
}