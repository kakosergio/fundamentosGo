package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	// grupo de espera que tem duas goroutines que fazem parte dela
	waitGroup.Add(2)

	go func() {
		go escrever("Olá Mundo!")
		waitGroup.Done() // -1
	}()

	go func() {
		escrever("Programando em Go!")
		waitGroup.Done() // -1 no waitGroup.add(2)
	}()

	waitGroup.Wait()
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
