package main

import (
	"fmt"
	"os"
)

func main() {
	config, err := LoadConfig()
	exitOnError(err)

	files, err := GetContentToUpdate(config.BlogPath)
	exitOnError(err)

	if config.Debug {
		showContentToPublish(files)
	}

	PublishContent(config, files)
}

func showContentToPublish(files []*Content) {
	fmt.Printf("%d files should be upload to S3\n\n", len(files))
	for _, file := range files {
		fmt.Printf("File: %s\n", file.BlogPath)
	}
}

func exitOnError(err error) {
	if err != nil {
		os.Exit(0)
	}
}
