package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	// grupo de espera que tem duas goroutines que fazem parte dela
	waitGroup.Add(4)

	go func() {
		go escrever("Olá Mundo!")
		waitGroup.Done() // -1 - waitGroup.Add(3)
	}()

	go func() {
		escrever("Programando em Go!")
		waitGroup.Done() // -1 - waitGroup.add(2)
	}()

	go func() {
		go escrever("Eita Preula!")
		waitGroup.Done() // -1 - waitGroup.add(1)
	}()

	go func() {
		escrever("Bora fazer back-end!")
		waitGroup.Done() // -1 - waitGroup.add(0)
	}()

	waitGroup.Wait() //  0 - Espera acabar as goroutines que foram indicadas no waitGroup.add(4), até que todas terminem
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
