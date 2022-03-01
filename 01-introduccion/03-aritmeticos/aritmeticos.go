package main

import "fmt"

func main() {
	a := 20
	b := 10
	//comentario linea
	/*
		comentario bloque
	*/
	result := a + b
	fmt.Println("Suma:", result)

	result = a - b
	fmt.Println("Resta:", result)

	result = a * b
	fmt.Println("Multiplicación:", result)

	result = a / b
	fmt.Println("Divición:", result)

	result = 3 / b
	fmt.Println("Divición 1:", result)

	var div float64 = 3.0 / 10.0
	fmt.Println("Divición 1:", div)

	result = a % b
	fmt.Println("Modulo:", result)
}
