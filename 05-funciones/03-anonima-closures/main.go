package main

import (
	"fmt"
	"strings"
)

//Closure funcion que retorna una función
func repeat(n int) func(cadena string) string {

	return func(cadena string) string {
		return strings.Repeat(cadena, n)
	}
}
func main() {

	repeat3 := repeat(3)
	fmt.Println(repeat3("Hola"))
	repeat5 := repeat(5)
	fmt.Println(repeat5("Daniela"))
	/*
		//Función anonima
		func() {
			fmt.Println("Hola desde la función anónima")
		}()

		myfunc := func(nombre string) string {
			return (fmt.Sprintf("Hola, %s desde la función anónima2", nombre))
		}

		fmt.Println(myfunc("Daniela"))

	*/
}
