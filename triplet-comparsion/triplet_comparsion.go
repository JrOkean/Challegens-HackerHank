package main

import "fmt"

/*
 * Complete the 'compareTriplets' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER_ARRAY b
 */

func compareTriplets(a []int32, b []int32) []int32 {
	var apoints int32
	var bpoints int32

	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			apoints += 1
		} else if a[i] < b[i] {
			bpoints += 1
		}
	}

	return []int32{apoints, bpoints}
}

func main() {
	var a = []int32{1, 2, 3}
	var b = []int32{3, 2, 1}

	result := compareTriplets(a, b)

	fmt.Println(result)
}
