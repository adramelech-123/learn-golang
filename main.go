package main

import (
	"fmt"
)



func main() {
	carInventory := map[string]int{
		"Sedan": 25,
		"SUV": 7,
		"Convertible": 10,
		"Hacthback": 8,
	}

	// Access both bodyType and count
	for bodyType, count := range carInventory {
		fmt.Printf("Key: %v -> Value: %v\n", bodyType, count)
	}

	// Access the values 
	totalInventory := 0
	for _, count := range carInventory {
		totalInventory += count
	}

	fmt.Printf("We have %v cars in total\n", totalInventory)
} 