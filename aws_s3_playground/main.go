package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func main() {
	filename := "testfors3.xlsx"
	result := uploadFile(filename)

	fmt.Println("\nRESULT => ", result.Location)
}

func uploadFile(filename string) *s3manager.UploadOutput {
	// The session the S3 Uploader will use
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("failed to open file %q, %v", filename, err)
	}

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("nominas-calculadas"),
		Key:    aws.String(filename),
		Body:   f,
	})
	if err != nil {
		fmt.Printf("failed to upload file, %v", err)
	}

	return result

}
