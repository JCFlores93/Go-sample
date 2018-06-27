package main

import "fmt"

const helloWorld string = "Hola %s %s, bienvenido al fascinante mundo de GO \n"
const testConst = "Test"

/*
	go run main.go to execute a file with go
	go build main.go create a binary file to be execute called ./main(example)
*/
func main() {
	var name string
	name = "Sin nombre"
	//Declaration and assignment, Go infers the kind of var
	lastName := "<Modificar>"
	var number = 100
	var (
		a = 1
		b = 2
		c = 3
	)
	//println ==> line's jump
	fmt.Print("Ingresa tu nombre: ")
	//Let us ask a value and use it.
	fmt.Scanf("%s", &name)
	fmt.Printf(helloWorld, name, lastName)
	//Let us print in console
	//fmt.Print("Hola mundo")
	//Let us format the text
	fmt.Printf("Hola %s, bienvenido al fascinante mundo de Go. \n", name)
	fmt.Println("Hola mundo")

	fmt.Println(number, a, b, c)
}
