package main

import "fmt"

func broadcast() {
	fmt.Println("Woi info info")
}

func runApps() {
	// defer fungsi non blocking yang memungkinkan code yang berada dibawah di eksekusi terlebih dahulu tanpa harus menunggu code di atas nya selesai dijalankan
	defer broadcast()
	fmt.Println("Applikasi sudah berjalan")
	// blocking tanpa defer broadcast()
}

func main() {
	runApps()
}