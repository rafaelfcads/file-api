package helper

import (
	"bytes"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func PublishToS3(key string, buffer *bytes.Buffer) (string, error) {

	creds := credentials.Value{AccessKeyID: os.Getenv("BUCKETEER_AWS_ACCESS_KEY_ID"), SecretAccessKey: os.Getenv("BUCKETEER_AWS_SECRET_ACCESS_KEY")}

	sess, sessErr := session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("BUCKETEER_AWS_REGION")),
		Credentials: credentials.NewStaticCredentialsFromCreds(creds),
	})

	if sessErr != nil {
		return "", fmt.Errorf("failed to create file %q, %v", key, sessErr)
	}

	result, err := s3manager.NewUploader(sess).Upload(&s3manager.UploadInput{
		Bucket: aws.String(os.Getenv("BUCKETEER_BUCKET_NAME")),
		Key:    aws.String(key),
		Body:   buffer,
	})

	return result.Location, err
}

func GetS3AsBuffer(key string) (*aws.WriteAtBuffer, error) {
	creds := credentials.Value{AccessKeyID: os.Getenv("BUCKETEER_AWS_ACCESS_KEY_ID"), SecretAccessKey: os.Getenv("BUCKETEER_AWS_SECRET_ACCESS_KEY")}

	sess, sessErr := session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("BUCKETEER_AWS_REGION")),
		Credentials: credentials.NewStaticCredentialsFromCreds(creds),
	})

	if sessErr != nil {
		return nil, fmt.Errorf("failed to create sess %q, %v", key, sessErr)
	}

	// f, err := os.Create("./" + key + ".xlsx")
	// if err != nil {
	// 	return -1, fmt.Errorf("failed to create file %q, %v", key, sessErr)
	// }

	buff := &aws.WriteAtBuffer{}

	_, err := s3manager.NewDownloader(sess).Download(buff, &s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("BUCKETEER_BUCKET_NAME")),
		Key:    aws.String(key),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to write file %q, %v", key, err)
	}

	return buff, err
}
