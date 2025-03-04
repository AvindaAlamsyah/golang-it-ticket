package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Avinda"
	names[1] = "Renaldi"
	names[2] = "Alamsyah"

	fmt.Println(names[1])

	names[2] = "test"
	fmt.Println(names[2])

	values := [4]int{1, 2, 3}
	fmt.Println(values)
	fmt.Println(len(values))

	numbers := [...]int{1, 5, 2}
	fmt.Println(numbers)
}
