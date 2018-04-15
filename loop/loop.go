package main

import (
	"fmt"
)

func main() {
	cara1()
	cara2()
	cara3()
	cara4()
}

func cara1() {
	//deklarasi nilai n
	n := 10

	// nilai awal x = 0, x kurang dari n, x selalu bertambah 1 (increment)
	for x := 0; x < n; x++ {
		//cetak x
		fmt.Print(x)

	}
	// hasilnya adalah mencetak nilai 0-9
}

func cara2() {
	//deklarasi nilai n adalah 10
	n := 10
	// dibuat variable x karena x belum diketahui nilainya
	var x int

	//ketika nilai x kurang dari nilai n, maka
	for x < n {
		//cetak x
		fmt.Print(x)
		//increment x, supaya x mempunyai nilai
		x++
	}
	// hasilnya adalah mencetak nilai 0-9
}

func cara3() { // pengganti while
	//deklarasi nilai x adalah 1
	x := 1
	// for tanpa kondisi (infinite loop)
	for {

		//cetak x
		fmt.Print(x)
		// x increment
		x++
		//jika x = 10
		if x == 10 {
			//maka berhenti
			break
		}
	}
}

func cara4() {

	//deklarai nilai x = 10
	x := 1
	//for tanpa kondisi
	for {
		//x increment
		x++
		//jika x bernilai 5
		if x == 5 {
			//maka lewati
			continue
		}
		//cetak nilai x
		fmt.Print(x)
		//jika nilai x = 10
		if x == 10 {
			//maka berhenti
			break
		}
	}
}

