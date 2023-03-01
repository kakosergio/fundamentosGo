package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)

	go escrever("Olá mundo!", canal)

	//? Método de receber mensagens pelo canal com um for infinito
	// for {
	// 	mensagem, aberto := <-canal
	// 	if !aberto{
	// 		break
	// 	}
	// 	fmt.Println(mensagem)
	// }

	// Método de receber mensagens pelo canal utilizando ele como range no for
	// Para cada mensagem que chegar no canal, ele printa a mensagem, até que o canal seja fechado, então não precisa
	// testar a variável aberto. Ele faz isso por conta própria =D
	for mensagem := range canal{
		fmt.Println(mensagem)
	}
	fmt.Println("Fim do Programa!")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}
	close(canal)
}
