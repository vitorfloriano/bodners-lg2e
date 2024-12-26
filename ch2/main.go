package main

import 	"fmt"

func Exercise1() {
	var i int = 20
	 f := float64(i)
	 fmt.Println(i, f)
}


func Exercise2() {
	const value = 10
	var i int = value
	var f float64 = value
	fmt.Println(i, f)
}


func Exercise3() {
	var (
		b 	byte
		smallI 	int32
		bigI 	uint64
	)
	b = 127
	smallI = 2147483647
	bigI = 18446744073709551615
	fmt.Println(b, smallI, bigI)
}

func main() {
	Exercise1()
	Exercise2()
	Exercise3()
}
