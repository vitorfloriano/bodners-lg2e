package main

import (
	"os"
	"fmt"
	"io/ioutil"
)

//go:embed english_rights.txt 
var englishRights string

//go:embed portuguese_rights.txt
var portugueseRights string

func main
