package main

import "fmt"

func random() any{
	return true
}

func main(){
	result := random()
	// resultString := result.(string)
	// fmt.Println(resultString)

	// resultInt := result.(int) //panic
	// fmt.Println(resultInt)

	switch value := result.(type){
	case string:
		fmt.Println("string", value)
	case int:
		fmt.Println("int", value)
	default:
		fmt.Println("unknown", value)
	}
}