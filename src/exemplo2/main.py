import queue
import threading

canal = queue.Queue()  # fila atua como canal de mensagens


def produtor():
    for i in range(5):
        canal.put(f"msg {i}")  # envia mensagem
        print(f"Produtor enviou msg {i}")
    canal.put(None)  # mensagem de encerramento


def consumidor():
    while True:
        msg = canal.get()  # recebe mensagem
        if msg is None:
            break  # termina
        print(f"Consumidor recebeu {msg}")


threading.Thread(target=produtor).start()
threading.Thread(target=consumidor).start()
