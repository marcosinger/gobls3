package main

import (
	"io/ioutil"
	"log"
	"os"
)

// GetContentToUpdate functions for get list of updated static blog files
func GetContentToUpdate() []string {
	files, date := getFilesWithModDate()
	return getUpdatedFiles(date, files)
}

func getFilesWithModDate() (map[string]int64, int64) {
	files, err := ioutil.ReadDir(os.Getenv("BLOG_PATH") + "/public")
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}

	var lastUpdate int64
	var blogFiles = make(map[string]int64)
	for _, f := range files {
		name := f.Name()
		if name == ".DS_Store" {
			continue
		}
		timeStamp := f.ModTime().Unix()
		if lastUpdate < timeStamp {
			lastUpdate = timeStamp
		}
		blogFiles[name] = timeStamp
	}

	return blogFiles, lastUpdate
}

func getUpdatedFiles(date int64, files map[string]int64) []string {
	var updated []string
	for name, lastUpdated := range files {
		if lastUpdated == date {
			updated = append(updated, os.Getenv("BLOG_PATH")+"/public/"+name)
		}
	}
	return updated
}
