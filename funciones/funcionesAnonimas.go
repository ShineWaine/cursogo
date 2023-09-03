package funciones

import "fmt"

// Implementacion funciones anónimas.
// Como usar una variable con una función como valor.

func Calculos() {

	suma := func(numero1 int, numero2 int) int {
		return numero1 + numero2
	}

	fmt.Println(suma(5, 30))

}
