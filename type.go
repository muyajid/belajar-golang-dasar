// typee adalah kemampuan golang dalam membuat type data baru dari type data yang ada 
// type bisa juga disebut sebagai alias untuk type data yang sudah ada 

// cara membuat type
package main

import "fmt"

func main() {
	type kata string

	var namaKu kata = "Muhammad Yazid Arsy"
	fmt.Println("variable namaKu dengan type data kata : ", namaKu)

	// Konversi string ke type data yang baru yang sudah dibuat
	myPacarStr := "Alaia mirdad"

	var myPacar kata = kata(myPacarStr)
	fmt.Println("Konversi string ke kata", myPacar)

	fmt.Println("Konversi byte ke kata", kata(myPacar[1]))
}