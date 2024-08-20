package main

import "fmt"

// Function to add two numbers
func add(input1 int, input2 int) int {
	return input1 + input2
}

func main() {
	// Assign the add function to a variable called "variableHavingFunction"
	variableHavingFunction := add

	outPut := variableHavingFunction(10, 20) // Call the function using that variable
	fmt.Println("Output:", outPut)
}
