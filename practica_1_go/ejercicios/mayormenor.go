package ejercicios

import (
	"fmt"
)

func MayorMenor() int {
	fmt.Print("Ingrese el primer número: ")
	var num1 int
	fmt.Scanln(&num1)
	fmt.Print("Ingrese el segundo número: ")
	var num2 int
	fmt.Scanln(&num2)
	if num1 >= num2 {
		return num1 / num2
	} else {
		return num2 / num1
	}
}

func MenorMayor() uint {
	fmt.Print("Ingrese el primer número: ")
	var num1 uint
	fmt.Scanln(&num1)
	fmt.Print("Ingrese el segundo número: ")
	var num2 uint
	fmt.Scanln(&num2)
	if num1 >= num2 {
		return num1 / num2
	} else {
		return num2 / num1
	}
}

func MayorMenorFlotante() float32 {
	fmt.Print("Ingrese el primer número: ")
	var num1 float32
	fmt.Scanln(&num1)
	fmt.Print("Ingrese el segundo número: ")
	var num2 float32
	fmt.Scanln(&num2)
	if num1 >= num2 {
		return num1 / num2
	} else {
		return num2 / num1
	}
}
