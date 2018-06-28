package main

import (
	"fmt"
	"strings"
)

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
	stringUTF8 := getUnicode()
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
	fmt.Println("Cadena con UTF8: ", stringUTF8)
	fmt.Println(string("hola"[0]))
	fmt.Println("Cantidad de letras que tiene Hola ==> ", len("hola"))
	getArray()
	getSlice()
	ifTest()
	forTest()
	strings2()
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

func getUnicode() string {
	return "欢迎来到丛林!"
}

func getArray() {
	//Arrays have an static space
	var array1 [2]string
	arr2 := [3]int{1, 2, 3}
	array1[0] = "array"
	array1[1] = "array2"
	fmt.Println(array1, arr2)
}

func getSlice() {
	//Slices are dynamic
	var slice1 []string
	slice1 = append(slice1, "mi", "slice", "1")
	fmt.Println(slice1)
}

func ifTest() {
	var number = 0
	fmt.Print("Ingresa un número del 1 al 10: ")
	fmt.Scanf("%d", &number)
	if number%2 == 0 {
		fmt.Println("El número es par")
	} else {
		fmt.Println("El número es impar")
	}

	if number2 := 3; number2 == 3 {
		fmt.Println("Entró al condicional")
	}
}

func forTest() {
	//Doesn't exist while,
	for i := 0; i < 5; i++ {
		fmt.Println("[FOR] ", i)
	}
	c := 100
	for c > 0 {
		c -= 100
		fmt.Println("[FOR] ", c)
	}
	s := 1000
	for {
		s--
		if s == 0 {
			fmt.Println("Termia el for 'infinito'")
		}
	}
}

func strings2() {
	var text = "Hello Platzi, Hello Go, Hello Platzi"
	fmt.Println(strings.ToUpper(text))
	fmt.Println(strings.ToLower(text))
	fmt.Println(strings.Replace(text, "Hello", "Hola", -1))
	fmt.Println(strings.Split(text, " "))
}

func switchTest() {
	var number = 0
	fmt.Print("[Switch] Ingresa un número del 1 al 10: ")
	fmt.Scanf("%d", &number)
	switch number {
	case 1:
		fmt.Println("El número es 1 ")
	default:
		fmt.Println("El número no es 1")
	}

	switch {
	case number%2 == 0:
		fmt.Println("El número es par")
	default:
		fmt.Println("E número no es par")
	}
}
