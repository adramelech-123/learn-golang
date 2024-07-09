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

---

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

---

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
---

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
---

### Type conversion

Type conversion in Go is a mechanism to convert a variable from one type to another. This is particularly useful when different types need to be operated on together. 

#### Integer to Float Conversion

To convert an integer to a float, use the syntax `float64(variable)`, where `variable` is the integer you want to convert.

```go
var tempInt int = 10
var tempFloat float64 = float64(tempInt)
fmt.Printf("Integer to float: %.2f \n", tempFloat)
```

In this example:
- `tempInt` is an integer with a value of 10.
- `tempFloat` is the float representation of `tempInt`.

#### Checking Variable Types

The `%T` format specifier in `fmt.Printf` can be used to print the type of a variable.

```go
fmt.Printf("Type of tempInt is %T \n", tempInt)
fmt.Printf("Type of tempFloat is %T \n", tempFloat)
```

#### Integer to String Conversion

There are different ways to convert an integer to a string. Using `fmt.Sprint` and `strconv.Itoa` are two common methods.

##### Using `fmt.Sprint`

`fmt.Sprint` converts the integer to a string.

```go
str := fmt.Sprint(123)
fmt.Printf("Type of variable str is %T \n", str)
```

##### Using `strconv.Itoa`

`strconv.Itoa` is a function from the `strconv` package specifically designed to convert an integer to a string.

```go
str2 := strconv.Itoa(10)
fmt.Printf("Type of variable str2 is %T \n", str2)
```

#### Note

Do not use `string(integer)` for conversion as it converts the integer to its corresponding Unicode character rather than a numeric string.

#### String to Float Conversion

To convert a string to a float, use `strconv.ParseFloat`. This function returns the converted float and an error value which indicates if the conversion was successful.

```go
var myFloatStr string = "3.141549"
var myFloatFromString, _ = strconv.ParseFloat(myFloatStr, 64)
fmt.Println("myFloatFromString:", myFloatFromString)
fmt.Printf("Type of variable myFloatFromString is %T \n", myFloatFromString)
```

In this example:
- `myFloatStr` is a string representation of a float.
- `myFloatFromString` is the float64 representation of the string.


- **Integer to Float**: Use `float64(variable)`.
- **Check Variable Type**: Use `%T` in `fmt.Printf`.
- **Integer to String**: Use `fmt.Sprint` or `strconv.Itoa`.
- **String to Float**: Use `strconv.ParseFloat`.
- **Avoid `string(integer)`**: This converts to a Unicode character, not a numeric string.

---


### Constants

Constants in Go are declared using the `const` keyword. They represent fixed values that cannot be changed during the execution of a program. Constants are evaluated at compile-time and are often used for defining fixed values like mathematical constants, configuration values, or enumerations.


```go
const Pi = 3.14
const Greeting = "Hello, World!"
```

`iota` is a special identifier in Go that simplifies the creation of enumerated constants. It is used within a `const` block and automatically increments starting from 0.

#### Usage of `iota`

In the provided code example, `iota` is used to create a set of related constants for car categories:

```go
func main() {
	const (
		Economy = iota
		Compact
		Standard
		FullSize
		Luxury
	)

	fmt.Println("Economy:", Economy)
	fmt.Println("Compact:", Compact)
	fmt.Println("Standard:", Standard)
	fmt.Println("FullSize:", FullSize)
	fmt.Println("Luxury:", Luxury)
}
```


- `const` block starts the sequence of constants.
- `iota` starts at 0 and increments by 1 for each line within the same `const` block.
- This automatically assigns:
  - `Economy` = 0
  - `Compact` = 1
  - `Standard` = 2
  - `FullSize` = 3
  - `Luxury` = 4

#### Output

```plaintext
Economy: 0
Compact: 1
Standard: 2
FullSize: 3
Luxury: 4
```

#### Advantages of `iota`

- **Readability**: Improves code readability by avoiding manual assignment of incremental values.
- **Maintenance**: Easier to maintain and extend. Adding or removing constants in the middle of the list automatically adjusts the subsequent values.
- **Consistency**: Ensures consistent and predictable values for enumerated constants.

#### Complex Usage

`iota` can be used in more complex scenarios involving bit shifts or arithmetic operations. For example, defining powers of two:

```go
const (
	_ = 1 << (iota * 10)
	KB
	MB
	GB
	TB
	PB
)
```

This example will create constants representing kilobyte, megabyte, gigabyte, etc., with values 1KB = 1024, 1MB = 1048576, and so on.


