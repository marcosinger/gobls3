package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// PublishContent function for put blog static files to AWS S3
func PublishContent(config *Config, content []*Content) {

}

// GetListOfBuckets function for get list of s3 buckets
func GetListOfBuckets(config *Config) error {

	creds := credentials.NewStaticCredentials(config.AwsKey, config.AwsSecret, "")
	if _, err := creds.Get(); err != nil {
		fmt.Printf("bad credentials: %s", err)
	}

	cnf := aws.NewConfig().WithRegion("us-west-1").WithCredentials(creds)
	svc := s3.New(session.New(), cnf)
	input := &s3.ListBucketsInput{}

	result, err := svc.ListBuckets(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		}
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return err
	}
	fmt.Println(result)

	return nil
}
