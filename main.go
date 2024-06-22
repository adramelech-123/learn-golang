package main

import (
	"fmt"
	"strings"
)


func main() {
	
	str := "Hello, World!"
	length := len(str) //Get length of a string
	fmt.Println("Length is ",length)

	str1 := "Go Programming"
	str2 := "go programming"

	//strings.EqualFold() checks if two strings are equal. The EqualFold method from the strings library is case insensitive
	fmt.Println(strings.EqualFold(str1, str2))

	str3 := "Hello, Worldly!"
	wIndex := strings.Index(str3, "W") //Provides the index of the first instance of the letter or string provided
	fmt.Println("Searched Index is:",wIndex)

	substr := str3[wIndex:12] // Slicing by index and extracting e.g Slice str3 starting at the index of wIndex and ending at index 12
	fmt.Println("Sliced String is:", substr)

	str4 := "Hello, Go, Go loser ranger!"
	fmt.Println(strings.Replace(str4, "Go", "Golang", 1)) // Replace returns a copy of the string s with the first n non-overlapping instances of old text replaced by new text.

	fmt.Println(strings.ToUpper(str2)) // String to uppercase
	fmt.Println(strings.ToLower(str1)) // String to lowercase

	fmt.Println(strings.Contains(str1, "Go")) // Check if variable contains a certain string
}

	



