package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func cajero(id int, fila chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for cliente := range fila {
		tiempoEspera := rand.Intn(5) + 1 // Espera entre 1 y 5 segundos
		fmt.Printf("Cajero %d atendiendo a %s durante %d segundos\n", id, cliente, tiempoEspera)
		time.Sleep(time.Duration(tiempoEspera) * time.Second)
		fmt.Printf("Cajero %d ha terminado de atender a %s\n", id, cliente)
		fmt.Println()
		fmt.Printf("Cajero %d estara listo para atender el siguiente cliente en %d segundos...\n", id, tiempoEspera)
		fmt.Println()
		time.Sleep(time.Duration(tiempoEspera) * time.Second)
	}
}

func cajeroPaja(id int, fila chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		cliente := <-fila
		tiempoEspera := rand.Intn(5) + 1 // Espera entre 1 y 5 segundos
		fmt.Printf("Cajero %d atendiendo a %s durante %d segundos\n", id, cliente, tiempoEspera)
		time.Sleep(time.Duration(tiempoEspera) * time.Second)
		fmt.Printf("Cajero %d ha terminado de atender a %s\n", id, cliente)
		fmt.Println()
		fmt.Printf("Cajero %d estara listo para atender el siguiente cliente en %d segundos...\n", id, tiempoEspera)
		fmt.Println()
		time.Sleep(time.Duration(tiempoEspera) * time.Second)
	}
	fmt.Println("Cajero", id, "ha terminado de atender a sus 3 clientes.")
}

func mainnnnn() {
	var wg sync.WaitGroup
	fila := make(chan string, 10) // Canal con capacidad para 10 clientes

	// simular 1 cajero
	wg.Add(1)
	go cajeroPaja(0, fila, &wg)

	// Simular llegada de clientes
	clientes := []string{"Cliente 1", "Cliente 2", "Cliente 3", "Cliente 4", "Cliente5"}
	for _, cliente := range clientes {
		fila <- cliente
		fmt.Printf("%s ha llegado a la fila\n", cliente)
		fmt.Println()
		time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second) // Tiempo entre llegadas de clientes
	}

	// Simular 3 cajeros
	for i := 1; i < 3; i++ {
		wg.Add(1)
		go cajeroPaja(i, fila, &wg)
	}

	clientesNombres := []string{"ey", "dou", "chi", "bye", "ayi", "fay", "guy", "huy", "jey", "kuy", "ley", "mey", "ney", "oey", "pey", "qey", "rey", "sey", "tey", "uey", "vey", "wey", "xey", "yey", "zey"}
	for _, c := range clientesNombres {
		fila <- c
		fmt.Printf("%s ha llegado a la fila\n", c)
		fmt.Println()
		time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second) // Tiempo entre llegadas de clientes
	}

	close(fila) // Cerrar el canal para indicar que no hay más clientes
	fmt.Println("Todos los clientes han llegado a la fila, esperando a que los cajeros terminen.")
	fmt.Println()

	if len(fila) == 0 {
		fmt.Println("Todos los cajeros han terminado de atender a los clientes.")
		fmt.Println()
	} else {
		fmt.Println("Aún quedan clientes en la fila.")
		c := rand.Intn(3) // generar un número aleatorio entre 0 y 2
		fmt.Printf("Cajero %d se encargará de los clientes restantes.\n", c)
		wg.Add(1)
		go cajero(c, fila, &wg)
		wg.Wait() // Esperar a que el cajero termine
	}
	fmt.Println("Fin del programa.")
}
