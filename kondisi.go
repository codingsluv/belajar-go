package main

import "fmt"

func main() {
	// var point = 8
	// if point == 10 {
	// 	fmt.Println("lulus dengan nilai yg bagus")
	// } else if point > 5 {
	// 	fmt.Println("lulus")
	// } else if point == 4 {
	// 	fmt.Println("hampir lulus")
	// } else {
	// 	fmt.Println("tidak lulus, nilai anda %d\n", point)
	// }

	name := "een"

	if name == "een" {
		fmt.Println("hallo een")
	} else {
		fmt.Println("salah orang")
	}

	// ? short if

	if length := len(name); length > 3 {
		fmt.Println("nama terlalu panjang")
	}
}
