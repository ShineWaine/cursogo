package mapas

import "fmt"

/* Los mapas son colecciones de datos solo que no se almacenan ni
se consultan con un índice, sino que son pares de datos clave:valor.
Basicamente lo que se conoce en otros lenguajes como diccionarios.
NO TIENEN ORDEN DE MUESTREO
*/

func MostrarMapas() {
	//La definición de un mapa es muy parecida a un slice
	paises := make(map[string]string)
	//map : indica que es un mapa
	//[string]: Definicion del tipo de dato para la clave
	//string : Definición del tipo de dato para el valor

	//Asignación de datos al mapa
	paises["Mexico"] = "D.F."
	paises["Argentina"] = "Buenos Aires"

	//Asigna una clave y un valor. Automáticamente se almacena.

	//Lista todo el contenido del mapa
	fmt.Println(paises)

	//Lista solo de la clave su valor.
	fmt.Println(paises["Argentina"])

	//Creación de mapa por asignación directa
	campeonato := map[string]int{
		"Barcelona":    39,
		"Real Madrid":  38,
		"Chivas":       37,
		"Boca Juniors": 30,
	}

	//Listado total del mapa
	fmt.Println(campeonato)

	//Recorrer un mapa
	for index, valor := range campeonato {
		fmt.Printf("Equipo: %s, tiene %d puntos.\n", index, valor)
	}

	//Eliminar un elemento de un mapa
	delete(campeonato, "Real Madrid")
	//delete : orden para borrar
	//campeonato: indica el nombre del mapa
	//"Real Madrid": indica el indice de dato a borrar. SE BORRA
	//TODO, TANTO INDICE COMO VALOR

	fmt.Println(campeonato)

	// Consulta de la existencia de una clave
	// Cada consulta a un mapa, devuelve dos valores:
	//- El valor de la clave
	//- Un bool que dice si la clave existe o no
	valor, existe := campeonato["Real Madrid"]
	fmt.Printf("El valor es: %d, y el equipo existe = %t \n", valor, existe)
	valor, existe = campeonato["Chivas"]
	fmt.Printf("El valor es: %d, y el equipo existe = %t \n", valor, existe)

}