- **Constants**: Fixed values that do not change during program execution.
- **`iota`**: A special identifier used to create enumerated constants with incrementing values starting from 0.
- **Benefits of `iota`**: Enhances readability, simplifies maintenance, and ensures consistency.


## 2. Getting User Input

Getting user input in Go involves reading data from the standard input (usually the keyboard). The `fmt` package provides functions for this purpose, such as `fmt.Scan` and `fmt.Scanln`.

The `fmt.Scanln` function is commonly used to read a single line of input from the user. This function stops reading input at the newline character.

```go
func main() {
	fmt.Println("What's your name?")
	var name string
	fmt.Scanln(&name) //The &name syntax passes a pointer to the name variable, allowing fmt.Scanln to modify the variable's value directly.

	fmt.Printf("Hello,%s 👻 \n", name)
}
```

## 3. Pointers

Pointers in Golang allow us to work with memory addresses directly, giving us the ability to manipulate the actual data stored in those addresses rather than creating copies of the data. This is especially useful when we need to modify the value of a variable inside a function and have that change reflected outside the function.


```go
package main

import (
	"fmt"
)

func sayHello(greeting string) {
	greeting = "Hello, World"
}

func sayHelloPointer(greeting *string) {
	*greeting = "Hello, World!"
}

func main() {
	// Pointers - These hold the memory address of a variable
	var greeting string = "Hello, GO"

	sayHello(greeting) // Function gets a copy of greeting, the copy is updated to "Hello, World"
	fmt.Println("After calling sayHello:", greeting) // However, the original greeting still remains the same

	sayHelloPointer(&greeting) // The function gets greeting's memory location and updates greeting to "Hello, World!"
	fmt.Println("After calling sayHelloPointer:", greeting) // greeting is now "Hello, World!" because we changed the original greeting not a copy

	// Pointers help us identify variables by their original location and we can manipulate the data in the original variable rather than creating an entirely different copy of the variable
}
```

### Explanation:

1. **Declaring and Assigning Variables:**
   ```go
   var greeting string = "Hello, GO"
   ```
   Here, we declare a variable `greeting` of type `string` and assign it the value "Hello, GO".

2. **Function without Pointers:**
   ```go
   func sayHello(greeting string) {
       greeting = "Hello, World"
   }
   ```
   In this function, `greeting` is passed by value, meaning a copy of `greeting` is created within the function scope. Changes made to `greeting` inside `sayHello` do not affect the original variable.

3. **Calling Function without Pointers:**
   ```go
   sayHello(greeting)
   fmt.Println("After calling sayHello:", greeting)
   ```
   When `sayHello` is called, it modifies the copy of `greeting` to "Hello, World". However, the original `greeting` remains unchanged, so the output is still "Hello, GO".

4. **Function with Pointers:**
   ```go
   func sayHelloPointer(greeting *string) {
       *greeting = "Hello, World!"
   }
   ```
   This function takes a pointer to a string (`*string`). The asterisk (`*`) is used to denote that `greeting` is a pointer. Inside the function, `*greeting` dereferences the pointer, allowing us to update the value at the memory address to which the pointer points.

5. **Calling Function with Pointers:**
   ```go
   sayHelloPointer(&greeting)
   fmt.Println("After calling sayHelloPointer:", greeting)
   ```
   The `&` operator is used to get the memory address of `greeting` and pass it to `sayHelloPointer`. Inside the function, the value at the address is updated to "Hello, World!". This change is reflected outside the function as well, so the output becomes "Hello, World!".

### Zero Values of Variables and Pointers:

- **Zero Value of Variables:**
  - For basic types like `int`, `float64`, `bool`, and `string`, the zero values are `0`, `0.0`, `false`, and `""` (empty string), respectively.
  - Example:
    ```go
    var num int       // Zero value is 0
    var name string   // Zero value is ""
    ```

- **Zero Value of Pointers:**
  - The zero value of a pointer is `nil`, which means it doesn't point to any memory address.
  - Example:
    ```go
    var ptr *int      // Zero value is nil
    if ptr == nil {
        fmt.Println("Pointer is nil")
    }
    ```

In summary, pointers in Golang are used to store memory addresses of variables. This allows functions to modify the actual data rather than working with copies. Understanding the zero values of variables and pointers helps in recognizing uninitialized states and preventing potential runtime errors like dereferencing `nil` pointers.

## 4. Control Structures

This is a basic Rock, Paper, Scissors game using control structures in Golang:

