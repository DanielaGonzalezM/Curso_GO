package main

import "fmt"

func main() {

	//Bucle infinito
	for {
		fmt.Println("Hola")
		break //sin eso sigue eternamente
	}
	//Bucle tipo while
	numeros := 12355321
	c := 0
	for numeros > 0 {
		numeros /= 10
		c++
	}

	fmt.Println("Cantidad de digitos es:", c)
	//Bucle for
	for i := 0; i <= 10; i++ {
		if i%2 == 0 { //numeros pares
			fmt.Println(i)
		}
	}

}
