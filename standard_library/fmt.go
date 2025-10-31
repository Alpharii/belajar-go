package main

import "fmt"

// Fungsi	Keterangan
// fmt.Print()	Menulis teks ke stdout tanpa newline
// fmt.Println()	Menulis teks dengan newline
// fmt.Printf()	Menulis dengan format string
// fmt.Sprintf()	Mengembalikan string terformat tanpa menulis ke stdout
// fmt.Scan() / Scanln()	Membaca input dari stdin
// fmt.Errorf()	Membuat error dengan format string

func main(){
	fisrtName := "bintang"
	age := 18
	fmt.Printf("nama: %s, umur: %d\n", fisrtName, age)

	lastName := "alphari"
	fmt.Printf("hello %s %s\n", fisrtName, lastName)
}