# Golang Tutorial

## 1. Go Basics

```go
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var pl = fmt.Println

func main() {
	pl("What is your name?")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')

	if err == nil {
		pl("Hello", name)
	} else {
		log.Fatal(err)
	}
}
```

The first thing we need to do is define our package as `package main` and then import the individual libraries we want to use e.g. `fmt`, `bufio`, `log`, `os` etc.

### Variables

```go
func main() {
	var name string = "Mr Sakamoto"
	fmt.Printf("Welcome, %v\n", name)

	var totalCars int = 50
	fmt.Println("Your fleet consists of", totalCars, "cars 🚗")

	startingPrice := 29.99
	fmt.Printf("Our prices start at $%v\n", startingPrice)

	var insuranceIncluded bool = false
	fmt.Printf("Do our Prices include insurance? : %t", insuranceIncluded)
}
```

Variables can be explicitly be defined in the `var variableName variableType = assignedValue` format. In the case where we want our variables to be defined implicitly we use the `:=`.

### String methods

Various string methods can be applied to variables of type string. These methods can be accessed using the [strings](https://pkg.go.dev/strings) package which you'll have to import in the `main.go` file. Examples of string methods are shown below:

```go
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

	// EqualFold() checks if two strings are equal. Note: The EqualFold method from the strings library is case insensitive
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
```

### Math Operations

We can use the [math]() library to access math functions in golang. Some math functions can be written as follows:

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	
	var temperatureC float64 = 12.15
	fmt.Println("Temperature:", temperatureC)
	fmt.Println("Temperature(Round):", math.Round(temperatureC)) // Round up or down to nearest whole number
	fmt.Println("Temperature(Ceil):", math.Ceil(temperatureC)) // Round Up
	fmt.Println("Temperature(Floor):", math.Floor(temperatureC)) // Round Down

	fmt.Println("Absolute Temperature:", math.Abs(-5.5)) // Absolute Value
	fmt.Println("Power:", math.Pow(3, 2)) // To the power of
	fmt.Println("Square Root:", math.Sqrt(16)) // Square Root
}

```

### Type conversion

