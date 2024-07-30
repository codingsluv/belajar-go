package main

import "fmt"

func main() {
	// membuat array dengan jumlah elemen = 10
	var n [10]int
	var i, j int

	// inisialiasi elemen array n ke 0
	for i = 0; i < 10; i++ {
		// set elemen pada lokasi i = i + 3
		// jadi memulai pada elemen 0 yaitu 3
		// 3,4,5,6,7,8,9,10....,17
		// sampai elemen ke 15
		n[i] = i + 3
	}

	// menampilkan hasil elemen pada array
	for j = 0; j < 10; j++ {
		fmt.Printf("elemen[%d] = %d\n", j, n[j])
	}
}
