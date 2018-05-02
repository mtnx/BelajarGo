/*

Map mirip seperti array atau slice
tetapi Map memiliki Key dan Elemen


Bentuk Map :

map[string]int{}
make(map[string]int)
*new(map[string]int)

*/

package main

import (
	"fmt"
)

func main() {
	// var mon bernilai map, dengan key string dan value/elemen int
	var mon = map[string]int{
		"jan": 10,
		"feb": 20,
		"mar": 30,
	}
	// ke : keynya, va: value atau elemennya
	for ke, va := range mon {
		fmt.Println(ke, " : ", va)
	}
	//mengecek apakah di key tersebut mempunyai value
	var val, cek = mon["mar"]
	if cek {
		fmt.Println(val)
	} else {
		fmt.Println("raono bro..")
	}
}

/*

Hasilnya adalah :

jan  :  10
feb  :  20
mar  :  30

*/
