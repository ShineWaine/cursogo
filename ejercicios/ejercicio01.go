package ejercicios

import (
	"strconv"
)

func Cadenaaentero(cadena string) (int, string) {
	entero, err := strconv.Atoi(cadena)
	if err != nil {
		return 0, "Error de conversión \n" + err.Error() + "\n"
	}
	if entero > 100 {
		return entero, "Es mayor a 100 \n"
	} else {
		return entero, "Es menor a 100 \n"
	}
}
