package s3

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type Downloader struct {
	manager           s3manager.Downloader
	bucket, key, dest string
}

func (downloader Downloader) Download() (string, error) {
	return "", nil
}

func NewDownloader(s *session.Session, bucket, key, dest string) *Downloader {
	return &Downloader{
		manager: *s3manager.NewDownloader(s),
		bucket:  bucket,
		key:     key,
		dest:    dest,
	}
}
