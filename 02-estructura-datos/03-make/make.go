package main

import "fmt"

func main() {

	// tipo de dato, longitud, capacidad
	numeros := make([]int, 3, 3)

	numeros[0] = 100
	numeros[1] = 200
	numeros[2] = 300

	numeros = append(numeros, 400)
	fmt.Println(numeros, len(numeros), cap(numeros))

	numeros2 := make([]int, 0, 3)

	numeros2 = append(numeros2, 400)
	fmt.Println(numeros2, len(numeros2), cap(numeros2))
}
