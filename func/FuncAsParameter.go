/*
Function As

*/

package main

import (
	"fmt"
)

func main() {
	//membuat fungsi sebagai value
	//fungsi disini harus sama persis dengan fungsi di parameter match
	ef := func(v string) bool {
		return v == "Golang"
	}

	result := match("Golang", ef)

	fmt.Println(result)

}

//fungsi ini mempunyai 2 buah argument
func match(v string, f func(string) bool) bool {
	//fungsi f membutuhka parameter pertama dari match yaitu v
	//karena type kembaliannya adalah bool, maka fungsi f langsung bisa di panggil di scope if ini
	if f(v) {
		return true
	}
	return false
}

