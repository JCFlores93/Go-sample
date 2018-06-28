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
