package main

import "fmt"

type Customer struct {
	Name			string
	Address			string
	Age				int
}

func (customer Customer) sayHello(name string){
	fmt.Println("hello", name, "my name is", customer.Name)
}

func main(){
	var eko Customer
	eko.Name = "eko kurniawan"
	eko.Address = "indonesia"
	eko.Age =  30

	fmt.Println(eko)
	fmt.Println(eko.Name)
	fmt.Println(eko.Address)
	fmt.Println(eko.Age)

	joko := Customer{
		Name: "joko",
		Address: "indonesia",
		Age: 20,
	}
	fmt.Println(joko)

	budi := Customer{"budi", "indonesia", 25}
	fmt.Println(budi)

	budi.sayHello("agus")
	eko.sayHello("joko")
	joko.sayHello("budi")
}