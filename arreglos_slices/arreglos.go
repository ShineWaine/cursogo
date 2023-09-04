package arreglos_slices

import "fmt"

/* Los arreglos o arrays tienen un número de elementos concretos y
definidos. Dicho 'tamaño' no puede variarse dentro del proceso que ocupa.

IMPORTANTE: El indice de elementos de una tabla inicia con el 0.
*/

// var slices []int	//Define un slice. No consta el num.de elementos
var tabla [10]int //Define una tabla. Consta que tiene 10 elementos y
// 					la inicializa con 0.

//creación y asignación directa
// var tabla [10]int = [10]int{10, 0, 59} //asigna valor a los tres primeros
//			elementos del array, los otros se inicializan con 0

// Creación de una MATRIZ
var matriz [20][30]int

func MuestroArreglos() {
	// asignación de valores a las posiciónes del array
	tabla[2] = 33
	tabla[4] = 98

	//crea una nueva tabla y asigna valores, el resto se llenan con ""
	tabla2 := [10]string{"Pablo", "Pepito", "Jesus"}

	fmt.Println(tabla)
	fmt.Println(tabla2)

	//Recorrer una tabla
	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}

	//Relleno de valores a una matriz. Lo mismo que a un vector pero
	// usando los dos índices
	matriz[3][25] = 25
	matriz[5][15] = 35
	matriz[10][10] = 10
	matriz[18][7] = 7

	for filas := 0; filas < len(matriz); filas++ {
		fmt.Printf("Fila: %d - %d \n", filas, matriz[filas])
		//for columnas := 0; columnas < len(matriz[filas]); columnas++ {
		//	fmt.Printf("Columna: %d. - %d \n", columnas, matriz[filas])
	}
}
