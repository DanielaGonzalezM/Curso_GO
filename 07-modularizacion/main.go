package main

/*este import no funciona correctamente debido a los - en la ruta por lo que se debe crear un go.mod para los paquetes
import (
	"Curso1/07-modularizacion/mensajes/"
)*/

//go mod init paquetes
import (
	"fmt"
	//"paquetes/figuras"
	"paquetes/models"

	"github.com/DanielaGonzalezM/figuras"
	"github.com/donvito/hellomod"
)

func main() {
	/*
		mensajes.Hola()
		mensajes.Imprimir()
	*/
	cua1 := figuras.Cuadrado{Lado: 8}
	figuras.Medidas(&cua1)
	cir := figuras.Circulo{Radio: 5}
	figuras.Medidas(&cir)

	p1 := models.Persona{}
	p1.Constructor("Daniela", 30)
	fmt.Println(p1)
	p1.SetNombre("Dany")
	p1.SetEdad(29)
	fmt.Println(p1)
	hellomod.SayHello()
}

//go get github.com/donvito/hellomod
//obtener paquetes de terceros en go.mod
