//inicia con el package main
package main

//se importa las librerias tanto para imprimir,convertir y para darle un tiempo
import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var suma int = 10 + 10                     // se crea la variable suma
	var divicion float32 = 200 / 3             // se crea la variable divicion con un float
	var resta int = 10 - 10                    // se crea la variables resta
	var multiplicacion float32 = 15 * 3        // se crea la variables multiplicacion
	var NombreDelUsuario string = "Santiago"   // se crea la variables para el nombre del usuario
	var ApellidoDelUsuario string = "Ferreyra" //se crea la variable para el apellido del usuario
	const DNI float32 = 40835432               //variable que es constante en un float para el dni
	//todo este codigo esta escrito con un salto de linea
	time.Sleep(time.Second * 10)                                                     //tiempo de 10 segundos
	convertir := strconv.Itoa(suma + resta)                                          //convierte las variables que estan en int a string
	fmt.Println(convertir)                                                           //imprime el la variable ya convertida
	fmt.Println(divicion)                                                            //imprime la divicion
	fmt.Println(multiplicacion)                                                      //imprime la multiplicacion
	fmt.Println("El nombre del usuario es " + NombreDelUsuario + ApellidoDelUsuario) //imprime el nombre y apellido del usuario
	fmt.Println(DNI)                                                                 //imprime el dni
	//esta parte del codigo es lo mismo que la anterior pero todo en el mismo renglon
	fmt.Print(convertir)
	fmt.Print(divicion)
	fmt.Print(multiplicacion)
	fmt.Print("El nombre del usuario es " + NombreDelUsuario + ApellidoDelUsuario)
	fmt.Print(DNI)

}
