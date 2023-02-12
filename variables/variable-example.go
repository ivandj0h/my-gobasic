package variables

import "fmt"

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
