package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"sync"
	"time"
)

// isPrime devuelve true si el número es primo
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 3; i <= sqrtN; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func primesSingleThread(n int) []int {
	var primes []int
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

// IMPORTANTE: LAS GOROUTINES NO PUEDEN DEVOLVER VALORES DIRECTAMENTE, POR ESO USAMOS UN CANAL PARA RECIBIR LOS RESULTADOS
func primesConcurrent(n, numWorkers int, canalPrimos chan []int) {
	var wg sync.WaitGroup
	primesChan := make(chan []int, numWorkers)
	chunkSize := n / numWorkers

	for i := 0; i < numWorkers; i++ {
		start := i*chunkSize + 1
		end := (i + 1) * chunkSize
		if i == numWorkers-1 {
			end = n
		}

		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			var localPrimes []int
			for i := start; i <= end; i++ {
				if isPrime(i) {
					localPrimes = append(localPrimes, i)
				}
			}
			primesChan <- localPrimes
		}(start, end)
	}

	wg.Wait()
	close(primesChan)

	var allPrimes []int
	for partial := range primesChan {
		allPrimes = append(allPrimes, partial...)
	}

	canalPrimos <- allPrimes
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run programa.go <N>")
		return
	}

	n, err := strconv.Atoi(os.Args[1]) // Convertir el argumento a entero
	// Validar que el número sea positivo
	if err != nil || n <= 0 {
		fmt.Println("Ingrese un número entero positivo válido.")
		return
	}

	fmt.Println("Buscando primos hasta:", n)

	// Versión con una goroutine
	start := time.Now()
	primos1 := primesSingleThread(n)
	duration1 := time.Since(start)
	fmt.Printf("Versión 1 hilo: %v ms\n", duration1.Milliseconds())
	fmt.Println("Primos encontrados:", primos1)

	// Versión concurrente
	numWorkers := 4 // Puedes ajustarlo
	canalPrimos := make(chan []int)
	start = time.Now()
	go primesConcurrent(n, numWorkers, canalPrimos)
	primos2 := <-canalPrimos // Recibir los resultados del canal
	close(canalPrimos)       // Cerrar el canal después de recibir los resultados
	duration2 := time.Since(start)
	fmt.Printf("Versión %d goroutines: %v ms\n", numWorkers, duration2.Milliseconds())
	fmt.Println("Primos encontrados:", primos2)

	// Calcular Speed-up
	if duration2 > 0 {
		speedup := float64(duration1.Milliseconds()) / float64(duration2.Milliseconds())
		fmt.Printf("Speed-up con %d goroutines: %.2f\n", numWorkers, speedup)
	}
}
