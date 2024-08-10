package main

import "fmt"

func simpleArraySum(ar [6]int32) int32 {
	var sum int32
	sum = 0
	for i := 0; i < len(ar); i++ {
		sum += ar[i]
	}
	return sum
}

func main() {
	var ar [6]int32 = [6]int32{1, 2, 3, 4, 10, 11}

	result := simpleArraySum(ar)

	fmt.Println(result)
}
