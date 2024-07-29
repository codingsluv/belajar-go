package main

import "fmt"

func main() {
	// var b int = 15
	// var a int

	// // deklarasi array (pembahasan lebih lanjut pada materi 10_array)
	// angka := [6]int{1, 2, 3, 5}

	// // Cara Pertama
	// /*
	//    1. untuk melakukan perulangan pada Go harus di awali dengan keyword for
	//    2. deklarasi variabel baru (variabel a pada looping pertama tidak sama dengan variabel a diatas)
	//    3. a<10, merupakan kondisi logika yang akan dicek setiap kali perulangan
	//    4. a++ (shorthand dari a + 1), merupakan penambahan nilai ke variabel a di for-loop. Jika setiap kali perulangan menginginkan penambahan dua atau lebih maka digunakan (a + 2), dst. Sebaliknya konsep yang sama pada pengurangan (a--) atau (a - 1).
	//    5. statement, merupakan blok kode yang akan dijalankan.
	// */
	// for a := 0; a < 10; a++ {
	// 	fmt.Printf("Value dari a: %d\n", a)
	// }

	// // Cara Kedua
	// // setelah perulangan diatas selesai maka akan dilanjutkan dengan perulangan dibawah
	// // variabel yang digunakan pada kondisi adalah dua variabel awal diatas
	// // dan penambahan nilai variabel terjadi pada badan perulangan (merupakan alternatif dari cara pertama)
	// // cara ini merupakan while loop pada bahasa pemrograman lain
	// for a < b {
	// 	a++
	// 	fmt.Printf("Value dari a: %d\n", a)
	// }

	// // Cara Ketiga
	// // penggunaan keyword break dan continue
	// // continue akan meng-skip sebuah perulangan
	// // sedangkan break akan menghentikan perulahan bahkan sebelum kondisi sampai false
	// for i := 10; i > 0; i-- {
	// 	if i%2 == 0 {
	// 		fmt.Printf("skip value dari i = %d\n", i)
	// 		continue
	// 	}

	// 	fmt.Printf("value dari i = %d\n", i)

	// 	if i == 5 {
	// 		fmt.Printf("value dari i = %d, perulangan ketiga berhenti!\n", i)
	// 		break
	// 	}
	// }

	// // Cara Keempat
	// // for - range biasa digunakan untuk melakukan perulangan pada data bertipe array
	// // range mengembalikan (return) dua nilai (value) yaitu index (i) dan data dari index tersebut (x)
	// for i, x := range angka {
	// 	fmt.Printf("value dari x = %d di %d\n", x, i)
	// }

	var i, j int

	for i = 2; i < 100; i++ {
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				break
			}
		}
		if j > (i / j) {
			fmt.Printf("%d adalah bilangan prima\n", i)
		}
	}

	var n int = 3
	fmt.Printf("%d! = %d", n, factorial_number(n))
}

func factorial_number(n int) int {
	fmt.Printf("bilangan faktorial %d\n", n)

	if n == 0 {
		return 1
	}

	return n * factorial_number(n-1)
}
