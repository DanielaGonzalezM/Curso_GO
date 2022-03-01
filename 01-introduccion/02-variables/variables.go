package main

import "fmt"

func main() {
	var nombre1 string
	nombre1 = "Daniela"

	var nombre2 string = "Andrea"

	edad := 30

	fmt.Println(nombre1, nombre2, edad)

	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	//no se puede modificar despues
	const pi = 3.141592

	fmt.Println(pi)
}
