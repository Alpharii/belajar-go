package main

import "fmt"

// +
// -
// /
// *
// %

func main(){
	a := 10
	b := 20
	c := a + b
	fmt.Println(c)

	//augmented assigment
	i := 10
	i += 10 // i = i + 10
	fmt.Println(i)

	i += 5 // i = i + 5
	fmt.Println(i)


	//unary operator
	j := 1
	j++ // j = j+1
	fmt.Println(j)
	j++
	fmt.Println(j)

	j-- // j = j-1
	fmt.Println(j)
	j--
	fmt.Println(j)
}