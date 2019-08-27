package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func GetArticleService(c *gin.Context, db *sql.DB) []Article {
	var articles []Article

	results := GetArticleDao(db)

	article := Article{}
	for results.Next() {
		err := results.Scan(&article.ID, &article.UUID, &article.TITLE, &article.CONTENT)
		if err != nil {
			panic(err.Error())
		} else {
			articles = append(articles, article)
		}
	}
	return articles
}

func GetSingleArticleService(c *gin.Context, db *sql.DB) Article {

	article, rows := GetSingleArticleDao(c, db)

	for rows.Next() {
		imageName := ImageName{}
		err := rows.Scan(&imageName.NAME)
		if err != nil {
			panic(err.Error())
		}
		article.IMAGENAMES = append(article.IMAGENAMES, imageName)
	}

	return article
}

func DeleteArticleService(c *gin.Context, db *sql.DB) {
	DeleteArticleDao(c, db)
}

func PostService(c *gin.Context, db *sql.DB) string {
	u, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
	}
	uu := u.String()
	var article Article
	c.BindJSON(&article)
	PostDao(db, article, uu)

	return uu
}

func PostImageService(c *gin.Context) []ImageName {
	creds := credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), "")

	cfg := aws.NewConfig().WithRegion("ap-northeast-1").WithCredentials(creds)
	svc := s3.New(session.New(), cfg)

	form, _ := c.MultipartForm()

	files := form.File["images[]"]

	var imageNames []ImageName
	imageName := ImageName{}

	for _, file := range files {

		f, err := file.Open()

		if err != nil {
			log.Println(err)
		}

		defer f.Close()

		size := file.Size
		buffer := make([]byte, size)

		u, err := uuid.NewRandom()
		if err != nil {
			fmt.Println(err)
		}
		uu := u.String()

		f.Read(buffer)
		fileBytes := bytes.NewReader(buffer)
		fileType := http.DetectContentType(buffer)
		path := "/media/" + uu
		params := &s3.PutObjectInput{
			Bucket:        aws.String("article-s3-jpskgc"),
			Key:           aws.String(path),
			Body:          fileBytes,
			ContentLength: aws.Int64(size),
			ContentType:   aws.String(fileType),
		}
		resp, err := svc.PutObject(params)

		fmt.Printf("response %s", awsutil.StringValue(resp))

		imageName.NAME = uu

		imageNames = append(imageNames, imageName)
	}

	return imageNames
}

func PostImageToDBService(c *gin.Context, db *sql.DB) {
	PostImageToDBDao(c, db)

}
