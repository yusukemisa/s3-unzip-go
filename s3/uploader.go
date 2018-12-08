package s3

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type Uploader struct {
	manager  *s3manager.Uploader
	src, dst string
}

func NewUploader(s *session.Session, src, dst string) *Uploader {
	return &Uploader{
		manager: s3manager.NewUploader(s),
		src:     "",
		dst:     "",
	}
}

func (u *Uploader) Upload() error {
	return nil
}
