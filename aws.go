package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// PublishContent function for put blog static files to AWS S3
func PublishContent(config *Config, content []*Content) {
	creds := credentials.NewStaticCredentials(config.AwsKey, config.AwsSecret, "")
	if _, err := creds.Get(); err != nil {
		fmt.Printf("bad credentials: %s", err)
	}

	cnf := aws.NewConfig().WithRegion(config.AwsBucketRegion).WithCredentials(creds)
	scv := s3.New(session.New(cnf))

	messages := make(chan string)
	go func() {
		for _, file := range content {
			upload(scv, file, config, messages)
		}
	}()

	files := len(content)
	for i := 1; i <= files; i++ {
		if config.Debug {
			fmt.Print(<-messages)
		}
	}
}

func upload(scv *s3.S3, file *Content, config *Config, c chan string) {
	_, err := scv.PutObject(&s3.PutObjectInput{
		Bucket:        aws.String(config.AwsBucket),
		Key:           aws.String(file.BlogPath),
		ACL:           aws.String("public-read"),
		Body:          file.Body,
		ContentLength: aws.Int64(file.Size),
		ContentType:   aws.String(file.Type),
	})
	if err != nil {
		c <- fmt.Sprintf("Failed: %s. Error: %s\n", file.BlogPath, err.Error())
		return
	}
	c <- fmt.Sprintf("Uploaded: %s\n", file.BlogPath)
}
