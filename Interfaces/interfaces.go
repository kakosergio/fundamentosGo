package main

import (
	"fmt"
	"math"
)

// cria uma interface forma para ser usada por outras formas
type forma interface {
	area() float64
}

// Função que escreve no terminal a área da forma que a chamar
func escreverArea(f forma) {
	fmt.Printf("A área da forma é %0.2f\n", f.area())
}

// struct de uma forma retangular
type retangulo struct {
	altura  float64
	largura float64
}

// implementa a interface
func (r retangulo) area() float64 {
	return r.altura * r.largura
}

// struct de uma forma circular
type circulo struct {
	raio float64
}

// implementa a interface
func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func main() {

	r := retangulo{15, 30}
	escreverArea(r)

	c := circulo{17}
	escreverArea(c)
}
