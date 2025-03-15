package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func Ups() any {
	//return 1
	//return true
	return "Ups"
}

func main() {
	person := Person{Name: "Avinda Alamsyah"}
	SayHello(person)

	animal := Animal{Name: "Leopard"}
	SayHello(animal)

	var emptyValue any = Ups()
	fmt.Println(emptyValue)
}
