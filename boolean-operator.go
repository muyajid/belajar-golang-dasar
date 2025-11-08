package main

import "fmt"

func main() {
	name := "Muhammad Yazid Arsy"
	inpName := "Muhammad Yazid Arsy"

	results := name == inpName
	fmt.Println("name === inpName : ", results)

	name = "Zacky Arya Satya"
	inpName = "Ara Ara"

	results = name != inpName
	fmt.Println("name != inpName : ", results)

	numberA := 10
	numberB := 5

	results = 10 < 5
	fmt.Println("numberA < numberB : ", results)

	results = numberA > numberB
	fmt.Println("numberA > numberB : ", results)

	results = numberA != numberB
	fmt.Println("numberA != numberB : ", results)

	logA := numberA == numberB
	logB := numberA < numberB

	results = logA && logB
	fmt.Println("logA && logB : ", results)

	logB = numberA > numberB
	results = logA || logB
	fmt.Println("logA || logB : ", results)


}