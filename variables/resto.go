package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoVariables() {
	Nombre = "Carles"
	Estado = true
	Sueldo = 1352.50
	Fecha = time.Now()

	fmt.Println("Nombre: ", Nombre)
	fmt.Println("Estado: ", Estado)
	fmt.Println("Sueldo: ", Sueldo)
	fmt.Println("Fecha: ", Fecha)
}

func Convierteanumero(numero int) (bool, string) {
	texto := strconv.Itoa(numero)
	fmt.Println("Estado conversión: ", true)
	fmt.Println("Texto convertido: ", texto)
	return true, texto
}
