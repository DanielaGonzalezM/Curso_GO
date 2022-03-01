package main

import "fmt"

func main() {
	//not
	fmt.Println(!true)
	//and
	fmt.Println(true && true)
	//or
	fmt.Println(false || true)

	b := 2
	r := b == 2 || !(4 > b)

	fmt.Println(r)

	b = 2
	r = b == 2 && !(4 > b)

	fmt.Println(r)
}
