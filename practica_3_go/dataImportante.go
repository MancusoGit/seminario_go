package main

import (
	"fmt"
	"sync"
	"time"
)

// ✅ Función que simula una tarea que corre en una goroutine.
func tarea(id int, sem chan struct{}, wg *sync.WaitGroup, mu *sync.Mutex, contador *int, resultados chan string) {
	defer wg.Done() // 🔁 Siempre marcar la goroutine como finalizada (aunque haya errores)

	// 🚦 Adquirir "permiso" para ejecutar usando un semáforo simulado
	// Esto bloquea si ya hay 3 tareas en ejecución
	sem <- struct{}{} // Enviar un valor al canal bloquea si está lleno

	fmt.Printf("🔄 Tarea %d: comenzó\n", id)

	// 💤 Simulamos trabajo usando sleep con una duración variable
	time.Sleep(time.Duration(1+id%3) * time.Second)

	// 🔒 Sección crítica: actualizamos el contador protegido con un Mutex (Es como un bloquear esquina)
	BloquearEsquina(mu)
	*contador++
	LiberarEsquina(mu)

	// 📡 Enviamos un mensaje con el resultado al canal
	resultado := fmt.Sprintf("✅ Tarea %d: finalizada", id)
	resultados <- resultado

	// 🚦 Liberar "permiso" del semáforo
	<-sem // Libera un lugar para otra goroutine
}

func BloquearEsquina(m *sync.Mutex) {
	// Simulamos un bloqueo de esquina
	m.Lock()
}

func LiberarEsquina(m *sync.Mutex) {
	// Simulamos una liberación de esquina
	m.Unlock()
}

func main() {
	var wg sync.WaitGroup // Para esperar a todas las goroutines
	var mu sync.Mutex     // Para proteger el acceso al contador compartido (monitor de acceso)
	contador := 0         // Contador de tareas finalizadas

	const maxConcurrentes = 3
	// 🧮 Semáforo con capacidad máxima de 3 tareas concurrentes
	sem := make(chan struct{}, maxConcurrentes)

	// 📮 Canal para recibir mensajes de resultados desde las goroutines
	resultados := make(chan string)

	// 📦 Canal para señal de terminación del receptor de resultados
	done := make(chan struct{})

	// 📻 Goroutine que actúa como "monitor" de resultados
	go func() {
		for {
			select {
			case msg := <-resultados:
				// 💬 Imprimimos los mensajes a medida que llegan
				fmt.Println(msg)
			case <-done:
				// 🛑 Cuando recibimos señal de cierre, salimos del bucle
				return
			}
		}
	}()

	// 🚀 Lanzamos 10 tareas concurrentes
	for i := 1; i <= 10; i++ {
		wg.Add(1) // Indicamos que vamos a esperar otra goroutine
		go tarea(i, sem, &wg, &mu, &contador, resultados)
	}

	// ⏳ Esperamos a que todas las goroutines terminen
	wg.Wait()

	// ✅ Cuando terminan todas, señalamos al receptor que ya no habrá más resultados
	done <- struct{}{}

	// 🎉 Imprimimos el resumen
	fmt.Println("🎉 Todas las tareas completadas.")
	fmt.Printf("🔢 Tareas finalizadas correctamente: %d\n", contador)
}

/*
monitores:
	esta el mutex y el rwmutex que sirve para leer y escribir
	con el mutex no se puede leer y escribir al mismo tiempo esta el lock y unlock exclusivo
	si tengo un rwmutex si hago un rlock puedo hacer multiples rlocks (lecturas)


semaforo:

el struct{} es un tipo de dato vacío en Go, que se utiliza comúnmente como señal o marcador.

sem := make(chan struct{}, 3) creo un semáforo con capacidad 3
sem <- struct{}{} // señal: "ocupo un lugar"
<-sem             // señal: "libero un lugar"
