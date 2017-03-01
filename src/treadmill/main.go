package main

import (
	"bytes"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	sess := session.Must(session.NewSession())

	svc := s3.New(sess, aws.NewConfig().WithRegion("ap-southeast-2"))

	headParams := &s3.HeadObjectInput{
		Bucket: aws.String("falafel-test"),
		Key:    aws.String("ObjectKey404"),
	}
	headResp, headErr := svc.HeadObject(headParams)
	if awsErr, ok := headErr.(awserr.Error); ok {
		if awsErr.Code() == "NotFound" {
			headErr = nil
		}
	}
	if headErr != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(headErr.Error())
		return
	}

	fmt.Printf("\n\nheadResp = %#v\n\n", headResp)

	paramsPut := &s3.PutObjectInput{
		Bucket:      aws.String("falafel-test"), // Required
		Key:         aws.String("ObjectKey"),    // Required
		Body:        bytes.NewReader([]byte("PAYLOAD")),
		ContentType: aws.String("application/json"),
		Metadata: map[string]*string{
			"Key":     aws.String("MetadataValue"),
			"PrevKey": aws.String("MetadataValue"),
			"NextKey": aws.String("MetadataValue"),
		},
	}
	respPut, errPut := svc.PutObject(paramsPut)

	if errPut != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(errPut.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(respPut)

	paramsGet := &s3.GetObjectInput{
		Bucket: aws.String("falafel-test"), // Required
		Key:    aws.String("ObjectKey"),    // Required
		//IfMatch:                    aws.String("IfMatch"),
		//IfModifiedSince:            aws.Time(time.Now()),
		//IfNoneMatch:                aws.String("IfNoneMatch"),
		//IfUnmodifiedSince:          aws.Time(time.Now()),
		//PartNumber:                 aws.Int64(1),
		//Range:                      aws.String("Range"),
		//RequestPayer:               aws.String("RequestPayer"),
		//ResponseCacheControl:       aws.String("ResponseCacheControl"),
		//ResponseContentDisposition: aws.String("ResponseContentDisposition"),
		//ResponseContentEncoding:    aws.String("ResponseContentEncoding"),
		//ResponseContentLanguage:    aws.String("ResponseContentLanguage"),
		//ResponseContentType:        aws.String("ResponseContentType"),
		//ResponseExpires:            aws.Time(time.Now()),
		//SSECustomerAlgorithm:       aws.String("SSECustomerAlgorithm"),
		//SSECustomerKey:             aws.String("SSECustomerKey"),
		//SSECustomerKeyMD5:          aws.String("SSECustomerKeyMD5"),
		//VersionId:                  aws.String("ObjectVersionId"),
	}
	respGet, errGet := svc.GetObject(paramsGet)

	if errGet != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(errGet.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(respGet)
}
