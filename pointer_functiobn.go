package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeAddressToIndonesia(address *Address){
	address.Country = "indonesia"
}

func main(){
	address := Address{"subang", "jawa barat", ""}
	ChangeAddressToIndonesia(&address)

	fmt.Println(address)
}