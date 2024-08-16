package main

import "fmt"

/*
 * Complete the 'aVeryBigSum' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts LONG_INTEGER_ARRAY ar as parameter.
 */

func aVeryBigSum(ar []int64) int64 {
	var sum int64
	sum = 0
	for i := 0; i < len(ar); i++ {
		sum += ar[i]
	}
	return sum
}

func main() {
	ar := []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}

	result := aVeryBigSum(ar)

	fmt.Printf("%d\n", result)

}

/* Resolução exige entender o range de diferentes tipos de dados, considerando onde
o `int32` e `int64` são diferentes*/
