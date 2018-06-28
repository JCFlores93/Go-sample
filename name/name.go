package name

import "fmt"

// la letra mayúscula es para método públicos y minúscula para privados
//GetName obtiene y retorna el nomrbe de la persona
func GetName() string {
	var name string
	name = "Sin nombre"
	fmt.Print("Ingresa tu nombre: ")
	fmt.Scanf("%s", &name)
	return name
}
