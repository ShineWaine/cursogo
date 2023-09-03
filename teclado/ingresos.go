package teclado

import (
	"bufio" //librería de entrada/salida teclado, archivos, etc
	"fmt"   //libreria salida datos por defecto
	"os"    //librería para interacciónes con el S.O.
	"strconv"
)

var num1 int
var num2 int
var leyenda string
var err error //Variable err del tipo error

func IngresoNumeros() {
	scaner := bufio.NewScanner(os.Stdin) //inicializa objeto scaner
	// usa os.Stdin para decir que la entrada proviene del teclado

	fmt.Println("Entra un valor (numero 1): ")
	if scaner.Scan() { //Si se entró algo por teclado
		num1, err = strconv.Atoi(scaner.Text()) //Por defecto las entradas
		//de bufio siempre son strings, por lo que hay que convertir el string
		//a número.
		if err != nil {
			//panic("mensaje") --> aborta la ejecución del script
			panic("Introducción de dato Erronea" + err.Error())
		}
	}

	fmt.Println("Entra un valor (numero 2): ")
	if scaner.Scan() { //Si se entró algo por teclado
		num2, err = strconv.Atoi(scaner.Text()) //Por defecto las entradas
		//de bufio siempre son strings, por lo que hay que convertir el string
		//a número.
		if err != nil {
			//panic("mensaje") --> aborta la ejecución del script
			panic("Introducción de dato Erronea" + err.Error())
		}

		fmt.Println("Entra un comentario: ")
		if scaner.Scan() { //Si se entró algo por teclado
			leyenda = scaner.Text()
		}

		fmt.Println(leyenda, num1*num2)
	}
}
