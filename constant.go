package main

import "fmt"

func main(){
	// const firstName string = "bintang"
	// const lastName = "alphari"

	const (
		firstName string = "bintang"
		lastName = "alphari"
	)

	// firstName = "abc" //tidak bisa di ubah
	// lastName = "error"

	fmt.Println(firstName)
	fmt.Println(lastName)
}