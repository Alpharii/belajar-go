package main

import "fmt"

func main(){
	person := map[string]string {
		"name": "bintang",
		"addres": "palembang",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["addres"])

	fmt.Println(person["salah"]) //kosong mengikuti default valuenya

	book := make(map[string]string)
	book["title"] = "Buku golang"
	book["author"] = "eko kurniawa"
	book["wrong"] = "ups"

	delete(book, "wrong")

	fmt.Println(book)
}