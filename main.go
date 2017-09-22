package main

import (
	"fmt"
	"os"
)

func main() {
	config, err := LoadConfig()
	if err != nil {
		os.Exit(0)
	}

	files := GetContentToUpdate(config.BlogPath)
	for _, content := range files {
		fmt.Printf("Content: %v\n", content)
	}

	//GetListOfBuckets(config)
}
