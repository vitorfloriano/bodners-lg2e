package ex3

import "fmt"

func Exercise3() {
	var (
		b 	byte
		smallI 	int32
		bigI 	uint64
	)
	b = 127
	smallI = 2147483647
	bigI = 18446744073709551615
	fmt.Println(b++, smallI++, bigI++)
}
