package main

import (
	"log"
	"os"
)

func main() {
	config, err := LoadConfig()
	exitOnError(err)

	files, err := GetContentToUpdate(config.BlogPath)
	exitOnError(err)

	if config.Debug {
		log.Printf("Detected %d files that should be upload to S3\n\n", len(files))
	}

	PublishContent(config, files)
}

func exitOnError(err error) {
	if err != nil {
		os.Exit(0)
	}
}
