/*
Go es fuertemente tipado, a diferencia de otros lenguajes el tipo se declara al final
tipos numéricos:
	uint8,uint16,uint32,uint64
	int8,int16,int32,int64
	float32,float64
	complex64,complex128
	byte es un alias para uint8
	rune es un alias para int32
	uint --> puede ser 32 o 64
	int -->puede ser 32 o 64
Resto de tipos:
	bool
	string
*/

package main

import "fmt"

//constantes
const dos int = 2
const(
	Pi = 3.14
	ElSentidoDeLaVida = 42
	)

func main(){
	//vamos a hacer 5 formas de hacer lo mismo
	//metodo 1
	var a1 int
	var b1, c1 int // los dos son int
	a1 = 2
	b1 = -15
	fmt.Println(a1, b1, c1)

	//metodo 2
	var a2 int = 2
	var b2, c2 int = -15, 0
	var d2 bool = true
	var s2 string = "hola"
	fmt.Println(a2, b2, c2, d2, s2)

	//metodo 3, el tipo se infiere de la inicialización
	var a3 = 2
	var b3, c3 = -15, 0
	var d3 = true
	var s3 = "hola"
	fmt.Println(a3, b3, c3, d3, s3)

	//metodo 4, eliminamos la palabra var por :, pero sigue siendo estático
	a4 := 2
	b4, c4 := -15, 0
	d4 := true
	s4 := "hola"
	fmt.Println(a4, b4, c4, d4, s4)

	//metodo 5,agrupación de declaraciones
	var(
		a5 = 2
		b5, c5 = -15, 0
		d5 = true
		s5 = "hola"
	)
	fmt.Println(a5, b5, c5, d5, s5)

}