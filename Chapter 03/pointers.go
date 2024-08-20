package main

import "fmt"

func main() {
	// creating the variable using short hand technique
	testVariable := 500
	testPointer := &testVariable // testPointer is a pointer to the memory address of testVariable

	// It will print the value of the variable
	fmt.Println("Value of testVariable:", testVariable)

	// It will print the memory address of testVariable
	fmt.Println("Value of testPointer:", testPointer)

	// Using * with variable name is used to print the value of that variable
	fmt.Println("Value pointed by testPointer:", *testPointer)

	// Using * with pointerVariable name with = assignment operator we can change the value
	// of the actual variable
	*testPointer = 600
	fmt.Println("Updated value pointed by testPointer:", *testPointer)

}
