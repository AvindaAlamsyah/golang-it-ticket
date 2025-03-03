package main

import "fmt"

func main() {
	const (
		firstName string = "Avinda"
		lastName         = "Alamsyah"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

	// Uncomment below will generate error message
	// firstName = "Budi"
	// lastName = "Joko"
}
