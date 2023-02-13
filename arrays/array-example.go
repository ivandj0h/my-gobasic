package arrays

import "fmt"

func ExampleArray() {

	// Shorthand declaration of array
	arr := [4]string{"one", "two", "three", "four"}
	fmt.Println("This is Example of the Array :")
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

func GetTheLengthOfArray() {

	// Shorthand declaration of array
	arrOne := [4]string{"one", "two", "three", "four"}
	arrTwo := [...]string{"one", "two", "three", "four", "five", "six"}
	fmt.Println("The length of the array One is :", len(arrOne))
	fmt.Println("The length of the array Two is :", len(arrTwo))
}
