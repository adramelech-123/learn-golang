package main

import (
	"fmt"
)


func main() {
	fmt.Println("What's your name?")
	var name string
	fmt.Scanln(&name) //The &name syntax passes a pointer to the name variable, allowing fmt.Scanln to modify the variable's value directly.

	fmt.Printf("Hello,%s 👻 \n", name)
}


	



