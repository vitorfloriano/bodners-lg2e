package main

import "fmt"

type Person struct {
	FirstName string
	LastName string
	Age int
}

func MakePerson(firstName, lastName string, age int) Person {
	p := Person{firstName, lastName, age}
	return p
}

func MakePersonPointer(firstName, lastName string, age int) *Person {
	p := &Person{firstName, lastName, age}
	return p
}

func main() {
	p := MakePerson("vitor", "floriano", 32)
	fmt.Println(p)
	p2 := MakePersonPointer("joao", "pedro", 34)	
	fmt.Println(p2)
	fmt.Println(p)
}
