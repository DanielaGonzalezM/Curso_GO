package figuras

import "fmt"

type Geometrica interface {
	area() float64
	perimetro() float64
}

func Medidas(g Geometrica) {
	fmt.Println(g)
	fmt.Println("Area:", g.area(), "Perimetro:", g.perimetro())
}
