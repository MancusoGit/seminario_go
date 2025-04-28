package ejercicios

import (
	"fmt"
)

func Imprimir250() {
	for i := 0; i <= 250; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func Revprint250() {
	fmt.Println()

	for i := 250; i >= 0; i-- {
		if i%2 == 0 {
			fmt.Println("Inverti en $libra. la siguiente cantidad:", i)
		}
	}
	fmt.Println("Era todo una estafa hay que vender la casa.")
}
