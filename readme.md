# Trabalho de SO -- Troca de mensagens

Esse repositório contém os códigos e os documentos referentes a apresentação
de trabalho da matéria de Sitemas Operacionais do IFPE Campus Garanhuns.
Dos alunos: Iasmin, Raul, Oliveira e Juan.

## Exemplo de código em C

Esse trecho de código é puramente explicativo, não tem funcionalidade real.
Foi tirado do livro _Sistemas Operacionais Modernos 4° edição._ De Tanembaum.

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
