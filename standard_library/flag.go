package main

import (
	"flag"
	"fmt"
)

func main(){
	name := flag.String("name", "anon", "masukan nama")
	age := flag.Int("age", 0, "masukan umur")

	fmt.Println(*name, *age)

	
	verbose := flag.Bool("verbose", false, "Tampilkan log detail")
	flag.Parse()
	
	if *verbose {
		fmt.Println("Mode verbose aktif")
	}
	fmt.Println(*name, *age)
	fmt.Printf("halo %s, umur kamu %d tahun\n", *name, *age)
}