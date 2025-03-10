package main

import "fmt"

func endApps() {
	fmt.Println("Application is stopped")
	message := recover()
	fmt.Println("Panic is invoke:", message)
}

func runApps(error bool) {
	defer endApps()
	if error {
		panic("Something is not right")
	}
}

func main() {
	runApps(true)
	fmt.Println("Avinda Alamsyah")
}
