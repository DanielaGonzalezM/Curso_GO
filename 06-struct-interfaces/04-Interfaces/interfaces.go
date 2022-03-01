package main

import "fmt"

type Animal interface {
	mover() string
}
type Perro struct {
}
type Pez struct {
}
type Pajaro struct{}

func (*Perro) mover() string {
	return "soy perro y camino"
}
func (*Pez) mover() string {
	return "soy pex y nado"
}
func (*Pajaro) mover() string {
	return "soy pajaro y vuelo"
}
func moverAnimal(animal Animal) {
	fmt.Println(animal.mover())
}
func main() {
	perro := Perro{}
	moverAnimal(&perro)
	pez := Pez{}
	moverAnimal(&pez)
	pajaro := Pajaro{}
	moverAnimal(&pajaro)

}
