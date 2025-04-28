package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func extrapolar(palabra, destino string) string {

	//esta funcion lo que hace es recibir dos strings, palabra y destino
	//la funcion debe retornar la palabra con las letras que esten en origen convertidas a las letras de destino

	chars := []rune(palabra)  // convierte la palabra en un slice de runes - chars
	result := []rune(destino) // convierte la palabra en un slice de runes - result

	for i := 0; i < len(chars) && i < len(result); i++ {
		if chars[i] >= 'A' && chars[i] <= 'Z' { // si la letra es mayuscula
			result[i] = rune(strings.ToUpper(string(result[i]))[0]) // convierte la letra a mayuscula
		}
	}
	return string(result) // convierte el slice de runes a string
}

func Ocurrencias(frase, ocurrencia, reemplazo string) string {

	//esta funcion lo que hace es recibir una frase y 2 palabras claves, una sera la ocurrencia y la otra el reemplazo
	//cada vez que en la frase aparezca la ocurrencia, se debe reemplazar por el reemplazo respetando los caracteres mayusculas y minusculas

	var resultado strings.Builder // creo un buffer de strings

	palabras := strings.Fields(frase) //convierto la frase en un slice de palabras

	for i, palabra := range palabras { // recorro las palabras de la frase
		if strings.EqualFold(palabra, ocurrencia) { // comparo si la palabra es igual a la ocurrencia sin distinguir mayusculas
			palabraNueva := extrapolar(palabra, reemplazo) //llamo a la funcion extrapolar
			if i > 0 {
				resultado.WriteString(" ")
			}
			resultado.WriteString(palabraNueva)
		} else {
			if i > 0 {
				resultado.WriteString(" ")
			}
			resultado.WriteString(palabra)
		}
	}

	return resultado.String() // retorna el resultado

}

func CambiarPalabras() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese una frase: ")
	frase, _ := reader.ReadString('\n')
	fmt.Println()
	fmt.Println("frase original: ", frase)
	fmt.Println("frase modificada: ", Ocurrencias(frase, "gato", "perro"))
	fmt.Println()
}
