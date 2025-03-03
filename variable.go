package main

import "fmt"

func main() {
	var text string = "initial"
	fmt.Println(text)

	var number1, number2 int = 1, 2
	fmt.Println(number1, number2)

	var decare int
	fmt.Println(decare)

	name := "Avinda Renaldi"
	fmt.Println(name)

	name = "Avinda Alamsyah"
	fmt.Println(name)

	var (
		firstName  = "Avinda"
		middleName = "Renaldi"
		lastName   = "Alamsyah"
	)

	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}
