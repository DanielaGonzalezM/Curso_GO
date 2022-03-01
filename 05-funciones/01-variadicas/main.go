package main

import "fmt"

//no sabemos cuandos numeros se entregarán a la función
func sumar(numeros ...int) int {
	var total int
	for _, num := range numeros {
		total += num
	}
	return total
}

//no sabemos cuandos numeros se entregarán a la función y se devuelven más de un valor
func sumar2(nombre string, numeros ...int) (string, int) {
	mensaje := fmt.Sprintf("La suma de %s es: ", nombre)
	var total int
	for _, num := range numeros {
		total += num
	}
	return mensaje, total
}
func main() {

	result := sumar(4, 5, 6, 7, 8)
	fmt.Println(result)

	mensaje, resultado := sumar2("Daneila", 4, 5, 6, 7, 8)
	fmt.Println(mensaje, resultado)
}
