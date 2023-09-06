package middleware

import "fmt"

//funciones del mismo tipo, por lo que son susceptible de que un solo
// middleware pueda gestionarlas.

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

/*
La función MiMiddleware lo que hace es recoger el resultado de lo que seria el llamado
a las funciones suma,resta,multiplicacion y lo mostraría, vendría a ser como un
lanzador de funciones, solo que en lugar de llamar directamente a la función de cada
operación, lo que hace es interponer una función de middleware para 'dar otro tipo de
funcionalidad' a cada función de operaciones, pero sin tocar ni modificar nada de estas.
MiMiddelware llamará a la función middleware pasando la propia función de operación con
todos los parámetros necesarios para que dicha función de operación... funcione.
*/

func MiMiddleware() {
	fmt.Println("Inicio Middleware")

	resultado := operacionesMidd(sumar)(2, 3)
	fmt.Println(resultado)

	resultado = operacionesMidd(restar)(10, 6)
	fmt.Println(resultado)

	resultado = operacionesMidd(multiplicar)(2, 4)
	fmt.Println(resultado)
}

// Esta función, recibe los mismos parámetros que las funciones a quién
// audita, (suma, resta, multiplicacion) y devuelve exactamente lo mismo.
// f como variable y recibe una funcion con 2 int (aqui solo tenemos que indicar el tipo,
// no hace falta el nombre de la variable, y 1 int de salida, y repite
// los valores para la salida, una funcion con 2 parametros de entrada y un int de salida.
func operacionesMidd(f func(int, int) int) func(int, int) int {
	//realizaría aqui los cambios necesarios y retorna la función que sea
	//sumar, restar, etc...
	return func(a, b int) int {
		fmt.Println("Inicio de operación y o modificación")
		return f(a, b)
	}
}
