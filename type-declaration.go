package main

import "fmt"

func main() {
	type noKTP string

	var myKTP noKTP = "3200091928852394"

	var another string = "359203994213497"
	var anotherKTP noKTP = noKTP(another)

	fmt.Println(myKTP)
	fmt.Println(anotherKTP)
}
