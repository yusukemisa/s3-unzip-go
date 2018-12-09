package s3

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type Downloader struct {
	manager           *s3manager.Downloader
	bucket, key, dest string
}

func NewDownloader(s *session.Session, bucket, key, dest string) *Downloader {
	return &Downloader{
		manager: s3manager.NewDownloader(s),
		bucket:  bucket,
		key:     key,
		dest:    dest,
	}
}

func (d *Downloader) Download() (string, error) {
	f, err := os.Create(d.dest)
	if err != nil {
		return "", err
	}
	defer f.Close()

	numBytes, err := d.manager.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(d.bucket),
		Key:    aws.String(d.key),
	})

	if err != nil {
		return "", nil
	}

	log.Println("Downloaded ", f.Name(), numBytes, " bytes")

	return f.Name(), nil
}
