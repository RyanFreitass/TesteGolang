package main

import "fmt"

func mainListas() {
	
	// Array
	var array [2]string
	array[0] = "Hello"
	array[1] = "World"
	//fmt.Println(array[0], array[1])
	//fmt.Println(array)

	numPrimos := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(numPrimos)


	// Slices
	//var slice []string
	slice := make([]string, 5)
	slice[0] = "Hello"
	slice[1] = "World"
	fmt.Println(slice[0], slice[1])
	fmt.Println(slice)
	fmt.Println(len(slice))

	numPares := []int{2, 4, 6, 8, 10, 12}
	fmt.Println(numPares)

	numPares = append(numPares, 14)
	fmt.Println(numPares)

}