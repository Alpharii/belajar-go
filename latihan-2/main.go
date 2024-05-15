package main
import "fmt"

//Belajar tipe data dan variabel

//var (variablenName) (type) = (value)

func main(){
	var saya string = "Muhammad Bintang Alphari" // string
	var tempatTinggal = "Palembang" //inferred
	var umur = 17 // inferred

	fmt.Println("Hallo nama saya " + saya + " saya tinggal di " + tempatTinggal + " umur saya adalah")
	fmt.Println(umur)

	//---------------------------------------------------------------------------------------//

	var a string // ""
	var b int // 0
	var c bool // false

	fmt.Println(a)
  	fmt.Println(b)
  	fmt.Println(c)

	//---------------------------------------------------------------------------------------//

	var name string
	name = "John"
	fmt.Println(name)
}