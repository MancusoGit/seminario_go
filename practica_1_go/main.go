package main

import (
	"fmt"
	"seminario/ejercicios"
)

func mainn() {

	/*
		 		Ejercicios Opcionales:
		 		fmt.Println("Ejercicio 1 : ")
				ejercicios.Imprimir250()
				fmt.Println()
				fmt.Println("Ejercicio 2 : ")
				num := 78
				resultado := ejercicios.Evaluar(num)
				fmt.Println("El resultado de evaluar el número", num, "es:", resultado)
				fmt.Println()
				fmt.Println("Ejercicio 3 : ")
				fmt.Println("El resultado de la función MayorMenor es:", ejercicios.MayorMenor())
				fmt.Println("El resultado de la función MenorMayor es:", ejercicios.MenorMayor()) // no se puede dividir por 0 (si se ingresa un numero negativo da error)
				fmt.Println("El resultado de la función MayorMenorFlotante es:", ejercicios.MayorMenorFlotante())
				fmt.Println()
	*/

	fmt.Println("Ejercicio Obligatorio 1  ")
	ejercicios.CambiarPalabras()
	fmt.Println()
	fmt.Println("Ejercicio Obligatorio 2  ")
	ejercicios.EjecutarEjOb2()
	fmt.Println()
	fmt.Println("Ejercicio Obligatorio 3  ")
	ejercicios.ProbarOcurrencias()
	fmt.Println()
}
