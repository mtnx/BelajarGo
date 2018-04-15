package main

import (
	"fmt"
)

func main() {
	//===== FUNGSI ADD =====
	//deklarasi nilai x dan y
	x := 5
	y := 3

	//variable z buat menmpung nilai dari fungsi add
	z := add(x, y)

	//cetak fungsi add
	fmt.Println(z)

	//===== FUNGSI HELLO =====
	//deklarasi nilai name
	name := "Hyu"
	//variable result untuk menampung nilai dari fungsi hello
	result := hello(name)
	//cetak nilai fungsi hello
	fmt.Println(result)

	//===== FUNGSI MULTI =====
	//deklarasi nilai mulA dan mulB

	aa := "hey"
	bb := 47

	vaa, vbb := multi(aa, bb)

	fmt.Println(vaa, vbb)

}

//fungsi add mempunysi parameter x dan y type data int, dan kembalian int
func add(x int, y int) int {
	//saat dikembalikan x langsung dikurangi y
	return x - y
}

//fungsi hello mempunyai oarameter  name tipedata string dan lembalian nilai string
func hello(name string) string {
	//mengembalikan nilai name dan langsung di gabung dengan typedata lain
	return fmt.Sprint("Hello ", name)
}

//fungsi multi mempunyai nilai multiple value
func multi(a string, b int) (string, int) {
	return a, b
}

