package main

import (
	"fmt"
	//"github.com/ShineWaine/cursogo/variables"
	"github.com/ShineWaine/cursogo/ejercicios"
)

func main() {
	/*variables.NumerosEnteros()
	variables.RestoVariables()
	variables.Convierteanumero(1780)*/
	entero, frase := ejercicios.Cadenaaentero("500")
	fmt.Printf("%d, %s", entero, frase)
}
