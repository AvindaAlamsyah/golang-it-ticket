package main

import "fmt"

func main() {
	name := "Avinda"

	if name == "Avinda" {
		fmt.Println("Hi, Avinda")
	}

	if name == "Alamsyah" {
		fmt.Println("Hi, Alamsyah")
	} else {
		fmt.Println("Hello, what is your name?")
	}

	if name == "Alamsyah" {
		fmt.Println("Hi, Alamsyah")
	} else if name == "Avinda" {
		fmt.Println("Hi, guys")
	} else {
		fmt.Println("Hello, what is your name?")
	}

	if length := len(name); length > 2 {
		fmt.Println("Its ok")
	}
}
