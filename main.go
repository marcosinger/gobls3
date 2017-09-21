package main

import (
	"fmt"
)

func main() {
	content := GetContentToUpdate()
	for _, file := range content {
		fmt.Printf("Filepath: %s\n", file)
	}
}
