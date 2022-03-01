package main

import (
	"errors"
	"fmt"
)

func division(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("No es posible dividir entre 0")
	} else {
		return dividendo / divisor, nil
	}

}

func main() {

	result, error := division(4, 0)
	if error == nil {
		fmt.Println("Resultado:", result)
	} else {
		fmt.Println(error)
	}

}
