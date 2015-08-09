package main 

import (
	"fmt"
	"os"
	)

func main(){
	/*leemos el fichero que contiene números y los sumna*/
	f, err := os.Open("fichero.txt")
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}

	suma := 0
	for {
		var a int
		_, err := fmt.Fscan(f,&a)
		if err != nil{ // uno de los errores es el fin del fichero
			break
		}
		suma += a
	}
	fmt.Println(suma)

	/*Escribir en un fichero*/
	file, error := os.Create("libro.txt")
	if != nil{
		fmt.Println(error)
		os.Exit(1)
	}
	fmt.Fprintln(file, "Escribiendo en fichero")
	file.Close()
}