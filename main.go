package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"os"
)

func main() {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("eu-central-1")}))

	DownloadFile(sess,
		os.Getenv("s3.bucket.file.name"),
		aws.String(os.Getenv("s3.bucket")),
		aws.String(os.Getenv("s3.bucket.file.key")),
	)
}
