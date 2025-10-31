package main

import "fmt"

func main(){
	name:= "ekossss"

	switch name {
	case "eko":
		fmt.Println("hi eko")
	case "budi":
		fmt.Println("hi budi")
	default:
		fmt.Println("hi boleh kenalan?")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("nama kepanjangan")
	case false :
		fmt.Println("nama sudah benar")
	}

	length := len(name)
	switch{
	case length > 10:
		fmt.Println("nama terlalu panjang")
	case length > 5:
		fmt.Println("nama lumayan panjang")
	default:
		fmt.Println("nama sudah benar")
	}
}	