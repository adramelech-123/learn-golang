package main

import (
	"fmt"
)


func main() {
	var tempInt int = 10
	var tempFloat float64 = float64(tempInt)
	fmt.Printf("Integer to float: %.2f \n", tempFloat)

	fmt.Printf("Type of tempInt is %T \n", tempInt) // Check type of a variable we use the %T
	fmt.Printf("Type of tempFloat is %T \n", tempFloat)

	str := fmt.Sprint(123) // Convert Integer to string. Note: We dont use string(123) to convert to string, if we do this we get a rune type instead
	fmt.Printf("Type of variable str is %T \n", str)
}


	



