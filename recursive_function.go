package main

import "fmt"

func factorial(value int) int {
	result := 1
	for i := value; i > 0; i--{
		result *= i
	}
	return  result
}

func factorialRecursice(value int) int{
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursice(value-1)
	}
}

func main(){
	fmt.Println(factorial(10))
	fmt.Println(factorialRecursice(10))
}