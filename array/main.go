/*

Perbedaan Array dan Slice hanya pada penentuan elemen,
- array ditetapkan elemennya		-->  [2]
- slice tidak ditetapkan elemennya	-->  []
*/

package main

func main() {

	var FruitsA = []string{"apple", "banana"}      // Sice
	var FruitsB = [2]string{"mango", "watermelon"} //Array
	var FruitsC = [...]string{"grape", "orange"}   //Array

}
