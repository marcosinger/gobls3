package main

import (
	"fmt"
	"os"
	"strconv"
)

// Config contain the configuration
type Config struct {
	BlogPath string
	Aws      ConfigAmazon
	Debug    bool
	Database ConfigDatabase
}

// ConfigDatabase contain the database configuration
type ConfigDatabase struct {
	Host, Name string
}

// ConfigAmazon contain the aws configuration
type ConfigAmazon struct {
	Key, Secret, Bucket, Region string
}

// LoadConfig function for get configuration from file
func LoadConfig() *Config {
	debug, err := strconv.ParseBool(env("DEBUG"))
	if err != nil {
		panic("Invalid parameter type")
	}

	return &Config{
		BlogPath: env("BLOG_PATH"),
		Aws:      getAwsConfig(),
		Debug:    debug,
		Database: getDatabaseConfig(),
	}
}

func getAwsConfig() ConfigAmazon {
	return ConfigAmazon{
		Key:    env("AWS_ACCESS_KEY"),
		Secret: env("AWS_SECRET_ACCESS_KEY"),
		Bucket: env("AWS_S3_BUCKET"),
		Region: env("AWS_S3_BUCKET_REGION"),
	}
}

func getDatabaseConfig() ConfigDatabase {
	return ConfigDatabase{
		Host: env("DB_HOST"),
		Name: env("DB_NAME"),
	}
}

func env(name string) string {
	val := os.Getenv(name)
	if val == "" {
		panic(fmt.Sprintf("Environment variable %s was not set", name))
	}
	return val
}
