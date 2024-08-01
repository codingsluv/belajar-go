package main

import (
	"fmt"
	"math"
)

// type HasName interface {
// 	GetName() string
// }

// func SayHello(value HasName) {
// 	fmt.Println("Hello", value.GetName())
// }

// type Person struct {
// 	Name string
// }

// func (person Person) GetName() string {
// 	return person.Name
// }

// type Animal struct {
// 	Name string
// }

// func (animal Animal) GetName() string {
// 	return animal.Name
// }

// func main() {
// 	person := Person{Name: "Een"}
// 	SayHello(person)

// 	animal := Animal{Name: "Kucing"}
// 	SayHello(animal)
// }

type hitung interface {
	luas() float64
	keliling() float64
}

type lingkaran struct {
	diameter float64
}

func (l lingkaran) jariJari() float64 {
	return l.diameter / 2
}

func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jariJari(), 2)
}

func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

type persegi struct {
	sisi float64
}

func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (p persegi) keliling() float64 {
	return p.sisi * 4
}

func main() {
	var bangunDatar hitung

	bangunDatar = persegi{10.0}
	fmt.Println("===== persegi")
	fmt.Println("luas      :", bangunDatar.luas())
	fmt.Println("keliling  :", bangunDatar.keliling())

	bangunDatar = lingkaran{14.0}
	fmt.Println("===== lingkaran")
	fmt.Println("luas      :", bangunDatar.luas())
	fmt.Println("keliling  :", bangunDatar.keliling())
	fmt.Println("jari-jari :", bangunDatar.(lingkaran).jariJari())
}
