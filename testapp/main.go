package main

import (
	"fmt"
	"os"

	"github.com/dotpy3/godocdown"
)

// This file replicates the behaviour of the traditional godocdown in the current directory
func main() {
	doc, err := godocdown.GenerateDocumentation(".", "")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(doc)
}
