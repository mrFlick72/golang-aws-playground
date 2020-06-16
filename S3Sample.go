package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
)

func DownloadFile(session *session.Session, fileName string, bucket *string, bucketKey *string) {
	downloader := s3manager.NewDownloader(session)

	file, err := os.Create(fileName)
	fmt.Println("error", err)

	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: bucket,
			Key:    bucketKey,
		})

	fmt.Println("error", err)
	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")

	defer file.Close()
}
