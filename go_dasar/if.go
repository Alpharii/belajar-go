package main

import "fmt"

func main(){
	name := "eko"

	if name == "eko"{
		fmt.Println("hello eko")
	} else if name == "joko"{
		fmt.Println("hello joko")
	} else {
		fmt.Println("hi, boleh kenalan?")
	}

	if length := len(name); length > 5{
		fmt.Println("nama kepanjangan")
	} else {
		fmt.Println("nama sudah benar")
	}
}