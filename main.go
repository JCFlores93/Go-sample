package main

import "fmt"

const helloWorld string = "Hola %s %s, bienvenido al fascinante mundo de GO \n"
const testConst = "Test"

/*
	go run main.go to execute a file with go
	go build main.go create a binary file to be execute called ./main(example)
*/
func main() {
	name := getName()
	//Declaration and assignment, Go infers the kind of var
	lastName := "<Modificar>"
	number := sum(50, 50)
	a, b, c := getVariables()
	f32, f64 := getFloat()
	//println ==> line's jump
	//fmt.Print("Ingresa tu nombre: ")
	//Let us ask a value and use it.
	//fmt.Scanf("%s", &name)
	fmt.Printf(helloWorld, name, lastName)
	//Let us print in console
	//fmt.Print("Hola mundo")
	//Let us format the text
	fmt.Printf("Hola %s, bienvenido al fascinante mundo de Go. \n", name)
	fmt.Println("Hola mundo")

	fmt.Println(number, a, b, c)
	fmt.Println(f32, f64)
}

func getName() string {
	var name string
	name = "Sin nombre"
	fmt.Print("Ingresa tu nombre: ")
	fmt.Scanf("%s", &name)
	return name
}

func getVariables() (int, int32, int64) {
	//2147000000 = max. value for int32
	return 1, 2147000000, 211231231231231233
}

func sum(a int, b int) int {
	return a + b
}

func getFloat() (float32, float64) {
	return float32(0.1), float64(float32(0.1))
}
