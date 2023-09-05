package defer_panic

import (
	//"os"
	"fmt"
	"log"
)

/*
El comando 'defer' nos permite configurar la ejecución de la última
instrucción que se llevará a cabo antes de abandonar la función en curso.
Esta funcionalidad puede servirnos para prevenir una falta de cierre de
un archivo ante una salida abrupta del script, etc.

'Defer' solo permite usar una sola instrucción, o en su lugar una función
anónima, pero es posible usar varios defer en una misma función, lo que
ocasionará que el orden de ejecución sea según el orden en que estén
programados. (posición)

'Defer' se ejecutara aun siendo la salida de la función una salida
provocada por un 'panic'.

Formato de uso:

	defer 'comando a ejecutar'
*/
func VemosDefer() {
	fmt.Println("Este es el PRIMER mensaje.")
	defer fmt.Println("Este es el mensaje FINAL.")
	fmt.Println("Este es el SEGUNDO mensaje")
}

/*
'panic' es un comando que aborta el proceso del script, mostrando un
mensaje previamente programado. Es util ante posibles problemas del
script, donde podemos hacer caer el proceso antes de que él caiga
de forma accidental.

*/

func EjemploPanic() {
	a := 1
	if a == 1 {
		panic("Valor 1 encontrado.")
	}
}

/*
'recover´ es quizá el comando contrario al comando 'panic'.
'recover' es la comando para recuperar el proceso de un script que
ha roto por la ejecución de un 'panic'.

recover captura los códigos de salida de 'panic' en forma de 'log'
para poder paliar y solventar la causa de ese 'panic'.

ES OBLIGATORIO USAR EL 'RECOVER' JUNTO CON EL 'DEFER' ya que si no es
imposible que el 'recover' capture las salidas de log del 'panic'.

'LOG' (necesario importar la libreria 'log') es como un Println pero
informa en todo momento de la fecha y hora de cada publicación, por
lo que es muy importante a la hora de hacer auditorias, etc.

'log' contiene muchos comandos, los cuales permiten hacer distintas
acciónes a parte de mostrar un mensaje fechado y controlar los códigos
de error generados por 'panic'.
*/

//VEASE LA IMPLEMENTACIÓN DE FUNCIÓN ANÓNIMA EN EL DEFER
// hay que indicar '()' al final de la definición de la función anónima
// como salida de parámetro, ya que defer devuelve una función.

func EjemploRecover() {
	defer func() {
		//esta función anónima siempre se ejecuta al salir de la función.
		//si no hubo 'panic', reco = nil (todo correcto).
		//si hubo 'panic', reco contendrá un código de error el cual podremos
		//evaluar.
		reco := recover()
		if reco != nil {
			//'Fatalf' equivale a Printf seguido de os.Exit() salida del script
			//'El contenido de 'reco' será el mensaje que programamos como mensaje
			//en el comando 'panic'.
			log.Fatalf("Ocurrio un error generando PANIC.\n %v", reco)
		}
	}()

	a := 1
	if a == 1 {
		panic("Valor 1. Genera un panic")
	}

}
