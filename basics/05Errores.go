package main

import(
	"fmt"
	"os"
)

func main(){
	/*Programa para manejar errores de formato en lugar de parar la ejecución*/
	var a int
	n, err := fmt.Scanf("[%d]", &a) //Scan devuelve el num de parametros que se le pasaron y los errores
	if err == nil{
		fmt.Println(a)
	} else {
		fmt.Println(err.Error()) //.Error es el único método de la clase error, asi que lo puedes obviar
	}
	fmt.Println("n: ", n)

	/*Programa para obtener info de errores al abrir ficheros*/
	_, err := os.Open("No existe")
	if err != nil{
		fmt.Println(err)
	}

}