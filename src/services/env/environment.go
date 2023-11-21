package env

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	awsCred "github.com/aws/aws-sdk-go/aws/credentials"

	"github.com/redis/go-redis/v9"
)

const (
	redisAddrKey = "REDIS_ADDR"
	redisPassKey = "REDIS_PASS"

	awsCredNameKey   = "AWS_CRED_NAME"
	awsCredSecretKey = "AWS_CRED_SECRET"
	awsEndpointKey   = "AWS_ENDPOINT"
	awsRegionKey     = "AWS_REGION"
)

func getEnv(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		println(key, value)
		return value
	}
	println("NO ENV VAR: " + key)
	return ""
}

func NewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     getEnv(redisAddrKey),
		Password: getEnv(redisPassKey),
		DB:       0,
	})
}

func NewAwsConfig() *aws.Config {
	return &aws.Config{
		Credentials: awsCred.NewStaticCredentials(
			getEnv(awsCredNameKey),
			getEnv(awsCredSecretKey),
			""),
		Endpoint:         aws.String(getEnv(awsEndpointKey)),
		Region:           aws.String(getEnv(awsRegionKey)),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
	}
}
