package main

import "fmt"

type BlackList func(string) bool

func registerUser(name string, blackList BlackList) {
	if blackList(name) {
		fmt.Println("kamu di blok", name)
	} else {
		fmt.Println("selamat datang", name)
	}
}

func main() {
	// blackList := func(name string) bool {
	// 	return name == "anjing"
	// }
	// registerUser("een", blackList)

	registerUser("een", func(name string) bool {
		return name == "een"
	})
}
