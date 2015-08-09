/*
Hay 18 funciones organizados con la siguiente nomenclatura:


Prefijo(origen de datos)     Func        Opciones
---------------------------------------------------
       --                                  --
   F(fichero)               scan             f(con formato)
    S(string)               print            ln(con salto de linea)


*/

package main 

import (
    "fmt"
    "math"
    "os"
    )

func main(){
    /*Programa que pide introducir un par de numeros
      e imprima su suma*/
    var a, b int      
    fmt.Scan(&a, &b) //se pasan como punteros
    fmt.Println(a + b)

    //variante con formato para la distancia entre 2 vectores
    var x1, y1 float64
    var x2, y2 float64
    fmt.Scanf("(%f,%f)", &x1, &y1) //el formato es como en C 
    fmt.Scanf("(%f,%f)", &x2, &y2) 
    dx := x1 - x2
    dy := y1 - y2
    fmt.Printf("La distancia es: %.3f\n",math.Sqrt(dx*dx + dy*dy))

    //variante que crea 10 ficheros con nombres consecutivos
    //y que contienen dentro la frase "este es el fichero "número""
    for i:= 1; i <= 10; i++{
        //creación de ficheros
        fichero, _ := os.Create(fmt.Sprintf("fichero%02d.txt", i)) //de lo que devuelve recojo solo una cosa en la vble fichero.
        //creación de contenido
        fmt.Fprintf(fichero, "Este es el fichero %d\n", i)
    }     
}