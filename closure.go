package main

import "fmt"

func main() {
	// Closure kemampuan sebuah function mengakses data di luar block nya 
	count := 0
	fmt.Println("Data Awal Count : ", count)
	
	increment := func() {
		fmt.Println("Increment")
		count++
	}

	increment()
	increment()
	increment()

	fmt.Println("Count : ", count)
}