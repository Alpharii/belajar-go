package main

import "fmt"

func main(){
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32) //jika melebihi kapastitas akan terjadi number overflow

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	name := "bintang alphari"
	e := name[0] //uint 8
	eString := string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
}