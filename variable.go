package main

import "fmt"

func main() {
	// Membuat variable tanpa menginisialisasi type data rumus nya var namaVariable typeDatannya
	var name string

	// Menginisialisasi variable harus sesuai type datanya 
	name = "Muhammad Yazid"
	fmt.Println(name)

	name = "Zaky Aryasatya"
	fmt.Println(name)

	var umur int

	umur = 15
	fmt.Println(umur)

	umur = 16
	fmt.Println(umur)

	var tinggi float32

	tinggi = 1.5
	fmt.Println(tinggi)

	tinggi = 1.7
	fmt.Println(tinggi)

	// Apabila kita membuat sebuah variable dan langsung mennginisialisasi nya maka tidak perlu disebutkan type datanya 

	var namaSiswa = "Zacky arya satya adinata"

	fmt.Println(namaSiswa)

	namaSiswa = "Michael Fabiansky Ambatukam"
	fmt.Println(namaSiswa)

	// namaSiswa = 111
	// fmt.Println(namaSiswa)
	// variable hanya boleh di deklarasi sekali dann untuk menginisialisasi ulang harus dengan type data yang sama

	// Membuat variable tanpa var bisa dengan namavariable := datanya untuk reinisialisasi data langsung pakai =
	namaLaptop := "Asus Vivobook"

	fmt.Println(namaLaptop)

	namaLaptop = "Thinkpad"

	fmt.Println(namaLaptop)

	var (
		firstName = "Muhammad"
		midleName = "Yazid"
		lastName = "Arsy"
	)

	fmt.Println("firstName : ", firstName)
	fmt.Println("midleName : ", midleName)
	fmt.Println("lastName  : ", lastName)


	const ulangTahun = 15012009
	fmt.Println("Ulang Tahun : ", ulangTahun)
}