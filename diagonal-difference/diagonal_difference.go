package main

import (
	"fmt"
)

/*
 * Complete the 'diagonalDifference' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

func diagonalDifference(arr [][]int32) int32 {
	// usando a diagonal de uma matriz
	var dL int32
	var dR int32

	for line := 0; line < len(arr); line++ {
		for columm := 0; columm < len(arr[]); columm++ {
			
		}
	}

}

func main() {
	array := [][]int32{{1, 2, 3}, {4, 5, 6}, {9, 8, 9}}

	result := diagonalDifference(array)
	fmt.Printf("%d\n", result)
	fmt.Printf("Hello, World")
}
