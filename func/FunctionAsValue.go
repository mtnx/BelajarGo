/*
Fungsi sebagai Value
Di Golang kita bisa membuat fungsi seperti kita membuat Variable

fitur yang sangat keren
*/

package main

import (
	"fmt"
)

func main() {

	//nama keren anonymous function
	add := func(x, y int) int {
		return x + y
	}

	hey := func(nama string) string {
		return fmt.Sprint("hello ", nama)
	}

	fmt.Println(add(5, 5))
	fmt.Println(hey("Hyu"))
}

