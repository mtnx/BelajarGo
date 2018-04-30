package main

import (
	"fmt"
	"strings" //library bawaan go untuk manipulasi String
)

func main() {
	//var mystring bernilai string
	var mystring = "Hello Hyu"

	//cetak mystring
	fmt.Println(mystring)

	//cetak jumlah karakter mystring
	fmt.Println(len(mystring))

	//var mystring gede dideklarasikan
	mystringGede := strings.ToUpper(mystring)
	//cetak mystringGede
	fmt.Println(mystringGede)

	//var mystringCilik dideklarasikan
	mystringCilik := strings.ToLower(mystring)
	//cetak mystring
	fmt.Println(mystringCilik)
}

