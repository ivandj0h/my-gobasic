package main

import (
	"fmt"
	"github.com/ivandi1980/my-basics/arrays"
	"github.com/ivandi1980/my-basics/variables"
)

func main() {

	/**
	 * Variable
	 * below is the example of the variable
	 */

	// Print the string "Hello, World!" to the console
	fmt.Println("Hello, World!")

	// Call the function "ExampleOfVariable" from the package "variables"
	variables.ExampleOfVariable()

	// Call the function "ExampleOfVariableWithParameter" from the package "variables"
	variables.ExampleOfVariableWithParameter("ivan")

	// Call the function "ExampleOfVariableWithMultipleParameters" from the package "variables"
	variables.ExampleOfVariableWithMultipleParameters("ivan", 43, "Jakarta")

	// Call the function "ExampleOfVariableWithMultipleParametersAndReturnValue" from the package "variables"
	name, age, address := variables.ExampleOfVariableWithMultipleParametersAndReturnValue(
		"ivan",
		43,
		"Jakarta",
	)
	fmt.Printf("Variable Example with parameter is : %s, %d, %s\n", name, age, address)

	// Call the function "PrintsAllVariables" from the package "variables"
	variables.PrintsAllVariables()

	/**
	 * Array
	 * below is the example of the Array
	 */

	// Call the function "ExampleOfArray" from the package "arrays"
	arrays.ExampleArray()

	// Call the function "getTheLengthOfArray" from the package "arrays"
	arrays.GetTheLengthOfArray()
}
