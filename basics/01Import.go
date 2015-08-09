/*
Import permite usar paquetes muy bien documentados en
	www.golang.com/pkg

Las funciones exportables de un paquete empiezan con mayúscula
Cuando metes un paquete que no usas go da un error de compilación
*/

package main

//import "fmt"
//import "math"
//Forma abreviada:
import(
	f "fmt"  //f ahora es un alias de fmt
	"math"
)

func main(){
	f.Println(math.Pi)
}