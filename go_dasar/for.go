package main

import "fmt"

func main(){
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("loop ke", counter)
	// 	counter++
	// }

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("loop ke", counter)
	}

	fmt.Println("selesai")

	names := []string{"eko", "kurniawan", "khanedy"}
	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	for index, value := range names{
		fmt.Println("index ke =", index, "value =", value)
	}

	for _, value := range names{
		fmt.Println("value =", value)
	}
}