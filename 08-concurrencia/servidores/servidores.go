package main

import (
	"fmt"
	"net/http"
	"time"
)

func revisarServidore(servidor string, canal chan string) {

	_, error := http.Get(servidor)
	if error != nil {
		//fmt.Println("Error en el servidor")
		canal <- "Error en el servidor"
	} else {
		//fmt.Println(servidor, ": OK")
		canal <- servidor + ": OK"
	}
}
func main() {
	inicio := time.Now()
	canal := make(chan string)
	servidores := []string{
		"http://oregoom.com/",
		"http://www.udemy.com/",
		"http://www.youtube.com/",
		"http://www.facebook.com/",
		"http://www.google.com/",
	}
	for _, servidor := range servidores {
		go revisarServidore(servidor, canal)
	}
	for i := 0; i < len(servidores); i++ {
		fmt.Println(<-canal)
	}
	tiempoPasado := time.Since(inicio)
	fmt.Println(tiempoPasado)
}
