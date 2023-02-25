package main

import "fmt"

func soma (numeros ... int) int {
	fmt.Println(numeros)

	total := 0

	for _, numero := range numeros{
		total += numero
	}
	return total
}
// Tem que ser sempre por último o argumento variático
func juntaTextoNumero (texto string, numeros ... int){
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main () {

	totalSoma := soma(1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20)

	fmt.Println(totalSoma)

	juntaTextoNumero("Hoje estou aqui, vai ter que decidir", 5,4,3,2,1,5,6,7,8,9,0,90)

}