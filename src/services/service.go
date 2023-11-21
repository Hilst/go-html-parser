package service

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	mdl "github.com/Hilst/go-ui-html-template/models"
	"github.com/Hilst/go-ui-html-template/services/env"

	"github.com/aws/aws-sdk-go/aws"
	awsSession "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"github.com/redis/go-redis/v9"
)

type Service struct {
	redis     *redis.Client
	awsConfig *aws.Config
}

func NewService(layoutRoot string) *Service {
	return &Service{
		env.NewRedisClient(),
		env.NewAwsConfig(),
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

func (s *Service) RequestLayout(layoutName string, ch chan mdl.LayoutResponse) {
	defer close(ch)
	session, ok := s.makeNewAWSSession(ch)
	if !ok {
		return
	}
	s3Client := s3.New(session)
	listObjsResult, ok := s.listS3Objects("screens", layoutName, s3Client, ch)
	if !ok {
		return
	}
	downloader := s3manager.NewDownloader(session)
	s.downloadS3Objs(downloader, listObjsResult, "screens", ch)
}

func (s *Service) makeNewAWSSession(ch chan mdl.LayoutResponse) (*awsSession.Session, bool) {
	if session, err := awsSession.NewSession(s.awsConfig); err == nil {
		return session, true
	} else {
		ch <- mdl.NewLayoutRespError(err)
		return nil, false
	}
}

func (s *Service) listS3Objects(bucket string, prefix string, client *s3.S3, ch chan mdl.LayoutResponse) ([]*s3.Object, bool) {
	listObjsInput := &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
		Prefix: aws.String(prefix),
	}
	listObjsResult, listObjsError := client.ListObjectsV2(listObjsInput)
	if listObjsError != nil {
		ch <- mdl.NewLayoutRespError(listObjsError)
		return nil, false
	}
	if *listObjsResult.KeyCount == 0 {
		ch <- mdl.NewLayoutRespError(errors.New("empty content"))
		return nil, false
	}
	return listObjsResult.Contents, true
}

func (s *Service) downloadS3Objs(downloader *s3manager.Downloader, s3Contents []*s3.Object, bucket string, ch chan mdl.LayoutResponse) bool {
	var buff *aws.WriteAtBuffer
	var getObjInput *s3.GetObjectInput
	var fileName string
	for _, object := range s3Contents {
		buff = aws.NewWriteAtBuffer([]byte{})
		getObjInput = &s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    object.Key,
		}
		_, err := downloader.Download(buff, getObjInput)
		if err != nil {
			ch <- mdl.NewLayoutRespError(err)
			return false
		}
		fileName = *object.Key
		fileName = strings.Split(fileName, "/")[1]
		fileName = strings.Trim(fileName, ".html")
		ch <- mdl.NewLayoutResp(string(buff.Bytes()), fileName)
	}
	return true
}
