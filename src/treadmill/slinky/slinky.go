package slinky

import (
	"bytes"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)


type Service interface {

  // InUse checks if the specified key is in use by Slinky
  InUse(key string) (bool, error)

  // Put stores the specified content in Slinky
  Put(key string, prevKey string, nextKey string, content []byte, contentType string) error

  // Get retrieves the specified content from Slinky
  Get(key string) (content []byte, contentType string, err error)

  // Stat returns the Info structure describing Slinky
  Stat() (Info, error)

}





func CreateService(region string, bucket string, env string) (Service, err) {
  sess, err := session.NewSession()
  if err != nil {
    return nil, fmt.Errorf("error creating aws session %v",err)
  }
  svc := s3.New(sess, aws.NewConfig().WithRegion(region))
  return svc{
    svc,
    region,
    bucket,
    env,
  }
}


type svc struct {
	s3svc *s3.S3
  region string
  bucket string
  env string
}


func (s Service) InUse(key string) (bool, error) {
  params := &s3.HeadObjectInput{
    Bucket:      aws.String(s.bucket),
		Key:         aws.String(key),
  }
  _, err := s.s3svc.HeadObject(params)
	if err != nil {
    if awsErr, ok := err.(awserr.Error); ok && awsErr.Code() == "NotFound" {
       return false, nil
    }
		return false, fmt.Printf("failed to check if this key is in use with error %v", err)
	}
  return true, nil
}

func (s Service) Put(key string, prevKey string, nextKey string, body []byte, contentType string) error {
	params := &s3.PutObjectInput{
		Bucket:      aws.String(s.bucket),
		Key:         aws.String(key),
		Body:        bytes.NewReader(body),
		ContentType: aws.String(contentType),
		Metadata: map[string]*string{
			"Key":     aws.String(key),
			"PrevKey": aws.String(prevKey),
			"NextKey": aws.String(nextKey),
		},
	}
	res, err := s.s3svc.PutObject(params)
	if err != nil {
		return err
	}
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
