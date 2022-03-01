package main

import "fmt"

func modificarValor(cadena *string) {
	fmt.Printf("%p\n", cadena) //imprime la referencia de la memoria
	*cadena = "Hola desde la funcion"
}

func main() {
	cadena := "Hola Mundo de Go"
	fmt.Printf("%p\n", &cadena)
	fmt.Println("Antes de la función cadena:", cadena)

	modificarValor(&cadena) //estamos pasando una copia sin el &

	fmt.Println("Despues de la función:", cadena)

}
