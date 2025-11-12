package main

import "fmt"

// membuat function denagn parameter array bukan variadic parameter
func printArr(arr [7]string) {
	for i := 0; i < len(arr); i++ {
		fmt.Println("Iteration : ", arr[i])
	}
}
// membuat function dengan variadic parameter 
// variadic parameter memungkinkan parameter untuk menerima data lebih dari satu, tanpa harus membuat parameter array
func sum(number ...int) {
	total := 0
	for _, numbers := range number {
		total += numbers
	}
	fmt.Println("SUM : ", total)
}
func mean(number ...int) {
	total := 0
	for _, numbers := range number {
		total += numbers
	}
	total = total / len(number)
	fmt.Println("MEAN : ", total)
}
func main() {
	presiden := [...]string{"Soekarno", "Soeharto", "Habibie", "Gus Dur", "Megawati", "SBY", "Jokowi"}
	printArr(presiden)
	sum(10, 10, 10)
	mean(4, 5, 10, 10, 8, 10)

	sales := []int{10, 30, 20, 10, 30, 20, 10, 15}
	mean(sales...)
}