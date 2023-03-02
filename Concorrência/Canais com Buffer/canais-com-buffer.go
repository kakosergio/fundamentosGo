package main

func main(){
	//? Cria canal que espera dois valores (por isso o 2 depois do chan string)
	canal := make(chan string, 2)
	canal <- "Olá Mundo!"
	canal <- "Programando em Go!"
	// canal <- "Programando em Go de novo!"

	mensagem := <- canal
	mensagem2 := <- canal

	println(mensagem)
	println(mensagem2)
}