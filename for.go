package main 

import "fmt"

func main() {

	// for tanpa statement melakukan bedasarkan kondisi true mirip while
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke : ", counter)
		counter++
	}

	fmt.Println("Perulangan tanpa statement selesai")

	// for dengan statement 
	for count := 0; count <= 10; count++ {
		fmt.Println("Perulangan dengan statement ke : ", count)
	}
	fmt.Println("Perulanagn dengan statement selesai")

	//  melakukan iterasi terhadap data colection atau modular 

	// iterasi data collection dengan for biasa 
	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	for i := 0; i < len(days); i++ {
		fmt.Println("Hari : ", days[i])
	}
	// for range mirip dengan foreach index, item := range data 
	names := []string{"Diandra Paramitha Sastrowardoyo", "Mariana Renata", "Anggelina Jolie"}
	for index, item := range names {
		fmt.Println("Index : ", index, "Item", item)
	}

	dailyMenus := map[string]string{
		"Senin": "Nasi Goreng",
		"Selasa": "Mie Goreng",
		"Rabu": "Mie Rebus",
		"Kamis": "Nasi Rendang",
		"Jumat": "Nasi Uduk",
		"Sabtu": "Nasi Ayam Goreng",
		"Minggu": "Nasi Rawon",
	}
	for days ,menus := range dailyMenus {
		fmt.Println("Hari : ", days, "Menus : ", menus)
	}
	fmt.Println("Perulangan untuk iterasi data collection selesai")

	// break dan continue 

	// break menghentikan perulangan
	books := []string{"Rich Dad Poor Dad", "Guerilia Warfare", "The Art Of War"}
	search := "Rich Dad Poor Dad"
	for _ , book := range books {
		if book == search {
			fmt.Println("Buku ditemukan", book)
			break
		}
		fmt.Println("Buku Tidak Ditemukan")
	}
	// continue melanjutkan perulangan 
	for i := 1; i < 10; i++ {
		if i % 2 == 0 {
			continue
		}
		fmt.Println("Ini angka ganjil : ",  i)
	}

	for i := 1; i < 10; i++ {
		if i % 2 != 0 {
			continue
		}
		fmt.Println("Ini angka genap : ", i)
	}

	// fizz buzz

	for i := 0; i < 50; i++ {
		if i % 3 == 0 {
			fmt.Println("Fizz : ", i)
		}

		if i % 5 == 0 {
			fmt.Println("Buzz : ", i)
		}

		if i % 3 != 0 && i % 5 != 0 {
			fmt.Println("Fizz  Buzz : ", i)
		}
	}
}