package testingunitario

func Suma(a, b int) int {
	return a + b
}
func GetMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

//go test -cover indica el porcentaje que no se ha testeado
//go test -coverprofile=coverange.out             ****************crea el archivo para los testeos
//go tool cover -func=coverange.out      indica las funciones no testeadas
//go tool cover -func=coverange.out    indica en rojo lo que no está testeado
//ERA coverage!!!

//go test -cpuprofile=cou.out //nos ayuda a testear el tiempo en que se demora nuestra aplicacion
//go tool pprof cou.out    //ver el archivo
//top						//ver el que más se demora
//list Fibonacci
//web   // se tuvo que instalar:sudo apt install graphviz
//pdf		//te descarga la info como pdf
//quit