```go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const rounds = 3
	
	fmt.Println("Let's play Rock, Paper, Scissors! You have", rounds, "rounds")

	// Game Rounds
	for i := 0; i < rounds; i++ {
		// Generate Computer's choice
		computerChoiceNum := rand.Intn(rounds)
		var computerChoice string 

		switch computerChoiceNum {
			case 0: computerChoice = "Rock"
			case 1: computerChoice = "Paper"
			case 2: computerChoice = "Scissors"
		}

		//  Read Player's Choice
		var playerChoice string
		fmt.Println("Enter your choice (Rock 🌑,Paper 🧻,Scissors ✂️ )")
		fmt.Scanln(&playerChoice)
		
		fmt.Println("Computer chose", computerChoice)
		switch {
			case playerChoice == computerChoice:
				fmt.Println("It's a tie!")
			case playerChoice == "Rock"	 && computerChoice == "Scissors",
				playerChoice == "Paper"	 && computerChoice == "Rock",
				playerChoice == "Scissors"	 && computerChoice == "Paper" :
				fmt.Println("You win this round!")	
			default:
				fmt.Println("Computer wins this round!")	
		}
	} 
} 

```

## 5. Data Structures (Arrays, Slices, Maps)

### Arrays

Arrays are declared in the following form:

```go
var variableName [arrayLength]varType
```

There are two ways to declare arrays arrays:

```go
func main() {
	// 1. Assign values by index
	var bodyTypes [3]string
	bodyTypes[0] = "Sedan"
	bodyTypes[1] = "SUV"
	bodyTypes[2] = "Convertible"

	fmt.Println(bodyTypes)

	// 2. Declare array and initialize
	colorSchemes := [3]string{"Army Green", "Matte Black", "Velvet"}
	fmt.Println(colorSchemes)
} 
```

### 2D Arrays

2D Arrays are essentially arrays of arrays. Each element in the 2D array can be accessed by two indices (row & column).

```go

func main() {
	// var variableName [rows][cols]varType
	var carFleet [3][2]string

	// We declare each of the 3 rows by index taking note that each row has two elements
	carFleet[0] = [2]string{"Sedan", "Matte Black"}
	carFleet[1] = [2]string{"SUV", "Velvet"}
	carFleet[2] = [2]string{"Convertieble", "Army Green"}

	
	// Print out each row
	for i := 0; i < len(carFleet); i++ {
		row := carFleet[i]
		
		for j := 0; j < len(row); j++ {
			fmt.Printf("%v", row[j])
		}
		fmt.Println()	
	}
} 
```

### Slices

Slices are a more flexible an powerful alternative to arrays. Empty Slices are initialized to `nil` and have a length of `0`.

```go
	fuelTypes := []string{"Electricity", "Gasoline", "Hybrid"}
	// fmt.Println("Length is:",len(fuelTypes))
	fmt.Println("Original Slice:",fuelTypes)

	// Add new element to a slice
	fuelTypes = append(fuelTypes, "Diesel", "Hydrogen")
	fmt.Println("Appended slice:", fuelTypes) //Note: This is a completely new slice
```

We can use the `make()` function to build objects of type slice, map, or chan (only) with a pre-defined length.

```go
fuelTypes := make([]string, 3)
	fmt.Println("Length is: ",len(fuelTypes))

	// Even though the slice is of size 3, we can assign any number of elements so long they dont exceed 3
	fuelTypes[0] = "Electric"
	fuelTypes[1] = "Hybrid"

	fuelTypes = append(fuelTypes, "Diesel", "Hydrogen", "Gasoline") // Since we did not initialize the 3rd element. It will remain empty, then the new elements we append will be added after the empty string on position 2

	fuelTypesCopy := make([]string, len(fuelTypes)) 
	copy(fuelTypesCopy, fuelTypes) // The copy() function creates a copy of a certain piece of data
	fmt.Println(fuelTypesCopy)

```

### Maps

Maps in golang are used for storing a collection of key-value pairs. They are the equivalent of dictionaries in python or objects in javascript.

A map consists of a key and a value and is declared as `map[keyType]valueType` for example:

```go
// We are creating a map of animals and their respective numbers at a zoo
animal := make(map[string]int)

animal["Dog"] = 23
animal["Cat"] = 5
animal["Tiger"] = 6

// Maps can also be declared in one line as follows

animals := map[string]int{
	"Lions":     25,
	"Ostriches": 15,
	"Dolphins" : 10,
}

```
Maps can be accessed with their keys:

