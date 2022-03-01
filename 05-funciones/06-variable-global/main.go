package main

import "fmt"

//Variable global
var mensaje string

func funcion1() {
	mensaje = "Hola desde funcion 1"
	fmt.Println(mensaje)
}

func funcion2() {
	mensaje = "Hola desde funcion 2"
	fmt.Println(mensaje)
}
func funcion3() {
	mensaje = "Hola desde funcion 3"
	fmt.Println(mensaje)
}

func main() {
	mensaje = "hola desde la funcion principal"
	fmt.Println(mensaje)
	defer funcion1() //defer hace que se ejecute al final
	funcion2()
	fmt.Println(mensaje)
}
