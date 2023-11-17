package service

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	mdl "github.com/Hilst/go-ui-html-template/models"
	"github.com/aws/aws-sdk-go/aws"
	awsCred "github.com/aws/aws-sdk-go/aws/credentials"
	awsSession "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/redis/go-redis/v9"
)

type Service struct {
	redis      redis.Client
	layoutRoot string
}

func NewService(layoutRoot string) *Service {
	rClient := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	return &Service{
		*rClient,
		layoutRoot,
	}
}

func (s *Service) RequestData(id string) mdl.DataResponse {
	ctx := context.Background()
	val, err := s.redis.Get(ctx, id).Result()
	if err != nil {
		return mdl.NewDataRespError(err)
	}
	var result mdl.MiddleDataResp
	err = json.Unmarshal([]byte(val), &result)
	return mdl.NewDataResp(result, err)
}

var s3Config = &aws.Config{
	Credentials:      awsCred.NewStaticCredentials("root", "password", ""),
	Endpoint:         aws.String("127.0.0.1:9000"),
	Region:           aws.String("us-east-1"),
	DisableSSL:       aws.Bool(true),
	S3ForcePathStyle: aws.Bool(true),
}

func (s *Service) RequestLayout(layoutName string, ch chan mdl.LayoutResponse) {
	defer close(ch)

	session, sessionErr := awsSession.NewSession(s3Config)
	if sessionErr != nil {
		ch <- mdl.NewLayoutRespError(sessionErr)
		return
	}
	s3Client := s3.New(session)
	listObjsInput := &s3.ListObjectsV2Input{
		Bucket: aws.String("screens"),
		Prefix: aws.String(layoutName),
	}
	listObjsResult, listObjsError := s3Client.ListObjectsV2(listObjsInput)
	if listObjsError != nil {
		ch <- mdl.NewLayoutRespError(listObjsError)
		return
	}
	if *listObjsResult.KeyCount == 0 {
		ch <- mdl.NewLayoutRespError(errors.New("empty content"))
		return
	}

	downloader := s3manager.NewDownloader(session)
	var buff *aws.WriteAtBuffer
	var getObjInput *s3.GetObjectInput
	var fileName string
	for _, object := range listObjsResult.Contents {
		buff = aws.NewWriteAtBuffer([]byte{})
		getObjInput = &s3.GetObjectInput{
			Bucket: aws.String("screens"),
			Key:    object.Key,
		}
		_, err := downloader.Download(buff, getObjInput)
		if err != nil {
			ch <- mdl.NewLayoutRespError(err)
			return
		}
		fileName = *object.Key
		fileName = strings.Split(fileName, "/")[1]
		fileName = strings.Trim(fileName, ".html")
		ch <- mdl.NewLayoutResp(string(buff.Bytes()), fileName)
	}
}
