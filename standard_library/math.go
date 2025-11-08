package main

import (
	"fmt"
	"math"
)

// math.Round(float64)			//Membulatkan float64 keatas atau kebawah, sesuai dengan yang paling dekat
// math.Floor(float64)			//Membulatkan float64 kebawah
// math.Ceil(float64)			//Membulatkan float64 keatas
// math.Max(float64, float64)	//Mengembalikan nilai float64 paling besar
// math.Min(float64, float64)	//Mengembalikan nilai float64 paling kecil

func main(){
	fmt.Println(math.Ceil(1.40))
	fmt.Println(math.Floor(1.60))
	fmt.Println(math.Round(1.5))
	fmt.Println(math.Max(1,2))
	fmt.Println(math.Min(1,2))


	fmt.Println(math.Pow(2, 3)) // 2 pangkat 3 = 8
}