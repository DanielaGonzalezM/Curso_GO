package main

import (
	"fmt"
	"strings"
)

func reverse(cadena string) string {

	arrayCadena := strings.Split(cadena, "")
	arrayInvertida := make([]string, 0)

	for i := len(arrayCadena) - 1; i >= 0; i-- {
		arrayInvertida = append(arrayInvertida, arrayCadena[i])
	}
	fmt.Println(arrayInvertida)

	return strings.Join(arrayInvertida, "")
}
func esPalindromo(palabra string) bool {

	palabra = strings.ToLower(palabra)
	fmt.Println(palabra)
	palabra = strings.ToUpper(palabra)
	fmt.Println(palabra)
	palabra = strings.ReplaceAll(palabra, " ", "")
	fmt.Println(palabra)
	palabraInvertida := reverse(palabra)
	fmt.Println(palabra)
	fmt.Println(palabraInvertida)
	return palabra == palabraInvertida

}
func main() {
	//ej. oso
	var palabra string
	fmt.Println("Ingrese palindromo")
	fmt.Scanln(&palabra)
	if esPalindromo(palabra) {
		fmt.Println(palabra, "es palindromo")
	} else {
		fmt.Println(palabra, "no es palindromo")
	}

}
