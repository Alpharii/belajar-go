package main

import "fmt"

func getGoodBye(name string) string {
	return  "good bye " + name
}

func main(){
	contoh := getGoodBye
	fmt.Println(contoh("eko"))
}