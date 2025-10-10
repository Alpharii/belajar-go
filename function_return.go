package main

import "fmt"

func getHello(name string) string{
	return "hello " + name
}

func main(){
	result := getHello("eko")
	fmt.Println(result)

	fmt.Println(getHello("budi"))
	fmt.Println(getHello("joko"))
	fmt.Println(getHello("bintang"))
}