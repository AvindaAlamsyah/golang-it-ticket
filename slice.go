package main

import (
	"fmt"
)

func main() {
	months := [...]string{"Jan", "Feb", "Mar", "Apr", "Mei", "Jun", "Jul", "Aug"}
	slice1 := months[4:7]

	fmt.Println(months)
	fmt.Println(slice1)

	slice2 := months[5:]
	slice3 := months[:2]
	slice4 := months[:]

	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)

	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	fullMonthSlice := months[6:]
	fmt.Println(fullMonthSlice)
	fullMonthSlice[0] = "July"
	fullMonthSlice[1] = "August"
	fmt.Println(months)

	fullMonthSlice2 := append(fullMonthSlice, "September")
	fmt.Println(fullMonthSlice2)
	fmt.Println(months)

	fullMonthSlice2[1] = "Aug"
	fmt.Println(fullMonthSlice2)
	fmt.Println(months)

	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "Avinda"
	newSlice[1] = "Renaldi"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Alamsyah")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Orarion"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	fromSlice := months[:]
	var toSlice []string = make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)
}
