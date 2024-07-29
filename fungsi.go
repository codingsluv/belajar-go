package main

import (
	"fmt"
)

// func main() {
// 	var names = []string{"jhon", "wick"}
// 	printMessage("hallo", names)
// }

// func printMessage(message string, arr []string) {
// 	var nameString = strings.Join(arr, " ")
// 	fmt.Println(message, nameString)
// }

// ? fungsi dengan return value / nilai balik
// var randomizer = rand.New(rand.NewSource(time.Now().Unix()))

// func main() {
// 	var randomValue int
// 	randomValue = randomWithRange(2, 10)
// 	fmt.Println("random number:", randomValue)

// 	randomValue = randomWithRange(2, 10)
// 	fmt.Println("random number:", randomValue)

// 	randomValue = randomWithRange(2, 10)
// 	fmt.Println("random number:", randomValue)
// }

// func randomWithRange(min, max int) int {
// 	var value = randomizer.Int()%(max-min+1) + min
// 	return value
// }

func main() {
	var a int = 10
	var b int = 5
	var hasil int

	hasil = hitung(a, b)
	fmt.Println("hasilnya adalah:", hasil)
}

func hitung(angka1, angka2 int) int {
	var hasil int
	hasil = angka1 + angka2
	return hasil
}
