package main

import (
	"fmt"
	"os"
	"strconv"
)

// Config contain the configuration
type Config struct {
	BlogPath,
	AwsKey,
	AwsSecret,
	AwsBucket,
	AwsBucketRegion string
	Debug    bool
	Database ConfigDatabase
}

// ConfigDatabase contain the database configuration
type ConfigDatabase struct {
	Host, Name string
}

// LoadConfig function for get configuration from file
func LoadConfig() *Config {
	debug, err := strconv.ParseBool(env("DEBUG"))
	if err != nil {
		panic("Invalid parameter type")
	}

	return &Config{
		BlogPath:        env("BLOG_PATH"),
		AwsKey:          env("AWS_ACCESS_KEY"),
		AwsSecret:       env("AWS_SECRET_ACCESS_KEY"),
		AwsBucket:       env("AWS_S3_BUCKET"),
		AwsBucketRegion: env("AWS_S3_BUCKET_REGION"),
		Debug:           debug,
		Database:        getDatabaseConfig(),
	}
}

func env(name string) string {
	val := os.Getenv(name)
	if val == "" {
		panic(fmt.Sprintf("Environment variable %s was not set", name))
	}
	return val
}

func getDatabaseConfig() ConfigDatabase {
	return ConfigDatabase{
		Host: env("DB_HOST"),
		Name: env("DB_NAME"),
	}
}
