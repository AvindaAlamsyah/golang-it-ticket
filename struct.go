package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hi", name, "my name is", customer.Name)
}

func main() {
	var avinda Customer
	fmt.Println(avinda)

	avinda.Name = "Avinda Renaldi Alamsyah"
	avinda.Address = "Magetan"
	avinda.Age = 28
	fmt.Println(avinda)
	fmt.Println(avinda.Name)
	fmt.Println(avinda.Address)
	fmt.Println(avinda.Age)

	sebastian := Customer{
		Name:    "Sebastian",
		Address: "Surabaya",
		Age:     30,
	}
	fmt.Println(sebastian)

	praditya := Customer{"Praditya", "Madiun", 22}
	fmt.Println(praditya)

	praditya.sayHello("Marc")
	avinda.sayHello("Marc")
	sebastian.sayHello("Marc")
}
