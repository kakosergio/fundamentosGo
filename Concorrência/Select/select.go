package main

import (
	"fmt"
	"time"
)

func main (){
	//? Cria os canais que irão receber dados
	canal1, canal2 := make(chan string), make(chan string)

	//? Cria uma função anônima que será uma goroutine para enviar dados para o canal a cada meio segundo
	go func(){
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
	}()

	//? Cria uma função anônima que será uma goroutine para enviar dados para o canal a cada dois segundos
	go func(){
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	//? Cria um for que ficará escutando esses canais e printando na tela os dados recebidos
	for {
		// Select é como um switch, só que para canais (streams)
		select {
		case mensagemCanal1 := <- canal1:
			fmt.Println(mensagemCanal1)
		case mensagemCanal2 := <- canal2:
			fmt.Println(mensagemCanal2)
		}
	}
}