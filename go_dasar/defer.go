package main

import "fmt"

func loging(){
	fmt.Println("loging")
}

func runApplication(){
	defer loging()
	fmt.Println("run application")
}

func main(){
	runApplication()
}