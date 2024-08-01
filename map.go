package main

import "fmt"

func main() {
	// m := make(map[string]int)

	// m["k1"] = 7
	// m["k2"] = 13

	// fmt.Println("map : ", m)

	person := map[string]string{
		"nama":   "een",
		"alamat": "tempat tidur",
		"kata2":  "belajar golang menyenangkan, paham!",
	}

	delete(person, "kata2")

	fmt.Println(person)
	fmt.Println("Data nama : ", person["nama"])
	fmt.Println("Data alamat : ", person["alamat"])
	fmt.Println("Data kata2 : ", person["kata2"])

	book := map[string]string{
		"title":  "Paham golang paham",
		"author": "een sudrajat",
		"ups":    "salah paham",
	}
	delete(book, "ups")

	fmt.Println(book)
	fmt.Println("title:", book["title"])
	fmt.Println("author:", book["author"])

}
