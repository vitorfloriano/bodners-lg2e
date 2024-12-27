package main

import "fmt"

func main(){
	type Employee struct{
		firstName string
		lastName string
		id int 
	}
	firstInst := Employee{
		"Rob",
		"Pike",
		1,
	}
	secondInst := Employee{
		firstName: "Ken",
		lastName: "Thompson",
		id: 2,
	}
	var thirdInst Employee
	thirdInst.firstName = "Robert"
	thirdInst.lastName = "Griesenemer"
	thirdInst.id = 3
	fmt.Println(firstInst)
	fmt.Println(secondInst)
	fmt.Println(thirdInst)
}
