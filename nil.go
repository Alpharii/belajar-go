package main

import "fmt"

func main(){
	fmt.Println(NewMap(""))

	data:= NewMap("")
	if data == nil {
		fmt.Println("data kosong")
	} else {
		fmt.Println(data)
	}
}

func NewMap(name string) map[string]string{
	if name == ""{
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}