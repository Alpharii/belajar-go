package main

import (
	"fmt"
	"strings"
)

func main()  {
	fmt.Println(strings.Contains("Bintang Alphari", "Alphari"))
	fmt.Println(strings.Split("Bintang Alphari", " "))
	fmt.Println(strings.ToLower("Bintang Alphari"))
	fmt.Println(strings.ToUpper("Bintang Alphari"))
	fmt.Println(strings.Trim("   Bintang Alphari      ", " "))
	fmt.Println(strings.ReplaceAll("Bintang Alphari", "Alphari", "Eko"))
}