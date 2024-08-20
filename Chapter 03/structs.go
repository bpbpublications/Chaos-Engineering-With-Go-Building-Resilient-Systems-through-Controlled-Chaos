package main

import "fmt"

// Define a struct named "Fruits" with two fields: "name" of type string and "price" of type int.
type Fruits struct {
	name  string
	price int
}

func main() {
	// Create a new variable of type Fruits and initialize its fields.
	apple := Fruits{
		name:  "Apple",
		price: 30,
	}

	// Create another variable of type Fruits using the field names and their corresponding values.
	mango := Fruits{
		name:  "Mango",
		price: 75,
	}

	// Access and modify the fields of a struct using dot notation.
	fmt.Println(apple.name, "is of price: ", apple.price)
	fmt.Println(mango.name, "is of price", mango.price)
}
