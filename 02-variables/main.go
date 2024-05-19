package main

import "fmt"

const LoginToken string = "ahdjgegf" //Public

func main() {
	var username string = "bintang"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	//Generally if you are working with integers you should just use the int type
	//uint should generally only be used for doing binary operations
	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)
	//Go has the following architecture-independent integer types:
	//uint8       unsigned  8-bit integers (0 to 255)
	//uint16      unsigned 16-bit integers (0 to 65535)
	//uint32      unsigned 32-bit integers (0 to 4294967295)
	//uint64      unsigned 64-bit integers (0 to 18446744073709551615)
	//int8        signed  8-bit integers (-128 to 127)
	//int16       signed 16-bit integers (-32768 to 32767)
	//int32       signed 32-bit integers (-2147483648 to 2147483647)
	//int64       signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
	//There are also a couple of alias number types, which assign useful names to specific data types:
	//byte        alias for uint8
	//rune        alias for int32

	var floatVal float64 = 255.2293013092389
	fmt.Println(floatVal)
	fmt.Printf("Variable is of type: %T \n", floatVal)
	//float is used in decimal number
	//float32 (32 bits)
	//float64 (64 bits) more accurate decimal number

	var coolNumber int = 7
	fmt.Println(coolNumber)
	fmt.Printf("Variable is of type: %T \n", coolNumber)

	//dafault values
	var intVariables int
	fmt.Println(intVariables)
	fmt.Printf("Variable is of type: %T \n", intVariables) //int defult value is 0

	var stringVariables string
	fmt.Println(stringVariables)
	fmt.Printf("Variable is of type: %T \n", stringVariables) //string defult value is ""

	//no var style
	numberOfUser :=3000.0
	fmt.Println(numberOfUser)
	fmt.Printf("Variable is of type: %T \n", numberOfUser)
	//this style is cannot used in outside the func

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}
