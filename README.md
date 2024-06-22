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