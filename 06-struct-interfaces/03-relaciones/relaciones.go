package main

import "fmt"

type User struct {
	nombre string
	email  string
	activo bool
}

type Student struct {
	user   User
	codigo string
}

type Curso struct {
	titulo string
	videos []Video
}

type Video struct {
	titulo string
	curso  Curso
}

func main() {

	//relación de 1 a 1
	/*
		alex := User{
			nombre: "Alex",
			email:  "alex@gmail.com",
			activo: true,
		}
		roel := User{
			nombre: "Roel",
			email:  "roel@gmail.com",
			activo: true,
		}

		alexStuden := Student{
			user:   alex,
			codigo: "asdfa",
		}

		fmt.Println(roel)
		fmt.Println(alexStuden.user.nombre)
	*/

	//relacion de uno a muchos

	v1 := Video{
		titulo: "Introducción",
	}
	v2 := Video{
		titulo: "Instalación",
	}

	curso := Curso{
		titulo: "Curso de Go",
		videos: []Video{v1, v2},
	}
	v1.curso = curso
	v2.curso = curso

	fmt.Println(v1.curso.titulo)

	for _, video := range curso.videos {
		fmt.Println(video.titulo)
	}
}
