package main

import "fmt"

func main() {
	for i := 0; i <= 100; i += 10 {
		fmt.Println("Loop :", i)
	}

	names := [3]string{"Avinda", "Renaldi", "Alamsyah"}
	for idx, val := range names {
		fmt.Println("Index", idx, "=", val)
	}
	for _, val := range names {
		fmt.Println("Name", val)
	}

	for idx, val := range names {
		if idx == 1 {
			break
		}
		fmt.Println("Break :", val)
	}
	for idx, val := range names {
		if idx == 1 {
			continue
		}
		fmt.Println("Continue :", val)
	}
}
