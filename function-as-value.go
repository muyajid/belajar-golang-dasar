package main 

import "fmt"

func sayHello() {
	fmt.Println("Hello Guys")
}
func goodBye() {
	fmt.Println("Good Bye")
}

func main() {
	// Menjadikan function sayHello menjadi value dari hello dan assalamualaikum
	var hello func()
	hello = sayHello
	assalamualaikum := sayHello
	hello()
	assalamualaikum()

	var bye = goodBye
	waalaikumsallam := goodBye
	bye()
	waalaikumsallam() 
}