package main

import "fmt"

//Struct persona
type Persona struct {
	nombre string
	edad   int
}

//metodos
func (p *Persona) imprimir() {
	fmt.Printf("Nombre: %s\nEdad: %d\n", p.nombre, p.edad)

}
func (p *Persona) editNombre(nombre string) {
	p.nombre = nombre

}

//herencia

type Empleado struct {
	Persona
	sueldo float64
}

func main() {
	p1 := Persona{"Daniela", 30}
	p1.imprimir()
	p1.editNombre("Dany")
	p1.imprimir()

	p2 := Persona{
		nombre: "Andrea",
		edad:   30,
	}

	p2.imprimir()

	em1 := Empleado{
		sueldo: 500,
	}
	em1.nombre = "Juanito"
	em1.edad = 30
	em1.imprimir()
	fmt.Println(em1)
}
