package variables

import (
	"fmt"
	"strconv"
)

// ExampleOfVariable /**
//   - This function is used to show how to use variable in Go
//   - This function does not have any parameter
func ExampleOfVariable() {
	fmt.Println("Variable Example")
}

// ExampleOfVariableWithParameter /**
//   - This function is used to show how to use variable in Go
//   - This function has a parameter
//   - @param name string
func ExampleOfVariableWithParameter(name string) {
	fmt.Printf("Variable Example with parameter is : %s\n", name)
}

// ExampleOfVariableWithMultipleParameters /**
//   - This function is used to show how to use variable in Go
//   - This function has multiple parameters
func ExampleOfVariableWithMultipleParameters(name string, age int, address string) {
	fmt.Printf("Variable Example with parameter is : %s, %d, %s\n", name, age, address)
}

// ExampleOfVariableWithMultipleParametersAndReturnValue /**
//   - This function is used to show how to use variable in Go
//   - This function has multiple parameters
//   - This function has a return value
func ExampleOfVariableWithMultipleParametersAndReturnValue(
	name string,
	age int,
	address string,
) (string, int, string) {
	return name, age, address
}

// PrintsAllVariables /**
// This is another way to create variable in Go
var (
	name1    string
	age1     int
	address1 string
)

// This is another way to create variable in Go
var name2, age2, address2 string

// This is another way to create variable in Go
var name3 string
var age3 int
var address3 string

// This is another way to create variable in Go
var (
	name4    = "ivan"
	age4     = 43
	address4 = "Jakarta"
)

func PrintsAllVariables() {

	// Add value to the variable 1
	name1, age1, address1 = "ivan", 11, "Jakarta"

	// Add value to the variable 2
	name2, age2, address2 = "ivan", strconv.Itoa(22), "Jakarta"

	// Add value to the variable 3
	name3, age3, address3 = "ivan", 33, "Jakarta"

	// Add value to the variable 4
	name4, age4, address4 = "ivan", 44, "Jakarta"

	fmt.Println(name1, age1, address1)
	fmt.Println(name2, age2, address2)
	fmt.Println(name3, age3, address3)
	fmt.Println(name4, age4, address4)
}
