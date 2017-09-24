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
	for _, file := range files {
		fmt.Printf("Filepath: %s, Blogpath: %s\n", file.Path, file.BlogPath)
	}
}

func exitOnError(err error) {
	if err != nil {
		os.Exit(0)
	}
}
