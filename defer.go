package main

import "fmt"

func logging() {
	fmt.Println("Complete calling function")
}

func runApps() {
	defer logging()
	fmt.Println("Application is running")
}

func main() {
	runApps()
}
