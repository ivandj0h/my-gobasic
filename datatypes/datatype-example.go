package datatypes

import "fmt"

func GetStringDataType() {

	name := "Jane"
	age := 17
	res := fmt.Sprintf("%s is %d years old", name, age)
	fmt.Println(res)

	fmt.Printf("The type of the variable %s is %T\n", name, name)
}

func GetIntDataType() {

	age := 17
	res := fmt.Sprintf("The age of the person is %d", age)
	fmt.Println(res)
	fmt.Printf("The type of the variable %d is %T\n", age, age)
}
