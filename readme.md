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

## Exemplo em Go

```go
const BUFFER_SIZE = 5
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
```

## Exemplo em Python

```py
import threading
import queue

canal = queue.Queue()   # fila atua como canal de mensagens

def produtor():
  for i in range(5):
    canal.put(f"msg {i}") # envia mensagem
    print(f"Produtor enviou msg {i}")
  canal.put(None)         # mensagem de encerramento

def consumidor():
  while True:
    msg = canal.get()     # recebe mensagem
    if msg is None:
      break               # termina
    print(f"Consumidor recebeu {msg}")

threading.Thread(target=produtor).start()
threading.Thread(target=consumidor).start()
```
