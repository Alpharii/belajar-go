package main

import "fmt"


type HasName interface{
	GetName()	string
}

type Person struct{
	Name	string
}

type Animal struct{
	Name	string
}

func (person Person) GetName() string{
	return person.Name
}

func (animal Animal) GetName() string{
	return  animal.Name
}

func sayHello(hasName HasName){
	fmt.Println("hello", hasName.GetName())
}

func main(){
	person := Person{Name: "eko"}
	sayHello(person)

	animal := Animal{Name: "kucing"}
	sayHello(animal)
}	