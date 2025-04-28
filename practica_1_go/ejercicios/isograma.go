package ejercicios
import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)
// Isograma verifica si una palabra es un isograma, es decir, si no tiene letras repetidas

func Isograma(palabra string) bool {
	letras := []rune(palabra) // creo un slice de las letras
	mapa := make(map[rune]bool) // creo un mapa para almacenar las letras
	for _, letra := range letras { // recorro las letras
		if unicode.IsLetter(letra) { // si la letra es una letra
			
}