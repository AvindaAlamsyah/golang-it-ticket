package main

import "fmt"

func main() {
	var (
		a = 10
		b = 8
		c = 5
		d = 2
	)

	var e = a + b
	fmt.Println(e)

	var f = e - b
	fmt.Println(f)

	var g = f / c
	fmt.Println(g)

	var h = g * d
	fmt.Println(h)

	h += a
	fmt.Println(h)

	h -= b
	fmt.Println(h)

	h %= d
	fmt.Println(h)

	h++
	fmt.Println(h)

	h--
	fmt.Println(h)
}
