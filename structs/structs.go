package structs

import "fmt"

//Get all elements
func GetArray() {
	//Arrays have an static space
	var array1 [2]string
	arr2 := [3]int{1, 2, 3}
	array1[0] = "array"
	array1[1] = "array2"
	fmt.Println(array1, arr2)
}

//Get a dynamic array
func GetSlice() {
	//Slices are dynamic
	var slice1 []string
	slice1 = append(slice1, "mi", "slice", "1")
	fmt.Println(slice1)
}

//a method of Platzicourse
func (p PlatziCourse) Subscribe(name string) {
	fmt.Printf("La persona %s se ha registrado al curso %s", name, p.Name)
}

//This is an struct
type PlatziCourse struct {
	Name   string
	Slug   string
	Skills []string
}

//add the struct of platzicourse
type PlatziCareer struct {
	PlatziCourse
}

//register a person to a career
func (p PlatziCareer) Subscribe(name string) {
	fmt.Printf("La persona %s se ha registrado a la carrera  %s", name, p.Name)
}
