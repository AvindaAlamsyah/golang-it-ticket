package main

import "fmt"

func main() {
	var priceA = 1200
	var priceB = 6000
	var money = 5000

	fmt.Println(priceA < money && priceB < money)
	fmt.Println(priceA < money || priceB < money)
	fmt.Println(!(priceB < money))
}
