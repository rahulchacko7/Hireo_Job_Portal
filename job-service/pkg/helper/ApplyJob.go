package helper

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"

	cfg "Auth/pkg/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type Helper struct {
	cfg cfg.Config
}

func NewHelper(cfg cfg.Config) *Helper {
	return &Helper{cfg: cfg}
}

func (h *Helper) AddImageToAwsS3(file []byte, filename string) (string, error) {
	config, err := cfg.LoadConfig()
	if err != nil {
		return "", err
	}

	fmt.Println("pppppppp", config.DBHost)
	fmt.Println("print1", config.AWSRegion)
	fmt.Println("print2", config.Access_key_ID)
	fmt.Println("print3", config.Secret_access_key)

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(config.AWSRegion),
		Credentials: credentials.NewStaticCredentials(
			config.Access_key_ID,
			config.Secret_access_key,
			"",
		),
	})
	if err != nil {
		fmt.Println("erorrrr here", err)
		return "", err
	}

	uploader := s3manager.NewUploader(sess)
	bucketName := "hireojobbucket"

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filename),
		Body:   bytes.NewReader(file),
	})

	if err != nil {
		fmt.Println("erroorrrr 2", err)
		return "", err
	}

	url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucketName, config.AWSRegion, filename)
	return url, nil
}

func GenerateVideoCallKey(userID, oppositeUser int) (string, error) {
	currentTime := strconv.FormatInt(time.Now().UnixNano(), 10)
	key := strconv.Itoa(userID) + "_" + strconv.Itoa(oppositeUser) + "_" + currentTime
	hash := md5.Sum([]byte(key))
	keyString := hex.EncodeToString(hash[:])

	return keyString, nil
}
