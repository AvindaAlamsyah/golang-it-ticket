package main

import "fmt"

func main() {
	var (
		firstName = "Avinda"
		lastName  = "Alamsyah"
		number1   = 23
		number2   = 6
	)

	fmt.Println(firstName == lastName)
	fmt.Println(firstName != lastName)
	fmt.Println(number1 > number2)
	fmt.Println(number1 < number2)
}
