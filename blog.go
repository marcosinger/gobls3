package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Content is the blog content file structure
type Content struct {
	Body *bytes.Reader
	Type string
	Size int64
	Path string
}

// GetContentToUpdate functions for get list of updated static blog files
func GetContentToUpdate(blogPath string) []*Content {
	files, date := getFilesWithModDate(blogPath)
	return getUpdatedFiles(date, files, blogPath)
}

// todo : should be refactored for recursion possibility
func getFilesWithModDate(blogPath string) (map[string]int64, int64) {
	files, err := ioutil.ReadDir(blogPath + "/public")
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}

	var lastUpdate int64
	var blogFiles = make(map[string]int64)
	for _, f := range files {
		name := f.Name()
		// todo : make it recursive here
		if name == ".DS_Store" || f.IsDir() {
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

func getUpdatedFiles(date int64, files map[string]int64, blogPath string) []*Content {

	var final []*Content

	//var updated []string
	for name, lastUpdated := range files {
		if lastUpdated == date {
			//updated = append(updated, blogPath+"/public/"+name)
			content, _ := getContent(blogPath + "/public/" + name)
			final = append(final, content)
		}
	}
	return final
}

func getContent(path string) (*Content, error) {
	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		fmt.Printf("Blog file cannot be open: %s\n", path)
		return &Content{}, err
	}

	fileInfo, _ := file.Stat()
	size := fileInfo.Size()

	buffer := make([]byte, size)
	file.Read(buffer)
	fileBytes := bytes.NewReader(buffer)
	fileType := http.DetectContentType(buffer)

	return &Content{Body: fileBytes, Size: size, Type: fileType, Path: path}, nil
}
