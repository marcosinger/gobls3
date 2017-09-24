package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config contain the configuration
type Config struct {
	BlogPath  string `json:"blog_path"`
	AwsKey    string `json:"aws_access_key_id"`
	AwsSecret string `json:"aws_secret_access_key"`
	AwsBucket string `json:"aws_s3_bucket"`
	Debug     bool   `json:"debug"`
}

// LoadConfig function for get configuration from file
func LoadConfig() (*Config, error) {
	file, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
		return &Config{}, err
	}

	decoder := json.NewDecoder(file)
	config := Config{}

	if err = decoder.Decode(&config); err != nil {
		fmt.Println(err)
		return &Config{}, err
	}
	return &config, nil
}
