package main

import "fmt"

func main() {
	nombre := [...]string{"Alex", "Roel", "Juan"}

	for i := 0; i < len(nombre); i++ {
		fmt.Println(nombre[i])
	}

	for indice, elemento := range nombre {
		fmt.Println(indice, elemento)
	}

	for _, elemento := range nombre {
		fmt.Println(elemento)
	}

}
