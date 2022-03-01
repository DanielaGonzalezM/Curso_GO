package main

import "fmt"

func main() {
	numeros := []int{1, 2, 3}

	fmt.Println(numeros, len(numeros))

	numeros = append(numeros, 4)
	numeros = append(numeros, 5)
	fmt.Println(numeros, len(numeros))

	subSlicen := numeros[:2]

	numeros[0] = 100

	fmt.Println(subSlicen)
	fmt.Println(numeros)

	meses := []string{"Enero", "Febrero", "Marzo"}

	fmt.Printf("Len: %v, Cap: %v, %p \n", len(meses), cap(meses), meses)

	meses = append(meses, "Abril")
	fmt.Printf("Len: %v, Cap: %v, %p \n", len(meses), cap(meses), meses)

}
