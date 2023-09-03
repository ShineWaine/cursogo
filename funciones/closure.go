package funciones

import "fmt"

//Implementación de CLOSURE.
//Destino, ofuscación de código, de tal manera que se puede conocer el
//dato inicial pero no los intermedios.

// Definición de tabla, funcion LOCAL con entero como entrada y como salida
// una función entera.
func tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	fmt.Println("secuencia: " + string(secuencia) + " valor: " + string(numero))
	return func() int {
		secuencia++
		return numero * secuencia
	}
}

// Definición funcion PUBLICA que asigna un argumento que usa la función
// tabla para generar un resultado, de tal manera que lo que tabla devuelve
// no es un valor, sino una función, sin visionar los valores intermedios que
// dicha función va generando.
func LlamarClosure() {
	tabladel := 1
	MiTabla := tabla(tabladel)

	for i := 1; i <= 10; i++ {
		fmt.Println(MiTabla())
	}

}
