package main

import (
	"fmt"
	"sync"
	"time"
)

// Capacidade do buffer
const BUFFER_SIZE = 5

// Número total de itens a produzir
const NUM_ITEMS = 10

func producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for item := 0; item < NUM_ITEMS; item++ {
		fmt.Println("PRODUTOR: Gerando item", item)

		ch <- item

		fmt.Println("PRODUTOR: Enviou item", item)
		time.Sleep(100 * time.Millisecond)
	}

	close(ch)
	fmt.Println("PRODUTOR: Canal fechado.")
}

func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for item := range ch {
		fmt.Println("CONSUMIDOR: Recebeu item", item)
		time.Sleep(300 * time.Millisecond)
	}

	fmt.Println("CONSUMIDOR: Canal fechado e vazio. Encerrando.")
}

func main() {
	ch := make(chan int, BUFFER_SIZE)

	var wg sync.WaitGroup

	wg.Add(2)

	go producer(ch, &wg)
	go consumer(ch, &wg)

	fmt.Println("MAIN: Aguardando o término das goroutines...")
	wg.Wait()

	fmt.Println("MAIN: Todas as goroutines terminaram.")
}
