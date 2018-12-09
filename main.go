package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/yusukemisa/s3-unzip-go/zip"

	"github.com/yusukemisa/s3-unzip-go/s3"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/google/uuid"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
)

var (
	id string
	// zip file path
	zipContetPath string
	// unzip dest path
	unzipContentPath string
	// unziped file upload for
	destBucket string
)

const (
	tempArtifactPath = "/tmp/artifact/"
	tempZipPath      = tempArtifactPath + "zipped/"
	tempUnzipPath    = tempArtifactPath + "unzipped/"
	tempZip          = "temp.zip"
	dirPerm          = 0777
	region           = endpoints.ApNortheast1RegionID
)

func init() {
	destBucket = os.Getenv("UNZIPPED_ARTIFACT_BUCKET")
}

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, s3Event events.S3Event) {
	if lc, ok := lambdacontext.FromContext(ctx); !ok {
		log.Printf("requestID=%s", lc.AwsRequestID)
	}

	r := s3Event.Records[0]
	log.Printf("%+v", r)
	bucket := r.S3.Bucket.Name
	key := r.S3.Object.Key
	log.Printf("bucket=%s,key=%s", bucket, key)

	if err := prepareDirectory(); err != nil {
		log.Fatal(err)
	}

	//認証情報
	s := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))

	downloader := s3.NewDownloader(s, bucket, key, zipContetPath+tempZip)
	downloadedZipPath, err := downloader.Download()
	if err != nil {
		log.Fatal(err)
	}

	if err := zip.Unzip(downloadedZipPath, unzipContentPath); err != nil {
		log.Fatal(err)
	}

	uploader := s3.NewUploader(s, tempUnzipPath, destBucket)
	if err := uploader.Upload(); err != nil {
		log.Fatal(err)
	}

	log.Printf("%s unzipped to S3 bucket: %s", downloadedZipPath, destBucket)
}

// Lambdaの実行環境では/tmpに対するwriteが可能。ただし512MB上限
// 実行環境自体はリクエストごとに作成される場合もあれば、既存のものを使い回される可能性もあるので/tmp配下に依存はできない。
func prepareDirectory() error {
	id = uuid.New().String()
	zipContetPath = fmt.Sprintf("%s/%s/", tempZipPath, id)
	unzipContentPath = fmt.Sprintf("%s/%s/", tempUnzipPath, id)

	if _, err := os.Stat(tempArtifactPath); err != nil {
		if err := os.RemoveAll(tempArtifactPath); err != nil {
			return err
		}
	}
	if err := os.MkdirAll(zipContetPath, dirPerm); err != nil {
		return err
	}
	if err := os.MkdirAll(unzipContentPath, dirPerm); err != nil {
		return err
	}
	return nil
}
