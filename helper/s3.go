package helper

import (
	"bytes"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func PublishToS3(key string, buffer *bytes.Buffer) (string, error) {

	const bucketName = "bucketeer-2e671b59-4d1e-47d6-bfd3-f1fa4e3d7f84"
	const region = "us-east-1"
	const accessKeyID = "AKIAVZH4SBSYZJFN4IVH"
	const secretAccessKey = "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"

	creds := credentials.Value{AccessKeyID: accessKeyID, SecretAccessKey: secretAccessKey}

	sess, sessErr := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentialsFromCreds(creds),
	})

	if sessErr != nil {
		return "", fmt.Errorf("failed to create file %q, %v", key, sessErr)
	}

	result, err := s3manager.NewUploader(sess).Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
		Body:   buffer,
	})

	return result.Location, err
}
