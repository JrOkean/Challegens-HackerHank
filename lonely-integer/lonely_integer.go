package main

import (
	"fmt"
)

func lonelyinteger(a []int32) int32 {
	var b int32
	b = a[0]
	//len_vector = len(a)

	return b
}

func main() {
	s := `Hello
					playground
							  hasld`
	fmt.Printf("%v\n%T\n", s, s)
	b := "Hello, playground"
	bb := []byte(b)
	for _, v := range bb {
		fmt.Printf("%v - %T - %#U - %#x", v, v, v, v)
	}
}

/* REMEMBERING HOW TO USE GO */
