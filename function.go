package main

import "fmt"

func sayHello() {
	fmt.Println("Say hello")
}

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello,", firstName, lastName)
}

func sumOperation(firstNumber int, secondNumber int) int {
	return firstNumber + secondNumber
}

func multipleOperation(firstNumber int, secondNumber int) (string, int) {
	result := firstNumber * secondNumber

	return "multiple result = ", result
}

func completeName() (firstName, middleName, lastName string) {
	firstName = "Avinda"
	middleName = "Renaldi"
	lastName = "Alamsyah"

	return firstName, middleName, lastName
}

func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func getGoodBye(name string) string {
	return "Good bye " + name
}

func sayHelloWithFilter(name string, filter func(string) string) {
	filteredName := filter(name)

	fmt.Println("Hi,", filteredName)
}

func spamFilter(name string) string {
	if name == "Noob" {
		return "..."
	} else {
		return name
	}
}

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func factorial(value int) int {
	if value == 1 {
		return value
	} else {
		return value * factorial(value-1)
	}
}

func main() {
	sayHello()

	sayHelloTo("Avinda", "Alamsyah")

	fmt.Println(sumOperation(12, 32))

	text, result := multipleOperation(10, 3)
	fmt.Println(text, result)

	_, onlyResult := multipleOperation(4, 2)
	fmt.Println(onlyResult)

	firstName, middleName, lastName := completeName()

	fmt.Println(firstName, middleName, lastName)

	fmt.Println(sumAll(1, 2, 3, 634, 3))

	goodBye := getGoodBye
	fmt.Println(goodBye("Avinda Alamsyah"))

	sayHelloWithFilter("Avinda", spamFilter)
	sayHelloWithFilter("Noob", spamFilter)

	blacklist := func(name string) bool {
		return name == "Feeder"
	}
	registerUser("Avinda", blacklist)
	registerUser("Noob", func(name string) bool {
		return name == "Noob"
	})

	fmt.Println(factorial(6))
}
