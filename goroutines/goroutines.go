package goroutines

import (
	"fmt"
	"strings" //librería de funciones para trabajo con strings
	"time"
)

/*
Las goroutines son procesos asincronos, es decir, son procesos que requieren
un cierto tiempo para ser ejecutados y no por eso, el script ha de pararse
hasta que esos procesos hayan concluido. El script continua la ejecución a
pesar de que los procesos lanzados antes no hayan concluido.
IMPORTANTE: Todos los procesos lanzados de forma asincrona finalizarán
cuando finalize el script principal, sin contar que hayan o no terminado
su propio proceso.

La forma de lanzar una gorutina es simplemente preceder a la llamada de
la función con 'go'.

	llamada a función normal: personal.adduser()
	llamada a funcion asincrona: go personal.adduser()
*/

// funcion recibe un string y lo va mostrando a razón de letra por segundo.
func MiNombreLento(nombre string) {
	//Split: comando que separa una cadena por el signo indicado generando
	// un slice como salida.
	// al indicar "" como separador, lo que hace es usar cada letra de
	// la palabra como elemento separado.
	letras := strings.Split(nombre, "")

	for _, letra := range letras {
		//sleep cuenta milisegundos
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)

	}

}

// Esta vez llamamos a la funcion de forma asincrona pero pasando un
// parametro que es 'canal1' del tipo 'chan' (canal) y que gestionará
// un valor 'bool'
func MiNombreConChanel(nombre string, canal1 chan bool) {
	//Vamos a repetir la función pero esta vez controlando el final
	//de su proceso para no dejarlo "tirado" al finalizar el script
	//que lo llama. Para eso usaremos chanel (canal1).
	letras := strings.Split(nombre, "")
	for _, letra := range letras {
		//sleep cuenta milisegundos
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
	// Cuando la rutina finalice, cambiaremos el estado de canal1
	// para indicar al proceso que lanzo esta función de forma asincrona
	// que ya ha terminado su proceso. Para hacerlo usaremos '<-'.
	canal1 <- true

}
