package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Cria um canal que recebe uma função que tem mais de um canal, e recebe tudo no mesmo canal
	canal := multiplexar(escrever("Olá world!"), escrever("Porgameiro em Golem!"))

	for i := 0; i < 30; i++ {
		fmt.Println(<-canal)
	}

}

//? Cria a função multiplexar que recebe dois canais e retorna um canal
//? Recebe informações nos dois canais e transmite pra um canal só de saída
func multiplexar(canalDeEntrada1, canalDeEntrada2 <-chan string) <-chan string {
	canalDeSaida := make(chan string)

	// Aqui temos uma função anônima que escuta quem enviou e quem recebeu pra poder enviar para o canal de saída
	// e esse canal de saída tá sendo recebido lá no main
	go func() {
		for {
			select {
			case mensagem := <-canalDeEntrada1:
				canalDeSaida <- mensagem
			case mensagem := <-canalDeEntrada2:
				canalDeSaida <- mensagem
			}
		}
	}()
	return canalDeSaida
}

// ? Esse é o padrão GENERATOR, onde você encapsula uma goroutine dentro de uma função que retorna um canal
func escrever(texto string) <-chan string {
	// Cria um canal
	canal := make(chan string)

	// Cria uma função anônima como goroutine
	go func() {
		for {
			canal <- fmt.Sprintf("Valor Recebido: %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()
	// Retorna o canal de string
	return canal
}
