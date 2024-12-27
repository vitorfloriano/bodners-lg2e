package main

import "fmt"

func main() {
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	firstSub := greetings[:2]
	secondSub := greetings[1:4]
	thirdSub := greetings[3:]
	fmt.Println(greetings)
	fmt.Println(firstSub)
	fmt.Println(secondSub)
	fmt.Println(thirdSub)
}
