package main

import "fmt"

func getFullName() (string, string){
	return "bintang", "eko"
}

// func main(){
// 	firstName, lastName := getFullName()
// 	fmt.Println(firstName, lastName)
// }

func main(){
	firstName, _ := getFullName()
	fmt.Println(firstName)
}