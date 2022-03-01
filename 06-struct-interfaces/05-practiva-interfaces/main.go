package main

import (
	"fmt"
	"math"
)

/*
En esta práctica vamos a sacar área y perímetro de un cuadrado y círculo utilizando interfaces.
Cuadrado
área = ancho * altura
perimetro = 2*ancho + 2*altura
Círculo
área = pi * (radio * radio)
perimetro = 2*pi * radio
*/
type Geometrica interface {
	area() float64
	perimetro() float64
}
type Cuadrado struct {
	lado float64
}
type Circulo struct {
	radio float64
}

func (cua *Cuadrado) area() float64 {
	return cua.lado * cua.lado
}
func (cua *Cuadrado) perimetro() float64 {
	return cua.lado * 4
}

func (cir *Circulo) area() float64 {
	return math.Pi * (cir.radio * cir.radio)
}

func (cir *Circulo) perimetro() float64 {
	return 2 * math.Pi * cir.radio
}
func medidas(g Geometrica) {
	fmt.Println(g)
	fmt.Println(g.area(), g.perimetro())
}
func main() {
	cua1 := Cuadrado{lado: 4}
	cir1 := Circulo{radio: 5}

	medidas(&cua1)
	medidas(&cir1)
}
