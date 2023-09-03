package funciones

import "fmt"

func Exponente(valor int) {
	if valor > 100000000 {
		return
	}

	fmt.Println(valor)
	Exponente(valor * 2)
}
