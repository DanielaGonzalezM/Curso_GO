package main

import "fmt"

func main() {

	a := 4
	b := 5
	var r bool

	r = a == b

	fmt.Printf("%d es igual que %d?: %t \n ", a, b, r)

	r = a != b
	fmt.Printf("%d es distinto que %d?: %t \n ", a, b, r)

	r = a > b
	fmt.Printf("%d es mayor que %d?: %t \n ", a, b, r)

	r = a < b
	fmt.Printf("%d es menor que %d?: %t \n ", a, b, r)

	r = a >= b
	fmt.Printf("%d es mayot o igual que %d?: %t \n ", a, b, r)

	r = a <= b
	fmt.Printf("%d es menor o igual que %d?: %t \n ", a, b, r)
}
