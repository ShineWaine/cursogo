package main

import (
	e "github.com/ShineWaine/cursogo/ejer_interfaces"
	"github.com/ShineWaine/cursogo/modelos"
)

func main() {
	/*variables.NumerosEnteros()
	variables.RestoVariables()
	variables.Convierteanumero(1780)
	entero, frase := ejercicios.Cadenaaentero("500")
	fmt.Printf("%d, %s", entero, frase)
	teclado.IngresoNumeros() */
	//ejercicios.TablaMultiplicar()
	//files.GrabaTabla()
	//files.AnexaTabla()
	//files.LeerDato()
	//funciones.Calculos()
	//funciones.LlamarClosure()
	//funciones.Exponente(1)
	//arreglos_slices.MuestroArreglos()
	//arreglos_slices.MuestroSlices()
	//arreglos_slices.Capacidad()
	//mapas.MostrarMapas()
	//users.AltaUsuario()
	Pedro := new(modelos.Hombre)
	e.HumanosRespirando(Pedro)

	Maria := new(modelos.Mujer)
	e.HumanosRespirando(Maria)
}
