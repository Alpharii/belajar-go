// ## `container/list` di Golang

// Package `container/list` adalah library bawaan Go yang menyediakan struktur data **doubly linked list** (list berantai ganda).
// Berbeda dengan array atau slice yang menggunakan indeks, `list` berisi node/element yang saling terhubung dan dapat bergerak maju (`Next`) atau mundur (`Prev`).

// ### Kapan digunakan?
// `container/list` cocok ketika:
// - Data sering **ditambah atau dihapus di tengah list**
// - Perlu traversal maju/mundur dengan efisien
// - Butuh struktur seperti **queue**, **stack**, **deque**, atau antrian lainnya

// ### Struktur yang Disediakan
// - `List` → container utama
// - `Element` → item yang disimpan
// - Akses elemen:
//   - `l.Front()` → elemen pertama
//   - `l.Back()` → elemen terakhir
//   - `e.Next()` → elemen setelahnya
//   - `e.Prev()` → elemen sebelumnya

package main

import (
	"container/list"
	"fmt"
)

func main(){
	lists:= list.New()

	//add data
	lists.PushBack("halo")
	lists.PushBack("bintang")
	lists.PushBack("alphari")

	fmt.Println(lists.Front())
	fmt.Println(lists.Front().Next().Next().Value)

	for e := lists.Front(); e!= nil; e = e.Next(){
		fmt.Println(e.Value)
	}
}