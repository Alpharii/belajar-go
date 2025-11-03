package main

import (
	"fmt"
	"strconv"
)

// strconv.parseBool(string)			//Mengubah string ke bool
// strconv.parseFloat(string)			//Mengubah string ke float
// strconv.parseInt(string)			//Mengubah string ke int64
// strconv.FormatBool(bool)			//Mengubah bool ke string
// strconv.FormatFloat(float, … )		//Mengubah float64 ke string
// strconv.FormatInt(int, … )			//Mengubah int64 ke string


func main(){
	boolean, err := strconv.ParseBool("true")
	if err == nil{
		fmt.Println(boolean)
	} else {
		fmt.Println(err)
	}

	intvar := strconv.FormatInt(100, 10)
	fmt.Printf("intvar %T\n", intvar)
}