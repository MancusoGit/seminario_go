package main

import "fmt"

//	"fmt"

type listita *nodito

type OptimumSlice struct {
	lista             listita
	ultimoNodo        listita
	cantidadElementos int
}

type nodito struct {
	numero      int
	apariciones int
	sig         listita
}

func New(s []int) OptimumSlice {
	var pi OptimumSlice
	pi.lista = nil
	pi.cantidadElementos = 0
	numAct := s[0]
	cant := 0
	for _, num := range s {
		if num == numAct {
			cant++
		} else {
			agregarElemento(&pi, numAct, cant)
			numAct = num
			cant = 1
		}
	}
	agregarElemento(&pi, numAct, cant)
	return pi
}

func Mostrar(o OptimumSlice) {
	if o.lista == nil {
		fmt.Println("La lista está vacía.")
		return
	}
	actual := o.lista
	for actual != nil {
		fmt.Printf("Valor: %d, Cantidad: %d\n", actual.numero, actual.apariciones)
		actual = actual.sig
	}
}

func FrontElement(o OptimumSlice) int {
	return o.lista.numero
}

func LastElement(o OptimumSlice) int {
	return o.ultimoNodo.numero
}

func agregarElemento(pi *OptimumSlice, num, cant int) {
	nue := new(nodito)
	nue.numero = num
	nue.apariciones = cant
	nue.sig = nil
	if pi.lista == nil {
		pi.lista = nue
	} else {
		pi.ultimoNodo.sig = nue
	}
	pi.ultimoNodo = nue
	pi.cantidadElementos++
}

func IsEmpty(o OptimumSlice) bool {
	return o.lista == nil
}

func Len(o OptimumSlice) int {
	cant := 0
	list := o.lista
	for {
		if list == nil {
			break
		} else {
			cant++
			list = list.sig
		}
	}
	return cant
}

func SliceArray(o OptimumSlice) []int {
	list := o.lista
	var vec []int
	for list != nil {
		for i := 0; i < list.apariciones; i++ {
			vec = append(vec, list.numero)
		}
		list = list.sig
	}
	return vec
}

func main() {
	vec := []int{1, 1, 1, 2, 2, 2, 2, 2, 3, 3, 3, 3, 6, 44, 4}
	slice := New(vec)
	fmt.Println(vec)
	fmt.Println("la cantida de elementos del slice es: ", Len(slice))
	fmt.Println("el primer elemento del slice es: ", FrontElement(slice))
	fmt.Println("el ultimo elemento del slice es: ", LastElement(slice))
	Mostrar(slice)
	vec = SliceArray(slice)
	fmt.Println(vec)
}
