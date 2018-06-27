package main

import "fmt"

/*
	go run main.go to execute a file with go
	go build main.go create a binary file to be execute called ./main(example)
*/
func main() {
	var name string
	//println ==> line's jump
	fmt.Print("Ingresa tu nombre: ")
	//Let us ask a value and use it.
	fmt.Scanf("%s", &name)
	//Let us print in console
	//fmt.Print("Hola mundo")
	//Let us format the text
	fmt.Printf("Hola %s, bienvenido al fascinante mundo de Go. \n", name)
	fmt.Println("Hola mundo")
}
