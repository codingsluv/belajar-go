package main

import "fmt"

type Person struct {
	name, address string
	age           int
}

func (person Person) sayHello() {
	fmt.Println("Hello, my name is", person.name)
	fmt.Println("I'm from", person.address)
	fmt.Println("I'm", person.age, "years old")
	fmt.Println()
}

func main() {
	// een := Person{
	// 	name:    "een",
	// 	address: "Indonesia",
	// 	age:     21,
	// }
	// fmt.Println(een)
	een := Person{
		name:    "een",
		address: "Indonesia",
		age:     21,
	}
	een.sayHello()
}
