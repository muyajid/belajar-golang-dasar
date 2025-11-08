package main

import "fmt"

func main() {
	var nilai32 int32 = 32767
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai64)

	fmt.Println("nilai32 : ", nilai32)
	fmt.Println("nilai64 : ", nilai64)
	fmt.Println("nilai16 : ", nilai16)

	negara := "Amerika Serikat"

	negara = string(negara[0])
	fmt.Println("Byte To String : ", negara)
}