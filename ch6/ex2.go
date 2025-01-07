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

func UpdateSlice(strsl []string, str string) {
	strsl[len(strsl)-1] = str
	fmt.Println("in UpdateSlice: ", strsl)
}	

func GrowSlice(strsl []string, str string) {
	strsl = append(strsl, str)
	fmt.Println("in GrowSlice: ", strsl)
}

func main() {
	p := MakePerson("vitor", "floriano", 32)
	fmt.Println(p)
	p2 := MakePersonPointer("joao", "pedro", 34)	
	fmt.Println(p2)
	slice := []string{"1","2","3"}
	fmt.Println("Current slice: ", slice)
	UpdateSlice(slice, "4")
	fmt.Println("After UpdateSlice: ", slice)
	GrowSlice(slice, "4")
	fmt.Println("After GrowSlice: ", slice)
}
