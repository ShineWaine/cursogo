package ejer_interfaces

import (
	"fmt"

	"github.com/ShineWaine/cursogo/interfaces"
)

// funcion HumanosRespirando que recibe parámetro 'hu' del tipo de
// estructura interfaces.Humano
func HumanosRespirando(hu interfaces.Humano) {
	hu.Respirar()
	fmt.Printf("Soy un/a %s, y estoy respirando \n", hu.Sexo())
}
