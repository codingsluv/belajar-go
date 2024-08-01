package main

import "fmt"

func main() {
	names := [...]string{"budi", "udiin", "wawan", "septo", "ucup"}
	slices := names[3:]
	slices[0] = "wawan baru"
	slices[1] = "ucup baru"

	fmt.Println(names)
	fmt.Println("Slices : ", slices)

	dataBaru := append(slices, "nama baru")
	fmt.Println("Data Baru : ", dataBaru)

	newSlices := make([]string, 2, 5)
	newSlices[0] = "een"
	newSlices[1] = "oke"
	// newSlices[2] = "mantap" || ini akan error, jika ingin menambah data kedalam array harus menggunakan append()

	fmt.Println("New Slices : ", newSlices)
	fmt.Println(len(newSlices))
	fmt.Println(cap(newSlices))

	slicesBaru := append(newSlices, "mantap")
	fmt.Println("Slices Baru : ", slicesBaru)

	fromSlice := newSlices[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)

	fmt.Println("From Slice : ", fromSlice)
	fmt.Println("To Slice : ", toSlice)
}
