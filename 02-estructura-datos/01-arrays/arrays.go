package main

import "fmt"

func main() {
	var numeros [5]int
	fmt.Println(numeros)

	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30
	fmt.Println(numeros)
	fmt.Println(numeros[2])

	nombres := [3]string{"Daniela", "Andrea", "Bryam"}
	fmt.Println(nombres)

	colores := [...]string{
		"rojo", "verde", "negro",
	}
	fmt.Println(colores, len(colores))

	monedas := [...]string{
		0: "Dolar",
		2: "Soles",
		3: "Pesos",
		6: "Euro",
	}
	fmt.Println(monedas, len(monedas))

	subMoneda := monedas[0:3]
	fmt.Println(subMoneda)
	subMoneda = monedas[:3]
	fmt.Println(subMoneda)

}