```go
numberOfLions := animals["Lions"]
fmt.Printf("The Zoo has %v Lions\n", numberOfLions)

// Since Crocodiles dont exist in the map, the value for animals["Crocodiles"] will be 0
fmt.Printf("The Zoo has %v Crocodiles\n", animals["Crocodiles"])

```

We can delete items from a mapby using the `delete()` function:

```go
delete(animals, "Ostriches")
fmt.Println("Animals: ", animals)
```

Checking the existence of a certain key in a map is easy. If a certain key does not exist in a map, the value initializes to zero and the key:value pair initializes to false.

```go
	numDolphins, dolphinsFound := animals["Dolphins"]
	numCats, catsFound := animals["Cats"]

	fmt.Println("Dolphins found:", dolphinsFound)

	if dolphinsFound {
		fmt.Printf("The zoo has %v Dolphins\n", numDolphins)
	}

	fmt.Println("Cats found:", catsFound)
	if !catsFound {
		fmt.Printf("The zoo has %v Cats\n", numCats)
	}
```

There is a better way of checking the existence of a key in a map by using the following pattern:

```go
	if numLions, ok := animals["Lions"]; ok {
		fmt.Printf("The zoo has %v Lions\n", numLions)
	}
```

The statement `numLions, ok := animals["Lions"]` attempts to retrieve the value associated with the key "Lions" from the animals map. `numLions` will hold the the value associated with the key "Lions". `ok` is a boolean that will be `true` if the key "Lions" exists in the map and `false` otherwise.

The conditional check `if ok {...}` will then follow to check if the key "Lions" was found in the map. If `ok` is `true` then the code in the `if ok {...}` block will be executed.


### for range

For range loops can be useful when working with arrays, slices and maps. for range loops use the following syntax:

`
	for index, item := range variable {
		...
	}
`

If you only want to access the index you can exclude the item as follows:

`
	for index:= range variable {
		...
	}
`

If you intend to only access the items without the index you can exclude the index as follows:

`
	for _, item := range variable {
		...
	}
`

Lets see the example code below:

```go
func main() {
	bodyTypes := [3]string{"Sedan", "SUV", "Convertible"}

	for i, bodyType := range bodyTypes {
		fmt.Printf("Index: %v. Item: %v\n", i, bodyType)
	}

	// If you only want the item, but not the index
	for _, bodyType := range bodyTypes {
		fmt.Printf("Item: %v\n", bodyType)
	}

	// If you only want to access the index and not the items
	for i := range bodyTypes {
		fmt.Printf("Index: %v\n", i)
	}
} 
```

Iterating over a map using for range can be done as follows:

```go
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
```

### Palindrome Exercise

We are going to write a small program to check if a string is a palindrome

```go
func main() {
	// Palindrome exercise

	input := "Madam, I'm Adam"

	// Remove non-letter characters and convert to lowercase
	// Reminder: A rune is a data type that stores codes that represent Unicode characters
	// Unicode is actually the collection of all possible ccharacters present in the whole world
	// In Unicode, each of these characters is assigned a unique number called the Unicode code point
	// This Unicode code point is what we store in a rune data type
	normalized := ""
	for _, r := range input {
		// Check if a rune r represents a letter
		// Note: When we are checking r, we are not checking the actual letter but we are checking the unicode code point representing the letter stored in the rune.
		if unicode.IsLetter(r) {
			normalized += string(unicode.ToLower(r))
		}
	}

	var isPalindrome bool = true

	length := len(normalized)

	for i := 0; i < length/2; i++ {
		if normalized[i] != normalized[length-1-i] {
			isPalindrome = false
			break
		}
	}

	if isPalindrome {
		fmt.Printf("%v is a palindrome\n", input)
	} else {
		fmt.Printf("%v is not a Palindrome\n", input)
	}
} 
```

## 6. Functions

An example of a function that adds cars to the carInventory:

```go
package main

import (
	"fmt"
)

var carInventory = map[string]int{}

func addToInventory(bodyType string, count int) {
	carInventory[bodyType] += count
	fmt.Println(bodyType, "added. New Count:", carInventory[bodyType])
}

func main() {
	addToInventory("Sedan", 10)
	addToInventory("Convertible", 25)
	addToInventory("Sedan", 2)

	fmt.Println(carInventory)
} 

```

We can also get the count of a certain `bodyType` by returning types. In this case the `getCount` function returns an integer

