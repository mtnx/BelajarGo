/*
Closure
FUngsi yang mengembalikan fungsi

*/

package main

import (
	"fmt"
)

func main() {
	//secara kasat mata ini adalah fungsi juga
	nextValue := genNumber()

	//cetak nilai NextValue
	fmt.Println(nextValue())
	//nilai selanjutnya
	fmt.Println(nextValue())
	//nilai selanjutnya
	fmt.Println(nextValue())

	//secara kasat mata ini adalah function juga
	//yang bernilai func love("wahyu") juga "love"
	lv := love("Wahyu")
	//mencetak func lv dan juga "Golang"
	fmt.Println(lv("Golang"))
	//mencetak func lv dan juga "Girls"
	fmt.Println(lv("Girls"))

}

// genNumber untuk generate angka, dan mengembalikan fungsi lain
// func() adalah anonymous function bertype data int
func genNumber() func() int {
	//x bernilai 0
	x := 0

	//kembalikan nilai func anonymous
	return func() int {
		//di dalam anonymous func ini nilai x selalu ditambah satu
		x = x + 1
		//kita kembalikan lagi nilai x
		return x
	}
}

/*

hasilnya adalah 1,2,3

*/

//func love mempunyai parameter name dengan typedata string dan mempunyai kembalian func lagi
//sedangkan func anonym mempunyai parameter string dan kembalian nilai string
func love(name string) func(string) string {
	//kembalikan nilai func love
	return func(things string) string {
		// kembalikan func anonym
		// Sprint = concat atau menggabungkan beberapa typedata
		return fmt.Sprintf("%s love %s", name, things)
	}
}

/*
Hasilnya adalah :

Wahyu love Golang
Wahyu love Girls

*/

