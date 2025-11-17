# Troca de mensagens

Esse repositório contém os códigos e os documentos referentes a apresentação
de trabalho da matéria de Sitemas Operacionais do IFPE Campus Garanhuns.
Dos alunos: Iasmin, Raul Holanda, Carolaine Ferreira e Juan.

## Exemplo de código em C

Esse trecho de código é puramente explicativo, não tem funcionalidade real.
Foi tirado do livro _Sistemas Operacionais Modernos 4° edição._ De Tanenbaum.

```c
#define N 100 /* numero de lugares no buffer */

void producer(void) {
  int item;
  message m; /* buffer de mensagens */
  while (TRUE) {
    item = produce_item(); /* gera alguma coisa para colocar no buffer */
    receive(consumer, &m); /* espera que uma mensagem vazia chegue */
    build_message(&m, item); /* monta uma mensagem para enviar */
    send(consumer, &m); /* envia item para consumidor */
  }
}

void consumer(void) {
  int item, i;
  message m;
  for (i = 0; i < N; i++); /* envia N mensagens vazias */
  while (TRUE) {
    receive(producer, &m); /* pega mensagem contendo item */
    item = extract_item(&m); /* extrai o item da mensagem */
    send(producer, &m); /* envia a mensagem vazia como resposta */
    consume_item(item); /* faz alguma coisa com o item */
  }
}
```

```go
func produtor(buffer chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 5; i++ {
		fmt.Printf("[Produtor] Enviando: %d\n", i)
		buffer <- i
		time.Sleep(500 * time.Millisecond)
	}
}

func consumidor(buffer <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 5; i++ {
		item := <-buffer
		fmt.Printf("\t[Consumidor] Recebido: %d\n", item)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	buffer := make(chan int, 2)
	var wg sync.WaitGroup
	wg.Add(2)
	go produtor(buffer, &wg)
	go consumidor(buffer, &wg)
	wg.Wait()
	close(buffer)
	fmt.Println("[Main] Produtor e Consumidor terminaram.")
}
```
