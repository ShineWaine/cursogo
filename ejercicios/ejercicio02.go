package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Entradanumero() {
	entrada := bufio.NewScanner((os.Stdin))
	for {
		fmt.Println("Introduce valor para tabla: ")
		if entrada.Scan() {
			numero, err := strconv.Atoi(entrada.Text())
			if err == nil {
				fmt.Printf("Tabla del número: %d \n", numero)
				for i := 1; i <= 10; i++ {
					fmt.Printf("%d x %d = %d \n", numero, i, numero*i)
				}
				break
			} else {
				fmt.Println("ERROR: Introducción erronea. Introduzca valor numérico. \n")
			}
		}

	}

}
