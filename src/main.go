package main

import (
	"fmt"
	"sync"
	"time"
)

// O canal 'buffer' funciona como a região crítica (memória compartilhada).
// Em Go, a troca de mensagens pelo canal gerencia implicitamente
// a sincronização e a exclusão mútua, substituindo mutexes explícitos.

// produtor envia dados para o canal.
func produtor(buffer chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() // Avisa que esta goroutine terminou

	for i := 1; i <= 5; i++ {
		fmt.Printf("[Produtor] Enviando: %d\n", i)
		// A operação de envio (<-) bloqueia se o canal (buffer) estiver cheio.
		// Esta é a sincronização.
		buffer <- i
		time.Sleep(500 * time.Millisecond) // Simula trabalho
	}
}

// consumidor recebe dados do canal.
func consumidor(buffer <-chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Avisa que esta goroutine terminou

	for i := 1; i <= 5; i++ {
		// A operação de recebimento (<-) bloqueia se o canal estiver vazio.
		// Esta é a sincronização.
		item := <-buffer
		fmt.Printf("\t[Consumidor] Recebido: %d\n", item)
		time.Sleep(1 * time.Second) // Simula consumo
	}
}

func main() {
	// Cria o buffer (região crítica) com capacidade para 2 itens.
	// Se a capacidade fosse 0 (não bufferizado), produtor e consumidor
	// teriam que se encontrar (rendezvous) a cada troca.
	buffer := make(chan int, 2)

	// WaitGroup é usado para esperar as goroutines terminarem.
	var wg sync.WaitGroup

	// Adicionamos 2 ao WaitGroup (1 para o produtor, 1 para o consumidor).
	wg.Add(2)

	// Inicia as goroutines
	go produtor(buffer, &wg)
	go consumidor(buffer, &wg)

	// Espera ambas as goroutines chamarem wg.Done()
	wg.Wait()

	// Opcional: Fechar o canal (não é estritamente necessário neste
	// exemplo de contagem fixa, mas é boa prática se o consumidor
	// usasse 'for range').
	close(buffer)
	fmt.Println("[Main] Produtor e Consumidor terminaram.")
}