```go
func getCount(bodyType string) int {
	fmt.Printf("Looking up %v count\n", bodyType)
	count := carInventory[bodyType]
	return count //count will be returned as an integer
}

func main() {
	addToInventory("Coupe", 3)
	fmt.Println("Found:", getCount("Coupe"))
} 

```

**Note:** Functions can also return multiple values


### Variadic Functions

A varidic function is function which we can pass as many arguments/parameters as we like.

```go
func sum(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}

	return total
}

func main() {
	fmt.Println("Sum: ", sum(1,2,3,4,5))
	fmt.Println("Sum: ", sum(100, 200, 300))
} 
```

### Anonymous Functions

```go
package main

import (
	"fmt"
	"math"

)

// A function can accept other functions as parameters
func processNumber(num int, operation func(int) int) int {
	return operation(num)
}

// We have two functions double and half which we can input as parameters in the processNumber function
func double(num int) int {
	return num *2
}

func half(num int) int {
	return num /2
}

func main() {
	first_result := processNumber(5, half)
	fmt.Println(first_result)

	//Anonymous function

	result := processNumber(6, func(i int) int {
		return int(math.Pow(float64(i), 3))
	})

	fmt.Println(result)
} 
```

### Closures
 
A function that returns a function. This style of function is called a closure.
 
```go
package main

import (
	"fmt"

)

func sequence() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := sequence()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
}
```

We can run the fibonnacci sequnce using closures as follows:

```go
func fib() func() int {
	a, b := 0, 1

	return func() int {
		a, b = b, a+b
		return a
	}

}

func main() {
	f := fib()

	fmt.Println(f(), f(), f(), f())
}
```

## 7. Structs

We use structs to define a new type. A struct is a prototype that groups variables of different types under a single name.

You declare a struct by using the `type` and `struct` keywords, kind of like how we do it with typescript:

```go
type AuditInfo struct {
	CreatedAt time.Time
	LastModified time.Time
}

type BankAccount struct {
	AccountNumber string
	Balance       float64  
	AuditInfo     // Embedding Audit info struct into BankAccount struct
}

// METHODS TO PERFORM ON STRUCTS


// Value Reciever
// Use when the method does not modify the original struct, If the struct is small an if you want to maintain immutability of the struct
func (bal BankAccount) DisplayBalance() {
	fmt.Printf("Account Number: %s, Balance: $%2f\n", bal.AccountNumber, bal.Balance)
}

//  Pointer Reciever
// Use if, the method modifies the original struct and if the struct is large
func (bal *BankAccount) Deposit(amt float64) {
	bal.Balance += amt
}

// Naive approach with no error handling (Amateur)

// func (bal *BankAccount) Withdraw(amt float64) {
// 	if bal.Balance > amt {
// 		bal.Balance -= amt
// 		fmt.Printf("You have withdrawn $%2f, your current balance is %2f\n", amt, bal.Balance)
// 	} else {
// 		fmt.Printf("Insufficient funds! Your balance is $%2f\n", bal.Balance)
// 	}
// }


// A better approach using error-handling
func (bal *BankAccount) Withdraw(amt float64) error {
	
	if bal.Balance < amt {
		fmt.Printf("Insufficient funds! Your balance is $%2f\n", bal.Balance)
		return errors.New("You have insufficient funds!")
	}

	bal.Balance -= amt
	fmt.Printf("You have withdrawn $%2f, your current balance is %2f\n", amt, bal.Balance)

	return nil
	
}




func main() {
	
	account := BankAccount{
		AccountNumber: "123456789",
		Balance: 1000.00,
		AuditInfo: AuditInfo{
			CreatedAt: time.Now(), 
			LastModified: time.Now(),
		},
	}


	account.DisplayBalance()
	account.Deposit(1000)
	account.DisplayBalance()
	account.Withdraw(1000)
}
```

Here we will build out a Banks with Customers who own bank accounts using structs and applying functions on the different structs.

