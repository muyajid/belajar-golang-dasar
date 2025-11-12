package main

import "fmt"

type Blacklist func(string) bool

func registerAccount(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("Kamu tidak diizinkan register account", name)
	} else {
		fmt.Println("Selamat datang", name)
	}
}

func main() {
	// membuat anonymus function yang tersimpan di variable
	filter := func(name string) bool {
		return name == "Joko"
	}
	// panggil function dengan parameter func
	registerAccount("Joko", filter)

	// membuat anonymus function langsung di param func 
	registerAccount("Bahlil", func(name string) bool {
		return name == "Bahlil"
	})
}