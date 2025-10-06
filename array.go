package main

import "fmt"

func main(){
	var names [3]string

	names[0]= "eko"
	names[1]= "kurniawan"
	names[2]= "khanedy"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [...]int {1, 2, 3}
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(len(values))
}