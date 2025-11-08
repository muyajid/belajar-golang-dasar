package main 

import "fmt"

func main() {

	// Map merupakan type data modular berbasis key and value, cara membuat map 

	// namavariable : map[type data untuk key]typedatauntukvalue{key:value}
	band := map[string]string{
		"vocalist": "Sting",
		"gitaris": "pay burman",
		"bass": "yuke sampurna",
		"drummer": "wong aksan",
		"kibor": "ahmad dhani",
	}
	fmt.Println("band : ", band)
	
	absen := map[int]string {
		1: "Aldhesta",
		2: "Aldiky",
		3: "Alfian",
	}
	fmt.Println("absen : ", absen)

	antrian := map[string]int {
		"aurel": 1,
		"zaky": 2,
		"osmar": 3,
	}
	fmt.Println("antrian : ", antrian)

	// Acces value dari map 
	fmt.Println("Nama Vocalist Band : ", band["vocalist"])
	fmt.Println("Absen Siswa Nomor Urut 2 : ", absen[2])
	fmt.Println("Si Aurel Antrian Ke Berapa Ya : ", antrian["aurel"])

	// mengubah data di map
	antrian["osmar"] = 100
	fmt.Println("antrian dengan key osmar : ", antrian["osmar"], " semua antrian : ",antrian )
	band["vocalist"] = "Diana Rose"
	fmt.Println("vocalist diganti : ", band)

	// Membuat map manual 
	menteri := make(map[string]string)
	menteri["keuangan"] = "Purbaya Yudhi Sadewa"
	menteri["pendidikan"] = "Anies Baswedan"

	fmt.Println("menteri : ", menteri)

	// map function
	fmt.Println("Panjang data antrian : ", len(antrian))

	// function delete data 
	delete(antrian, "osmar")
	fmt.Println("key osmar sudah dihapus : ", antrian)
 }