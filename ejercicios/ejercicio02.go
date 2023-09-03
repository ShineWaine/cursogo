package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Generación variable para acumulación de tabla para posterior grabado
// en archivo
var tabla string

func TablaMultiplicar() string {
	entrada := bufio.NewScanner((os.Stdin))
	for {
		fmt.Println("Introduce valor para tabla: ")
		if entrada.Scan() {
			numero, err := strconv.Atoi(entrada.Text())
			if err == nil {
				fmt.Printf("Tabla del número: %d \n", numero)
				tabla += fmt.Sprintf("Tabla del número: %d \n", numero)
				for i := 1; i <= 10; i++ {
					fmt.Printf("%d x %d = %d \n", numero, i, numero*i)
					tabla += fmt.Sprintf("%d x %d = %d \n", numero, i, numero*i)
				}
				tabla += "\n"
				break
			} else {
				fmt.Println("ERROR: Introducción erronea. Introduzca valor numérico. \n")
			}
		}

	}
	return tabla
}
