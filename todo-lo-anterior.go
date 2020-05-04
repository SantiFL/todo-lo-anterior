//inicia con el package main
package main

//se importa las librerias tanto para imprimir,convertir y para darle un tiempo
import (
	"fmt"
	"strconv"
	"time"
)

type conjuntodedatos struct { //se crea un struc que es como una cadena de caractares predeterminado
	numero1 int
	numero2 int
	numero3 int
	numero4 int
	numero5 int

	dato1 string
	dato2 string
	dato3 string
	dato4 string
	dato5 string

	numeroconcoma1 float32
	numeroconcoma2 float32
	numeroconcoma3 float32
	numeroconcoma4 float32
	numeroconcoma5 float32

	tipo1 bool
	tipo2 bool
	tipo3 bool
	tipo5 bool
	tipo6 bool
	tipo7 bool
}

func main() {
	var suma int = 10 + 10                     // se crea la variable suma
	var divicion float32 = 200 / 3             // se crea la variable divicion con un float
	var resta int = 10 - 10                    // se crea la variables resta
	var multiplicacion float32 = 15 * 3        // se crea la variables multiplicacion
	var NombreDelUsuario string = "Santiago"   // se crea la variables para el nombre del usuario
	var ApellidoDelUsuario string = "Ferreyra" //se crea la variable para el apellido del usuario
	const DNI float32 = 40.835432              //variable que es constante en un float para el dni
	
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
	
	//para imprimir un entero se usa el comando println
	var todoslosdatos = conjuntodedatos{ //se crea una variable nueva llamada todoslosdatos que es igual a conjuntodedatos con el valor de las variables ya escrita antes
		numero1:        1,
		numero2:        2,
		numero3:        3,
		numero4:        4,
		numero5:        5,
		dato1:          "nombre",
		dato2:          "segundo nombre",
		dato3:          "Apellido",
		dato4:          "Segundo apellido",
		dato5:          "apellido de casado",
		numeroconcoma1: 1.0,
		numeroconcoma2: 1.5,
		numeroconcoma3: 2.0,
		numeroconcoma4: 2.5,
		numeroconcoma5: 3.0,
		tipo1:          true,
		tipo2:          true,
		tipo3:          false,
		tipo5:          false,
		tipo6:          true,
		tipo7:          false,
	}
	fmt.Println(todoslosdatos)

}
