package main

import (
	"errors"
	"fmt"
)

func pembagian(nilai, pembagi int) (int, error){
	if pembagi == 0 {
		return 0, errors.New("error pembaginya 0")
	} else {
		return  nilai / pembagi, nil
	}
}

func main(){
	hasil, err := pembagian(100, 10)
	if err != nil{
		fmt.Println("error", err.Error())
	} else {
		fmt.Println(hasil)
	}
}