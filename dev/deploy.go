package main

import (
	"fmt"
	"mime"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const endpoint = "https://sfo2.digitaloceanspaces.com"
const bucket = "laniakea"
const keyPrefix = "nacelle/public"
const publicDir = "public"

var accessKey = os.Getenv("ACCESS_KEY")
var secretKey = os.Getenv("SECRET_KEY")

func main() {
	if err := mainErr(); err != nil {
		fmt.Fprint(os.Stderr, fmt.Sprintf("error: %s\n", err.Error()))
		os.Exit(1)
	}
}

func mainErr() error {
	s3Client := makeClient()

	for _, f := range []func(*s3.S3) error{clearAssets, uploadAssets} {
		if err := f(s3Client); err != nil {
			return err
		}
	}

	return nil
}

func makeClient() *s3.S3 {
	s3Config := &aws.Config{
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
		Endpoint:    aws.String(endpoint),
		Region:      aws.String("us-east-1"),
	}

	return s3.New(session.New(s3Config))
}

func clearAssets(s3Client *s3.S3) error {
	output, err := s3Client.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String(bucket),
		Prefix: aws.String(keyPrefix),
	})
	if err != nil {
		return err
	}

	var keys []string
	for _, object := range output.Contents {
		keys = append(keys, *object.Key)
	}

	return invokeParallel(keys, func(value string) error {
		_, err := s3Client.DeleteObject(&s3.DeleteObjectInput{
			Bucket: aws.String(bucket),
			Key:    &value,
		})

		return err
	})
}

func uploadAssets(s3Client *s3.S3) error {
	paths, err := gatherPaths()
	if err != nil {
		return err
	}

	return invokeParallel(paths, func(value string) error {
		f, err := os.Open(filepath.Join(publicDir, value))
		if err != nil {
			return err
		}
		defer f.Close()

		// Extract the mimetype from the inferred `<mimetype>; <encoding>`
		contentType := strings.Split(mime.TypeByExtension(filepath.Ext(value)), ";")[0]

		_, err = s3Client.PutObject(&s3.PutObjectInput{
			Bucket:      aws.String(bucket),
			Key:         aws.String(filepath.Join(keyPrefix, value)),
			Body:        f,
			ContentType: aws.String(contentType),
			ACL:         aws.String("public-read"),
		})

		return err
	})
}

func gatherPaths() (paths []string, _ error) {
	err := filepath.Walk(publicDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			paths = append(paths, strings.TrimPrefix(path, publicDir))
		}

		return err
	})

	return paths, err
}

func invokeParallel(values []string, f func(arg string) error) error {
	ch := make(chan string, len(values))
	for _, value := range values {
		ch <- value
	}
	close(ch)

	n := runtime.GOMAXPROCS(0)
	errs := make(chan error, n)
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for value := range ch {
				if err := f(value); err != nil {
					errs <- err
				}
			}
		}()
	}

	wg.Wait()
	close(errs)
	return <-errs
}
