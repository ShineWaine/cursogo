package funciones

import "fmt"

/*
Las funciones recursivas son las que se llaman a si mismas para completar
una función.
Muy importante diseñar la salida o el final de la llamada para no caer
en un bucle sin final.
*/

func Exponente(valor int) {
	if valor > 100_000_000 {
		return
	}

	fmt.Println(valor)
	Exponente(valor * 2)
}
