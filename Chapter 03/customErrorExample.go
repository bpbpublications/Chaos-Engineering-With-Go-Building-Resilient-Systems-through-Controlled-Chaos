package main

import (
	"fmt"
)

type UserDefinedError struct {
	errorMessage string
}

func (err UserDefinedError) Error() string {
	return err.errorMessage
}

func divison(input1, input2 float64) error {
	if input2 == 0 {
		return UserDefinedError{"This is a user defined error"}
	}
	return nil
}

func main() {
	err := divison(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("No error occured.")
	}
}
