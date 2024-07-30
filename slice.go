package main

import "fmt"

func main() {
	var makanan = []string{"ayam", "ikan", "bebek"}

	var makananSaya = makanan[0:3]

	fmt.Println(makananSaya)
}
