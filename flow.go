package main 

import "fmt"

func main() {
	name := "sastro"

	if name == "sastro" {
		fmt.Println("Halo ", name)
	} else {
		fmt.Println("Halo boleh kenalan nggak?")
	}

	// if else if else 

	umur := 10
	if umur >= 17 {
		fmt.Println("Kamu sudah 17 tahun")
	} else if umur >= 15 {
		fmt.Println("Sebentar lagi kamu 17 tahun")
	} else {
		fmt.Println("Kamu masih anak kecil")
	}

	// if short statement 
	if username := "fea"; len(username) < 5 {
		fmt.Println("Username kamu kurang 5 kata")
	} else {
		fmt.Println("Oke jumlah katanya udah pas")
	}

	if alamat := "kemang"; alamat != "malang" {
		fmt.Println("Salah Alamat Nih")
	} else {
		fmt.Println("Oke alamat nya udah bener")
	}

	// switch 
	hari := "Upss"
	switch hari {
	case "senin":
		fmt.Println("Hari senin")
	case "selasa":
		fmt.Println("Hari selasa")
	case "rabu":
		fmt.Println("Hari rabu")
	case "kamis":
		fmt.Println("Hari kamis")
	case "jumat":
		fmt.Println("Hari Jumat")
	case "sabtu":
		fmt.Println("Hari sabtu")
	case "minggu":
		fmt.Println("Hari minggu")
	default:
		fmt.Println("Ini hari apa ini")
	}
}