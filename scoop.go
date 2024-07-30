package main

import "fmt"

var angka int

// func main() {
// 	var a, b int
// 	a = 30
// 	b = 10
// 	angka = a + b

// 	fmt.Println("total penjumlah a dan b adalah", angka)
// }

func main() {
	var a int = 10
	var b int = 5
	var c int

	c = tambah(a, b)
	fmt.Println("hasilnya adalah:", c)
}

func tambah(a, b int) int {
	return a + b
}
