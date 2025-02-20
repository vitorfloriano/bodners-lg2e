package main

import (
	"os"
	"fmt"
	"embed"
	"strings"
)

//go:embed english_rights.txt portuguese_rights.txt 
var files embed.FS

var languages = map[string]string{
	"english": "english_rights.txt",
	"portuguese": "portuguese_rights.txt",
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <language>")
		return
	}

	language := strings.ToLower(os.Args[1])
	
	filename, exists := languages[language]
	if !exists {
		fmt.Println("Language not supported")
		return
	}
	
	content, err := files.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	
	fmt.Println(string(content))
}
