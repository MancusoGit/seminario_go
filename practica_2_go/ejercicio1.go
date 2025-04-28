package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	CIUDAD = "Bariloche"
	CORTE  = "pepe"
	añoMax = 2025
	añoMin = 1900
	diaMax = 31
	diaMin = 1
	mesMax = 12
	mesMin = 1
)

var codigosCarrera map[int]string

var inscriptosCarrera map[string]int

type fecha struct {
	dia int
	mes int
	año int
}

type Estudiante struct {
	apellido        string
	nombre          string
	ciudadOrigen    string
	fechaNacimiento fecha
	analitico       bool
	codigoCarrera   int
}

type lista *nodo

type nodo struct {
	data Estudiante
	sig  lista
}

func generarDatos() (fecha, bool, int) {
	rand.Seed(time.Now().UnixNano())
	var date fecha
	date.año = rand.Intn(añoMax-añoMin+1) + añoMin
	date.mes = rand.Intn(mesMax-mesMin+1) + mesMin
	date.dia = rand.Intn(diaMax-diaMin+1) + diaMin
	codigo := rand.Intn(3) + 1
	valor := rand.Intn(11)
	if (valor % 2) == 0 {
		return date, true, codigo
	} else {
		return date, false, codigo
	}
}

func leerEstudiante() Estudiante {
	rand.Seed(time.Now().UnixNano())
	lector := bufio.NewReader(os.Stdin)
	var e Estudiante
	fmt.Print("Ingrese el apellido del estudiante: ")
	ap, _ := lector.ReadString('\n')
	e.apellido = strings.TrimSpace(ap)

	if !strings.EqualFold(e.apellido, CORTE) {
		fmt.Print("Ingrese el nombre del estudiante: ")
		nom, _ := lector.ReadString('\n')
		e.nombre = strings.TrimSpace(nom)

		fmt.Print("Ingrese la ciudad de origen del estudiante: ")
		ciudad, _ := lector.ReadString('\n')
		e.ciudadOrigen = strings.TrimSpace(ciudad)

		e.fechaNacimiento, e.analitico, e.codigoCarrera = generarDatos()
	}
	fmt.Println()
	return e
}

func agregarAdelante(punteroInicial *lista, e Estudiante) {
	nue := new(nodo)
	nue.data = e
	nue.sig = *punteroInicial
	*punteroInicial = nue
}

func cargarLista(punteroInicial *lista) {
	e := leerEstudiante()
	for {
		if strings.EqualFold(e.apellido, CORTE) {
			break
		}
		agregarAdelante(punteroInicial, e)
		e = leerEstudiante()
	}
}

func inicializarInscriptos() {
	inscriptosCarrera = make(map[string]int)
	inscriptosCarrera["APU"] = 0
	inscriptosCarrera["LS"] = 0
	inscriptosCarrera["LI"] = 0
}

func inicializarCarreras() {
	codigosCarrera = make(map[int]string)
	codigosCarrera[1] = "APU"
	codigosCarrera[2] = "LS"
	codigosCarrera[3] = "LI"
}

func (e Estudiante) toString() {
	fmt.Println("apellido y nombre del estudiante: ", e.apellido, " ", e.nombre)
	fmt.Println("fecha de nacimiento: ", e.fechaNacimiento.dia, "/", e.fechaNacimiento.mes, "/", e.fechaNacimiento.año)
	fmt.Println("ciudad de origen: ", e.ciudadOrigen)
	fmt.Println("carrera a estudiar: ", codigosCarrera[e.codigoCarrera])
	fmt.Println("estado de entrega titulo universitario: ", e.analitico)
}

func imprimirLista(pi lista) {
	if pi != nil {
		pi.data.toString()
		fmt.Println()
		imprimirLista(pi.sig)
	}
}

func maximoAñoInscriptos(mapaAños map[int]int) int {
	añoMax := 0
	maxCant := 0
	for año, cant := range mapaAños {
		if cant >= maxCant {
			añoMax = año
			maxCant = cant
		}
	}
	return añoMax
}

func maximaCarreraInscriptos() string {
	maxCant := 0
	maxCarrera := ""
	for _, carrera := range codigosCarrera {
		if inscriptosCarrera[carrera] >= maxCant {
			maxCant = inscriptosCarrera[carrera]
			maxCarrera = carrera
		}
	}
	return maxCarrera
}

func eliminarEstudiantes(pi *lista) {
	act := *pi
	ant := act
	for {
		if act == nil {
			break
		}
		if act.data.analitico {
			ant = act
			act = act.sig
		} else {
			if act == *pi {
				*pi = act.sig
				act = *pi
			} else {
				ant.sig = act.sig
				act = act.sig
			}
		}
	}
}

func procesarLista(pi lista) {
	contadorAños := make(map[int]int)
	for {
		if pi == nil {
			break
		}
		if strings.EqualFold(pi.data.ciudadOrigen, CIUDAD) {
			pi.data.toString()
			fmt.Println()
		}
		inscriptosCarrera[codigosCarrera[pi.data.codigoCarrera]]++
		contadorAños[pi.data.fechaNacimiento.año]++
		pi = pi.sig
	}
	ingresantesMax := maximoAñoInscriptos(contadorAños)
	carreraMaxInscriptos := maximaCarreraInscriptos()
	fmt.Println("el año en el que más ingresantes nacieron es: ", ingresantesMax)
	fmt.Println()
	fmt.Println("la carrera que más inscriptos tuvo fue ", carreraMaxInscriptos, ", con un total de ", inscriptosCarrera[carreraMaxInscriptos])
	fmt.Println()
}

func main() {
	inicializarCarreras()
	inicializarInscriptos()
	var pi lista
	pi = nil
	fmt.Println()
	fmt.Println("Bienvenido al ejercicio 1, by MancuSoftware (C) 2025.")
	fmt.Println()
	cargarLista(&pi) // IMPORTANTE: pasar dirección de pi
	fmt.Println("Lista antes de procesar")
	fmt.Println()
	imprimirLista(pi)
	fmt.Println("Lista post procesar")
	fmt.Println()
	procesarLista(pi)
	eliminarEstudiantes(&pi)
	fmt.Println("Lista con los estudiantes sin entrega del analitico eliminados ")
	fmt.Println()
	imprimirLista(pi)
}
