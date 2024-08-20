package main

import "fmt"

func main() {
	// Array,for array we need to specify the length
	animalArray := [3]string{"cat", "dog", "lion"}
	fmt.Println("Printint Array: ", animalArray)

	// Slice, for slice we don't need to specify the length
	countries := []string{"India", "UK", "US"}
	fmt.Println("Printing Slice: ", countries)
}
