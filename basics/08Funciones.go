package main 

import(
	"fmt"
)

//FUNCIONES BASICAS
func pi(){
	fmt.Print("pi ")		
}

func piN(n int){
	/*Función que recive el numero de veces 
	que quieres mostrar "pi" por pantalla*/
	for i := 0; i < n; i++{
		pi()
	}
}

//FUNCIONES QUE DEVUELVE RESULTADO
func cubo(x float32) float32{
	/*función a la que se le pasa un float32 y devuelve un float32*/
	c := x * x * x
	return c
}

//FUNCIONES QUE DEVUELVEN MAS DE UN RESULTADO
func HHMM(n int)(int, int){
	/*Funcion que recibe una hora tal que 1440 y devuelve por separado la hora y los minutos*/
	h := n / 100
	m := n % 100
	return h, m
}

//FUNCION QUE SE LE PASA COMO PARAMETRO UN LITERAL A FUNCION
func repite(n int, f func(i int)){
	for i:=1; i <= n; i++{
		f(i)
	}
}


func main(){
	piN(5) //llamada a función
	a := cubo(5.0)
	b, c := HHMM(1440)

	//asignar una función a una vble
	var vble func(x float32) float32
	vble = cubo
	vble(5.0)

	//Literal de función/función anónima
	s := func(a, b int) int{
		return a + b
	}

	//función anónima con llamada, pasandole dos valores 4 y 5
	fmt.Println(func (a,b int) int{
		return a + b
		} (4, 5))

	//llamada a funcion que se le pasa literal como parametro
	repite(10, func(i int){
		fmt.Println("Iteración número: ", i)
		})
}