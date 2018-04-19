// go get -u github.com/aws/aws-sdk-go
// ListObjectsInput{Prefix}
// https://blog.csdn.net/zhonglinzhang/article/details/77040478
package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var Session = session.Must(session.NewSession(aws.NewConfig().
	WithCredentials(credentials.NewStaticCredentials("accessKey", "secretKey", "")).
	WithRegion(endpoints.ApNortheast1RegionID)))

func main() {
	svc := s3.New(Session)

	pageNum := 0
	err := svc.ListObjectsPages(&s3.ListObjectsInput{
		Bucket: aws.String("thisismyawss3"),
	}, func(p *s3.ListObjectsOutput, last bool) (shouldContinue bool) {
		fmt.Println("Page,", pageNum)
		pageNum++

		for _, obj := range p.Contents {
			fmt.Println("Object:", *obj.Key)
		}
		return true
	})
	if err != nil {
		fmt.Println("failed to list objects", err)
		return
	}
}
