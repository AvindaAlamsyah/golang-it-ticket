package main

import "fmt"

func main() {
	person := map[string]string{
		"name": "Avinda",
		"nick": "d_gato",
	}
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"]) //default value
	fmt.Println(len(person))

	person["address"] = "Magetan"
	fmt.Println(person["address"])
	fmt.Println(person)
	delete(person, "address")
	fmt.Println(person)
}
