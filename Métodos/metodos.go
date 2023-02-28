package main

import "fmt"

type usuario struct {
	name string
	idade uint8
}

func (u usuario) salvar(){
	fmt.Printf("Salvando dados do Usuário %s no banco de dados.\n", u.name)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade > 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main (){
	usuario1 := usuario{"Kako", 36}
	usuario1.salvar()

	usuario2 := usuario{"Cintya", 34}
	usuario2.salvar()

	maiorDeIdade := usuario2.maiorDeIdade()
	fmt.Println(maiorDeIdade)

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)

}