package main

import (
	"fmt"
)

/*
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int32) []float32 {
	arrRatios := make([]float32, 3) // Alocar um slice de tamanho 3
	a := float32(len(arr))          // Converter o comprimento para float32
	var positive, negative, zero int32

	for i := 0; i < len(arr); i++ {
		if arr[i] > 0 {
			positive += 1
		} else if arr[i] < 0 {
			negative += 1
		} else {
			zero += 1
		}
	}

	// Calcular as proporções
	arrRatios[0] = float32(positive) / a
	arrRatios[1] = float32(negative) / a
	arrRatios[2] = float32(zero) / a

	return arrRatios
}

func main() {
	var arr = []int32{-4, 3, -9, 0, 4, 1}
	arrResult := plusMinus(arr)
	for i := 0; i < len(arrResult); i++ {
		fmt.Printf("%.5f\n", arrResult[i])
	}
}
