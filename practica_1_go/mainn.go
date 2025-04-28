package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func mainnn() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese una frase: ")
	frase, _ := reader.ReadString('\n')
	fmt.Println()
	fmt.Print("Ingrese la palabra a cambiar: ")
	ocurrencia, _ := reader.ReadString('\n')
	fmt.Println()
	if len(os.Args) > 1 {
		frase = os.Args[1]
		ocurrencia = os.Args[2]
	}
	fmt.Println("Frase original: ", frase)
	fmt.Println("Palabra a cambiar: ", ocurrencia)
	fmt.Println("Frase modificada: ", cambiarMinusMayus(frase, ocurrencia))
}

func cambiarMinusMayus(palabras, ocurrencia string) string {
	frase := strings.Fields(palabras)          // creo un slice de las palabras
	var result strings.Builder                 //creo un buffer para construir la frase resultante
	ocurrencia = strings.TrimSpace(ocurrencia) //elimino espacios en blanco al principio y al final

	for i, palabra := range frase {
		if strings.EqualFold(palabra, ocurrencia) { //comparo sin tener en cuenta mayusculas y minusculas
			permutacionChars(&result, palabra)
		} else {
			result.WriteString(palabra)
		}
		if i < len(frase)-1 {
			result.WriteString(" ") //agrego espacios entre palabras
		}
	}
	return result.String() //devuelvo la frase resultante
}

// esta funcion recibe el buffer por referencia y la palabra de la frase e intercambia los caracteres
// entre mayusculas y minusculas
func permutacionChars(result *strings.Builder, palabraFrase string) {
	for _, char := range palabraFrase {
		if unicode.IsLower(char) {
			result.WriteRune(unicode.ToUpper(char)) // paso a mayuscula
		} else {
			result.WriteRune(unicode.ToLower(char)) // paso a minuscula
		}
	}
}
