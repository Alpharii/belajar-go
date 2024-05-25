package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating of our pizza:")

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for the rating,", input)
	fmt.Printf("type of this rating is %T", input)
}
