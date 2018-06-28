package main

import (
	"fmt"
	"strings"

	"github.com/JCFlores93/Go-sample/maps"
	// "github.com/JCFlores93/Go-sample/flow"
	// "github.com/JCFlores93/Go-sample/name"
	// "github.com/JCFlores93/Go-sample/numbers"
	// "github.com/JCFlores93/Go-sample/structs"
)

const helloWorld string = "Hola %s %s, bienvenido al fascinante mundo de GO \n"
const testConst = "Test"

/*
	go run main.go to execute a file with go
	go build main.go create a binary file to be execute called ./main(example)
*/
func main() {
	//name := getName()
	/*firstName := name.GetName()
	//Declaration and assignment, Go infers the kind of var
	lastName := "<Modificar>"
	number := numbers.Sum(50, 50)
	a, b, c := numbers.GetVariables()
	f32, f64 := numbers.GetFloat()
	stringUTF8 := name.GetUnicode()
	//println ==> line's jump
	//fmt.Print("Ingresa tu nombre: ")
	//Let us ask a value and use it.
	//fmt.Scanf("%s", &name)
	fmt.Printf(helloWorld, firstName, lastName)
	//Let us print in console
	//fmt.Print("Hola mundo")
	//Let us format the text
	fmt.Printf("Hola %s, bienvenido al fascinante mundo de Go. \n", firstName)
	fmt.Println("Hola mundo")

	fmt.Println(number, a, b, c)
	fmt.Println(f32, f64)
	fmt.Println("Cadena con UTF8: ", stringUTF8)
	fmt.Println(string("hola"[0]))
	fmt.Println("Cantidad de letras que tiene Hola ==> ", len("hola"))
	structs.GetArray()
	structs.GetSlice()
	flow.IfTest()
	forTest()
	flow.SwitchTest()
	strings2()*/

	fmt.Println(maps.GetMap("Freddy"))
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
