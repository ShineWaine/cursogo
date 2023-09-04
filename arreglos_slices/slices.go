package arreglos_slices

import "fmt"

/*
En la definición de un slice,a diferencia de un arreglo o array,
no se especifica la cantidad de elementos que contendrá, por lo que
es posible en el tiempo de ejecución, ampliar su dimensión y por tanto
su contenido.

Los slices manejan dos datos, uno es la cantidad de elementos que tienen
y la segunda la capacidad, o lo que es lo mismo, la cantidad de memoria
que se ha reservado a ese slice para guardar datos, que siempre es
superior a la cantidad usada.

// En el ejemplo la asignación de valores, se fija a 3.
var tablaS []int = []int{2, 5, 8}

*/

var tablaS []int = []int{2, 5, 8}
var arreglo [10]int = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

func MuestroSlices() {
	fmt.Println(tablaS)

	//funcion slicing de los arreglos-slice
	//Crea porcion como array con los elementos del 3 al final de arreglo
	porcion := arreglo[3:]
	//Crea nuevo array con elementos de arreglo del 0 al 5 (n-1)
	porcion2 := arreglo[:6]
	//Crea nuevo array con elementos del 4 al 8 (n-1)
	porcion3 := arreglo[4:9]
	fmt.Println("Tabla completa: ", arreglo)
	fmt.Println("elementos [3: ]: ", porcion)
	fmt.Println("elementos [ :6]: ", porcion2)
	fmt.Println("elementos [4:9]: ", porcion3)
}

//Gestión de capacidad de un slice
// la capacidad se gestiona con 'make'. Este comando permite
// dimensionar un slice con un número determinado de posiciones
//para elementos, pero aumentando esas posiciones en memoria para
//una futura expansión del tamaño del slice.
//Para el compilador es más fácil manipular la memoria guardada
// que gestionar memoria nueva que no está reservada.

func Capacidad() {
	//'make' permite definir un slice vacio con una cantidad definida
	// un una capacidad extra.
	elementos := make([]int, 5, 20)
	// []int = slice, de 5 elementos pero de capacidad 20

	//'cap()', permite averiguar la capacidad del slice. Tamaño guardado en memoria
	fmt.Printf("Largo %d, Capacidad %d \n", len(elementos), cap(elementos))

	//creacion de slice con 0 posicion y 0 capacidad
	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	//Go crea un 2 a la n elementos más de capacidad que la indicada en tamaño
	fmt.Printf("\nLargo %d, Capacidad: %d \n", len(nums), cap(nums))

}
