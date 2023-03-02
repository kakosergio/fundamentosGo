package main

import (
	"fmt"
	"time"
)

func main (){
	// Cria uma variável que recebe a função generator
	canal := escrever("Olá mundo")

	// Fica escutando o canal e printando a mensagem recebida
	for mensagem := range canal {
		fmt.Println(mensagem)
	}

}

//? Esse é o padrão GENERATOR, onde você encapsula uma goroutine dentro de uma função que retorna um canal
func escrever (texto string) <- chan string{
	// Cria um canal
	canal := make(chan string)

	// Cria uma função anônima como goroutine
	go func(){
		for {
			canal <- fmt.Sprintf("Valor Recebido: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()
	// Retorna o canal de string
	return canal
}