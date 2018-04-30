package main

import (
	"fmt"

)

func main(){
	outerLoop:
	for i :=0;i<5;i++{
		for j :=0;j<5;j++{
			if i== 3{
				break outerLoop
			}
			fmt.Println("Matriks [",i,"][",j,"]","\n")
		}
	}
}



/*



Hasilnya..



Matriks [ 0 ][ 0 ]

Matriks [ 0 ][ 1 ]

Matriks [ 0 ][ 2 ]

Matriks [ 0 ][ 3 ]

Matriks [ 0 ][ 4 ]

Matriks [ 1 ][ 0 ]

Matriks [ 1 ][ 1 ]

Matriks [ 1 ][ 2 ]

Matriks [ 1 ][ 3 ]

Matriks [ 1 ][ 4 ]

Matriks [ 2 ][ 0 ]

Matriks [ 2 ][ 1 ]

Matriks [ 2 ][ 2 ]

Matriks [ 2 ][ 3 ]

Matriks [ 2 ][ 4 ]






*/
