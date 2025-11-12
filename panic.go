package main

import "fmt"

func sendData() {
	fmt.Println("Hello worlds from sendData() func")
	message := recover()
	fmt.Println("Telah terjadi panic pada program : ", message)
}

func runApps(err bool) {
	defer sendData()
	if err {
		panic("Upss program berhenti")
	}
	fmt.Println("Applikasi berjalan")
}

func main() {
	runApps(true)
}