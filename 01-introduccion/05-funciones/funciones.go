package main

import "fmt"

func saludar(nombre string) {
	fmt.Println("Hola", nombre)

}
func sumar(a, b int) int {
	return a + b

}
func main() {
	saludar("Daniela")
	r := sumar(10, 20)
	println(r)
	println(sumar(10, 5))

}
