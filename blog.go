package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// Content is the blog content file structure
type Content struct {
	Body     *bytes.Reader
	Type     string
	Size     int64
	Path     string
	BlogPath string
}

// GetContentToUpdate functions for get list of updated static blog files
func GetContentToUpdate(blogPath string) ([]*Content, error) {
	var content []*Content
	files, date, err := getFiles(blogPath + "/public")
	if err != nil {
		return content, err
	}
	content = getUpdatedFiles(date, files, blogPath)
	return content, nil
}

func getFiles(path string) (map[string]int64, int64, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Printf("Cannot read directore by path: %s\n", path)
		return make(map[string]int64), 0, err
	}

	var lastUpdate int64
	var blogFiles = make(map[string]int64)

	for _, f := range files {
		name := f.Name()
		if name == ".DS_Store" {
			continue
		}
		if f.IsDir() {
			blogDirFiles, timeStamp, err := getFiles(path + "/" + name)
			if err != nil {
				return make(map[string]int64), 0, err
			}
			for k, v := range blogDirFiles {
				blogFiles[k] = v
			}
			if lastUpdate < timeStamp {
				lastUpdate = timeStamp
			}
		} else {
			timeStamp := f.ModTime().Unix()
			blogFiles[path+"/"+name] = timeStamp
			if lastUpdate < timeStamp {
				lastUpdate = timeStamp
			}
		}
	}

	return blogFiles, lastUpdate, nil
}

func getUpdatedFiles(date int64, files map[string]int64, blogPath string) []*Content {
	var final []*Content
	for path, lastUpdated := range files {
		if lastUpdated == date {
			content, _ := getContent(path, blogPath)
			final = append(final, content)
		}
	}
	return final
}

func getContent(path string, blogPath string) (*Content, error) {
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

	return &Content{
		Body:     fileBytes,
		Size:     size,
		Type:     fileType,
		Path:     path,
		BlogPath: strings.TrimPrefix(path, blogPath+"/public"),
	}, nil
}
