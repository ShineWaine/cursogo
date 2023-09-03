package files

import (
	"bufio"
	"fmt"

	//"bufio"
	"os"

	"github.com/ShineWaine/cursogo/ejercicios"
)

// Creación de constantes
var filename string = "./files/txt/tabla.txt"

// Funcion de creación de archivo y primera grabación de datos
func GrabaTabla() {
	//Cargamos en texto el retorno de la salida de la función TablaMultiplicar
	var texto string = ejercicios.TablaMultiplicar()

	//Creación de un archivo en el hd
	archivo, err := os.Create(filename)
	if err != nil {
		fmt.Println("ERROR de creación de archivo" + err.Error())
		return
	}

	//Inserción de datos en el archivo. Fprintln
	fmt.Fprintln(archivo, texto)

	//Cierre de archivo después de escritura
	archivo.Close()

}

// Función de actualización de archivo. Grabación de nuevos datos
func AnexaTabla() {
	//Cargamos en texto el retorno de la salida de la función TablaMultiplicar
	var texto string = ejercicios.TablaMultiplicar()

	//Verifica usando otra función si la grabación ha sido satisfactoria
	//if !Anexar(filename,texto) -->equivale a la linea inferior
	if Anexar(filename, texto) == false { //si el retorno de Anexar es false
		fmt.Println("ERROR al anexar contenido al archivo. ")
	}
}

func Anexar(archivo string, registro string) bool {
	//genera un objeto dato con la propiedades de anexar al archivo
	dato, err := os.OpenFile(archivo, os.O_WRONLY|os.O_APPEND, 644)

	if err != nil {
		fmt.Println("ERROR: Parece que existen problemas en la escritura del archivo." + err.Error())
		return false
	}

	_, err = dato.WriteString(registro)
	if err != nil {
		fmt.Println("ERROR de escritura." + err.Error())
		return false
	}
	dato.Close()
	return true
}

//Funcion de lectura de datos

func LeerDato() {
	archivo, err := os.Open(filename)
	if err != nil {
		fmt.Println("ERROR: Fallo en la lectura del archivo." + err.Error())
		return
	}

	//Creación de un objeto tipo bufio para captura de datos leidos procedentes de archivo
	scaner := bufio.NewScanner(archivo)

	for scaner.Scan() { //Mientras existan entradas en scaner
		registro := scaner.Text()
		fmt.Println(">> " + registro)
	}
	archivo.Close()

}
