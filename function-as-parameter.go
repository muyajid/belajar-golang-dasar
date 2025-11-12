package main 

import "fmt"

type Filter func(string) string
func sayHello(name string, filter Filter) {
	fmt.Println("Hello", filter(name))
}
func filterName(name string) string {
	if name == "Anjing" || name == "Mulyono" {
		return "...Nggak Sopan"
	}
	return name
}
func main() {
	sayHello("Nicholas", filterName)
	
	f := filterName
	sayHello("Anjing", f)
}