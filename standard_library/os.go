package main

import (
	"fmt"
	"os"
)

// os (Operating System) menyediakan fungsi untuk berinteraksi langsung dengan sistem operasi, meliputi:
// File dan direktori
// Environment variable
// Argumen program (CLI)
// Hostname dan informasi sistem
// Exit code dan signal
// Proses dan permission

func main(){
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <nama>")
		return
	}
	nama := os.Args[1]
	fmt.Println("Halo,", nama)

	hostname, err := os.Hostname()
	if err != nil{
		fmt.Println(err)
	}
	
	fmt.Println(hostname)

	os.Setenv("APP_MODE", "development")
	fmt.Println(os.Getenv("APP_MODE"))

}