```go

package main

import (
	"errors"
	"fmt"
	"time"
)

type AuditInfo struct {
	CreatedAt time.Time
	LastModified time.Time
}

type BankAccount struct {
	AccountNumber string
	Balance       float64  
	AuditInfo     
}

// Display Bank Balance
func (bal BankAccount) DisplayBalance() {
	fmt.Printf("Account Number: %s, Balance: $%2f\n", bal.AccountNumber, bal.Balance)
}

// Deposit into bank account
func (bal *BankAccount) Deposit(amt float64) {
	bal.Balance += amt
}

// Withdraw from bank account
func (bal *BankAccount) Withdraw(amt float64) error {
	
	if bal.Balance < amt {
		fmt.Printf("You have Insufficient funds! Your balance is $%2f\n", bal.Balance)
		return errors.New("insufficient funds")
	}

	bal.Balance -= amt
	fmt.Printf("You have withdrawn $%2f, your current balance is %2f\n", amt, bal.Balance)

	return nil
}

// Customer Type

type Customer struct {
	Name      string
	accounts  []BankAccount
	AuditInfo
}

// Add New account to customer
func (c *Customer) AddAccount(account BankAccount) {
	c.accounts = append(c.accounts, account)
}

// Display all acccounts of a customer
func (c Customer) DisplayAccounts() {
	fmt.Printf("Customer: %s\n", c.Name)
	for _, account := range c.accounts {
		account.DisplayBalance()
	}
}

// Bank Type
type Bank struct {
	Name string
	Customers []Customer
} 

// Add Customer to Bank
func (b *Bank) AddCustomer (customer Customer) {
	b.Customers = append(b.Customers, customer)
}

// Display all customers in the bank
func (b Bank) DisplayCustomers() {
	fmt.Printf("Bank: %s\n", b.Name)
	for _, customer := range b.Customers {
		customer.DisplayAccounts()
	}
}

func main() {
	bank1 := Bank{
		Name: "Hidden-Leaf Bank",
	}

	bank2 := Bank {
		Name: " Broke People's Bank",
	}

	customer1 := Customer {
		Name: "Sasuke Uchiha",
		AuditInfo: AuditInfo{CreatedAt: time.Now(), LastModified: time.Now()},
	}

	customer2 := Customer {
		Name: "Naruto Uzumaki",
		AuditInfo: AuditInfo{CreatedAt: time.Now(), LastModified: time.Now()},
	}

	customer3 := Customer {
		Name: "Johnny Bravo",
		AuditInfo: AuditInfo{CreatedAt: time.Now(), LastModified: time.Now()},
	}

	account1 := BankAccount {
		AccountNumber: "178001618844",
		Balance: 3467,
		AuditInfo: AuditInfo{CreatedAt: time.Now(), LastModified: time.Now()},
	}

	account2 := BankAccount {
		AccountNumber: "1780023418756",
		Balance: 5800.93,
		AuditInfo: AuditInfo{CreatedAt: time.Now(), LastModified: time.Now()},
	}

	account3 := BankAccount {
		AccountNumber: "1571059317788",
		Balance: 25.33,
		AuditInfo: AuditInfo{CreatedAt: time.Now(), LastModified: time.Now()},
	}


	// Adding accounts to customers
	customer1.AddAccount(account1)
	customer2.AddAccount(account2)
	customer3.AddAccount(account3)

	// Adding Customers to bank
	bank1.AddCustomer(customer1)
	bank1.AddCustomer(customer2)
	bank2.AddCustomer(customer3)

	bank1.DisplayCustomers()
	bank2.DisplayCustomers()

}
```

## 8. Constructor Functions

In the example code below we will create constructor functions as:

```go
// Constructor Functions

func NewBank(name string) *Bank {
	return &Bank{Name: name}
}

func NewCustomer(name string) *Customer {
	return &Customer{
		Name: name,
		AuditInfo: AuditInfo{CreatedAt: time.Now(), LastModified: time.Now()},
	}
}

func NewBankAccount(acc string) *BankAccount {
	return &BankAccount{
		AccountNumber: acc,
		AuditInfo: AuditInfo{CreatedAt: time.Now(), LastModified: time.Now()},
	}
}
```

Now we can refactor our code in the main func as follows:

```go
func main() {
	bank1 := NewBank("Hidden Leaf Bank")
	bank2 := NewBank("Broke People's Bank")

	customer1 := NewCustomer("Sasuke Uchiha")
	customer2 := NewCustomer("Naruto Uzumaki")
	customer3 := NewCustomer("Johhny Bravo")

	account1 := NewBankAccount("178001618844")
	account1.Deposit(3600)

	account2 := NewBankAccount("1780023418756")
	account2.Deposit(6850)

	account3 := NewBankAccount("1571059317788")
	account3.Deposit(13.5)


	// Adding accounts to customers
	customer1.AddAccount(*account1)
	customer2.AddAccount(*account2)
	customer3.AddAccount(*account3)

	// Adding Customers to bank
	bank1.AddCustomer(*customer1)
	bank1.AddCustomer(*customer2)
	bank2.AddCustomer(*customer3)

	bank1.DisplayCustomers()
	bank2.DisplayCustomers()
}
```
