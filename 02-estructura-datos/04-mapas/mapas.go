package main

import (
	"fmt"
)

func main() {

	dias := make(map[int]string)

	fmt.Println(dias)

	//agregando datos
	dias[1] = "Lunes"
	dias[2] = "Martes"
	dias[3] = "Miercoles"
	dias[4] = "Jueves"
	dias[5] = "Viernes"
	dias[6] = "Sabado"
	dias[7] = "Domingo"

	fmt.Println(dias)

	dias[7] = "DOMINGO"
	fmt.Println(dias)

	delete(dias, 1)
	fmt.Println(dias, len(dias))

	//nueva mapa
	estudieantes := make(map[string][]int)

	estudieantes["alex"] = []int{13, 15, 16}
	estudieantes["roel"] = []int{14, 13, 17}

	fmt.Println(estudieantes)

	fmt.Println(estudieantes["alex"][1])
}
