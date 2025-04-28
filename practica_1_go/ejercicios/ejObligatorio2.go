package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ImparReverb(frase string) string {
	// Esta funcion recibe una palabra y devuelve la misma palabra pero con las letras en orden inverso
	palabras := strings.Fields(frase) // convierte la palabra en un slice de palabras
	var resultado strings.Builder     // crea un buffer de strings
	for i := 0; i < len(palabras); i++ {
		if i%2 == 0 { // si es par
			revertirPalabra(palabras[i], &resultado)
		} else {
			resultado.WriteString(palabras[i]) // agrega la palabra al resultado
		}
		if i < len(palabras)-1 { // si no es la ultima palabra
			resultado.WriteString(" ") // agrega un espacio
		}
	}
	return resultado.String() // retorna el resultado
}

func revertirPalabra(palabra string, resultado *strings.Builder) {
	letras := []rune(palabra) // convierte la palabra en un slice de runes
	// Esta funcion recibe una palabra y la agrega al resultado pero en orden inverso
	for i := len(letras) - 1; i >= 0; i-- { // recorre la palabra al reves
		resultado.WriteRune(letras[i]) // agrega la letra al resultado
	}
}

func EjecutarEjOb2() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese una frase: ")
	frase, _ := reader.ReadString('\n')
	fmt.Println()
	fmt.Println("Frase original: ", frase)
	fmt.Println()
	fmt.Println("Frase modificada: ", ImparReverb(frase))
	fmt.Println()
}
