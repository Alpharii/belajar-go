package main

import "fmt"

func main(){
	names := [...]string{"bintang", "alphari", "eko", "kurniawan", "khanedy" ,"pzn"}
	slice := names[4:6]

	fmt.Println(slice)

	slice2:= names[:3]
	fmt.Println(slice2)	
	
	slice3:= names[3:]
	fmt.Println(slice3)

	days:= [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}

	daySlice1 := days[5:]
	fmt.Println(daySlice1)

	daySlice1[0] = "sabtu baru"
	daySlice1[1] = "minggu baru"
	fmt.Println(daySlice1)
	fmt.Println(days)

	daySlice2 := append(daySlice1, "libur baru")
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "bintang"
	newSlice[1] = "bintang"
	// newSlice[2] = "bintang" // error harusnya pakai apppend

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "eko")
	fmt.Println(newSlice)
	fmt.Println(newSlice2)

	newSlice2[0] = "budi"
	fmt.Println(newSlice)
	fmt.Println(newSlice2)

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	iniArray := [...]int{1,2,3,4,5}
	iniSlice := []int{1,2,3,4,5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}