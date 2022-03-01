package main

import "fmt"

func main() {
	var vocal string
	fmt.Println("Ingrese una letra:")
	fmt.Scanln(&vocal)

	switch {
	case vocal == "a" || vocal == "A":
		fmt.Println(vocal, "es una vocal")
	case vocal == "e" || vocal == "E":
		fmt.Println(vocal, "es una vocal")
	case vocal == "i" || vocal == "I":
		fmt.Println(vocal, "es una vocal")
	case vocal == "o" || vocal == "O":
		fmt.Println(vocal, "es una vocal")
	case vocal == "u" || vocal == "U":
		fmt.Println(vocal, "es una vocal")
	default:
		fmt.Println(vocal, "no es una vocal")
	}

	switch vocal {
	case "a", "A", "e", "E", "i", "I", "o", "O", "u", "U":
		fmt.Println(vocal, "es una vocal")
	default:
		fmt.Println(vocal, "no es una vocal")
	}

}
