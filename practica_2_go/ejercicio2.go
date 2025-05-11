package main

import "fmt"

type Values struct {
	numero      int
	apariciones int
}

type OptimumSlice struct {
	slice []Values
}

func New(s []int) OptimumSlice {
	var pi OptimumSlice
	pi.slice = nil
	if len(s) != 0 {
		numAct := s[0]
		cant := 0
		for _, num := range s {
			if num == numAct {
				cant++
			} else {
				pi.slice = append(pi.slice, Values{numAct, cant})
				numAct = num
				cant = 1
			}
		}
		pi.slice = append(pi.slice, Values{numAct, cant})
	}
	return pi
}

func (o OptimumSlice) IsEmpty() bool {
	return len(o.slice) == 0
}

func (o OptimumSlice) Len() int {
	return len(o.slice)
}

func (o OptimumSlice) FrontElement() int {
	if o.IsEmpty() {
		return 0
	}
	return o.slice[0].numero
}

func (o OptimumSlice) LastElement() int {
	if o.IsEmpty() {
		return 0
	}
	return o.slice[len(o.slice)-1].numero
}

// retorna 0 si salio todo bien, y -1 si hubo un fallo
func (o *OptimumSlice) Insert(element int, position int) int {
	if position < 0 || position > len(o.slice) {
		return -1
	}
	pos := o.Len() - position
	if o.slice[pos].numero == element {
		o.slice[pos].apariciones++
		return 0
	}
	primerVec := o.slice[:pos]
	segundoVec := o.slice[pos:]
	primerVec = append(primerVec, Values{element, 1})
	o.slice = append(primerVec, segundoVec...)
	return 0
}

func (o OptimumSlice) SliceArray() []int {
	var vec []int
	for _, v := range o.slice {
		for i := 0; i < v.apariciones; i++ {
			vec = append(vec, v.numero)
		}
	}
	return vec
}

func Mostrar(o OptimumSlice) {
	if o.IsEmpty() {
		fmt.Println("el slice esta vacio")
		return
	}
	for i := 0; i < len(o.slice); i++ {
		fmt.Printf("Valor: %d, Cantidad: %d\n", o.slice[i].numero, o.slice[i].apariciones)
	}
}

func main() {
	vec := []int{1, 1, 1, 2, 2, 2, 2, 2, 3, 3, 3, 3, 6, 44, 4}
	slice := New(vec)
	fmt.Println(vec)
	fmt.Println("la cantida de elementos del slice es: ", slice.Len())
	fmt.Println("el primer elemento del slice es: ", slice.FrontElement())
	fmt.Println("el ultimo elemento del slice es: ", slice.LastElement())
	slice.Insert(2, 4)
	Mostrar(slice)
	vec = slice.SliceArray()
	fmt.Println(vec)
}
