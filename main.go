package main

import (
	"fmt"
	"strings"
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

	//fmt.Println(maps.GetMap("Freddy"))
	//structs.InterfaceTest()
	// number, err := numbers.Sum("50", 50)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(number)
	//pointerTest()
	// go forGo(500)
	// go forGo(400)
	// time.Sleep(10000 * time.Millisecond)
	channels()
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

func pointerTest() {
	a := 100
	//creamos un puntero --> recibe una dirección de un espacio en memoria
	var b *int
	// & -> podmeos extrar la dirección en memoria de a
	b = &a
	// * -> extraemos el valor verdadero que está en esa dirección de memoria
	fmt.Println("Sin modificar")
	fmt.Println(a, *b)
	fmt.Println(&a, b)
	pointerModify(b)
	fmt.Println("Después de modificar")
	fmt.Println(a, *b)
	fmt.Println(&a, b)
}

//recibe una dirección de memoria
func pointerModify(c *int) {
	*c = 10
}

func helloGo(index int) {
	fmt.Println("Hola soy un print en la ", index)
}

func forGo(n int) {
	for i := 0; i < n; i++ {
		go helloGo(i)
	}
}

func channels() {
	ch := make(chan string)
	//funciones anónimas o closures
	go func() {
		//cerramos el canal
		defer close(ch)
		// <- es la forma como los channels reciben los datos
		ch <- "Hola channel"
	}()
	//si el canal está del lado izquierdo está recibiendo una info
	//si el canal está del lado derecho está enviando una info
	fmt.Println(<-ch)

	ch1 := make(chan int)
	go func() {
		defer close(ch1)
		ch1 <- 1
		ch1 <- 2
		ch1 <- 3
		ch1 <- 4
		ch1 <- 5
	}()
	for n := range ch1 {
		fmt.Printf("%d\n", n)
	}

	ch2 := make(chan int, 2)
	ch2 <- 1
	ch2 <- 2
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)
	ch2 <- 3
	fmt.Println(<-ch2)
}
