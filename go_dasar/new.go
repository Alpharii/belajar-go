package main

import "fmt"

type Addres struct {
	City, Province, Country string
}

func main(){
	// addres1:= Addres{"subang", "jawa barat", "indonesia"}
	// addres2 := &addres1

	// addres2.City = "bandung"
	// fmt.Println(addres1)
	// fmt.Println(addres2)

	var addres1 = new(Addres)
	var addres2 = addres1

	addres2.City = "bandung"
	fmt.Println(addres1)
	fmt.Println(addres2)
}
