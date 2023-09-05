package main

import (
	//e "github.com/ShineWaine/cursogo/ejer_interfaces"
	//"github.com/ShineWaine/cursogo/modelos"
	//"github.com/ShineWaine/cursogo/defer_panic"
	"fmt"

	"github.com/ShineWaine/cursogo/goroutines"
)

func main() {
	/*variables.NumerosEnteros()
	variables.RestoVariables()
	variables.Convierteanumero(1780)
	entero, frase := ejercicios.Cadenaaentero("500")
	fmt.Printf("%d, %s", entero, frase)
	teclado.IngresoNumeros() */
	//ejercicios.TablaMultiplicar()
	//files.GrabaTabla()
	//files.AnexaTabla()
	//files.LeerDato()
	//funciones.Calculos()
	//funciones.LlamarClosure()
	//funciones.Exponente(1)
	//arreglos_slices.MuestroArreglos()
	//arreglos_slices.MuestroSlices()
	//arreglos_slices.Capacidad()
	//mapas.MostrarMapas()
	//users.AltaUsuario()

	/*Pedro := new(modelos.Hombre)
	e.HumanosRespirando(Pedro)

	Maria := new(modelos.Mujer)
	e.HumanosRespirando(Maria) */

	//defer_panic.VemosDefer()

	//defer_panic.EjemplSoPanic()

	//defer_panic.EjemploRecover()

	/* MUESTRA GO ROUTINES

	go goroutines.MiNombreLento("Emilo José Zurita")

	fmt.Println("Estoy esperando: ")
	var x string
	//&x no es una variable, sino un puntero a la dirección de mem.x
	fmt.Scanln(&x)
	*/

	//Muestra de control de GOROUTINES con 'channel'
	// usando 'make', especificando el tipo de dato 'chan' y asociando
	// otro tipo que será del dato que canal va a controlar. (se puede
	// usar cualquier tipo de dato, no solo bool)
	canal1 := make(chan bool)

	fmt.Println("LLamando a la GO-Routina")
	//llamaremos a la funcion de forma asincrona adjuntando como argumento
	//un canal de control antes definido.
	go goroutines.MiNombreConChanel("Pepito Guzmán", canal1)
	defer func() { <-canal1 }()
	fmt.Println("Go-Routina en ejecución...")

	//establecemos un control de flag o cambio de estado de canal1
	// estado := <- canal1  Si queremos hacer pregunta sobre el estado del flag
	//usando canal de esta manera, el proceso se clava aqui, en su lugar
	// si usamos un defer (como ejemplo de arriva), el proceso de este script
	// continua y cuando acaba, solo finaliza si y solo si la goruotina
	// ha mandado la señal de final.

	//<-canal1

	// Podemos usar el control de flag en lugar de dejarlo al final
	// como colgado, es usar un 'defer'
	// defer func(){ <- canal1 }()
	fmt.Println("Señal de Go-ruotina acabada")
	fmt.Println("Script finalizado")

}
