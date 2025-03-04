package main

import "fmt"

func main() {
	name := "Avinda"

	switch name {
	case "Alamsyah":
		fmt.Println("Hi, Alamsyah")
	case "Avinda":
		fmt.Println("Hello, Avinda")
	default:
		fmt.Println("Hi, buddy")
	}

	switch length := len(name); length > 4 {
	case true:
		fmt.Println("Ok, accept")
	case false:
		fmt.Println("Unacceptable")
	}

	switch name {
	case "Alamsyah", "Avinda":
		fmt.Println("Hy, guys")
	default:
		fmt.Println("Hmmm....")
	}
}
