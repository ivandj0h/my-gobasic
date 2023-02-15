package functions

import "fmt"

/** Function is a block of code that performs a specific task.
	Functions are used to perform certain actions, and they are also known as methods.
	Functions are used to reuse code: Define the code once, and use it many times.
	You can pass data, known as parameters, into a function.
**/

func Add(x int, y int) {
	fmt.Println("The sum of the two numbers is :", x+y)
}

func AddWithReturnValue(x int, y int) int {
	return x + y
}

func AddWithMultipleReturnValues(x int, y int) (int, int) {
	return x + y, x - y
}

func Swap(x, y string) (string, string) {
	return y, x
}
