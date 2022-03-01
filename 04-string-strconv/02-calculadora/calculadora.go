package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sumar(expresion string) int {

	valores := strings.Split(expresion, "+")
	var result int

	for _, valor := range valores {
		num, error := strconv.Atoi(valor)
		if error != nil {
			fmt.Println(error)
			fmt.Println("solo debe ingresar numeros enteros a sumar")
		} else {
			result += num
		}

	}

	return result

}
func main() {

	var expresion string
	var result int
	fmt.Print("=>")
	fmt.Scanln((&expresion))
	result = sumar(expresion)

	fmt.Println(result)

}
