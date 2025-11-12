package main

import "fmt"

// Void function tidak mengembalikan apa apa 
func sayHello() {
	fmt.Println("Hello Worlds")
}
// function dengan parameter 
func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}
func sayGoodByeTo(firstName, lastName string) {
	fmt.Println("Goodbye", firstName, lastName)
}
func calculate(numberOne float32, ops string, numberTwo float32) {
	var results float32
	switch ops {
	case "+":
		results = numberOne + numberTwo
		fmt.Println("Hasil = ", results)
	case "-":
		results = numberOne - numberTwo
		fmt.Println("Hasil = ", results)
	case "*":
		results = numberOne * numberTwo
		fmt.Println("Hasil = ", results)
	case "/":
		if numberOne == 0 || numberTwo == 0 {
			fmt.Println("Tidak dapat membagi dengan angka 0")
		} else {
			results = numberOne / numberTwo
			fmt.Println("Hasil = ", results)
		}
	default:
		fmt.Println("Operator tidak valid")
	}
}
// return function function yang mengembalikan suatu nilai 
func getRespect() string {
	return "Info Respect Woi"
}
// multipe returning value function
func getFullName() (string, string) {
	return "Jokowi", "Dodo"
}
func checkAge(age int) (results string) {
	if age < 17 {
		results = "Tidak boleh pacaran"
	} else {
		results = "Boleh kok pacaran asal pacar nya cantik"
	}
	return results
}

func main() {
	// Memanggil function 
	sayHello()
	sayHello()
	sayHello()
	
	fmt.Println("Function sayHello() selesai di jalankan")
	
	// Memanggil dan memasukan argument ke dalam function parameter 
	sayHelloTo("Yazid", "Arsy")
	sayHelloTo("Budi", "Nugraha")
	sayHelloTo("Elon", "Musk")
	sayGoodByeTo("Raden", "Adika")
	sayGoodByeTo("Firdaus", "Naysaburi")
	
	fmt.Println("Function sayHelloTo() sayGoodByeTo() selesai di jalankan")

	// Memanggil return function menyimpan nilai nya ke variable dan print variable yang berisi value itu 

	rispect := getRespect()
	fmt.Println("rispect : ", rispect)
	fmt.Println("rispect() : ", getRespect())

	firstName, lastName := getFullName()
	fmt.Println("First Name : ", firstName)
	fmt.Println("Last Name : ", lastName)

	firstNameTwo, _ := getFullName()
	fmt.Println("First Name : ", firstNameTwo)

	fmt.Println(checkAge(9))
	calculate(10, "-", 5)
}