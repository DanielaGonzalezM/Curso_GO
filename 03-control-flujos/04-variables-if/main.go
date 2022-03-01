package main

import "fmt"

func main() {

	//estas variables solo se pueden utilizar dentro del if
	if nombre, edad := "Alexx", 26; nombre == "Alex" {
		fmt.Println("Hola,", nombre)
	} else {
		fmt.Printf("Nombre: %s - Edad: %d\n", nombre, edad)
	}

	users := make(map[string]string)

	users["Alex"] = "alex@gmail.com"
	users["Roel"] = "roel@gmail.com"
	users["Juan"] = "juan@gmail.com"

	if correo, ok := users["Aalex"]; ok {
		fmt.Println(correo, ok)
	} else {
		fmt.Println("Error")
	}

}
