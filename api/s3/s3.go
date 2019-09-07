package s3

import (
	"article/api/util"
	"bytes"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
)

type S3 struct {
	APPID  string
	SECRET string
}

type DaoInterface interface {
	PostImageToS3(c *gin.Context) []util.ImageName
	DeleteS3Image(imageNames []util.ImageName)
}

func NewS3(appid, secret string) *S3 {
	objs := &S3{APPID: appid, SECRET: secret}
	return objs
}

func (objs *S3) PostImageToS3(file *multipart.FileHeader, imageName string) error {
	creds := credentials.NewStaticCredentials(objs.APPID, objs.SECRET, "")

	cfg := aws.NewConfig().WithRegion("ap-northeast-1").WithCredentials(creds)
	svc := s3.New(session.New(), cfg)

	// form, _ := c.MultipartForm()

	// files := form.File["images[]"]

	// var imageNames []util.ImageName
	// imageName := util.ImageName{}

	// for _, file := range files {

	f, err := file.Open()

	if err != nil {
		log.Println(err)
	}

	defer f.Close()

	size := file.Size
	buffer := make([]byte, size)

	// u, err := uuid.NewRandom()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// uu := u.String()

	f.Read(buffer)
	fileBytes := bytes.NewReader(buffer)
	fileType := http.DetectContentType(buffer)
	path := "/media/" + imageName
	params := &s3.PutObjectInput{
		Bucket:        aws.String("article-s3-jpskgc"),
		Key:           aws.String(path),
		Body:          fileBytes,
		ContentLength: aws.Int64(size),
		ContentType:   aws.String(fileType),
	}
	resp, err := svc.PutObject(params)

	fmt.Printf("response %s", awsutil.StringValue(resp))

	//imageName.NAME = uu

	//imageNames = append(imageNames, imageName)
	// }

	return err
}

func (objs *S3) DeleteS3Image(imageName util.ImageName) error {

	creds := credentials.NewStaticCredentials(objs.APPID, objs.SECRET, "")

	cfg := aws.NewConfig().WithRegion("ap-northeast-1").WithCredentials(creds)
	svc := s3.New(session.New(), cfg)

	// for _, imageName := range imageNames {

	path := "/media/" + imageName.NAME

	input := &s3.DeleteObjectInput{
		Bucket: aws.String("article-s3-jpskgc"),
		Key:    aws.String(path),
	}

	result, err := svc.DeleteObject(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
		return err
	}

	fmt.Println(result)
	return err
	// }
}
