package main

import (
	"fmt"
)

type Printable interface {
	fmt.Stringer
	~int | ~float64 
}

type PrintableInt int
type PrintableFloat float64

func (pi PrintableInt) String() string {
	return fmt.Sprintf("My Int: %d", pi)
}

func (pf PrintableFloat) String() string {
	return fmt.Sprintf("My Float: %f", pf)
}

func PrintValue[T Printable](p T) {
	fmt.Println(p.String())
}

func main() {
	var myInt PrintableInt = 1
	var myFloat PrintableFloat = 2.0
	PrintValue(myInt)
	PrintValue(myFloat)
}	
