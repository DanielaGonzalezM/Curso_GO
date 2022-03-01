package mensajes

import "fmt"

//mayuscula es publica la función, con minuscula es privada
func Hola() {
	fmt.Println("Hola desde el paquete mensajes")

}

const mensaje = "Hola desde mi constante"

func funcionPrivada() {
	fmt.Println("Hola desde funcion privada")
}
func Imprimir() {
	fmt.Println(mensaje)
	funcionPrivada()

}
