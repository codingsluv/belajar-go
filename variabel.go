package main

import (
	"fmt"
)

func main() {
	// // ? Variabel tipe statis
	// var angka float32
	// angka = 5.7

	// fmt.Println(angka)
	// fmt.Printf("tipe variabela angka adalah %T\n", angka) // float32

	// // ? Variabel tipe dinamis

	// // * membuat variabel dengan tipe data float32
	// var angka_saya = 29.9
	// // * deklarasi tipe data dengan operator walrus
	// angka_lain := 32

	// fmt.Println(angka_saya)
	// fmt.Println(angka_lain)

	// //* melihat tipe data dari
	// //* kedua variabel
	// fmt.Printf("tipe data angka_saya adalah %T\n", angka_saya)
	// fmt.Printf("tipe data angka_lain adalah %T\n", angka_lain)

	// // ? deklarasi variabel tipe campuran
	// var a, b, c = 4, 7, "een"
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)

	// fmt.Printf("a adalah tipe %T\n", a)
	// fmt.Printf("b adalah tipe %T\n", b)
	// fmt.Printf("c adalah tipe %T\n", c)

	// name := "een"
	// fmt.Println(name)
	// umur := 21
	// fmt.Println(umur)

	const (
		text = "belajar golang menyenangkan"
		jam  = 22.50
	)

	fmt.Println(text, jam)
}
