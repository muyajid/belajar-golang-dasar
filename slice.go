package main

import "fmt"

func main() {
	normalArray := [...]string{"John", "Paul", "Ringgo", "Brian", "Jeff"}
	fmt.Println("Normal Array : ", normalArray)
	fmt.Println("Normal Array Len : ", len(normalArray))
	fmt.Println("Normal Array Cap : ", cap(normalArray))

	// Slice adalah type data yang menyimpan potongan data array 
	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	fmt.Println("days array yang akan di potong : ", days)

	// Membuat slice dari array 
	daysSlice := days[5:] //Memotong array dari index 5 sampai index selanjutnya
	fmt.Println("daysSlice : ", daysSlice)
	fmt.Println("daysSlice len : ", len(daysSlice))
	fmt.Println("daysSlice cap : ", cap(daysSlice)) //len tidak boleh lebih dari cap, cap slice di dapat dari index yang di potong dan index sisanya 
	
	// Slice merupakan referensi dari array sumber nya jadi ketika kita ubah index slice maka array sumber juga terubah
	daysSlice[0] = "Sabtu Baru"
	daysSlice[1] = "Minggu Baru"
	// daysSlice[2] = "Ups" //err
	fmt.Println("daysSlice potongan dari array days : ", daysSlice)
	fmt.Println("days sumber dari slice daysSlice   : ", days)

	// append digunakan untuk tambah data ke dalam slice dengan cara membuat slice baru lagi
	daysSlice2 := append(daysSlice, "Ini Hari Baru")
	fmt.Println("Ini daysSlice2 ref => daysSlice : ",daysSlice2)
	fmt.Println("Ini daysSlice : ", daysSlice)
	fmt.Println("Ini days sumber => daysSlice : ", days)

	// Merubah data dari daysSlice 2 
	daysSlice2[1] = "Upss"
	fmt.Println("Ini daysSlice2 setelah index 2 dirubah : ", daysSlice2)
	fmt.Println("Ini daysSlice : ", daysSlice)

	// Cara membuat slice asli bukan dari array 
	number := []int{1,2,3,4,5,6,7,8,9,10} //Beda utama array dengan slice, slice bisa di append di tamabh datanya sedang array constan datannya
	fmt.Println("Slice asli : ", number)
	// append slice 
	number = append(number, 2)
	fmt.Println("Slice asli setelah di append : ", number)

	// Nah yang ini array ... itu buat hitung otomatis biar nggak [2]string
	huruf := [...]string{"A", "B", "C", "D", "E", "F", "G", "H"}
	fmt.Println("Ini array : ", huruf)
	// err huruf = append(huruf, "2")
} 