package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main(){
	// addres1:= Addres{"subang", "jawa barat", "indonesia"}
	// addres2 := &addres1

	// addres2.City = "bandung"
	// fmt.Println(addres1)
	// fmt.Println(addres2)

	var addres1 Address = Address{"subang", "jawa barat", "indonesia"}
	var addres2 *Address = &addres1

	addres2.City = "bandung"
	fmt.Println(addres1)
	fmt.Println(addres2)
	
	// addres2 = &Address{"jakarta", "dki jakarta", "indonesia"}
	*addres2 = Address{"jakarta", "dki jakarta", "indonesia"}
	fmt.Println(addres1)
	fmt.Println(addres2)

}	