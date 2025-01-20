package main

import (
	"fmt"
)

type Numbers interface {
	int | int8 | int16 | int32 | int64 |
 	uint | uint8 | uint16 | uint32 | uint64 | uintptr |
	float32 | float64
}

func Double[T Numbers](number T) T {
	return number * 2
}

func main() {
	fmt.Println(Double(2))
	fmt.Println(Double(2.5))
}	
