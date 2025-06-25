package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
)

const (
	bufferMax    = 20
	NUM          = 10
	pathPrioCero = "prioridad0.txt"
	pathPrioUno  = "prioridad1.txt"
)

type valor struct {
	num       int
	prioridad int
}

var prioTres, contadorPrioCero int

var wg sync.WaitGroup

var semaforoArchivoPrioridadCero, semaforoArchivoPrioridadUno, semaforoPrioridadTres, semaforoContadorPrioCero sync.Mutex

func descomponer(n int) int {
	total := 0
	for n != 0 {
		digito := n % 10
		total = total + digito
		n /= 10
	}
	return total
}

func invertir(n int) int {
	invertido := 0
	for n != 0 {
		digito := n % 10
		invertido = invertido*10 + digito
		n /= 10
	}
	return invertido
}

func prioridadCero(numero valor, path string) {
	semaforoArchivoPrioridadCero.Lock()
	arch, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer arch.Close()
	_, err = fmt.Fprintf(arch, "(%d,%d)", numero.prioridad, descomponer(numero.num))
	if err != nil {
		panic(err)
	}
	semaforoArchivoPrioridadCero.Unlock()
}

func prioridadUno(numero valor, path string) {
	semaforoArchivoPrioridadUno.Lock()
	arch, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer arch.Close()
	_, err = fmt.Fprintf(arch, "(%d,%d)", numero.prioridad, invertir(numero.num))
	if err != nil {
		panic(err)
	}
	semaforoArchivoPrioridadUno.Unlock()
}

func prioridadDos(numero valor) {
	fmt.Printf("el numero con prioridad %d luego de su proceso es: %d\n", numero.prioridad, numero.num*NUM)
	fmt.Println()
}

func prioridadTres(numero valor) {
	semaforoPrioridadTres.Lock()
	prioTres += numero.num
	fmt.Printf("total acumulado en la prioridad %d es: %d", numero.prioridad, prioTres)
	fmt.Println()
	semaforoPrioridadTres.Unlock()
}

func ElegirTarea(numero valor) {
	switch numero.prioridad {
	case 0:
		semaforoContadorPrioCero.Lock()
		contadorPrioCero++
		semaforoContadorPrioCero.Unlock()
		prioridadCero(numero, pathPrioCero)
		semaforoContadorPrioCero.Lock()
		contadorPrioCero--
		semaforoContadorPrioCero.Unlock()
	case 1:
		prioridadUno(numero, pathPrioUno)
	case 2:
		prioridadDos(numero)
	case 3:
		prioridadTres(numero)
	default:
		fmt.Println("prioridad no reconocida")
	}
}

func laburante(id int, canalesLaburantes []chan valor) {
	for {
		v := <-canalesLaburantes[id]
		ElegirTarea(v)
		fmt.Printf("tarea de prioridad %d ejecutada por el laburante %d\n", v.prioridad, id)
		wg.Done()
	}
}

func scheduler(canalesValores []chan valor, canalesLaburantes []chan valor) {
	for {
		var v valor
		v.prioridad = NUM
		select {
		case v = <-canalesValores[0]:
		default:
			if contadorPrioCero == 0 {
				select {
				case v = <-canalesValores[1]:
				default:
					select {
					case v = <-canalesValores[2]:
					default:
						select {
						case v = <-canalesValores[3]:
						default:
							fmt.Println("no se recibio un valor...")
						}
					}
				}
			}
		}

		//agarra el primer laburante disponible
		if v.prioridad != NUM {
			select {
			case canalesLaburantes[0] <- v:
				fmt.Println("worker 0 trabajando")
			case canalesLaburantes[1] <- v:
				fmt.Println("worker 1 trabajando")
			case canalesLaburantes[2] <- v:
				fmt.Println("worker 2 trabajando")
			case canalesLaburantes[3] <- v:
				fmt.Println("worker 3 trabajando")
			}
		}
	}
}

func init() {
	os.WriteFile(pathPrioCero, []byte(""), 0644)
	os.WriteFile(pathPrioUno, []byte(""), 0644)
}

func main() {

	var canalValores [4]chan valor
	var canalTareas [4]chan valor

	for i := 0; i < 4; i++ {
		canalValores[i] = make(chan valor, bufferMax)
		canalTareas[i] = make(chan valor, bufferMax)
	}

	for i := 0; i < bufferMax; i++ {
		v := valor{num: rand.Intn(900), prioridad: rand.Intn(4)}
		fmt.Printf("se genero el numero %d, con prioridad %d\n", v.num, v.prioridad)
		fmt.Println()
		canalValores[v.prioridad] <- v
	}

	wg.Add(bufferMax)

	//generador -> scheduler
	//scheduler -> tareas

	for i := 0; i < 4; i++ {
		go laburante(i, canalTareas[:])
	}

	go scheduler(canalValores[:], canalTareas[:])

	wg.Wait()
	fmt.Println("Procesos finalizados")
	fmt.Println("valor final de prioridad tres", prioTres)
}
