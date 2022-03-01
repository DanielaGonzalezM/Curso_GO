package main

import "fmt"

func main() {
	hola := "Hola"
	mundo := "Mundo"

	fmt.Print(hola)
	fmt.Println(hola)
	fmt.Println(mundo)

	nombre := "Daniela"
	edad := 30

	//string y entero
	fmt.Printf("Hola, %s y su edad es %d \n", nombre, edad)

	//cuando no se sabe el tipo de dato
	fmt.Printf("Hola, %v y su edad es %v \n", nombre, edad)

	mensaje := fmt.Sprintf("Hola, %s y su edad es %d \n", nombre, edad)

	fmt.Println(mensaje)

	fmt.Printf("nombre: %T \n", nombre)
	fmt.Printf("nombre: %T \n", edad)
	fmt.Println("Ingrese otro nombre: ")
	fmt.Scanln(&nombre)
	fmt.Println("Otro nombre:", nombre)

}
