package main

import "fmt"

func main() {
	var a = 10
	var b = 10
	var c = a + b
	fmt.Println("Penjumlahan : ", c)

	a = 20 
	b = 5
	c = a - b
	fmt.Println("Pengurangan : ", c)

	a = 5 
	b = 2
	c = a * b
	fmt.Println("Perkalian : ", c)

	a = 10
	b = 2
	c = a / b
	fmt.Println("Pembagian : ", c)

	a = 30
	b = 7
	c = a % b 
	fmt.Println("Modulus : ", c)
}