package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Diandra paramitha sastrowardoyo"
	names[1] = "Mariana Renata"2
	names[2] = "Alaia Al Mirdad"

	fmt.Println("Print seluruh data array : ", names)
	fmt.Println("Print array index ke 0 : ", names[0])
	fmt.Println("Print array index ke 1 : ", names[1])
	fmt.Println("Print array indes ke 2 : ", names[2])

	arr := [3]string{
		"Nasi Goreng",
		"Mie Goreng",
		"Mie Rebus",
	}
	fmt.Println("arr : ", arr)

	values := [...]string{"Soekarno", "Soeharto", "BJ Habibie", "Gus Dur", "Megawati"}
	fmt.Println("values : ", values)
}